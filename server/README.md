# aldb-server

| 文件夹/文件     | 说明          | 描述                |
|------------|-------------|-------------------|
| api/v1     | api层v1版本接口  | api层v1版本接口        |
| config     | 配置包         | 从json文件中读取配置信息    |
| docs       | swagger文档目录 | swagger自动生成的api文档 |
| middleware | 中间件层        | 用于存放gin中间件代码      |
| model      | 模型层         | 绑定前端数据，与数据库交互     |
| --database | 数据库层        | 封装数据库具体操作         |
| router     | 路由层         | 构建路由信息            |
| util       | 工具包         | 包含响应代码等           |
| main.go    | 执行入口        | 执行入口              |

## Setup

```bash
git clone https://github.com/SukiEva/aldb.git
cd server
go generate
go build -o server main.go # windows: go build -o server.exe main.go )
# run
./server # windows: server.exe
```


