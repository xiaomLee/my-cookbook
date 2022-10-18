# etcd

ETCD是一个分布式、可靠的key-value存储的分布式系统，用于存储分布式系统中的关键数据；当然，它不仅仅用于存储，还提供配置共享及服务发现；基于Go语言实现 。

## 应用场景

**服务注册与发现**

- 利用lease 绑定key进行ttl的自动续保实现服务注册、以及心跳机制
- 利用watch + prefix 实现服务的自发现

[example](https://www.cnblogs.com/ricklz/p/15059497.html)

**分布式配置系统**

- 利用key-value机制实现配置的分布式存储与发布，保证配置中心高可用
- 应用服务监听相应的配置Key，实现配置的热更新，可用于实现灰度发布、ABTest等应用场景

**分布式通知与协调**

- 利用对同一个目录的读写监听实现事件通知机制
- 任务管理器 master-worker模型， 用于master负责任务生产及调度，worker监听消费
- 监控系统检测服务健康状态

**分布式锁**

分布式锁应该具备哪些条件。

- 互斥性：在任意时刻，对于同一个锁，只有一个客户端能持有，从而保证一个共享资源同一时间只能被一个客户端操作；

- 安全性：即不会形成死锁，当一个客户端在持有锁的期间崩溃而没有主动解锁的情况下，其持有的锁也能够被正确释放，并保证后续其它客户端能加锁；

- 可用性：当提供锁服务的节点发生宕机等不可恢复性故障时，“热备” 节点能够接替故障的节点继续提供服务，并保证自身持有的数据与故障节点一致。

- 对称性：对于任意一个锁，其加锁和解锁必须是同一个客户端，即客户端 A 不能把客户端 B 加的锁给解了。

ETCD分布式锁的实现机制

- Lease 机制： 为key绑定租约，可主动续期，亦可防止客户端意外宕机无法主动释放锁的情况发生
- Revision 机制：没进行一次事务revision便会＋1，全局唯一；于是通过revision便可知道写入的顺序，在实现分布式锁时，多个客户端同时抢锁，根据 Revision 号大小依次获得锁
- Prefix 机制：目录机制，不同客户端对同一个prefix进行写入操作，结合watch机制，根据返回的revision号进行是否抢锁成功的判断
- Watch 机制：支持对前一个持有锁的客户端key进行监听，当其释放之后，当前客户端便自动拥有锁

实现源码位于：client/v3/concurrency/mutex.go

分布式锁demo
```go
package main

import (
	"context"
	"fmt"
	"log"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/concurrency"
)

func main() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Close()
	ctx := context.Background()
	// m1来抢锁
	go func() {
		s1, err := concurrency.NewSession(cli)
		if err != nil {
			log.Fatal(err)
		}
		defer s1.Close()
		m1 := concurrency.NewMutex(s1, "/my-lock/")

		// acquire lock for s1
		if err := m1.Lock(ctx); err != nil {
			log.Fatal(err)
		}
		fmt.Println("m1---获得了锁")

		time.Sleep(time.Second * 3)

		// 释放锁
		if err := m1.Unlock(ctx); err != nil {
			log.Fatal(err)
		}
		fmt.Println("m1++释放了锁")
	}()

	// m2来抢锁
	go func() {
		s2, err := concurrency.NewSession(cli)
		if err != nil {
			log.Fatal(err)
		}
		defer s2.Close()
		m2 := concurrency.NewMutex(s2, "/my-lock/")
		if err := m2.Lock(ctx); err != nil {
			log.Fatal(err)
		}
		fmt.Println("m2---获得了锁")

		// mock业务执行的时间
		time.Sleep(time.Second * 3)

		// 释放锁
		if err := m2.Unlock(ctx); err != nil {
			log.Fatal(err)
		}

		fmt.Println("m2++释放了锁")
	}()

	time.Sleep(time.Second * 10)
}
```


**Leader选主**

- 利用分布式锁的机制实现选主

实现源码：client/v3/concurrency/election.go


## raft原理

[raft原理详解](./raft.md)

## etcd-raft实现

[etcd实现raft源码解读](https://www.cnblogs.com/ricklz/p/15155095.html)


## etcd读写的线性一致性

**CAP的权衡**

根据定理，分布式系统只能满足三项中的两项而不可能满足全部三项。

AP wihtout C

允许分区下的高可用，就需要放弃一致性。一旦分区发生，节点之间可能会失去联系，为了高可用，每个节点只能用本地数据提供服务，而这样会导致全局数据的不一致性。

CA without P

如果不会出现分区，一直性和可用性是可以同时保证的。但是我们现在的系统基本上是都是分布式的，也就是我们的服务肯定是被多台机器所提供的，所以分区就难以避免。

CP without A

如果不要求A（可用），相当于每个请求都需要在Server之间强一致，而P（分区）会导致同步时间无限延长，如此CP也是可以保证的。

**线性一致性写**

所有的写操作，都要经过leader节点，一旦leader被选举成功，就可以对客户端提供服务了。客户端提交每一条命令都会被按顺序记录到leader的日志中，每一条命令都包含term编号和顺序索引，然后向其他节点并行发送AppendEntries RPC用以复制命令(如果命令丢失会不断重发)，当复制成功也就是大多数节点成功复制后，leader就会提交命令，即执行该命令并且将执行结果返回客户端，raft保证已经提交的命令最终也会被其他节点成功执行。

因为日志是顺序记录的，并且有严格的确认机制，所以可以认为写是满足线性一致性的。

**线性一致性读**

ReadIndex算法：

每次读操作的时候记录此时集群的commited index，当状态机的apply index大于或等于commited index时才读取数据并返回。由于此时状态机已经把读请求发起时的已提交日志进行了apply动作，所以此时状态机的状态就可以反应读请求发起时的状态，符合线性一致性读的要求。

etcd可通过参数控制是否启用顺序读(serializable read)
- 不启用，直接将当前节点的数据返回给client，此时从状态机读取的数据可能是落后已提交的数据的(apply index < commited index)
- 启用，将所有读请求转发给leader，leader通过ReadIndex算法读取状态机数据

Leader执行ReadIndex大致的流程：
1. 记录当前的commit index，称为ReadIndex；
2. 向 Follower 发起一次心跳，如果大多数节点回复了，那就能确定现在仍然是Leader；
3. 等待状态机的apply index大于或等于commited index时才读取数据；
4. 执行读请求，读取状态机数据，将结果返回给Client；

**小结**

etcd中对于写的请求，因为所有的写请求都是通过leader的，leader的确认机制将会保证消息复制到大多数节点中；

对于只读的请求，同样也是需要全部转发到leader节点中，通过ReadIndex算法，来实现线性一致性读；

raft执行ReadIndex大致的流程如下：

1. 记录当前的commit index，称为ReadIndex；
2. 向Follower发起一次心跳，如果大多数节点回复了，那就能确定现在仍然是Leader；
3. 等待状态机的apply index大于或等于commited index时才读取数据；
4. 执行读请求，将结果返回给Client。

关于状态机数据的读取，首先从treeIndex获取版本号，然后在buffer是否有对应的值，没有就去boltdb查询对应的值

## etcd存储实现(状态机的存储模型)

**v2 vs v3**

v2的问题
- 纯内存型key-value数据库，使用树形结构在内存中使用滑动窗口保存了最近的 1000 条变更事件，内存占用大，只适应小数据量场景
- 定时针对全量内存数据的快照消耗大量cpu io等资源
- 不支持多key事务，不支持范围查询
- 客户端协议是Http/1.x + json, 在有大量watcher的场景下，会创建大量连接，server负载大
- 没有lease机制， 需为每个key设置ttl，增加资源消耗


v3针对以上问题作出改进
- 引入mvcc+boltdb实现多key的事务机制，同时支持更大的历史数据回溯，range数据等操作
- 内存中只保留用户key到boltdb-key的映射关系，减少内存资源消耗
- 模块化key存储，与实际的磁盘存储实现，解耦各个组件
- 客户端协议使用GRPC + protobuf，数据传输更高效，资源消耗更少

**MVCC**

MVCC 机制是基于多版本技术实现的一种乐观锁机制，它乐观地认为数据不会发生冲突，但是当事务提交时，具备检测数据是否冲突的能力。

在 MVCC 数据库中，你更新一个 key-value 数据的时候，它并不会直接覆盖原数据，而是新增一个版本来存储新的数据，每个数据都有一个版本号，版本号是一个逻辑时钟，不会因为服务器时间的差异而受影响。

悲观锁是一种事先预防机制，它悲观地认为多个并发事务可能会发生冲突，因此它要求事务必须先获得锁，才能进行修改数据操作。但是悲观锁粒度过大、高并发场景下大量事务会阻塞等，相对来说会导致服务性能较差。

**ETCD-MVCC实现**

核心组件分成三个模块：TreeIndex, Backend, Compactor

TreeIndex：内存key索引，底层结构使用[google/b-tree](https://github.com/google/btree)，用于索引用户key到backend存储的数据

Backend：实际数据的存储端，使用[BoltDB](https://github.com/etcd-io/bbolt)实现

Compactor：异步压缩组件，针对lazy delete的数据执行压缩，释放/回收相应的内存，以及磁盘页

**TreeIndex**

在 treeIndex 中，每个节点的 key 是一个 keyIndex 结构，etcd 就是通过它保存了用户的 key 与版本号的映射关系。

keyIndex数据结构
```go
// etcd/server/mvcc/key_index.go

type keyIndex struct {
	key         []byte // 用户的key名称
	modified    revision // 最后一次修改key时的etcd版本号
	generations []generation // generation保存了一个key若干代版本号信息，每代中包含对key的多次修改的版本号列表
}

// generation contains multiple revisions of a key.
// generations 表示一个 key 从创建到删除的过程，每代对应 key 的一个生命周期的开始与结束。
// 当你第一次创建一个 key 时，会生成第 0 代，后续的修改操作都是在往第 0 代中追加修改版本号。
// 当你把 key 删除后，它就会生成新的第 1 代，一个 key 不断经历创建、删除的过程，它就会生成多个代。
type generation struct {
	ver     int64 // 表示此key的修改次数
	created revision // 表示generation结构创建时的版本号
	revs    []revision // 每次修改key时的revision追加到此数组
}

// A revision indicates modification of the key-value space.
// The set of changes that share same main revision changes the key-value space atomically.
type revision struct {
	// 一个全局递增的主版本号，随put/txn/delete事务递增，一个事务内的key main版本号是一致的
	main int64

	// 一个事务内的子版本号，从0开始随事务内put/delete操作递增
	sub int64
}
```

1. 更新Key
	
	执行 put 操作的时候首先从 treeIndex 模块中查询 key 的 keyIndex 索引信息。
	如果首次操作，也就是 treeIndex 中找不到对应的，etcd 会根据当前的全局版本号（空集群启动时默认为 1）自增，生成 put 操作对应的版本号 revision{2,0}，这就是 boltdb 的 key。
	如果能找到，在当前的 keyIndex append 一个操作的 revision。
	源码参见：etcd/server/mvcc/index.go func treeIndex.PUT

2. 查询Key

	在读事务中，它首先需要根据 key 从 treeIndex 模块获取版本号，如果未带版本号，默认是读取最新的数据。treeIndex 模块从 B-tree 中，根据 key 查找到 keyIndex 对象后，匹配有效的 generation，返回 generation 的 revisions 数组中最后一个版本号给读事务对象。
	读事务对象根据此版本号为 key，通过 Backend 的并发读事务（ConcurrentReadTx）接口，优先从 buffer 中查询，命中则直接返回，否则从 boltdb 中查询此 key 的 value 信息。
	上面是查找最新的数据，如果我们查询历史中的某一个版本的信息呢？处理过程是一样的，只不过是根据 key 从 treeIndex 模块获取版本号，不是获取最新的，而是获取小于等于 我们指定的版本号 的最大历史版本号，然后再去查询对应的值信息。

3. 删除key

	etcd 中的删除操作，是延期删除模式，和更新 key 类似.
	1、生成的 boltdb key 版本号追加了删除标识（tombstone, 简写 t），boltdb value 变成只含用户 key 的 KeyValue 结构体；
	真正删除 treeIndex 中的索引对象、boltdb 中的 key 是通过压缩 (compactor) 组件异步完成。

例子： 

put(key, 1), put(key, 2), put(key, 3), del(key), put(4), put(5), del(key), put(key, 6)

generations{
	{1.0, 2.0, 3.0(t)},
	{4.0, 5.0(t)},
	{6.0, 7.0},
}


**Backend**

通过Backend抽象，很好的封装了存储引擎的实现细节，为上层提供一个一致的接口，同时也方便做扩张替换
```go
// etcd/server/mvcc/backend/backend.go
type Backend interface {
	// ReadTx 返回一个读事务。它被主数据路径中的 ConcurrentReadTx 替换
	ReadTx() ReadTx
	BatchTx() BatchTx
	// ConcurrentReadTx returns a non-blocking read transaction.
	ConcurrentReadTx() ReadTx

	Snapshot() Snapshot
	Hash(ignores func(bucketName, keyName []byte) bool) (uint32, error)
	// Size 返回后端物理分配的当前大小。
	Size() int64
	// SizeInUse 返回逻辑上正在使用的后端的当前大小。
	SizeInUse() int64
	OpenReadTxN() int64
	Defrag() error
	ForceCommit()
	Close() error
}
```

提供了两个核心操作：并发读(ReadTx) 并发读写(BatchTx)


**Compactor**

ETCD支持人工压缩和自动压缩两种策略。

人工压缩：通过client提供的API

根据配置提供两种自动压缩方式
- 周期性压缩
- 版本号压缩

周期性压缩：启动时会创建periodic Compactor，它会异步的获取、记录过去一段时间的版本号。会将间隔参数划分成10个区间，比如period=1h 则每个区间6分钟。

版本号压缩：创建 revision Compactor，根据配置保留多少个历史版本号，auto-compaction-retention 为 10000。

压缩后boltdb并不会将释放的磁盘空间归还给系统，它会通过一个freelist page来记录空间的磁盘空间，当新的写请求来临时将首先去freelist里面请求空间。


**总结**

1. treeIndex 模块基于 Google 开源的 btree 库实现，它的核心数据结构 keyIndex，保存了用户 key 与版本号关系。每次修改 key 都会生成新的版本号，生成新的 boltdb key-value。boltdb 的 key 为版本号，value 包含用户 key-value、各种版本号、lease 的 mvccpb.KeyValue 结构体。

2. 如果我们不带版本号查询的时候，返回的是最新的数据，如果携带版本号，将会返回版本对应的快照信息；

3. 删除一个数据时，etcd 并未真正删除它，而是基于 lazy delete 实现的异步删除。删除原理本质上与更新操作类似，只不过 boltdb 的 key 会打上删除标记，keyIndex 索引中追加空的 generation。真正删除 key 是通过 etcd 的压缩组件去异步实现的；


## etcd对比Consul和zooKeeper如何选型

**选型对比**

1、并发原语：etcd 和 ZooKeeper 并未提供原生的分布式锁、Leader 选举支持，只提供了核心的基本数据读写、并发控制 API，由应用上层去封装，consul 就简单多了，提供了原生的支持，通过简单点命令就能使用；

2、服务发现：etcd 和 ZooKeeper 并未提供原生的服务发现支持，Consul 在服务发现方面做了很多解放用户双手的工作，提供了服务发现的框架，帮助你的业务快速接入，并提供了 HTTP 和 DNS 两种获取服务方式；

3、健康检查：consul 的健康检查机制，是一种基于 client、Gossip 协议、分布式的健康检查机制，具备低延时、可扩展的特点。业务可通过 Consul 的健康检查机制，实现 HTTP 接口返回码、内存乃至磁盘空间的检测，相比 etcd、ZooKeeper 它们提供的健康检查机制和能力就非常有限了；

etcd 提供了 Lease 机制来实现活性检测。它是一种中心化的健康检查，依赖用户不断地发送心跳续租、更新 TTL

ZooKeeper 使用的是一种名为临时节点的状态来实现健康检查。当 client 与 ZooKeeper 节点连接断掉时，ZooKeeper 就会删除此临时节点的 key-value 数据。它比基于心跳机制更复杂，也给 client 带去了更多的复杂性，所有 client 必须维持与 ZooKeeper server 的活跃连接并保持存活。

4、watch 特性：相比于 etcd , Consul 存储引擎是基于Radix Tree实现的，因此它不支持范围查询和监听，只支持前缀查询和监听，而 etcd 都支持, ZooKeeper 的 Watch 特性有更多的局限性，它是个一次性触发器。

5、线性读。etcd 和 Consul 都支持线性读，而 ZooKeeper 并不具备。

6、权限机制比较。etcd 实现了 RBAC 的权限校验，而 ZooKeeper 和 Consul 实现的 ACL。

7、事务比较。etcd 和 Consul 都提供了简易的事务能力，支持对字段进行比较，而 ZooKeeper 只提供了版本号检查能力，功能较弱。

8、多数据中心。在多数据中心支持上，只有 Consul 是天然支持的，虽然它本身不支持数据自动跨数据中心同步，但是它提供的服务发现机制、Prepared Query功能，赋予了业务在一个可用区后端实例故障时，可将请求转发到最近的数据中心实例。而 etcd 和 ZooKeeper 并不支持。

**总结**

- consul 提供了原生的分布式锁、健康检查、服务发现机制支持，让业务可以更省心，同时也对多数据中心进行了支持；

- etcd 和 ZooKeeper 也都有相应的库，也能很好的进行支持，但是这两者不支持多数据中心；

- ZooKeeper 在 Java 业务中选型使用的较多，etcd 因为是 go 语言开发的，所以如果本身就是 go 的技术栈，使用这个也是个不错的选择，Consul 在国外应用比较多，中文文档及实践案例相比 etcd 较少；


## 参考
[etcd的使用场景](https://www.cnblogs.com/ricklz/p/15033193.html)

[etcd中的线性一致性实现](https://www.cnblogs.com/ricklz/p/15204381.html)

[etcd状态机存储](https://wingsxdu.com/posts/database/etcd/#%E7%8A%B6%E6%80%81%E6%9C%BA%E5%AD%98%E5%82%A8)

[ETCD中的存储实现](https://boilingfrog.github.io/2021/09/10/etcd%E4%B8%AD%E7%9A%84%E5%AD%98%E5%82%A8%E5%AE%9E%E7%8E%B0/)

[etcd对比Consul和zooKeeper如何选型](https://www.cnblogs.com/ricklz/p/15292306.html)