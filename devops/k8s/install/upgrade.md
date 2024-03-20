# 升级kubernetes集群

- [升级kubernetes集群](#升级kubernetes集群)
  - [1. 升级控制平面（主节点）](#1-升级控制平面主节点)
    - [1.1 升级kubeadm](#11-升级kubeadm)
    - [1.2 升级控制平面组件 kube-apiserver kube-controller-manager kube-scheduler kube-proxy 等](#12-升级控制平面组件-kube-apiserver-kube-controller-manager-kube-scheduler-kube-proxy-等)
    - [1.3 升级控制平面的kubelet组件](#13-升级控制平面的kubelet组件)
    - [1.4 升级其他控制平面节点](#14-升级其他控制平面节点)
  - [2 升级worker节点](#2-升级worker节点)
    - [2.1 升级kubeadm](#21-升级kubeadm)
    - [2.2 升级kubelet](#22-升级kubelet)
    - [2.3 重复升级其他worker节点](#23-重复升级其他worker节点)
  - [3. 原理总结](#3-原理总结)

**执行的步骤**
1. 升级控制平面
2. 升级集群中的节点
3. 升级kubectl之类的客户端
4. 根据新Kubernetes版本带来的API变化，调整清单文件和其他资源，视具体应用而定

下述文档使用kubeadm进行集群的升级。

[参考](https://kubernetes.io/zh/docs/tasks/administer-cluster/kubeadm/kubeadm-upgrade/)

## 1. 升级控制平面（主节点）
任何k8s集群的升级都是先升级master节点，如果有多个master节点，需一个个全部升级完成后再执行worker节点的升级动作。

控制平面的升级首先需要升级kubeadm，之后再根据kubeadm upgrade plan执行kubelet等相关组件的升级。

### 1.1 升级kubeadm
```bash
# 查看最新支持的kubeadm版本
[root@k8s-master ~]# yum list --showduplicates kubeadm --disableexcludes=kubernetes |tail -5
kubeadm.x86_64                       1.22.1-0                        kubernetes 
kubeadm.x86_64                       1.22.2-0                        kubernetes 
kubeadm.x86_64                       1.22.3-0                        kubernetes 
kubeadm.x86_64                       1.22.4-0                        kubernetes 
kubeadm.x86_64                       1.23.0-0                        kubernetes 

# 将kubeadm升级到最新的1.23.0
[root@k8s-master ~]# yum install -y kubeadm-1.23.0-0 --disableexcludes=kubernetes

# 验证是否升级成功
[root@k8s-master ~]# kubeadm version
kubeadm version: &version.Info{Major:"1", Minor:"23", GitVersion:"v1.23.0", GitCommit:"ab69524f795c42094a6630298ff53f3c3ebab7f4", GitTreeState:"clean", BuildDate:"2021-12-07T18:15:11Z", GoVersion:"go1.17.3", Compiler:"gc", Platform:"linux/amd64"}
```

### 1.2 升级控制平面组件 kube-apiserver kube-controller-manager kube-scheduler kube-proxy 等
```bash
# 查看升级计划
[root@k8s-master ~]# kubeadm upgrade plan 
[upgrade/config] Making sure the configuration is correct:
[upgrade/config] Reading configuration from the cluster...
[upgrade/config] FYI: You can look at this config file with 'kubectl -n kube-system get cm kubeadm-config -o yaml'
[preflight] Running pre-flight checks.
[upgrade] Running cluster health checks
[upgrade] Fetching available versions to upgrade to
[upgrade/versions] Cluster version: v1.22.0
[upgrade/versions] kubeadm version: v1.23.0
[upgrade/versions] Target version: v1.23.0
[upgrade/versions] Latest version in the v1.22 series: v1.22.4

Components that must be upgraded manually after you have upgraded the control plane with 'kubeadm upgrade apply':
COMPONENT   CURRENT       TARGET
kubelet     3 x v1.22.0   v1.22.4

Upgrade to the latest version in the v1.22 series:

COMPONENT                 CURRENT   TARGET
kube-apiserver            v1.22.0   v1.22.4
kube-controller-manager   v1.22.0   v1.22.4
kube-scheduler            v1.22.0   v1.22.4
kube-proxy                v1.22.0   v1.22.4
CoreDNS                   v1.8.4    v1.8.6
etcd                      3.5.0-0   3.5.1-0

You can now apply the upgrade by executing the following command:

	kubeadm upgrade apply v1.22.4

_____________________________________________________________________

Components that must be upgraded manually after you have upgraded the control plane with 'kubeadm upgrade apply':
COMPONENT   CURRENT       TARGET
kubelet     3 x v1.22.0   v1.23.0

Upgrade to the latest stable version:

COMPONENT                 CURRENT   TARGET
kube-apiserver            v1.22.0   v1.23.0
kube-controller-manager   v1.22.0   v1.23.0
kube-scheduler            v1.22.0   v1.23.0
kube-proxy                v1.22.0   v1.23.0
CoreDNS                   v1.8.4    v1.8.6
etcd                      3.5.0-0   3.5.1-0

You can now apply the upgrade by executing the following command:

	kubeadm upgrade apply v1.23.0

_____________________________________________________________________


The table below shows the current state of component configs as understood by this version of kubeadm.
Configs that have a "yes" mark in the "MANUAL UPGRADE REQUIRED" column require manual config upgrade or
resetting to kubeadm defaults before a successful upgrade can be performed. The version to manually
upgrade to is denoted in the "PREFERRED VERSION" column.

API GROUP                 CURRENT VERSION   PREFERRED VERSION   MANUAL UPGRADE REQUIRED
kubeproxy.config.k8s.io   v1alpha1          v1alpha1            no
kubelet.config.k8s.io     v1beta1           v1beta1             no
_____________________________________________________________________

# 跟据上述输出提示可知
# 升级小版本 执行 kubeadm upgrade apply v1.22.4
# 升级大版本 执行 kubeadm upgrade apply v1.23.0

# 执行大版本升级 成功后可看到SUCCESS字样
[root@k8s-master ~]# kubeadm upgrade apply v1.23.0
...
[upgrade/successful] SUCCESS! Your cluster was upgraded to "v1.23.0". Enjoy!

[upgrade/kubelet] Now that your control plane is upgraded, please proceed with upgrading your kubelets if you havent already done so

# 再次执行kubeadm upgrade plan可看到已经是最新版本，无可用升级计划
[root@k8s-master ~]# kubeadm upgrade plan
[upgrade/config] Making sure the configuration is correct:
[upgrade/config] Reading configuration from the cluster...
[upgrade/config] FYI: You can look at this config file with 'kubectl -n kube-system get cm kubeadm-config -o yaml'
[preflight] Running pre-flight checks.
[upgrade] Running cluster health checks
[upgrade] Fetching available versions to upgrade to
[upgrade/versions] Cluster version: v1.23.0
[upgrade/versions] kubeadm version: v1.23.0
[upgrade/versions] Target version: v1.23.0
[upgrade/versions] Latest version in the v1.23 series: v1.23.0

```
### 1.3 升级控制平面的kubelet组件
```bash
# 1. 腾空节点
[root@k8s-master ~]# kubectl drain k8s-master --ignore-daemonsets --delete-emptydir-data
node/k8s-master already cordoned
WARNING: ignoring DaemonSet-managed Pods: calico-system/calico-node-fwzfq, kube-system/kube-proxy-v47qc
evicting pod kube-system/metrics-server-7795cfc6ff-tl2g9
evicting pod calico-apiserver/calico-apiserver-54b98ccf5b-slfsz
evicting pod kube-system/coredns-6d8c4cb4d-fdkqn
pod/calico-apiserver-54b98ccf5b-slfsz evicted
pod/metrics-server-7795cfc6ff-tl2g9 evicted
pod/coredns-6d8c4cb4d-fdkqn evicted
node/k8s-master evicted

# 2. 升级kubelet 和 kubectl
[root@k8s-master ~]# yum install -y kubelet-1.23.0-0 kubectl-1.23.0-0 --disableexcludes=kubernetes

# 3. 重启kubelet 验证版本
[root@k8s-master ~]# systemctl daemon-reload 
[root@k8s-master ~]# systemctl restart kubelet
[root@k8s-master ~]# kubectl version
Client Version: version.Info{Major:"1", Minor:"23", GitVersion:"v1.23.0", GitCommit:"ab69524f795c42094a6630298ff53f3c3ebab7f4", GitTreeState:"clean", BuildDate:"2021-12-07T18:16:20Z", GoVersion:"go1.17.3", Compiler:"gc", Platform:"linux/amd64"}
Server Version: version.Info{Major:"1", Minor:"23", GitVersion:"v1.23.0", GitCommit:"ab69524f795c42094a6630298ff53f3c3ebab7f4", GitTreeState:"clean", BuildDate:"2021-12-07T18:09:57Z", GoVersion:"go1.17.3", Compiler:"gc", Platform:"linux/amd64"}

# 4. 解除节点的保护
# 解除前status=SchedulingDisabled 解除后status=ready
[root@k8s-master ~]# kubectl get nodes
NAME           STATUS                     ROLES                  AGE   VERSION
k8s-master     Ready,SchedulingDisabled   control-plane,master   20h   v1.23.0
k8s-worker-1   Ready                      worker                 20h   v1.22.0
k8s-worker-2   Ready                      worker                 20h   v1.22.0     
[root@k8s-master ~]# kubectl uncordon k8s-master 
node/k8s-master uncordoned
[root@k8s-master ~]# kubectl get nodes
NAME           STATUS   ROLES                  AGE   VERSION
k8s-master     Ready    control-plane,master   20h   v1.23.0
k8s-worker-1   Ready    worker                 20h   v1.22.0
k8s-worker-2   Ready    worker                 20h   v1.22.0
```
### 1.4 升级其他控制平面节点
与第一个控制面节点相同，但是使用：

sudo kubeadm upgrade node
而不是：

sudo kubeadm upgrade apply
此外，不需要执行 kubeadm upgrade plan

## 2 升级worker节点
工作节点上的升级过程应该一次执行一个节点，或者一次执行几个节点， 以不影响运行工作负载所需的最小容量。

### 2.1 升级kubeadm
```bash
# 升级kubeadm
[root@k8s-worker-1 ~]# yum install -y kubeadm-1.23.0-0 --disableexcludes=kubernetes

# 对于工作节点，下面的命令会升级本地的 kubelet 配置
[root@k8s-worker-1 ~]# sudo kubeadm upgrade node
[upgrade] Reading configuration from the cluster...
[upgrade] FYI: You can look at this config file with 'kubectl -n kube-system get cm kubeadm-config -o yaml'
[preflight] Running pre-flight checks
[preflight] Skipping prepull. Not a control plane node.
[upgrade] Skipping phase. Not a control plane node.
[kubelet-start] Writing kubelet configuration to file "/var/lib/kubelet/config.yaml"
[upgrade] The configuration for this node was successfully updated!
[upgrade] Now you should go ahead and upgrade the kubelet package using your package manager.

```

### 2.2 升级kubelet
```bash
# 1. 腾空节点 
[root@k8s-master ~]# kubectl drain k8s-worker-1 --ignore-daemonsets  --delete-emptydir-data
# 2. 升级kubelet kubectl
[root@k8s-worker-1 ~]# yum install -y kubelet-1.23.0-0 kubectl-1.23.0-0
# 3. 重启
[root@k8s-worker-1 ~]# systemctl daemon-reload 
[root@k8s-worker-1 ~]# systemctl restart kubelet
[root@k8s-worker-1 ~]# kubectl get nodes
NAME           STATUS                     ROLES                  AGE   VERSION
k8s-master     Ready                      control-plane,master   21h   v1.23.0
k8s-worker-1   Ready,SchedulingDisabled   worker                 20h   v1.23.0
k8s-worker-2   Ready                      worker                 20h   v1.22.0
[root@k8s-worker-1 ~]# kubectl uncordon k8s-worker-1
node/k8s-worker-1 uncordoned
[root@k8s-worker-1 ~]# kubectl get nodes
NAME           STATUS   ROLES                  AGE   VERSION
k8s-master     Ready    control-plane,master   21h   v1.23.0
k8s-worker-1   Ready    worker                 20h   v1.23.0
k8s-worker-2   Ready    worker                 21h   v1.22.0
```
### 2.3 重复升级其他worker节点


## 3. 原理总结
工作原理 
kubeadm upgrade apply 做了以下工作：

检查你的集群是否处于可升级状态:

API 服务器是可访问的

所有节点处于 Ready 状态

控制面是健康的

强制执行版本偏差策略。

确保控制面的镜像是可用的或可拉取到服务器上。

如果组件配置要求版本升级，则生成替代配置与/或使用用户提供的覆盖版本配置。

升级控制面组件或回滚（如果其中任何一个组件无法启动）。

应用新的 CoreDNS 和 kube-proxy 清单，并强制创建所有必需的 RBAC 规则。

如果旧文件在 180 天后过期，将创建 API 服务器的新证书和密钥文件并备份旧文件。

kubeadm upgrade node 在其他控制平节点上执行以下操作：

从集群中获取 kubeadm ClusterConfiguration。

（可选操作）备份 kube-apiserver 证书。

升级控制平面组件的静态 Pod 清单。

为本节点升级 kubelet 配置

kubeadm upgrade node 在工作节点上完成以下工作：

从集群取回 kubeadm ClusterConfiguration。

为本节点升级 kubelet 配置。