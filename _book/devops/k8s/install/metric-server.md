## Metrics API
通过 Metrics API，你可以获得指定节点或 Pod 当前使用的资源量。例如容器 CPU 和内存使用率，可通过 Metrics API 在 Kubernetes 中获得。 这些指标可以直接被用户访问，比如使用 kubectl top 命令行，或者被集群中的控制器 （例如 Horizontal Pod Autoscalers) 使用来做决策。
此 API 不存储指标值，因此想要获取某个指定节点 10 分钟前的 资源使用量是不可能的。

## Metrics服务器
[metrics-server](https://github.com/kubernetes-sigs/metrics-server)

通过部署组件components.yaml来部署。
```bash
# 下载components.yaml， 在git仓库的release页面可找到最新版本
wget https://github.com/kubernetes-sigs/metrics-server/releases/download/v0.5.2/components.yaml

# 添加--kubelet-insecure-tls参数。 对于自签的kubernetes证书，需添加此参数跳过证书认证检测
# 此外因网络原因替换镜像 k8s.gcr.io/metrics-server/metrics-server:v0.5.2
# k8s.gcr.io域下的镜像都可在registry.cn-hangzhou.aliyuncs.com/google_containers找到对应的
vim components.yaml
[root@k8s-master metric-server]# egrep -C 5 'kubelet-insecure-tls'  components.yaml 
    spec:
      containers:
      - args:
        - --cert-dir=/tmp
        - --secure-port=4443
        - --kubelet-insecure-tls
        - --kubelet-preferred-address-types=InternalIP,ExternalIP,Hostname
        - --kubelet-use-node-status-port
        - --metric-resolution=15s
        image: registry.cn-hangzhou.aliyuncs.com/google_containers/metrics-server:v0.5.2
        imagePullPolicy: IfNotPresent

# 如不执行上述镜像替换的动作，请在每个node节点执行手动下载镜像的动作
docker pull registry.cn-hangzhou.aliyuncs.com/google_containers/metrics-server:v0.5.2
docker tag registry.cn-hangzhou.aliyuncs.com/google_containers/metrics-server:v0.5.2 k8s.gcr.io/metrics-server/metrics-server:v0.5.2

```