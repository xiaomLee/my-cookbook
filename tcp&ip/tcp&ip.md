## tcp/ip协议簇

### 三次握手

第一次次握手：客户端将标志位SYN置为1，随机产生一个值seq=J，并将该数据包发送给服务器端，客户端进入SYN_SENT状态，等待服务器端确认。

第二次握手：服务器端收到数据包后由标志位SYN=1知道客户端请求建立连接，服务器端将标志位SYN和ACK都置为1，ack=J+1，随机产生一个值seq=K，
并将该数据包发送给客户端以确认连接请求，服务器端进入SYN_RCVD状态。

第三次握手：客户端收到确认后，检查ack是否为J+1，ACK是否为1，如果正确则将标志位ACK置为1，ack=K+1，
并将该数据包发送给服务器端，服务器端检查ack是否为K+1，ACK是否为1，如果正确则连接建立成功，客户端和服务器端进入ESTABLISHED状态，完成三次握手。随后客户端与服务器端之间可以开始传输数据了。
![三次握手](./images/三次握手.jpg "三次握手")
   
### 四次挥手

中断连接端可以是客户端，也可以是服务器端。

CLIENT SEND FIN -> CLIENT FIN_WAIT_1 -> SERVER RCV AND SNED ACK -> CLIENT FIN_WAIT_2
SERVER SEND FIN -> SERVER LAST_ACK -> CLIENT RCV AND SEND ACK -> TIME_WAIT

第一次挥手：客户端发送一个FIN=M，用来关闭客户端到服务器端的数据传送，客户端进入FIN_WAIT_1状态。
意思是说"我客户端没有数据要发给你了"，但是如果你服务器端还有数据没有发送完成，则不必急着关闭连接，可以继续发送数据。

第二次挥手：服务器端收到FIN后，先发送ack=M+1，告诉客户端，你的请求我收到了，但是我还没准备好，请继续你等我的消息。
这个时候客户端就进入FIN_WAIT_2 状态，继续等待服务器端的FIN报文。

第三次挥手：当服务器端确定数据已发送完成，则向客户端发送FIN=N报文，告诉客户端，好了，我这边数据发完了，准备好关闭连接了。
服务器端进入LAST_ACK状态。

第四次挥手：客户端收到FIN=N报文后，就知道可以关闭连接了，但是他还是不相信网络，怕服务器端不知道要关闭，
所以发送ack=N+1后进入TIME_WAIT状态，如果Server端没有收到ACK则可以重传。
服务器端收到ACK后，就知道可以断开连接了。客户端等待了2MSL后依然没有收到回复，则证明服务器端已正常关闭，那好，我客户端也可以关闭连接了。

最终完成了四次握手。
![四次挥手](./images/四次挥手.jpg "四次挥手")

**time_wait问题**  
time_wait的存在是为了让对方准确收到最后一次ack。在高并发环境下，可能对导致系统可用socket不足的情况，无法为新到的请求分配端口。

解决方案
1. 设置系统参数，加快time_wait状态连接的回收
    ```
    vim /etc/sysctl.conf
    
    #time wait 最高的队列数
    tcp_max_tw_buckets = 256000
    
    #FIN_WAIT_2到TIME_WAIT的超时时间
    net.ipv4.tcp_fin_timeout = 30
    
    #表示开启重用
    net.ipv4.tcp_tw_reuse = 1 允许将TIME-WAIT sockets重新用于新的TCP连接，默认为0，表示关闭；
    
    #表示开启TCP连接中TIME-WAIT sockets的快速回收，默认为0，表示关闭
    net.ipv4.tcp_tw_recycle = 1
    ```
2. 使用长连接，比如grpc

[参考](https://developer.51cto.com/art/201906/597961.htm)