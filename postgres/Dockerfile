FROM postgres:latest

RUN mkdir -p /tmp/psql_data/

COPY setup.sql /docker-entrypoint-initdb.d/10-init.sql
