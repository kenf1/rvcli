FROM golang:1.24.5-alpine3.22
RUN go telemetry off

RUN apk update && \
	apk add --no-cache curl git make