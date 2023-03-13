# aldb

*MongoDB-Based Annotated Image Data Storage System*

## Setup

### Download

Download by [releases](https://github.com/SukiEva/aldb/releases)

### Config

Create or rename `server/config/config.json.example` to `config.json`

### Run

```bash
chmod 0777 server
export GIN_MODE=release
./server
```

## API Docs

```bash
cd server
swag init
```

API Docs will automatically generated at `/swagger/index.html`

## Deployment

[deployment](https://github.com/SukiEva/aldb/tree/master/deployment)

## Compile 

### server

[server](https://github.com/SukiEva/aldb/tree/master/server)

### Web

[web](https://github.com/SukiEva/aldb/tree/master/web)

