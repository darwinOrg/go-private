FROM registry.cn-shanghai.aliyuncs.com/star_base/golang-x86:1.19.1 as builder

LABEL stage=gobuilder

# 环境变量
ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct \
    CGO_ENABLED=0 \
    GOOS=linux    \
    GOARCH=arm64

WORKDIR /application

COPY . .
RUN go env
RUN go mod download

RUN go build -ldflags "-s -w" -o /application/build/go-private main.go

FROM registry.cn-shanghai.aliyuncs.com/star_base/golang-x86:1.19.1

WORKDIR /target

# 复制编译后的程序
COPY --from=builder /application/build/go-private /target/go-private
COPY --from=builder /application/resources/ /target/resources
COPY --from=builder /application/certs/ /target/certs
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

EXPOSE 443
ENTRYPOINT ["/target/go-private"]