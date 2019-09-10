# Build Go Server Binary
FROM golang:1.13

ARG SERVICE_NAME
ARG VERSION

WORKDIR /project

COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go install -v \
            -ldflags="-w -s -X main.version=${VERSION} -X main.serviceName=${SERVICE_NAME}" \
            ./cmd/notify

FROM alpine:latest
COPY --from=0 /go/bin/notify /bin/notifyer
ENTRYPOINT ["/bin/notifyer"]