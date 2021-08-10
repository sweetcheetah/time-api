FROM golang:1.16 as builder
WORKDIR /src

# Enable Go's DNS resolver to read from /etc/hosts
RUN echo "hosts: files dns" > /etc/nsswitch.conf.min

# Create a minimal passwd so we can run as non-root in the container
RUN echo "nobody:x:65534:65534:Nobody:/:" > /etc/passwd.min

# Fetch latest CA certificates
RUN apt-get update && \
    apt-get install -y ca-certificates

# Only download Go modules (improves build caching)
COPY go.mod go.sum ./
RUN go mod download

COPY . /src

RUN go test ./... -v -cover
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags='-s -w' -o time-api .

FROM scratch
WORKDIR /

COPY --from=debian /usr/share/zoneinfo/ /usr/share/zoneinfo/

COPY --from=builder /etc/nsswitch.conf.min /etc/nsswitch.conf
COPY --from=builder /etc/passwd.min /etc/passwd
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

COPY --from=builder /src/env/dev.env /.env
COPY --from=builder /src/swagger-ui/ /swagger-ui
COPY --from=builder /src/time-api /time-api

EXPOSE 8080
ENV PORT=8080
ENV TIMEZONE="America/New_York"
USER nobody

ENTRYPOINT ["/time-api"]
