FROM golang:1.20.4-bullseye AS build

WORKDIR /workspace

COPY ./go.mod ./go.mod
COPY ./main.go ./main.go

RUN go build \
  -ldflags="-linkmode \"external\" -extldflags=-static" \
  -o go-minimal-web-server

FROM scratch
COPY --from=build /workspace/go-minimal-web-server /usr/bin/go-minimal-web-server
CMD ["/usr/bin/go-minimal-web-server"]
