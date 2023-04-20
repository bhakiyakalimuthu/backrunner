FROM golang:1.19-alpine3.17 as builder
ARG VERSION

WORKDIR /build
COPY go.mod ./
COPY go.sum ./
RUN go mod download
ADD . .

RUN apk add --no-cache
RUN --mount=type=cache,target=/root/.cache/go-build CGO_ENABLED=0 go build -trimpath -ldflags "-s -X main.buildVersion=${VERSION}" -v -o backrunner ./cmd/main.go


FROM alpine:latest
WORKDIR /app
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /build/backrunner /app/backrunner
RUN chmod +x /app/backrunner
CMD ["/app/backrunner"]
