# syntax=docker/dockerfile:1
FROM postgres:14-alpine

ENV POSTGRES_HOST_AUTH_METHOD trust

COPY init.sql /docker-entrypoint-initdb.d/

EXPOSE 5432