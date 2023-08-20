# ==== build stage ====
FROM golang:1.21-bullseye as build
WORKDIR /app
COPY . .
RUN apt-get update && apt-get install -y --no-install-recommends \
    build-essential wget unzip
ENV GOARCH=arm64
ENV GOOS=linux
RUN wget -qO /app/libduckdb.zip https://github.com/duckdb/duckdb/releases/download/v0.8.1/libduckdb-linux-aarch64.zip \
    && unzip -o /app/libduckdb.zip -d /app \
    && rm /app/libduckdb.zip
RUN CGO_ENABLED=1 CGO_LDFLAGS="-L/app" go build -tags=duckdb_use_lib -o server server.go

# ==== final stage ====
FROM debian:bullseye-slim
WORKDIR /var/task
COPY --from=build /app/server ./
COPY --from=build /app/libduckdb.so /lib/
# Set the library search path
ENV LD_LIBRARY_PATH=/lib
# set the home directory
ENV HOME=/tmp
ENTRYPOINT ["/var/task/server"]

