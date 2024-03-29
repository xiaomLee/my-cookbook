# HTTP/HTTPS

## HTTP

超文本传输协议，基于请求与响应模式的、是一种无状态、无连接的一种应用层协议。
无连接指的是在 HTTP1.0 版本中，每次建立起的 TCP 连接只处理一个请求，服务端在收到客户端的应答之后就立即断开链接。
   
**HTTP方法**
GET    ------ 获取资源  
POST   ------ 传输资源  
PUT    ------ 更新资源  
DELETE ------ 删除资源  
HEAD   ------ 获取报文首部  
TRACE  ------ 请求服务器回送收到的请求信息，主要用于测试或诊断  
CONNECT------ 保留将来使用  
OPTIONS------ 请求查询服务器的性能，或者查询与资源相关的选项和需求  

**HTTP状态码**  
1xx: 指示信息   =》 表示请求已接收，继续处理  
2xx: 成功       =》 表示请求已成功接收  
3xx: 重定向     =》 要完成请求必须更进一步操作  
4xx: 客户端错误 =》 请求语法错误或者请求无法实现，比如路径错误，资源不存在等  
5xx: 服务器错误 =》 服务器未能实现合法的请求  
   
**常见的HTTP状态码**  
200 ok: 客户端请求成功。  
206 Partail Content： 客户发生了一个带有Range头的get请求，服务器完成了它（通常在请求大的视频或音频时可能出现）。  
301 Moved Permanently: 请求的页面已经转移至新的url地址。  
302 Found: 请求的页面已经临时转移至新的url地址。  
304 Not Modified: 客户端有缓存的文档并发送了一个有条件的请求，服务器告诉客户端原来的缓存文档该可以继续使用。  
400 Bad Request: 客户端语法错误。  
401 Unauthorized: 请求未经授权，这个状态码必须和WWW-Authenticate报头域一起使用。 
403 Forbidden: 页面禁止被访问   
404 Not Found: 请求资源不存在。  
500 Internal Sever Error: 服务器发生不可预期的错误，原来缓存的文档还可以继续被使用。  
503 Server Unavaliable: 请求未完成，服务器临时过载或当机，一段时间后可能恢复正常。  
  
**keepalive**  
TCP的KeepAlive和HTTP的Keep-Alive是完全不同的概念，不能混为一谈。实际上HTTP的KeepAlive写法是Keep-Alive，跟TCP的KeepAlive写法上也有不同。

tcp的keepalive是侧重在保持客户端和服务端的连接，一方会不定期发送心跳包给另一方，当一方断掉的时候，没有断掉的定时发送几次心跳包，
如果间隔发送几次，对方都返回的是RST，而不是ACK，那么就释放当前链接。
tcp连接默认时长，一般默认是2小时，可由系统配置更改。通过keepalive可实现真正的长连接。

HTTP的keep-alive一般我们都会带上中间的横杠，普通的http连接是客户端连接服务端，然后结束请求后，
由客户端或者服务端进行http连接的关闭。下次再发送请求的时候，客户端再发起一个连接，传送数据，关闭连接。
侧重于tcp的连接复用。

二者的作用简单来说：
HTTP协议的Keep-Alive意图在于短时间内连接复用，希望可以短时间内在同一个连接上进行多次请求/响应。

TCP的KeepAlive机制意图在于保活、心跳，检测连接错误。
当一个TCP连接两端长时间没有数据传输时(通常默认配置是2小时)，发送keepalive探针，探测链接是否存活。

[参考](https://juejin.im/post/6844903789703462925)  
[参考](https://www.jianshu.com/p/9fe2c140fa52)  

## HTTPS

### 实现原理

**交互流程**

整体过程分为整数验证、数据传输两个阶段。
- 客户端发起请求，包含自身浏览器支持的加密算法 HASH算法等信息
- 服务端接收请求，选择一组加密算法与 hash算法，返回证书
- 判断证书合法性
- 若证书合法，客户端生成随机数
- 通过证书内的公钥对随机数进行加密，通过随机数对握手hash信息进行加密
- 将公钥加密后的随机数与随机数加密后的握手信息通过网络传输
- 服务端接收数据，通过私钥解密，获得随机数，通过随机数解密握手信息
- 服务端加密握手信息返回给客户端
- 客户端解密判断hash算法是否一致，相同则建立握手

![交互流程图](./images/https%E4%BA%A4%E4%BA%92%E6%B5%81%E7%A8%8B%E5%9B%BE.jpeg)

**CA证书合法性**

证书内容：
- 颁发机构信息
- 公钥
- 公司信息
- 域名
- 有效期
- 指纹
- 其他

证书合法性的依据：基于信任机制，即中心的权威的机构，机构需对其颁发的证书进行信用背书。

验证流程：
- 验证域名、有效期等信息
- 验证来源是否合法，即证书链是否完整
- 请求CA服务器，验证是否被篡改

[终于有人把 HTTPS 原理讲清楚了](https://cloud.tencent.com/developer/article/1601995)

[HTTPS 详解一](https://segmentfault.com/a/1190000021494676)