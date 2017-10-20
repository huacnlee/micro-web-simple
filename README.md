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
