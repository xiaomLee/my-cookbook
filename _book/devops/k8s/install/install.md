# Centos安装手册 

- [Centos安装手册](#centos安装手册)
  - [系统环境准备](#系统环境准备)
    - [1. 配置防火墙](#1-配置防火墙)
    - [2. 配置系统时区](#2-配置系统时区)
    - [3. 关闭swap缓存](#3-关闭swap缓存)
    - [4. 关闭Selinux](#4-关闭selinux)
    - [5. 安在ssh远程访问[可选]](#5-安在ssh远程访问可选)
    - [6. 安装docker运行时](#6-安装docker运行时)
      - [6.1 安装docker-ce](#61-安装docker-ce)
      - [6.2 配置docker网络模块开机自动加载](#62-配置docker网络模块开机自动加载)
      - [6.3 配置桥接流量对 iptables 可见](#63-配置桥接流量对-iptables-可见)
      - [6.4 配置 docker](#64-配置-docker)
      - [6.5 查看docker配置信息](#65-查看docker配置信息)
  - [部署kubernetes集群](#部署kubernetes集群)
    - [1. 安装kubelet kubeadm kubectl](#1-安装kubelet-kubeadm-kubectl)
    - [2. 启动kubelet服务](#2-启动kubelet服务)
    - [3. 初始化master节点服务](#3-初始化master节点服务)
      - [3.1 可选 kebernetes镜像预拉取](#31-可选-kebernetes镜像预拉取)
      - [3.2 初始化master](#32-初始化master)
      - [3.3 安装网络插件 calico](#33-安装网络插件-calico)
    - [4. 初始化node节点并加入集群](#4-初始化node节点并加入集群)

## 系统环境准备

服务器信息
| IP      | 主机名称 | cpu | 内存 | 说明 |
| :------ | :------ | :-- | :-- | :--- |
|172.16.41.134|k8s-master| 8 | 4g | 主节点|
|172.16.41.135|k8s-worker-1|8|4g|工作节点|
|172.16.41.136|k8s-worker-2|4|2g|工作节点|


软件版本
| 软件    | 版本   |
| :--    | :---- |
|CentOS|8|
|Kubernetes|v1.22.0|
|Docker|20.10.11|


环境正确性

| 说明 | 查看命令 | 修改命令 |
| :--- |:----   | :--    |
|集群各节点互通|ping | 192.168.33.10|
|MAC地址唯一|ip |link或ifconfig -a|
|集群内主机名唯一|hostnamectl status|hostnamectl set-hostname|
|系统产品uuid唯一|dmidecode -s | system-uuid|


端口开放

kube-master 节点
|协议|方向|端口|目标|
| :-- | :-- | :-- | :--|
|TCP|inbound|6443|kube-api-server|
|TCP|inbound|2379-2380|etcd API
|TCP|inbound|10250|Kubelet API
|TCP|inbound|10251|kube-scheduler|
|TCP|inbound|10252|kube-controller-manager|

kube-worker 节点
|协议|方向|端口|目标|
| :-- | :-- | :-- | :-- |
|TCP|inbound|10250|Kubelet API|
|TCP|Inbound|30000-32767|NodePort Services|


所有机器均需执行以下操作.
```bash
# 可选 替换yum源为阿里镜像源
yum install -y wget
# 备份
mv /etc/yum.repos.d/CentOS-Base.repo /etc/yum.repos.d/CentOS-Base.repo.backup
# 下载新的CentOS-Base.repo 到/etc/yum.repos.d/
wget -O /etc/yum.repos.d/CentOS-Base.repo http://mirrors.aliyun.com/repo/Centos-7.repo
# 之后运行yum makecache生成缓存
yum clean all 
yum makecache 
```

### 1. 配置防火墙
```bash
# 方案1
# 配置默认zone trusted
[root@k8s-master ~]# firewall-cmd --set-default-zone=trusted
success

# 方案2
# 此步骤也可改成关闭防火墙，暴力手法
[root@k8s-master ~]# systemctl disable firewalld.service 
Removed /etc/systemd/system/multi-user.target.wants/firewalld.service.
Removed /etc/systemd/system/dbus-org.fedoraproject.FirewallD1.service.
# 可在k8s安装完成后重新开启. systemctl enable表示加入开机自启
# --now参数表示立即启动, 亦可通过 systemctl start firewall 进行启动
[root@k8s-master ~]# systemctl enable --now firewalld.service 
Created symlink /etc/systemd/system/dbus-org.fedoraproject.FirewallD1.service → /usr/lib/systemd/system/firewalld.service.
Created symlink /etc/systemd/system/multi-user.target.wants/firewalld.service → /usr/lib/systemd/system/firewalld.service.

# 方案3
# 防火墙状态
firewall-cmd --state
# 防火墙开放的端口
firewall-cmd --zone=public --list-ports
# 开放端口
firewall-cmd --zone=public --add-port=6443/tcp --permanent
firewall-cmd --zone=public --add-port=10250/tcp --permanent
firewall-cmd --zone=public --add-port=10251/tcp --permanent
firewall-cmd --zone=public --add-port=10252/tcp --permanent
# 批量开放端口
firewall-cmd --permanent --zone=public --add-port=2379-2380/tcp
# 重新加载
firewall-cmd --reload

```
### 2. 配置系统时区
所有机器时钟信息必须一致，后续安装k8s相关软件时将会生成证书信息，若不一致将导致证书认证出错，出现类似509.x的错误
```bash
# 所有节点服务器执行
# 设置时区
[root@k8s-master ~]# timedatectl set-timezone Asia/Shanghai
[root@k8s-master ~]# systemctl enable --now chronyd
# 验证设置是否成功
[root@k8s-master ~]# date
2021年 12月 13日 星期一 20:51:51 CST
# 查看同步状态
[root@k8s-master ~]# timedatectl status
               Local time: 一 2021-12-13 20:52:09 CST
           Universal time: 一 2021-12-13 12:52:09 UTC
                 RTC time: 二 2021-12-14 03:25:06
                Time zone: Asia/Shanghai (CST, +0800)
System clock synchronized: no
              NTP service: n/a
          RTC in local TZ: no
# 输出结果中显示下列属性证明时钟同步正常, 此处可不关注NTP时钟同步
System clock synchronized: yes
              NTP service: active
# 将当前的UTC时间写入硬件时钟
timedatectl set-local-rtc 0
# 重启依赖于系统时间的服务
systemctl restart rsyslog && systemctl restart crond
```

### 3. 关闭swap缓存
kubelet暂缺少对swap的支持，同时为了提升pod运行性能，所以禁用swap缓存
```bash
# 查看已使用的交换设备， swapon --help可查看命令使用帮助
[root@k8s-master ~]# swapon -s
# free -h 命令可看到swap并不为0
[root@k8s-master ~]# free -h
              total        used        free      shared  buff/cache   available
Mem:          3.7Gi       1.6Gi       142Mi        33Mi       1.9Gi       1.7Gi
Swap:            20M          20M          0B
# 禁用 /proc/swaps 中的所有交换区
[root@k8s-master ~]# swapoff -a
# 注释 /etc/fstab 中的交换设备，防止重启后自动开启swap.
[root@k8s-master ~]# sed -i 's/.*swap.*/#&/' /etc/fstab
# 上面命令亦可通过 vim /etc/fstab 手动进行注释
#
# /etc/fstab
# Created by anaconda on Sun Dec 12 19:23:51 2021
#
# Accessible filesystems, by reference, are maintained under '/dev/disk/'.
# See man pages fstab(5), findfs(8), mount(8) and/or blkid(8) for more info.
#
# After editing this file, run 'systemctl daemon-reload' to update systemd
# units generated from this file.
#
UUID=c64fefa3-1470-4ada-afbe-6a983ce86292 /                       xfs     defaults        0 0
UUID=174af5bb-ee77-4fb8-904f-1f2a09f568f6 /boot                   ext4    defaults        1 2
#UUID=32b1506e-4b00-453b-80c2-c12a7d62f4ed swap                    swap    defaults        0 0

# 确认swap是否被禁用 free -h 可看到swap所有项都为0
[root@k8s-master ~]# free -h
              total        used        free      shared  buff/cache   available
Mem:          3.7Gi       1.6Gi       142Mi        33Mi       1.9Gi       1.7Gi
Swap:            0B          0B          0B
```

### 4. 关闭Selinux
[Selinux介绍](https://www.redhat.com/zh/topics/linux/what-is-selinux)
安全增强型 Linux（SELinux）是一种采用安全架构的 Linux® 系统，它能够让管理员更好地管控哪些人可以访问系统。它最初是作为 Linux 内核的一系列补丁，由美国国家安全局（NSA）利用 Linux 安全模块（LSM）开发而成。

特点：配置复杂，容易出错，kubelet暂不支持
```bash
# 查看当前系统selinux等级
[root@k8s-master ~]# getenforce 
Enforcing
# 关闭 usage:  setenforce [ Enforcing | Permissive | 1 | 0 ]
[root@k8s-master ~]# setenforce 0
# 更改配置防止系统重启后再次被启动 vim /etc/selinux/config SELINUX 改为 permissive
[root@k8s-master ~]# sed -i 's/^SELINUX=.*/SELINUX=permissive/' /etc/selinux/config
```

### 5. 安在ssh远程访问[可选]
安装 openssh openssh-server开启远程ssh连接
```bash
yum install -y openssh openssh-server
```

### 6. 安装docker运行时
通过yum 安装docker-ce，之后配置docker相关网络模块/iptables转发以及cgroup驱动等.
#### 6.1 安装docker-ce
```bash
# 所有节点服务器执行
# 安装必要依赖 yum-utils包里面包含有yum-config-manager工具，可用于管理yum源
yum install -y yum-utils device-mapper-persistent-data lvm2
# 添加aliyum docker-ce yum源
yum-config-manager --add-repo http://mirrors.aliyun.com/docker-ce/linux/centos/docker-ce.repo
# 重建yum缓存
yum makecache
# 所有节点服务器执行
yum install -y docker-ce
```
#### 6.2 配置docker网络模块开机自动加载
```bash
# 所有节点服务器执行
[root@k8s-master ~]# lsmod | grep overlay
[root@k8s-master ~]# lsmod | grep br_netfilter

# 若上面的命令无返回值输出或提示文件不存在，需要执行以下命令：
[root@k8s-master ~]#  cat > /etc/modules-load.d/docker.conf <<EOF
overlay
br_netfilter
EOF

# 显示加载， 之后再次执行lsmod |grep br_netfilter即可看到输出
[root@k8s-master ~]# modprobe overlay
[root@k8s-master ~]# modprobe br_netfilter

[root@k8s-master ~]# lsmod |grep br_netfilter
br_netfilter           24576  0
bridge                192512  1 br_netfilter
```

#### 6.3 配置桥接流量对 iptables 可见
```bash
# 所有节点服务器执行
cat > /etc/sysctl.d/k8s.conf <<EOF
net.bridge.bridge-nf-call-ip6tables = 1
net.bridge.bridge-nf-call-iptables = 1
EOF

# 重启system服务
sysctl --system

# 验证是否生效，下面两个命令结果需均返回 1
sysctl -n net.bridge.bridge-nf-call-iptables
sysctl -n net.bridge.bridge-nf-call-ip6tables

# this will turn things back on a live server
sysctl -w net.ipv4.ip_forward=1
# on Centos this will make the setting apply after reboot
echo net.ipv4.ip_forward=1 >> /etc/sysconf.d/10-ipv4-forwarding-on.conf

# 验证并生效
sysctl -p

```
#### 6.4 配置 docker
```bash
# 所有节点服务器执行
mkdir /etc/docker
# 修改cgroup驱动为systemd[k8s官方推荐]、限制容器日志量、修改存储类型，最后的docker根目录可修改
cat > /etc/docker/daemon.json <<EOF
{
  "exec-opts": ["native.cgroupdriver=systemd"],
  "log-driver": "json-file",
  "log-opts": {
    "max-size": "100m"
  },
  "registry-mirrors": [
      "https://hub-mirror.c.163.com",
      "https://mirror.baidubce.com",
      "https://registry.cn-hangzhou.aliyuncs.com"
    ]
}
EOF
# 添加开机自启动，立即启动
systemctl enable --now docker

```

#### 6.5 查看docker配置信息
```bash
docker info
```

## 部署kubernetes集群

### 1. 安装kubelet kubeadm kubectl
kubectl可只在master节点安装，用于连接集群执行命令.
```bash
# 添加kubernetes源
# 所有节点服务器均执行
cat > /etc/yum.repos.d/kubernetes.repo <<EOF
[kubernetes]
name=Kubernetes
baseurl=https://mirrors.aliyun.com/kubernetes/yum/repos/kubernetes-el7-x86_64/
enabled=1
gpgcheck=1
repo_gpgcheck=1
gpgkey=https://mirrors.aliyun.com/kubernetes/yum/doc/yum-key.gpg https://mirrors.aliyun.com/kubernetes/yum/doc/rpm-package-key.gpg
EOF
# 重建yum缓存，输入y添加证书认证
yum makecache

# 查看kubelet可选安装版本， tail -n可显示最近n个版本
[root@k8s-master ~]# yum list --showduplicates kubelet |tail -10
kubelet.x86_64                       1.21.4-0                        kubernetes 
kubelet.x86_64                       1.21.5-0                        kubernetes 
kubelet.x86_64                       1.21.6-0                        kubernetes 
kubelet.x86_64                       1.21.7-0                        kubernetes 
kubelet.x86_64                       1.22.0-0                        kubernetes 
kubelet.x86_64                       1.22.1-0                        kubernetes 
kubelet.x86_64                       1.22.2-0                        kubernetes 
kubelet.x86_64                       1.22.3-0                        kubernetes 
kubelet.x86_64                       1.22.4-0                        kubernetes 
kubelet.x86_64                       1.23.0-0                        kubernetes 

# 安装指定版本kubelet kubeadm kubectl, 都使用相同的版本号
yum install -y kubectl-1.22.0 kubeadm-1.22.0 kubectl-1.22.0 --disableexcludes=kubernetes
```

### 2. 启动kubelet服务
```bash
# 配置开机启动并立即启动kubelet
systemctl enable --now kubelet
# 查看服务状态 此时kubelet可能会启动不成功，每个1s会重启, 因为在等待api-server的服务被kubeadm拉起
[root@k8s-master ~]# systemctl status kubelet
● kubelet.service - kubelet: The Kubernetes Node Agent
   Loaded: loaded (/usr/lib/systemd/system/kubelet.service; enabled; vendor preset: disabled)
  Drop-In: /usr/lib/systemd/system/kubelet.service.d
           └─10-kubeadm.conf
   Active: active (running) since Mon 2021-12-13 14:10:37 CST; 9h ago
     Docs: https://kubernetes.io/docs/
 Main PID: 52232 (kubelet)
    Tasks: 22 (limit: 23804)
   Memory: 135.5M
   CGroup: /system.slice/kubelet.service

# 配置自动补全命令
# 所有节点服务器均执行
# 安装bash自动补全插件
yum install bash-completion -y
# 设置kubectl与kubeadm命令补全，下次login生效
kubectl completion bash >/etc/bash_completion.d/kubectl
kubeadm completion bash > /etc/bash_completion.d/kubeadm
source /usr/share/bash-completion/bash_completion

```
### 3. 初始化master节点服务

#### 3.1 可选 kebernetes镜像预拉取
```bash
# 由于国内网络因素，kubernetes镜像需要从mirrors站点或通过dockerhub用户推送的镜像拉取。
# 现在kubeadm提供了--image-repository参数，可通过指定该参数从阿里云拉取.
# 对于可连接外网服务器可跳过下述步骤。
# 所有节点服务器均执行
# 查看执行kubernetes版本需要哪些镜像
kubeadm config images list --kubernetes-version v1.22.0
# 结果如下
k8s.gcr.io/kube-apiserver:v1.22.0
k8s.gcr.io/kube-controller-manager:v1.22.0
k8s.gcr.io/kube-scheduler:v1.22.0
k8s.gcr.io/kube-proxy:v1.22.0
k8s.gcr.io/pause:3.5
k8s.gcr.io/etcd:3.5.0-0
k8s.gcr.io/coredns/coredns:v1.8.4
# 在/root/k8s目录下新建脚本get-k8s-images.sh，命令如下：

# 所有节点服务器均执行
cd /root/
mkdir k8s
cd k8s/
# 创建脚本文件，文件内容如下一代码段所示
vim get-k8s-images.sh

#!/bin/bash

KUBE_VERSION=v1.22.0
PAUSE_VERSION=3.5
CORE_DNS_VERSION=v1.8.4
ETCD_VERSION=3.5.0-0

# pull kubernetes images from hub.docker.com
docker pull kubeimage/kube-proxy-amd64:$KUBE_VERSION
docker pull kubeimage/kube-controller-manager-amd64:$KUBE_VERSION
docker pull kubeimage/kube-apiserver-amd64:$KUBE_VERSION
docker pull kubeimage/kube-scheduler-amd64:$KUBE_VERSION
# pull aliyuncs mirror docker images
docker pull registry.cn-hangzhou.aliyuncs.com/google_containers/pause:$PAUSE_VERSION
docker pull registry.cn-hangzhou.aliyuncs.com/google_containers/coredns:$CORE_DNS_VERSION
docker pull registry.cn-hangzhou.aliyuncs.com/google_containers/etcd:$ETCD_VERSION

# retag to k8s.gcr.io prefix
docker tag kubeimage/kube-proxy-amd64:$KUBE_VERSION  k8s.gcr.io/kube-proxy:$KUBE_VERSION
docker tag kubeimage/kube-controller-manager-amd64:$KUBE_VERSION k8s.gcr.io/kube-controller-manager:$KUBE_VERSION
docker tag kubeimage/kube-apiserver-amd64:$KUBE_VERSION k8s.gcr.io/kube-apiserver:$KUBE_VERSION
docker tag kubeimage/kube-scheduler-amd64:$KUBE_VERSION k8s.gcr.io/kube-scheduler:$KUBE_VERSION
docker tag registry.cn-hangzhou.aliyuncs.com/google_containers/pause:$PAUSE_VERSION k8s.gcr.io/pause:$PAUSE_VERSION
docker tag registry.cn-hangzhou.aliyuncs.com/google_containers/coredns:$CORE_DNS_VERSION k8s.gcr.io/coredns:$CORE_DNS_VERSION
docker tag registry.cn-hangzhou.aliyuncs.com/google_containers/etcd:$ETCD_VERSION k8s.gcr.io/etcd:$ETCD_VERSION

# untag origin tag, the images won't be delete.
docker rmi kubeimage/kube-proxy-amd64:$KUBE_VERSION
docker rmi kubeimage/kube-controller-manager-amd64:$KUBE_VERSION
docker rmi kubeimage/kube-apiserver-amd64:$KUBE_VERSION
docker rmi kubeimage/kube-scheduler-amd64:$KUBE_VERSION
docker rmi registry.cn-hangzhou.aliyuncs.com/google_containers/pause:$PAUSE_VERSION
docker rmi registry.cn-hangzhou.aliyuncs.com/google_containers/coredns:$CORE_DNS_VERSION
docker rmi registry.cn-hangzhou.aliyuncs.com/google_containers/etcd:$ETCD_VERSION

# 添加脚本执行权限
chmod +x get-k8s-images.sh
# 执行脚本 
./get-k8s-images.sh

# 其他节点服务器均执行上述脚本，或亦可手动打包所有镜像，上传至其他机器进行load
docker save \
  k8s.gcr.io/kube-apiserver:v1.22.0 \
  k8s.gcr.io/kube-controller-manager:v1.22.0 \
  k8s.gcr.io/kube-scheduler:v1.22.0 \
  k8s.gcr.io/kube-proxy:v1.22.0 \
  k8s.gcr.io/pause:3.5 \
  k8s.gcr.io/etcd:3.5.0-0 \
  k8s.gcr.io/coredns/coredns:v1.8.4 \
  -o k8s-images-all.tar.gz

# 导入镜像
docker load -i k8s-images-all.tar.gz

```

#### 3.2 初始化master
```bash
# master节点服务器执行
# WARNING是正常的。
kubeadm init phase preflight
# 原始命令：kubeadm init phase preflight [--config kubeadm-init.yaml]
# 命令执行结束如果出现warning是正常的，一般会出现防火墙、无法连接k8s站点的警告。
# 如果出现无法从k8s拉去镜像的错误属于正常的，在执行初始化时优先使用我们本地Docker中的镜像，如果本地镜像不存在才会从k8s站点拉取。

# master节点服务器执行
# --kubernetes-version=v1.22.0 指定安装master组件的版本
# --image-repository 指定k8s各组件的镜像仓库, 若执行了上述预拉取镜像的操作，请不要指定该参数，用默认的k8s.gcr.io即可
# --pod-network-cidr 是指定pod网络ip地址的分配域，后续配置calico等网络插件时请确保使用相同的IP段.
# --service-cidr 指定service网络地址的IP段，默认10.96.0.0/12 可不指定
# 详细参数说明kubeadm init --help进行查看
# 原始命令：kubeadm init --pod-network-cidr=10.244.0.0/16 --kubernetes-version=v1.22.0
kubeadm init --image-repository registry.aliyuncs.com/google_containers --kubernetes-version=v1.22.0 --pod-network-cidr=10.244.0.0/16 --service-cidr=10.96.0.0/12

# 常见初始化错误，及解决方案
# 若出现http://localhost:6443连接不上的问题， 且提示docker cgroupfs相关的错误，说明docker runtime的cgroup driver用的不是kubernetes默认的systemd，请参考上述docker安装部分内容进行cgroup驱动修改。需在/etc/docker/daemon.json添加 "exec-opts": ["native.cgroupdriver=systemd"]
# 执行过程出错需重新运行上述命令进行初始化，重新运行前需执行kubeadm reset命令进行环境清除重置，否则会报错。

# 执行成功后请及时记录init命令输出的 kubeadm join 相关的命令, 形式如下
kubeadm join ip:6443 --token mfc8vh.z6jh8jonw3sls557 \
  --discovery-token-ca-cert-hash sha256:34fdc17960414969a0d3fee3079e570bb60ff6ed03b83ba694caaf23fa3f934a

# 若上述命令遗失或过期(token默认时效24h), 通过下述命令重新生成.
[root@k8s-master ~]# kubeadm token create --print-join-command 
kubeadm join ip:6443 --token e679dm.mxfi3eojvqv3tkvz --discovery-token-ca-cert-hash sha256:34fdc17960414969a0d3fee3079e570bb60ff6ed03b83ba694caaf23fa3f934a

# tips: 查看历史token
[root@k8s-master ~]# kubeadm token list
TOKEN                     TTL         EXPIRES                USAGES                   DESCRIPTION                                                EXTRA GROUPS
e679dm.mxfi3eojvqv3tkvz   23h         2021-12-14T16:53:04Z   authentication,signing   <none>                                                     system:bootstrappers:kubeadm:default-node-token
mfc8vh.z6jh8jonw3sls557   13h         2021-12-14T06:10:37Z   authentication,signing   The default bootstrap token generated by 'kubeadm init'.   system:bootstrappers:kubeadm:default-node-token

# tips: master节点获取ca证书sha256编码hash值
[root@k8s-master ~]# openssl x509 -pubkey -in /etc/kubernetes/pki/ca.crt | openssl rsa -pubin -outform der 2>/dev/null | openssl dgst -sha256 -hex | sed 's/^.* //'
34fdc17960414969a0d3fee3079e570bb60ff6ed03b83ba694caaf23fa3f934a

# 根据init提示 保存kubectl的连接配置
mkdir -p $HOME/.kube
sudo cp -i /etc/kubernetes/admin.conf $HOME/.kube/config
sudo chown $(id -u):$(id -g) $HOME/.kube/config

# 测试kubectl命令行工具
[root@k8s-master ~]# kubectl cluster-info 
Kubernetes control plane is running at https://172.16.41.134:6443
CoreDNS is running at https://172.16.41.134:6443/api/v1/namespaces/kube-system/services/kube-dns:dns/proxy

To further debug and diagnose cluster problems, use 'kubectl cluster-info dump'.

# 查看节点状态 此时为NotReady状态，因为还未安装必须的网络插件
[root@k8s-master ~]# kubectl get nodes 
NAME           STATUS   ROLES                  AGE   VERSION
k8s-master     NotReady    control-plane,master   10h   v1.22.0

```

#### 3.3 安装网络插件 calico
[calico官网](https://projectcalico.docs.tigera.io/getting-started/kubernetes/quickstart)
参考官网quickstart进行安装
```bash
# 获取calico需要的deploy.yaml
wget https://docs.projectcalico.org/manifests/tigera-operator.yaml
wget https://docs.projectcalico.org/manifests/custom-resources.yaml

# 编辑custom-resources.yaml 修改cidr, 默认配置为192.168.0.0/16，需修改为kubeadm init 中指定的--pod-network-cidr 10.244.0.0/16

# 部署calico服务 等待服务部署成功
kubectl apply -f tigera-operator.yaml custom-resources.yaml

# 查看节点状态 此时应为Ready状态
[root@k8s-master ~]# kubectl get nodes 
NAME           STATUS   ROLES                  AGE   VERSION
k8s-master     Ready    control-plane,master   10h   v1.22.0

```
### 4. 初始化node节点并加入集群
```bash
# 参考上述内容进行环境初始化
# 系统初始化配置， docker安装及配置， kubelet kubeadm安装， 导入初始镜像[可选]
# 在node节点执行如下命令，等待服务初始化完成
kubeadm join ip:6443 --token mfc8vh.z6jh8jonw3sls557 \
  --discovery-token-ca-cert-hash sha256:34fdc17960414969a0d3fee3079e570bb60ff6ed03b83ba694caaf23fa3f934a

# 每个节点重复上述操作

# master上查看集群节点状态, 此时worker节点的ROLES显示为none，不影响pod调度与集群使用
[root@k8s-master ~]# kubectl get nodes 
NAME           STATUS   ROLES                  AGE   VERSION
k8s-master     Ready    control-plane,master   11h   v1.22.0
k8s-worker-1   Ready    <none>                 10h   v1.22.0
k8s-worker-2   Ready    <none>                 10h   v1.22.0

# 修改node节点role属性
kubectl label node k8s-worker1 node-role.kubernetes.io/worker=worker

# 删除node节点role属性
kubectl label nodes  k8s-node1 node-role.kubernetes.io/worker-

```

**至此, Kubernetes集群安装完成.**