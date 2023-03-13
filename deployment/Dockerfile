FROM alpine:latest
MAINTAINER SukiEva<dev.suki@outlook.com>
ENV VERSION 1.0

# 在容器根目录 创建一个 apps 目录
WORKDIR /apps

# 挂载容器目录
# VOLUME ["/apps"]

# 拷贝当前目录下可执行文件
COPY server /apps/server

# 拷贝配置文件到容器中
COPY config.json /apps/config.json

# 设置时区为上海
RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
RUN echo 'Asia/Shanghai' >/etc/timezone

# 设置编码
ENV LANG C.UTF-8

# 暴露端口
EXPOSE 8888

# 运行golang程序的命令
ENTRYPOINT /apps/server