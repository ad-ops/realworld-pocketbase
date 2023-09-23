FROM golang:1.21.1-bookworm as build

WORKDIR /go/src/app
COPY . .

RUN go mod download
RUN go vet -v
RUN go test -v

RUN CGO_ENABLED=0 go build -o /go/bin/app

# Minimal executable image
FROM gcr.io/distroless/static-debian12
COPY --from=build /go/bin/app /
CMD ["/app", "serve", "--http", "0.0.0.0:8080"]