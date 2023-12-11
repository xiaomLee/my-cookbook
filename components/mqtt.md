# MQTT

MQTT 是一种基于发布/订阅模式的轻量级消息传输协议，专门针对低带宽和不稳定网络环境的物联网应用而设计，可以用极少的代码为联网设备提供实时可靠的消息服务。MQTT 协议广泛应用于物联网、移动互联网、智能硬件、车联网、智慧城市、远程医疗、电力、石油与能源等领域。

MQTT 协议由 Andy Stanford-Clark （IBM）和 Arlen Nipper（Arcom，现为 Cirrus Link）于 1999 年发布。

全称为: MQ Telemetry Transport，是九十年代早期Arlen Nipper在参与 Conoco Phillips 公司的一个原油管道数据采集监控系统（pipeline SCADA system）时开发的一个实时数据传输协议。它的目的在于让传感器通过带宽有限的 VSAT ，与 IBM 的 MQ Integrator 通信。由于 Nipper 是遥感和数据采集监控专业出身，所以按业内惯例取了 MQ TT 这个名字。

现被广泛应用于物联网协议。物联网设备通信需解决几个关键问题：网络环境复杂而不可靠、内存和闪存容量小、处理器能力有限。
而MQTT协议正好契合了以上需求，经过多年发展已经基本具备了以下特点：
- 轻量高效，节省带宽
- 可靠的消息传递
- 海量连接支持
- 安全的双向通信
- 在线状态感知

[MQTT 5.0](https://www.emqx.com/zh/mqtt/mqtt5)

## MQTT协议详解

[MQTT协议](http://docs.oasis-open.org/mqtt/mqtt/)

### 轻量高效、低带宽

MQTT 将协议本身占用的额外消耗最小化，消息头部最小2字节，且荷载是二进制透明的，大大减小消息体大小，可稳定运行在带宽受限的网络环境下。同时，MQTT 客户端只需占用非常小的硬件资源，能运行在各种资源受限的边缘端设备上。

### 可靠的消息传递(QoS)

MQTT 协议提供了 3 种消息服务质量等级（Quality of Service），保证了在不同的网络环境下消息传递的可靠性。

QoS 0：消息最多传递一次。

如果当时客户端不可用，则会丢失该消息。发布者发送一条消息之后，就不再关心它有没有发送到对方，也不设置任何重发机制。

QoS 1：消息传递至少 1 次。

包含了简单的重发机制，发布者发送消息之后存储该PUBLISH报文(packetID标识)， 并等待接收者的 PUBACK；
若收到回复的PUBAC, 则将PACKET ID 置为可重用；
如果没收到 ACK 则重新发送消息， PACKET ID 不变， DUP=1；
因为存在PACKET ID 重用，这种模式能保证消息至少能到达一次，但无法保证消息重复。

QoS 2：消息仅传送一次。

设计了重发和重复消息发现机制，保证消息到达对方并且严格只到达一次。类似两阶段提交，需要与对端有两次交互。
1. 客户端发送 PUBLISH Qos=2 packectID=1024 msg=bytes 的报文，并缓存
2. 服务端接收并回复 PUBREC（pub recevice） packetID=1024，记录packetID
3. 客户端接收 PUBREC 删除本地 PUBLISH 缓存，并新增 PUBREL 缓存 (pub release)
4. 客户端发送 PUBREL packetID=1024
5. 服务接收 PUBREL，同时释放packetID, 并回复 PUBCOMP (pub complete) 
6. 客户端接收 PUBCOMP 删除本地 PUBREL 缓存， 释放重用packetID

- QoS 2 规定，发送方只有在收到 PUBREC 报文之前可以重传 PUBLISH 报文。
- 在收到对端回复的 PUBCOMP 报文确认双方都完成 Packet ID 释放之前，也不可以使用当前 Packet ID 发送新的消息。

[MQTT QoS（服务质量）介绍](https://www.emqx.com/zh/blog/introduction-to-mqtt-qos)

## 海量连接支持

大部分的 MQTT 服务器都可轻松具备高并发、高吞吐、高可扩展能力。目前开源 Server 实现中，EMQX是支持海量连接的代表，天然契合云原生领域，最具扩展性。

## 安全的双向通信

基于发布订阅模式，MQTT 允许在设备和云之间进行双向消息通信。发布订阅模式的优点在于：发布者与订阅者不需要建立直接连接，也不需要同时在线，而是由消息服务器负责所有消息的路由和分发工作。

MQTT 支持通过 TLS/SSL 确保安全的双向通信，同时 MQTT 协议中提供的客户端 ID、用户名和密码允许我们实现应用层的身份验证和授权。

## 在线状态感知

为了应对网络不稳定的情况，MQTT 提供了心跳保活（Keep Alive）机制。在客户端与服务端长时间无消息交互的情况下，Keep Alive 保持连接不被断开，若一旦断开，客户端可即时感知并立即重连。MQTT客户端可以设置一个心跳间隔时间(Keep Alive Timer)，表示在每个心跳间隔时间内发送一条消息。如果在这个时间周期内，没有业务数据相关的消息，客户端会发一个PINGREQ消息，相应的，服务器会返回一个PINGRESP消息进行确认。如果服务器在一个半(1.5)心跳间隔时间周期内没有收到来自客户端的消息，就会断开与客户端的连接。

同时，MQTT 设计了遗愿（Last Will） 消息。客户端在连接 server 的时候，可以设置是否发送遗嘱消息标志(Will Flag、Will Qos、Will Retain)、遗嘱消息主题(Topic)、遗嘱消息内容(Payload)。设置了遗嘱消息消息的 MQTT 客户端异常下线时（客户端断开前未向服务器发送 DISCONNECT 消息），MQTT 消息服务器会发布该客户端设置的遗嘱消息。

[MQTT 协议 Keep Alive 详解](https://www.emqx.com/zh/blog/mqtt-keep-alive)

[MQTT 遗嘱消息（Will Message）的使用](https://www.emqx.com/zh/blog/use-of-mqtt-will-message)

[MQTT 协议快速体验](https://www.emqx.com/zh/blog/the-easiest-guide-to-getting-started-with-mqtt)