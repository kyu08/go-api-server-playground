FROM mysql:8.4

COPY ./internal/infrastructure/database/docker_init /docker-entrypoint-initdb.d/
COPY ./internal/infrastructure/database/schema/schema.sql /docker-entrypoint-initdb.d/100_schema.sql
