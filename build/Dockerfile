# 使用轻量级的 Alpine 镜像运行程序
FROM registry.cn-hangzhou.aliyuncs.com/daemon_public/alpine:latest

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories && \
    apk update && apk add tzdata && ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
    echo "Asia/Shanghai" > /etc/timezone

# 设置工作目录为 /app
WORKDIR /app

# 从 builder 镜像复制编译后的二进制文件到容器中
COPY ./basic /app/
COPY config.yml /app/config.yml

# 启动 Go 程序
CMD ["./basic"]
