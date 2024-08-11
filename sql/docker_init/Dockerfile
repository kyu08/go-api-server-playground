FROM mysql:8.4

COPY ./sql/docker_init /docker-entrypoint-initdb.d/
COPY ./sql/schema/schema.sql /docker-entrypoint-initdb.d/100_schema.sql
