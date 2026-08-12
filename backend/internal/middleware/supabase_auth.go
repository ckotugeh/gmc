package middleware

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// ---------------------------------------------------------------------------
// This file verifies Supabase Auth access tokens instead of your own
// hand-rolled JWTs (see auth.go). It expects your Supabase project to use
// the current, recommended asymmetric (ES256) JWT signing keys — not the
// legacy shared-secret (HS256) ones. See:
// https://supabase.com/docs/guides/auth/jwts
//
// It is NOT wired into main.go yet — see the note at the bottom of this
// file about the user-ID mismatch that needs a decision before this
// replaces internal/middleware/auth.go.
// ---------------------------------------------------------------------------

type jwk struct {
	Kty string `json:"kty"`
	Kid string `json:"kid"`
	Crv string `json:"crv"`
	X   string `json:"x"`
	Y   string `json:"y"`
}

type jwksResponse struct {
	Keys []jwk `json:"keys"`
}

type jwksCache struct {
	mu        sync.RWMutex
	keys      map[string]*ecdsa.PublicKey
	fetchedAt time.Time
}

var supabaseJWKS = &jwksCache{keys: map[string]*ecdsa.PublicKey{}}

const jwksTTL = 10 * time.Minute

func jwksURL() (string, error) {
	base := strings.TrimRight(os.Getenv("SUPABASE_URL"), "/")
	if base == "" {
		return "", fmt.Errorf("SUPABASE_URL is not set")
	}
	return base + "/auth/v1/.well-known/jwks.json", nil
}

func (c *jwksCache) getKey(kid string) (*ecdsa.PublicKey, error) {
	c.mu.RLock()
	key, ok := c.keys[kid]
	stale := time.Since(c.fetchedAt) > jwksTTL
	c.mu.RUnlock()

	if ok && !stale {
		return key, nil
	}

	if err := c.refresh(); err != nil {
		if ok {
			return key, nil // fall back to a possibly-stale key rather than fail outright
		}
		return nil, err
	}

	c.mu.RLock()
	defer c.mu.RUnlock()
	key, ok = c.keys[kid]
	if !ok {
		return nil, fmt.Errorf("no matching JWKS key for kid %q", kid)
	}
	return key, nil
}

func (c *jwksCache) refresh() error {
	url, err := jwksURL()
	if err != nil {
		return err
	}

	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("fetching Supabase JWKS: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("fetching Supabase JWKS: unexpected status %d", resp.StatusCode)
	}

	var parsed jwksResponse
	if err := json.NewDecoder(resp.Body).Decode(&parsed); err != nil {
		return fmt.Errorf("decoding Supabase JWKS: %w", err)
	}

	keys := make(map[string]*ecdsa.PublicKey, len(parsed.Keys))
	for _, k := range parsed.Keys {
		if k.Kty != "EC" || k.Crv != "P-256" {
			continue // only ES256 (Supabase's current default) is handled here
		}
		pub, err := ecPublicKeyFromJWK(k)
		if err != nil {
			continue
		}
		keys[k.Kid] = pub
	}

	if len(keys) == 0 {
		return fmt.Errorf("no usable ES256 keys in Supabase JWKS — project may still be on legacy HS256 (enable asymmetric JWT signing keys in Settings -> API -> JWT Keys)")
	}

	c.mu.Lock()
	c.keys = keys
	c.fetchedAt = time.Now()
	c.mu.Unlock()
	return nil
}

func ecPublicKeyFromJWK(k jwk) (*ecdsa.PublicKey, error) {
	xBytes, err := base64.RawURLEncoding.DecodeString(k.X)
	if err != nil {
		return nil, err
	}
	yBytes, err := base64.RawURLEncoding.DecodeString(k.Y)
	if err != nil {
		return nil, err
	}
	return &ecdsa.PublicKey{
		Curve: elliptic.P256(),
		X:     new(big.Int).SetBytes(xBytes),
		Y:     new(big.Int).SetBytes(yBytes),
	}, nil
}

// SupabaseClaims are the JWT claims Supabase Auth puts on a user's access token.
type SupabaseClaims struct {
	Email string `json:"email"`
	Role  string `json:"role"` // Postgres role, e.g. "authenticated" — not your app's user role
	jwt.RegisteredClaims
}

// SupabaseAuthMiddleware verifies a Supabase Auth access token and stores
// the caller's Supabase user UUID (claims.Subject) and email in the Gin
// context as "supabaseUserID" and "email". Requires SUPABASE_URL.
func SupabaseAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")
		if header == "" || !strings.HasPrefix(header, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing or malformed Authorization header"})
			return
		}
		tokenString := strings.TrimPrefix(header, "Bearer ")

		claims := &SupabaseClaims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodECDSA); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
			}
			kid, _ := t.Header["kid"].(string)
			return supabaseJWKS.getKey(kid)
		})
		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid or expired token"})
			return
		}

		c.Set("supabaseUserID", claims.Subject) // UUID string, matches auth.users.id
		c.Set("email", claims.Email)
		c.Next()
	}
}

// ---------------------------------------------------------------------------
// IMPORTANT — not wired in yet, needs a decision:
//
// Supabase Auth identifies users by a UUID (claims.Subject / auth.users.id).
// Your app's `users` table (and every doctor_id/patient_id/user_id column
// across ~30 tables) currently uses BIGINT ids. Those two ID spaces don't
// match, so simply swapping this middleware in for internal/middleware/auth.go
// isn't enough on its own. See the chat response for the options.
// ---------------------------------------------------------------------------
