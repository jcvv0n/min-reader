FROM golang:alpine3.17
ADD ./ /opt/app/
WORKDIR /opt/app
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories && apk update
RUN apk add build-base
RUN go env -w GOPROXY=https://proxy.golang.com.cn,https://goproxy.cn,direct
RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -installsuffix cgo -o readerapp .

FROM alpine:latest
MAINTAINER jcvv0n
WORKDIR /opt/app/
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories && apk update
RUN apk add ca-certificates tzdata
ENV TZ Asia/Shanghai
COPY --from=0 /opt/app/readerapp ./
COPY --from=0 /opt/app/db ./db/
COPY --from=0 /opt/app/template ./template/
EXPOSE 80
ENTRYPOINT ["./readerapp"]