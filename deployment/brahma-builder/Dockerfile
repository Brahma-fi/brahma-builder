FROM golang:1.23-alpine AS builder
ARG TAGS="prod"
ARG TOKEN

# Install git + SSL ca certificates.
# Git is required for fetching the dependencies.
# Ca-certificates is required to call HTTPS endpoints.
RUN apk update && apk add --no-cache make git tzdata ca-certificates && update-ca-certificates
# RUN apk update && apk add --no-cache make git tzdata
ENV USER=appuser
ENV UID=10001

WORKDIR $GOPATH/src/

# See https://stackoverflow.com/a/55757473/12429735
RUN adduser \
    --disabled-password \
    --gecos "" \
    --home "/nonexistent" \
    --shell "/sbin/nologin" \
    --no-create-home \
    --uid "${UID}" \
    "${USER}"

COPY ../.. .

# Fetch dependencies.# Using go mod with go >= 1.11
RUN make build-linux-release tags=${TAGS}

############################
# STEP 2 build a small image
############################
FROM scratch

# Import from builder.
# only if this api handle tls termination
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group

# Copy our static executable
COPY --from=builder /go/src/_bin/release/linux/brahma-builder /go/bin/brahma-builder

# Use an unprivileged user.
USER appuser:appuser

EXPOSE 8080
ENTRYPOINT ["/go/bin/brahma-builder"]
