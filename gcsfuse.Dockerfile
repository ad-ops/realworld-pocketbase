FROM golang:1.21.1-bookworm as build

WORKDIR /go/src/app
COPY . .

RUN go mod download
RUN go vet -v
RUN go test -v

RUN CGO_ENABLED=0 go build -o /go/bin/app

FROM golang:1.21.1-bookworm
RUN apt-get update && apt-get install -y \
    curl \
    gnupg \
    lsb-release \
    tini && \
    gcsFuseRepo=gcsfuse-`lsb_release -c -s` && \
    echo "deb http://packages.cloud.google.com/apt $gcsFuseRepo main" | \
    tee /etc/apt/sources.list.d/gcsfuse.list && \
    curl https://packages.cloud.google.com/apt/doc/apt-key.gpg | \
    apt-key add - && \
    apt-get update && \
    apt-get install -y gcsfuse && \
    apt-get clean

WORKDIR /app
COPY --from=build /go/bin/app ./
COPY --from=build /go/src/app/gcsfuse_run.sh ./
# Ensure the script is executable
RUN chmod +x /app/gcsfuse_run.sh

# Use tini to manage zombie processes and signal forwarding
# https://github.com/krallin/tini
ENTRYPOINT ["/usr/bin/tini", "--"]

# Pass the wrapper script as arguments to tini
CMD ["/app/gcsfuse_run.sh"]