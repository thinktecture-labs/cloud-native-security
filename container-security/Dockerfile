#build stage
FROM golang:alpine AS builder
RUN apk add --no-cache git
WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...
RUN go build -o /go/bin/app -v ./...

#final stage
FROM alpine:latest
## also install curl for healthcheck
RUN apk update && apk add ca-certificates curl

COPY --from=builder /go/bin/app /app

# non-root
RUN adduser --disabled-password bob && chown -R bob /app
USER bob
# end-non-root

# healthcheck
HEALTHCHECK CMD curl --fail http://localhost:3000/healthz/alive || exit 1
# end healthcheck

ENTRYPOINT /app
LABEL Name=containersecurity Version=0.0.2
EXPOSE 3000
