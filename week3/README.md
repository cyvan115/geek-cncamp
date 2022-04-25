### 步骤

#### 1. 打包 http_server

```shell
# alias linuxgo="GOOS=linux GOARCH=386 go"
linuxgo build -o httpserver
```

#### 2. 编写 Dockerfile

```Dockerfile
FROM alpine
# 最小镜像构建

COPY ./httpserver /app/
# copy 可执行文件

ENTRYPOINT /app/httpserver
# 指定入口操作
```

也可以参考多阶段构建，使用 `golang` 作为基础镜像，设置 `go proxy`，拉取 `github` 代码，`cd` 到 `week3` 后做对应构建。

在基于最小镜像构建 `httpserver` 的最终镜像

#### 3. 构建

```bash
docker build -t cyvan115/httpserver .
```

#### 4. 启动

```bash
docker run -d -p 8080:80 cyvan115/httpserver
```

因为 macOS 的 Docker Desktop for Mac 使用不了默认的 bridge network，所以通过指定端口映射的方式进行 `httpserver` 访问

#### 5. 测试

```bash
curl -X 'GET' 'http://127.0.0.1:8080/time'
# 2022-04-25 16:24:26.72652675 +0000 UTC m=+965.423497243

curl -X 'GET' 'http://127.0.0.1:8080/healthz'
```

#### 6. 查看 IP

```bash
# 查看 container id
docker ps | grep httpserver
# => 1d9548df8afb

# 查看容器 pid
docker inspect 1d9548df8afb| grep Pid
# => "Pid": 2517

# 通过 nsenter 查看 ip
nsenter -t 2517 -n ip a
```

#### 7. 推送镜像

```bash
docker login

docker push cyvan115/httpserver
```

[my docker image](https://hub.docker.com/repository/docker/cyvan115/httpserver)