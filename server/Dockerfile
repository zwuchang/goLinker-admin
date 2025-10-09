FROM alpine AS ca

RUN apk add -U --no-cache ca-certificates

FROM golang:1.23 AS build

ENV GOPROXY=https://goproxy.cn
ENV GO111MODULE=on

WORKDIR /go/cache

ADD go.mod .
ADD go.sum .
RUN go mod download

WORKDIR /build

ADD . .

RUN GOOS=linux CGO_ENABLED=0 go build -o server .

FROM busybox AS prod

COPY --from=ca /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
COPY --from=build /build/server /goLinker-admin/server
COPY --from=build /build/config.yaml /goLinker-admin/config/config.yaml
COPY --from=build /build/resource /goLinker-admin/resource

ENV STAGE=production

EXPOSE 8888/tcp

WORKDIR /goLinker-admin
CMD ["/goLinker-admin/server","-c","/goLinker-admin/config/config.yaml"]