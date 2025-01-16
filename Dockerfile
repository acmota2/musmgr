FROM postgres:17.2
ENV POSTGRES_USER docker
ENV POSTGRES_PASSWORD docker
ENV POSTGRES_DB docker
COPY ./backend/schema.sql /docker-entrypoint-initdb.d/

