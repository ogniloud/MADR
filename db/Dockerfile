# syntax=docker/dockerfile:1
FROM postgres:16-alpine

ENV POSTGRES_HOST_AUTH_METHOD trust

COPY init.sql /docker-entrypoint-initdb.d/

EXPOSE 5432