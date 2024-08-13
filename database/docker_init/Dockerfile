FROM mysql:8.4

COPY ./database/docker_init /docker-entrypoint-initdb.d/
COPY ./database/schema/schema.sql /docker-entrypoint-initdb.d/100_schema.sql
