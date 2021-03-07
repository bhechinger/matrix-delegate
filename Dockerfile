FROM golang:1.16-alpine AS build

RUN mkdir /app
ADD . /app/
WORKDIR /app
RUN CGO_ENABLED=0 go build -o service -ldflags "-w" app/well-known/main.go && \
    chmod 755 service

FROM scratch
MAINTAINER Brian Hechinger <wonko@4amlunch.net>
COPY --from=build /app/service /service

ENTRYPOINT ["/service"]