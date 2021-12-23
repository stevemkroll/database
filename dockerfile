FROM postgres

COPY ./internal/migration/000_config.sql /docker-entrypoint-initdb.d/
COPY ./internal/migration/001_accounts.sql /docker-entrypoint-initdb.d/

ENV POSTGRES_PASSWORD="password"
ENV POSTGRES_USER="user"
ENV POSTGRES_DB="db"
