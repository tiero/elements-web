###
## Build the Go App
###
FROM golang:1.18-buster AS builder

ARG TARGETOS
ARG TARGETARCH

WORKDIR /src

COPY . .

RUN GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -o /bin/web cmd/web/*

FROM debian:buster-slim
COPY --from=builder /bin/web .

COPY layout.html .
ADD static /static

EXPOSE 8080

CMD [ "./web" ]
