## 通用业务网关项目

api网关，负责公司的服务路由、登录鉴权、请求转发等需求。  
服务端，处理接受外部请求。集群模式，可水平扩展。  
配置管理端，配合平台管理服务元信息、服务api、负载均衡器、限流熔断的配置。 

![系统架构图](./images/gw-系统架构设计.jpg)
![主要模块](./images/gw-主要模块.jpg)
   
### 配置管理
 
通过配置文件加载后端所有微服务。

```
// api配置
{
    "module":"user",
    "url":"/user/login",
    "unionApi":false,
    "login":false,
    "ratelimit":[
        {
            "func":"iplimit",
            "creditsPerSecond":3
            "maxBalance":5
        },
        {
            "func":"apilimit",
            "creditsPerSecond":100
            "maxBalance":50
        },
    ]
    "beforeAction":{},
    "backend":[
        {
            "service":"user",
            "api":"login",
            "protocol":"grpc"
        }
    ],
    "afterAction":{
        "service":"common",
        "api":"userinfo/mask",
        "protocol":"grpc"
    }
}

// 微服务配置
{
    "service":"user",
    "loadbalance":"roundrobin",
    "breaker":{
        "timeout":5,    // Timeout is the period of the open state,
        "interval":3,   
        "maxrequest":5,    
        "failureCount":3,    
        "failureRatio":0.6,    
        "stateChangeCb":"http://xxx",    
    },
    "endpoints":{
        "grpc":[
            "10.11.11.1:1235",
            "10.11.11.2:1235"
        ],
        "http":[
            "10.11.11.1:1234",
            "10.11.11.2:1234"
        ]
    }
}
```

### 微服务框架
    
- Go Micro
    
    服务发现 - 应用程序自动注册到服务发现系统。  
    负载平衡 - 客户端负载平衡，用于平衡服务实例之间的请求。  
    同步通信 - 提供请求 / 响应传输层。  
    异步通信 - 内置发布 / 订阅功能。  
    消息编码 - 基于消息的内容类型头的编码 / 解码。  
    RPC 客户机 / 服务器包 - 利用上述功能并公开接口来构建微服务  
    
- Go Kit

    认证 - Basic 认证和 JWT 认证  
    传输 - HTTP、Nats、gRPC 等等。  
    日志记录 - 用于结构化服务日志记录的通用接口。  
    指标 - CloudWatch、Statsd、Graphite 等。  
    追踪 - Zipkin 和 Opentracing。  
    服务发现 - Consul、Etcd、Eureka 等等。  
    断路器 - Hystrix 的 Go 实现。  


### JWT鉴权

**JWT(json web token)的构成**  
第一部分我们称它为头部（header),第二部分我们称其为载荷（payload, 类似于飞机上承载的物品)，第三部分是签证（signature).

*header* 声明类型，加密算法；json串，base64加密  
*payload* 公共声明，过期时间，uid uname之类；json串，base64加密  
*signature* base64加密后的header和base64加密后的payload使用.连接组成的字符串，通过header中声明的加密方式进行加盐secret组合加密.
secret是保存在服务器端的，jwt的签发生成也是在服务器端的，secret就是用来进行jwt的签发和jwt的验证，所以，它就是你服务端的私钥，在任何场景都不应该流露出去。  


### 服务注册、发现

借鉴go-micro.register自己实现，支持etcd/consul后端存储可选
service启动的时候进行注册，关闭时deregister，10秒expire，2/3*expireTime秒主动同步心跳
http网关启动时拉取并更新保存一份本地缓存，watch监听配置变更.
可配合nginx使用，详情参考nginx-upstream配置

### 服务限流  

支持不同级别的限流，默认限流策略：gatewayLimit --> ipLimit --> serviceLimit --> apiLimit --> api级别的其他自定义限流器

限流算法以及实现：edis+lua实现的分布式限流器，采用令牌通算法实现
    
*扩展*([参考](https://www.infoq.cn/article/Qg2tX8fyw5Vt-f3HH673))

    漏桶算法
    将每个请求视作"水滴"放入"漏桶"进行存储；
    “漏桶"以固定速率向外"漏"出请求来执行如果"漏桶"空了则停止"漏水”；
    如果"漏桶"满了则多余的"水滴"会被直接丢弃。
    漏桶算法多使用队列实现，服务的请求会存到队列中，服务的提供方则按照固定的速率从队列中取出请求并执行，过多的请求则放在队列中排队或直接拒绝。
    漏桶算法的缺陷也很明显，当短时间内有大量的突发请求时，即便此时服务器没有任何负载，每个请求也都得在队列中等待一段时间才能被响应。
    
    令牌桶算法
    令牌以固定速率生成；
    生成的令牌放入令牌桶中存放，如果令牌桶满了则多余的令牌会直接丢弃，当请求到达时，会尝试从令牌桶中取令牌，取到了令牌的请求可以执行；
    如果桶空了，那么尝试取令牌的请求会被直接丢弃。
    令牌桶算法既能够将所有的请求平均分布到时间区间内，又能接受服务器能够承受范围内的突发请求，因此是目前使用较为广泛的一种限流算法。


### 负载均衡

可结合服务发现使用，动态负载均衡。

- 随机
-  
- 轮训
- 
- 带权重轮训
    ```
    var total int64
    l := len(servers)
    if 0 >= l {
        return nil
    }
    
    best := ""
    for i := l - 1; i >= 0; i-- {
        svr := servers[i]
        weight, err := strconv.ParseInt(svr.Metadata["weight"], 10, 64)
        if err != nil {
            return nil
        }
        id := svr.Id
    
        if _, ok := w.opts[id]; !ok {
            w.opts[id] = &weightRobin{
                node: svr,
                effectiveWeight: weight,
            }
        }
    
        wt := w.opts[id]
        wt.currentWeight += wt.effectiveWeight
        total += wt.effectiveWeight
    
        if wt.effectiveWeight < weight {
            wt.effectiveWeight++
        }
    
        if best == "" || w.opts[best] == nil || wt.currentWeight > w.opts[best].currentWeight {
            best = id
        }
    }
    
    if best == "" {
        return nil
    }
    
    w.opts[best].currentWeight -= total
    ```

- hash 可指定hash字符串 $http_refer $client_ip $uri $uid $phone $email 
    

### 服务降级(熔断)

hystrix-go实现的熔断器

其他熔断器实现：gobreaker 
   

### 服务监控

使用prometheus实现，各api请求信息统计、延时统计、错误统计、错误msg分类。
业务服务也可主动上报自身的业务监控指标。

使用promPushGateway的方式进行上报，各服务只需往pushgateway推送指标即可，promServer从pushgateway采集信息，减少业务入侵与网络隔离。

监控平台使用grafana

优化方向，使用metric包装监控信息接口，解耦prometheus的具体实现，比如可使用xxx实现


### 链路追踪

通过Opentracing包实现，后端存储使用jaeger

其他：Twitter公司开源zipkin， Google dapper Uber的 jaeger