## raft

### 原理
   
follower：所有结点都以follower的状态开始。如果没收到leader消息则会变成candidate状态。  
candidate：会向其他结点“拉选票”，如果得到大部分的票则成为leader。这个过程就叫做Leader选举(Leader Election)。  
leader：所有对系统的修改都会先经过leader。每个修改都会写一条日志(log entry)。leader收到修改请求执行日志复制(Log Replication)。  

```
复制日志到所有follower结点(replicate entry)
大部分结点响应时才提交日志
通知所有follower结点日志已提交
所有follower也提交日志
现在整个系统处于一致的状态
```

全局只有两种请求
leader发出的Append Entries，心跳与日志都是通过它进行发送的。
candidate发出的vote

[参考](http://thesecretlivesofdata.com/raft/)
   
### 选举算法

当follower在选举超时时间(election timeout)内未收到leader的心跳消息(append entries)，则变成candidate状态。
为了避免选举冲突，这个超时时间是一个150~300ms之间的随机数。

成为candidate的结点发起新的选举期(election term)去“拉选票”：
1. 重置自己的计时器
2. 投自己一票
3. 发送 Request Vote消息

如果接收结点在新term内没有投过票那它就会投给此candidate，并重置它自己的选举超时时间。  
candidate拉到大部分选票就会成为leader，并定时发送心跳——Append Entries消息，去重置各个follower的计时器。当前Term会继续直到某个follower接收不到心跳并成为candidate。

如果不巧两个结点同时成为candidate都去“拉票”怎么办？这时会发生Splite Vote情况。
两个结点可能都拉到了同样多的选票，难分胜负，选举失败，本term没有leader。
之后又有计时器超时的follower会变成candidate，将term加一并开始新一轮的投票。

### 日志复制

当发生改变时，leader会复制日志给follower结点，这也是通过Append Entries心跳消息完成的。  