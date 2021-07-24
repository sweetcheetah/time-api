FROM golang:1.16 as builder
COPY . /src
WORKDIR /src

RUN go test ./... -v -cover
RUN CGO_ENABLED=0 GOOS=linux go build -o time-api .

FROM scratch
WORKDIR /

COPY --from=builder /src/env/dev.env /.env
COPY --from=builder /src/swagger-ui/ /swagger-ui
COPY --from=builder /src/time-api /time-api

EXPOSE 8080
ENV PORT=8080

CMD ["/time-api"]
