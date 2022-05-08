# httpServer
> 若使用alpine基础镜像，编译源代码需要先把CGO_ENABLE设置为0，即export CGO_ENABLE=0,然后编译
### 推送镜像
1. 登录账号:```docker login 用户名```
2. 改标签:```docker tag 本地镜像名  用户名/本地镜像名:版本号```
3. 推送镜像:```docker push 用户名/本地镜像名:版本号```
### 查看容器ip
1. 进入容器:```docker exec -it 容器名 bash```
2. 查看ip: ```ip a```
