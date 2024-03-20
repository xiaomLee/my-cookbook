# k8s-cookbook

k8s是一个容器编排引擎
   
## Kubernetes 架构
   
从宏观上来看 Kubernetes 的整体架构，主要分为3大块：控制平面、工作节点、存储。 
![k8s架构图](./k8s架构.jpg)

**控制平面**，即 Master 主节点，负责控制整个 Kubernetes 集群的状态变化、以及状态存储。核心组件如下：
- API Server：主要提供资源操作的统一入口，这样就屏蔽了与 Etcd 的直接交互，；核心功能是封装对资源操作的API，供客户端以及集群其他组件使用；
- Controller：资源控制中心，监听资源的变化，根据变化事件作出反馈，以达到目标状态；比如rc就是监听deployment的资源变化，然后创建pod。
- Scheduler：管理、调度、运行 pod；核心逻辑是监听 pod 资源变化，然后执行调度逻辑，以达到目标状态。

**工作平面**，即 Node 工作节点，为整个集群提供计算力，是容器真正运行的地方。核心组件如下：
- Kubelet：主要工作包括管理容器的生命周期、结合 cAdvisor 进行监控、健康检查以及定期上报节点状态。
- Kube-proxy：主要利用 service 提供集群内部的服务发现和负载均衡，同时监听 service/endpoints 变化并刷新负载均衡。

### 从创建 Deployment 开始
Deployment 是用于编排 Pod 的一种控制器资源，这里以 Deployment 为例，来看看架构中的各组件在创建 Deployment 资源的过程中都干了什么。
![deployment-init](./deployment-init.jpg)

步骤如下：
1. 首先是 kubectl 发起一个创建 deployment 的请求。
2. apiserver 接收到创建 deployment 请求，将相关资源写入 etcd；之后所有组件与 apiserver/etcd 的交互都是类似的。
3. deployment controller list/watch 资源变化并发起创建 replicaSet 请求。
4. replicaSet controller list/watch 资源变化并发起创建 pod 请求。
5. scheduler 检测到未绑定的 pod 资源，通过一系列匹配以及过滤选择合适的 node 进行绑定。
6. kubelet list/watch 自己 node 上的pod事件，运行 pod 及后续生命周期管理。
7. kube-proxy 负责初始化 service 相关的资源，包括服务发现、负载均衡等网络规则。

**参考**
- [Kubernetes 技术架构深度剖析](https://www.infvie.com/ops-notes/kubernetes-in-depth-analysis-of-technical-architecture.html)
- [Kubernetes 架构](https://kubernetes.io/zh-cn/docs/concepts/architecture/)
- [深入浅出Kubernetes架构](https://cloud.tencent.com/developer/article/1663968)

## 资源类型

## 核心组件详解

### api-server

[kube-apiserver的设计与实现](https://blog.tianfeiyu.com/source-code-reading-notes/kubernetes/kube_apiserver.html)

### controller

### scheduler

### kubelet

### kube-proxy

## k8s网络

### pod网络

#### 单机容器网络
[kubernetes网络之浅谈单机容器网络](https://cvvz.github.io/post/container-network/)

#### 跨主机容器网络
[跨主机网络](https://cvvz.github.io/post/k8s-network-cross-host/)

[循序渐进理解CNI机制与Flannel工作原理](https://blog.yingchi.io/posts/2020/8/k8s-flannel.html)

### svc网络

[service网络](https://kubernetes.io/zh-cn/docs/concepts/services-networking/service/#virtual-ips-and-service-proxies)

### ingress网络

## 二次开发

### 自定义resource

### 自定义controller

### 自定义scheduler


 
[参考](https://cloud.tencent.com/developer/article/1663968)
