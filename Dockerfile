# Stage 1: Build frontend
FROM node:18-alpine AS frontend-builder

WORKDIR /app/frontend
COPY frontend/package*.json ./
RUN npm ci
COPY frontend/ .
RUN npm run build:prod

# Stage 2: Build backend
FROM golang:1.22-alpine AS backend-builder

ARG APP=joylivedashboard
ARG VERSION=v1.0.0
ARG RELEASE_TAG=${VERSION}
ARG GOPROXY="https://proxy.golang.org,direct"

# Install the required packages
RUN apk add --no-cache gcc musl-dev sqlite-dev

# Set CGO_CFLAGS to enable large file support
ENV CGO_CFLAGS="-D_LARGEFILE64_SOURCE"
ENV GOPROXY=${GOPROXY}

WORKDIR /go/src/${APP}
COPY . .

# Copy frontend build output
COPY --from=frontend-builder /app/frontend/dist ./frontend/dist

# Build the application
RUN CGO_ENABLED=1 go build -ldflags "-w -s -X main.VERSION=${RELEASE_TAG}" -o ./${APP} .

# Stage 3: Production image
FROM alpine
ARG APP=joylivedashboard

# Install ca-certificates, timezone data, and SQLite runtime library
RUN apk add --no-cache ca-certificates tzdata sqlite-libs

WORKDIR /app

# Copy binary
COPY --from=backend-builder /go/src/${APP}/${APP} /usr/bin/${APP}

# Copy frontend static files
COPY --from=frontend-builder /app/frontend/dist /app/dist

# Copy configuration files
COPY configs /app/configs

EXPOSE 8040

ENTRYPOINT ["/usr/bin/joylivedashboard", "start", "-d", "/app/configs", "-c", "prod", "-s", "/app/dist"]
