# 系统部署说明

如果不想手动编译，可前往项目的[Github发布页](https://github.com/SukiEva/aldb/releases)下载编译好的文件。

**注：下载仅限后端，前端需要修改文件自行编译**

## 编译

系统前后端分离，需要分别编译才能用于生产环境。

### 后端

后端项目文件在`server`目录，具体目录说明如下：

| 文件夹/文件 | 说明            | 描述                       |
| ----------- | --------------- | -------------------------- |
| api/v1      | api层v1版本接口 | api层v1版本接口            |
| config      | 配置包          | 从json文件中读取配置信息   |
| docs        | swagger文档目录 | swagger自动生成的api文档   |
| middleware  | 中间件层        | 用于存放gin中间件代码      |
| model       | 模型层          | 绑定前端数据，与数据库交互 |
| --database  | 数据库层        | 封装数据库具体操作         |
| router      | 路由层          | 构建路由信息               |
| util        | 工具包          | 包含响应代码等             |
| main.go     | 执行入口        | 执行入口                   |

**确保Go语言运行环境**

Go语言支持交叉编译，可以通过环境变量编译不同系统的可执行文件。

> - Windows:
    >   - `SET GOOS=windows`
          >     `SET GOARCH=amd64`
>
> - Mac:
    >
    >   - `SET GOOS=darwin`
          >     `SET GOARCH=amd64`
>
> - Linux:
    >
    >   - `SET GOOS=linux`
          >     `SET GOARCH=amd64`
>
> - Android:
    >
    >   - `SET GOOS=android`
          >     `SET GOARCH=arm64`

```shell
git clone https://github.com/SukiEva/aldb.git
cd server
go generate
go build -o server main.go # windows: go build -o server.exe main.go
# run
./server # windows: server.exe
```

### 前端

前端项目文件在`web`目录，具体目录说明如下：

| 文件夹/文件 | 说明        | 描述          |
| ----------- | ----------- | ------------- |
| api         | api层       | api层请求接口 |
| assets      | 静态文件    | 静态文件      |
| components  | 组件        | 包含Header等  |
| router      | 路由        | 构建路由信息  |
| styles      | 样式表      | 存放SCSS文件  |
| views       | 页面        | 包含Home等    |
| App.vue     | Vue文件入口 | Vue文件入口   |
| main.ts     | Vue执行入口 | Vue执行入口   |

**确保Nodejs运行环境**

在运行`npm run build`前，修改`web/.env.production`的API地址为服务器地址或域名。

```ABAP
ENV = 'production'
VITE_APP_BASE_API=http://101.35.145.144:8888 
```

```bash
# 安装依赖
npm install
# dev运行，用于测试环境
npm run dev
# 编译dist文件，用于生产环境
npm run build
```

<div STYLE="page-break-after: always;"></div>

## 部署

### 后端

**确保已有MongoDB数据库**

在编译生成的文件同目录下新建`config.json`配置文件

```json
{
  "port": 8888,
  "mongo": {
    "uri": "mongodb://localhost:27017",
    "db": "aldb"
  },
  "cos": {
    "bucket": "",
    "region": "",
    "secretID": "",
    "secretKey": "",
    "baseURL": "",
    "pathPrefix": "aldb"
  },
  "log": {
    "level": "info",
    "filename": "app.log",
    "maxsize": 200,
    "max_age": 7,
    "max_backups": 10
  }
}
```

#### 直接运行

后端编译生成的文件是Linux的可执行二进制文件，直接运行`./server`即可，可以配合守护进程后台启用。

```shell
chmod 0777 server
./server
```

<div STYLE="page-break-after: always;"></div>

#### Docker

- 新建`Dockerfile`，端口需要改为和配置文件中相同，默认`8888`。

```dockerfile
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
```

- 构建镜像`docker build -t aldb .`

- 移动文件，并授予执行权限

```shell
mkdir /data/aldb
mv server /data/aldb/server
mv config.json /data/aldb/config.json
chmod 0777 /data/aldb/server
```

- Docker运行，**端口修改与配置文件相同**

```shell
docker run -p 8888:8888 -v /data/aldb:/apps --name aldb --network host --env GIN_MODE=release -d aldb
```


<div STYLE="page-break-after: always;"></div>

### 前端

编译可以得到`dist`文件夹。

因为是纯静态页面，可以直接通过`nginx`或者`github page`部署。

