FROM mysql:8.4

COPY ./internal/database/docker_init /docker-entrypoint-initdb.d/
COPY ./internal/database/schema/schema.sql /docker-entrypoint-initdb.d/100_schema.sql
