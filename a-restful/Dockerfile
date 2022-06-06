FROM golang:1.16 as BUILDER
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN make test
RUN make build

FROM alpine:3.12
WORKDIR /app
COPY --from=BUILDER /app/main main
ENTRYPOINT ["./main"]