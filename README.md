演示基于 Micro 架构实现 Web 微服务
-------------------------------

## 准备

1. 启动 `consul agent`

```bash
$ consul agent -dev
```

2. 启动 `micro web`

```bash
$ micro web
```

3. 启动应用

```bash
$ go get github.com/huacnlee/micro-web-simple && micro-web-simple
```

4. 通过 Micro Web Proxy 的代理，访问应用的页面

http://localhost:8082/simple/

> simple 这点原理是，main.go 里面注册应用的服务名称是 `go.micro.web.simple`，micro web 会自动将这类服务变为一个路径 /<service>

将会通过 micro web proxy 访问到应用服务器的相应页面。
