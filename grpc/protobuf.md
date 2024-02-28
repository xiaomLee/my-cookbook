# protobuf

Protobuf 是由 Google 设计的一种高效、轻量级的信息描述格式, 起初是在 Google 内部使用, 后来被开放出来, 它具有语言中立、平台中立、高效、可扩展等特性, 它非常适合用来做数据存储、RPC数据交换等。与 json、xml 相比, Protobuf 的编码长度更短、传输效率更高, 其实严格意义上讲, json、xml 并非是一种「编码」, 而只能称之为「格式」, json、xml 的内容本身都是字符形式, 它们的编码采用的是 ASCII 编码, 本文讲述 Protobuf 的底层编码原理, 以便于了解 Protobuf 为什么编码长度短并且扩展性强, 与此同时我们也将了解到它有哪些不足

Protobuf 的一个典型应用场景便是做通信的数据交换格式, 它在通信管道上是以纯二进制的形式进行传输, 发送端使用编码器将数据序列化为二进制, 接收端使用解码器将收到的二进制流进行反序列化从而取得原始信息, 因此当对通信管道进行抓包时无法获知数据的详细内容, 事实上, 同一段 Protobuf 的二进制数据流, 在接收端使用不同的解码格式进行解码, 可能得到完全不同的信息。在了解 Protobuf 的底层编码细节之前, 需要首先了解 Protobuf 所用到的两种主要的编码方式, 它们分别是 Varints 编码和 Zigzag 编码。

## protobuf 编码原理





## 引用

[Protobuf编码规则详解](https://blog.csdn.net/mijichui2153/article/details/111475823)
[Protobuf编码](https://www.cnblogs.com/jialin0x7c9/p/12418487.html)
[Google Protobuf 编码原理](https://sunyunqiang.com/blog/protobuf_encode/)