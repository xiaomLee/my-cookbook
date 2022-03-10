## nginx

1. 高可用部署
    
    keepalived + vip
    
    高可用(HA, High Availability)：提供健康检查功能，基于 VRRP(Virtual RouterRedundancy Protocol) 协议实现多台机器间的故障转移服务；
    负载均衡(LB, Load Balancing)：基于 Linux 虚拟服务器(IPVS)内核模块，提供 Layer4 负载均衡。
    
   **keepalived原理**
    
   Keepalived 分为3个守护进程：
    
   父进程: 很简单，负责 fork 子进程，并监视子进程健康(图中 WatchDog 周期性发送检测包，需要的话重启子进程)；
   
   子进程A: 负责VRRP框架
   
   子进程B: 负责健康检查
   
    
   **VIP原理**
    
   利用arp缓存，将一个并未对于实际的主机的IP，动态绑定到一个mac地址。
