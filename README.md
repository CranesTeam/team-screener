# team-screener

Team screener

version 0.0.1-SNAPSHOT




## Init migrations:
1. Install migrate tool
2. run postgres
3. add uuid extention  
`docker exec -it 9a7b2d429cfc /bin/bash` . 
`psql -U postgres`  
` CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`  
4. run script 
`migrate -path ./schema -database 'postgres://postgres:pwd@localhost:5432/postgres?sslmode=disable' up`


---

It might help u to read more userful materials:
1. Create unique index in [postgres](https://postgrespro.ru/docs/postgresql/9.6/sql-createindex)
2. Generate uuid in [postgres](https://www.postgresql.org/docs/current/uuid-ossp.html)
3. Golang migrate [tool](https://github.com/golang-migrate/migrate)
4. Golang migrate [installation](https://github.com/golang-migrate/migrate/blob/master/cmd/migrate/README.md)
5. Docker compose with [Ppostgres](https://github.com/IliaEre/composes/blob/main/db/postgres-compose.yaml)
6. Postgres uuid extention [problem](https://stackoverflow.com/questions/22446478/extension-exists-but-uuid-generate-v4-fails) 
7. Generate uuid as default extrernal key. [link](https://dba.stackexchange.com/questions/122623/default-value-for-uuid-column-in-postgres)

--- 

# Enjoy!