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