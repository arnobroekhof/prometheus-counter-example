# syntax=docker/dockerfile:1
FROM golang:1.17 as gobuilder
WORKDIR /go/src/github.com/arnobroekhof/prometheus-counter-example/
ADD . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o prometheus-counter-example .


FROM debian:stable-slim
COPY --from=gobuilder /go/src/github.com/arnobroekhof/prometheus-counter-example/prometheus-counter-example /
RUN chmod +x /prometheus-counter-example


ENTRYPOINT ["/prometheus-counter-example"]
