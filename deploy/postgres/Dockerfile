FROM postgres:13.3-alpine

# Set password for postgres user
ENV POSTGRES_USER postgres
ENV POSTGRES_PASSWORD 12345

# Copy schema initialization file
COPY schema.sql /docker-entrypoint-initdb.d/1.sql

# Expose the PostgreSQL port
EXPOSE 5432

# Run postgres
CMD ["postgres"]
