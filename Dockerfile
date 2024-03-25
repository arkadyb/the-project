# build

FROM golang:1.22-alpine3.19 AS build

RUN apk update
RUN apk add make

COPY . /build
WORKDIR /build

RUN make build

# publish

FROM alpine:3.19

RUN apk add ca-certificates

COPY --from=build /build/www www
COPY --from=build /build/dist/app app

CMD ["/app"]
