## Test user's registration endpoints 

<pre>
curl -X POST http://localhost:8080/api/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "full_name":"Charles Otugeh",
    "username":"charles",
    "email":"charles@example.com",
    "password":"Password123"
  }'
  </pre>

  ## Testing user's Login endpoints
  <pre>
  
  curl -X POST http://localhost:8080/api/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "email":"ckotugeh@gmail.com",
    "password":"Password123"
  }'
  </pre>
  ## Implementing middleware auth.go
  <pre>
    What the file will handle: 
    1. Reads the Authorization header.
    2. Verifies the Bearer format.
    3. Validates the JWT.
    4. Extracts the claims.
    5. Stores the authenticated user's data in the Gin context.
  </pre>
  ## Testing functionallity of midleware auth
  <pre>
    curl http://localhost:8080/api/me # broken token
    curl http://localhost:8080/api/me \
    -H "Authorization: Bearer invalid-token" # Invalid token
    curl http://localhost:8080/api/me \
    -H "Authorization: Bearer YOUR_TOKEN" //Valid token
  </pre>
  ## Branck profile

  <pre>
    Sprint Goal

By the end of this branch, a doctor should be able to:

Register
Login
Authenticate with JWT
Create a professional profile
View their profile
Update their profile

After that, we'll build Communities.
  </pre>
  ## repository
  <pre>
    Repository responsibilities

At this layer, keep it focused on CRUD operations:

Create() → insert a profile.
GetByUserID() → fetch the logged-in doctor's profile.
GetByID() → fetch a public profile by ID.
Update() → save profile changes.
Delete() → remove a profile if needed.
  </pre>

  ## Create Community
  ```
curl -X POST http://localhost:8080/api/communities \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "name":"General Medicine",
    "description":"Community for physicians and medical professionals",
    "category":"Medicine",
    "is_private":false
  }'

  ```