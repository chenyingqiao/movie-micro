# Movie-micro

![](https://img.shields.io/badge/License-MIT-orange)
![AppVeyor](https://img.shields.io/badge/Docker-32M-yellowgreen)
![](https://img.shields.io/badge/Istio-1.6.5-red)
![](https://img.shields.io/badge/K8s-v1.18.3-yellowgreen)
![](https://img.shields.io/badge/minikube-v1.12.2-orange)

> 一个简单的微服务实例项目

主要应用的组件有 `k8s` `istio` `mongodb` `elasticsearch`
微服务通讯主要使用 `grpc` 进行通讯

**Demo地址**

[http://movie.chenyingqiao.top/](http://movie.chenyingqiao.top/)

![20200906153312](http://img.chenyingqiao.top/blog/20200906153312.png)
![20200906212813](http://img.chenyingqiao.top/blog/20200906212813.png)

|  服务名称 |  作用 |
|---|---|
| movie  |  用户前端 |
| auth  |  用户权限鉴定，jwt校验 |
| server  |  影片信息相关接口 |
| talk  |  电影聊天室 |

![20200921085053](http://img.chenyingqiao.top/blog/20200921085053.png)

|  任务名称 |  作用 |
|---|---|
| crawler-minute  |  最新电影更新爬虫，频率高 |
| crawler  |  全量电影更新爬虫，频率低 |


![20200906154322](http://img.chenyingqiao.top/blog/20200906154322.png)

**k8s描述文件结构图**

![20200906153952](http://img.chenyingqiao.top/blog/20200906153952.png)

# 快速安装项目

> 仅仅用于ubuntu系统

快速安装

```shell
curl -o init-system.sh https://raw.githubusercontent.com/chenyingqiao/movie-micro/master/init-system.sh&& chmod +x init-system.sh && ./init-system.sh
```

将 movie.chenyingqiao.com以及talk.chenyingqiao.com 添加到到hosts中，对应的ip可以通过

```shell
sudo kubectl get ingress -n movie
```

找到