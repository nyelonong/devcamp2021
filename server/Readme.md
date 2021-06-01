1. docker build -t devcamp-postgres-db ./
2. docker run -d --name devcamp-db -p 5432:5432 devcamp-postgres-db
3. docker exec it devcamp-db /bin/bash
4. psql -d devcamp -U postgres