# 启动编译环境
FROM golang:1.18-alpine AS builder

# 配置编译环境
RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.cn,direct

# 拷贝源代码到镜像中
COPY . /go/src/happycar/server

# 编译
WORKDIR /go/src/happycar/server
RUN go mod download && go mod verify
RUN go install ./blob/...

FROM alpine
COPY --from=builder /go/bin/blob /bin/blob
ENV ADDR=:8081

# 申明暴露的端口
EXPOSE 8081

# 设置服务入口
ENTRYPOINT [ "/bin/blob" ]
