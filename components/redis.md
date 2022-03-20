## redis

内存数据库，拥有非常高的性能，单个实例的 QPS 能够达到 10W 左右。
    
### redis持久化
    
   redis提供了两种持久化的方式，分别是**RDB（Redis DataBase）**和**AOF（Append Only File）。**  
   RDB，简而言之，就是在不同的时间点，将redis存储的数据生成快照并存储到磁盘等介质上；  
   AOF，则是换了一个角度来实现持久化，那就是将redis执行过的所有写指令记录下来，在下次redis重新启动时，
   只要把这些写指令从前到后再重复执行一遍，就可以实现数据恢复了。
   
   **RDB**  
   RDB方式，是将redis某一时刻的数据持久化到磁盘中，是一种快照式的持久化方法。
   单独fork一个子进程进行写文件，先写到临时文件，结束后替换上次的快照。
   数据恢复快，存在数据不完整的可能。
   手动命令：save，bgsave，flushall
   预置的配置  
   save 900 1 # 15分钟内至少有一个键被更改  
   save 300 10 # 5分钟内至少有10个键被更改  
   save 60 10000 # 1分钟内至少有10000个键被更改
   
   
   **AOF**  
   配置appendonly yes  
   AOF文件持续增长而过大时，会fork出一条新进程来将文件重写(也是先写临时文件最后再rename)，遍历新进程的内存中数据，每条记录有一条的set语句。重写aof文件的操作，并没有读取旧的aof文件，
   而是将整个内存中的数据库内容用命令的方式重写了一个新的aof文件，这点和快照有点类似。
   Redis会记录上次重写时的AOF大小，默认配置是当AOF文件大小是上次rewrite后大小的一倍且文件大于64M时触发。
   
   
   *每次修改同步*：appendfsync always 同步持久化,每次发生数据变更会被立即记录到磁盘,性能较差但数据完整性比较好  
   *每秒同步*：appendfsync everysec 异步操作，每秒记录,如果一秒内宕机，有数据丢失 默认策略  
   *不同步*：appendfsync no   从不同步  
   
   redis提供了redis-check-aof工具，可以用来进行日志修复
   

### 主从同步

   **一主多从模型**  
   从向主发送sync指令，主执行bgsave，同时将此间的写入指令都缓存与内存中。
   结束后将rdb文件传输给从，从接受到后执行数据恢复的操作，然后接受主内存中的缓存指令。

   新版本支持增量同步，主服务器中会维护一个从服务器的ID，以及上次同步的offset，用于下次的增量同步。PSYNC指令。
    
### 内存淘汰机制

   **过期策略**  
   Redis是使用定期删除+惰性删除两者配合的过期策略。
   
   *定期删除*：Redis默认每隔100ms就随机抽取一些设置了过期时间的key，检测这些key是否过期，如果过期了就将其删掉。  
   *惰性删除*：客户端要获取某个key的时候，Redis会先去检测一下这个key是否已经过期，如果没有过期则返回给客户端，如果已经过期了，执行删除。
   
   
   **内存淘汰机制**  
   Redis在使用内存达到某个阈值（通过maxmemory配置)的时候，就会触发内存淘汰机制，选取一些key来删除。
   内存淘汰有许多策略，下面分别介绍这几种不同的策略。
   ```
   # maxmemory <bytes> 配置内存阈值
   # maxmemory-policy noeviction 
   
   noeviction：当内存不足以容纳新写入数据时，新写入操作会报错。默认策略
   allkeys-lru：当内存不足以容纳新写入数据时，在键空间中，移除最近最少使用的key。
   allkeys-random：当内存不足以容纳新写入数据时，在键空间中，随机移除某个key。
   volatile-lru：当内存不足以容纳新写入数据时，在设置了过期时间的键空间中，移除最近最少使用的key。
   volatile-random：当内存不足以容纳新写入数据时，在设置了过期时间的键空间中，随机移除某个key。
   volatile-ttl：当内存不足以容纳新写入数据时，在设置了过期时间的键空间中，有更早过期时间的key优先移除。
   ```
   
   lru算法：最近最少使用，Least Recently Used.
   ```
   // 算法实现
   // 需实现Get Set两个操作
   
   map+双向链表实现，count, cap记录当前的数量及容量， head tail分别记录头尾
   
   Get：map.get --> remove --> addToHead
   Set：map.get --> update --> remove --> addToHead
               |
               |
                --> addToHead --> count++ --> count > cap --> poptail --> count--
   ```
   
   持久化写入和重载时，都会判断key是否过期，过期不载入。

### 底层数据编码结构

   在Redis中，这5种基本类型的对象都是封装在**robj**这个结构体中.
   ```
   typedef struct redisObject {
       // 类型
       unsigned type:4;
   
       // 编码
       unsigned encoding:4;
   
       // 对象最后一次被访问的时间
       unsigned lru:REDIS_LRU_BITS; /* lru time (relative to server.lruclock) */
   
       // 引用计数，用于内存回收与对象共享
       int refcount;
   
       // 指向实际值的指针
       void *ptr;
   } robj;
   
   // 通过调用createObject方法可以创建其对象
   robj *createObject(int type, void *ptr) {
       robj *o = zmalloc(sizeof(*o));
       o->type = type;
       o->encoding = OBJ_ENCODING_RAW;
       o->ptr = ptr;
       o->refcount = 1;
   
       /* Set the LRU to the current lruclock (minutes resolution), or
        * alternatively the LFU counter. */
       if (server.maxmemory_policy & MAXMEMORY_FLAG_LFU) {
           o->lru = (LFUGetTimeInMinutes()<<8) | LFU_INIT_VAL;
       } else {
           o->lru = LRU_CLOCK();
       }
       return o;
   }
   
   // type 该属性表示对象的类型
   #define REDIS_STRING 0
   #define REDIS_LIST 1
   #define REDIS_SET 2
   #define REDIS_ZSET 3
   #define REDIS_HASH 4
   
   // encoding 该属性表示该类型的对象具体的实现
   #define REDIS_ENCODING_RAW 0     /* Raw representation */    //简单动态字符串
   #define REDIS_ENCODING_INT 1     /* Encoded as integer */    // 
   #define REDIS_ENCODING_HT 2      /* Encoded as hash table */
   #define REDIS_ENCODING_ZIPMAP 3  /* Encoded as zipmap */     // 压缩列表
   #define REDIS_ENCODING_LINKEDLIST 4 /* Encoded as regular linked list */ //
   #define REDIS_ENCODING_ZIPLIST 5 /* Encoded as ziplist */
   #define REDIS_ENCODING_INTSET 6  /* Encoded as intset */
   #define REDIS_ENCODING_SKIPLIST 7  /* Encoded as skiplist */
   #define REDIS_ENCODING_EMBSTR 8  /* Embedded sds string encoding */
   ```
   
   **简单动态字符串**
   
   ```
   struct sdshdr {
    // buf 中已占用空间的长度
    int len;
   
    // buf 中剩余可用空间的长度
    int free;
   
    // 数据空间
    char buf[];
   };
   
   O(1)时间获取字符串长度
   自动扩容，预分配空间以减少内存重新分配次数
   惰性删除，缩容时至更新free的值，并不真正释放内存
   二进制安全
   ```
  
  **ziplist 压缩链表**
  
   指向一块连续的内存，遍历的时候从尾部开始向前遍历，通过pre_content_length计算前一个数据块的大小
  
   | zlbytes | zltail | zlen | entry1 | entry2 | ... | entryN | zlend |
   
   zlbytes：记录整个压缩列表占用的内存字节数，在压缩列表内存重分配，或者计算zlend的位置时使用
   
   zltail：记录压缩列表表尾节点距离压缩列表的起始地址有多少字节，通过该偏移量，可以不用遍历整个压缩列表就可以确定表尾节点的地址
   
   zllen：记录压缩列表包含的节点数量，但该属性值小于UINT16_MAX（65535）时，该值就是压缩列表的节点数量，否则需要遍历整个压缩列表才能计算出真实的节点数量
   
   entryX：压缩列表的节点
   
   zlend：特殊值0xFF（十进制255），用于标记压缩列表的末端
  
   entry结构
  
   | pre_content_length | encoding | content |
   
   previous_entry_ength：记录压缩列表前一个字节的长度
   
   encoding：节点的encoding保存的是节点的content的内容类型
   
   content：content区域用于保存节点的内容，节点内容类型和长度由encoding决定
  
  
  **linkedlist 双向链表**
  
  **hashtable 哈希表**
  
  ```
    
    // hash表数组
    typedef struct dictht {
        // 哈希表数组
        dictEntry **table;
        // 哈希表大小
        unsigned long size;
        // 哈希表大小掩码，用于计算索引值，等于size-1
        unsigned long sizemask;
        // 哈希表已有节点的数量
        unsigned long used;
    } dictht;
   
    // hash值链表
    typedef struct dictEntry {
        // 键
        void *key;
        // 值
        union {
            void *val;
            uint64_t u64;
            int64_t s64;
            double d;
        } v;
        // 指向下一个哈希表节点，形成链表
        struct dictEntry *next;
    } dictEntry;
    
    // hash结构体
    typedef struct dict {
        // 和类型相关的处理函数
        dictType *type;
        // 私有数据
        void *privdata;
        // 哈希表, 一般只是用ht[0]， ht[1]在数据量较大 需要rehash时使用
        dictht ht[2];
        // rehash 索引，当rehash不再进行时，值为-1
        long rehashidx; /* rehashing not in progress if rehashidx == -1 */
        // 迭代器数量
        unsigned long iterators; /* number of iterators currently running */
    } dict;
    
  ```
  
  dict --> dictht --> dictEntry
  
  **skiplist 跳跃表**
  
  多层链表结构，插入时数据的层高随机决定
  
### 数据结构实现

   **string**  
   有三种编码实现：REDIS_ENCODING_INT REDIS_ENCODING_EMBSTR REDIS_ENCODING_RAW  
   EMBSTR RAW 内存实现都是简单动态字符串
       
   **EMBSTR**  
   用来保存短字符串的编码方式。
   当字符串保存的是一个小于等于44个字节的字符串时，那么robj对象里的属性ptr就会指向一个SDS对象。
   embstr编码通过调用一次内存分配函数来创建一块连续的内存空间，即redisObject对象和它的ptr指针指向的SDS对象是连续的。
   不过embstr编码的字符串对象是只读性的，一旦对其指向APPEND命令追加字符串会导致其变为raw编码实现。
   
   **RAW**  
   当字符串对象保存的是一个超过44个字节的字符串时。
   raw编码的字符串对象是可读可写的，对其指向APPEND命令追加字符串会不会导致其实现改变，
   如果追加的字符串的长度超过其free属性值，会在追加前重新进行内存空间分配。
   
   **list**  
   两种编码实现：REDIS_ENCODING_ZIPLIST REDIS_ENCODING_LINKEDLIST
   数据元素数量不超过128时，使用压缩链表，反之使用双向链表
   
   **hash**  
   编码实现：REDIS_ENCODING_ZIPLIST REDIS_ENCODING_HT
   数据元素数量不超过128时，使用压缩链表，反之使用hash表
   
   **set**  
   编码实现：REDIS_ENCODING_INTSET REDIS_ENCODING_HT
   存储的元素时整数时，使用整数集合实现， 反之哈希表实现
   
   **zset/ sorted set**  
   编码实现：REDIS_ENCODING_ZIPLIST REDIS_ENCODING_SKIPLIST
   数据元素数量不超过128时，使用压缩链表，反之使用跳跃链表+dict
   
  
### 常见问题
    
**缓存穿透**  
    指大量用户同时访问不存在的key，导致请求直接穿透到数据库层。
    
    解决方案：
    1. 对内部的合法key实行一定的命名规范，非法key未满足正则表达式的直接返回
    2. 结合具体场景，数据敏感度不高的业务，可只读缓存。另一个线程负责将数据库中的数据更新到缓存
    3. 将数据库查询回来的值，不管是否查到，都写入缓存。
    4. 布隆过滤器，将所有合法key预先使用布隆过滤器加载到内存
        
**热点数据/缓存击穿**  
    指大量请求同时访问同一个数据，而缓存恰好失效，此时qps直接打到数据库层
    
    解决方案：
    1. 请求数据库时先获取一个分布式锁，拿到锁的才能请求数据库，未拿到的自循环等待一定时间/次数
    2. 请求入队列，排队请求数据
    3. 对数据延时性要求不高的场景，可设置key不过期，异步线程负责更新数据
    
**缓存雪崩**  
    同一时间大面积的key集体失效，所有请求直接打到数据库

    解决方案：
    1. 给缓存的失效时间，加上一个随机值，避免集体失效
    2. 双缓存，A缓存有过期时间，B缓存长期有效，异步线程负责更新
    
**缓存、数据库一致性**  
    只能保证最终一致性。先更新数据库，再删缓存。可能存在删除缓存失败的问题，提供一个补偿措施即可，例如利用消息队列。
    
    解决方案：
    1. 一致性要求高场景，实时同步方案，即查询redis，若查询不到再从DB查询，保存到redis；
    2. 结合kafka mysql.binlog，异步线程消费消息更新缓存
    3. mysql触发器的机制，对数据库压力大

### redis集群方案
    
   **主从复制模式**
   
   主从复制模式包含一主多从，主实例负责读写，从实例只负责读。
   
   具体工作机制为：
   1. slave启动后，向master发送SYNC命令，master接收到SYNC命令后通过bgsave保存快照（即上文所介绍的RDB持久化），并使用缓冲区记录保存快照这段时间内执行的写命令
   2. master将保存的快照文件发送给slave，并继续记录执行的写命令
   3. slave接收到快照文件后，加载快照文件，载入数据
   4. master快照发送完后开始向slave发送缓冲区的写命令，slave接收命令并执行，完成复制初始化
   5. 此后master每次执行一个写命令都会同步发送给slave，保持master与slave之间数据的一致性
   
   优点：
   1. master能自动将数据同步到slave，可以进行读写分离，分担master的读压力
   2. master、slave之间的同步是以非阻塞的方式进行的，同步期间，客户端仍然可以提交查询或更新请求
   
   缺点：
   1. 不具备自动容错与恢复功能，master或slave的宕机都可能导致客户端请求失败，需要等待机器重启或手动切换客户端IP才能恢复
   2. master宕机，如果宕机前数据没有同步完，则切换IP后会存在数据不一致的问题
   3. 难以支持在线扩容，Redis的容量受限于单机配置
   
   **Sentinel（哨兵）模式**
   
   哨兵模式基于主从复制模式，只是引入了哨兵来监控与自动处理故障。哨兵的配置文件为sentinel.conf
   
   其功能包括
   1. 监控master、slave是否正常运行
   2. 当master出现故障时，能自动将一个slave转换为master（大哥挂了，选一个小弟上位）
   3. 多个哨兵可以监控同一个Redis，哨兵之间也会自动监控
   
   优点：
   
   1. 哨兵模式基于主从复制模式，所以主从复制模式有的优点，哨兵模式也有
   2. 哨兵模式下，master挂掉可以自动进行切换，系统可用性更高
   
   缺点：
   1. 同样也继承了主从模式难以在线扩容的缺点，Redis的容量受限于单机配置
   2. 需要额外的资源来启动sentinel进程，实现相对复杂一点，同时slave节点作为备份节点不提供服务
   
   **redis-cluster模式**
   
   Cluster采用无中心结构,它的特点如下：
   1. 所有的redis节点彼此互联(PING-PONG机制),内部使用二进制协议优化传输速度和带宽
   2. 节点的fail是通过集群中超过半数的节点检测失效时才生效
   3. 客户端与redis节点直连,不需要中间代理层.客户端不需要连接集群所有节点,连接集群中任何一个可用节点即可
   
   Cluster模式的具体工作机制：
   1. 在Redis的每个节点上，都有一个插槽（slot），取值范围为0-16383
   2. 当我们存取key的时候，Redis会根据CRC16的算法得出一个结果，然后把结果对16384求余数，这样每个key都会对应一个编号在0-16383之间的哈希槽，通过这个值，去找到对应的插槽所对应的节点，然后直接自动跳转到这个对应的节点上进行存取操作
   3. 为了保证高可用，Cluster模式也引入主从复制模式，一个主节点对应一个或者多个从节点，当主节点宕机的时候，就会启用从节点
   4. 当其它主节点ping一个主节点A时，如果半数以上的主节点与A通信超时，那么认为主节点A宕机了。如果主节点A和它的从节点都宕机了，那么该集群就无法再提供服务了
   
   Cluster模式集群节点最小配置6个节点(3主3从，因为需要半数以上)，其中主节点提供读写操作，从节点作为备用节点，不提供请求，只作为故障转移使用。

   优点：
   1. 无中心架构，数据按照slot分布在多个节点。
   2. 集群中的每个节点都是平等的关系，每个节点都保存各自的数据和整个集群的状态。每个节点都和其他所有节点连接，而且这些连接保持活跃，这样就保证了我们只需要连接集群中的任意一个节点，就可以获取到其他节点的数据。
   3. 可线性扩展到1000多个节点，节点可动态添加或删除
   4. 能够实现自动故障转移，节点之间通过gossip协议交换状态信息，用投票机制完成slave到master的角色转换
   
   缺点：
   1. 客户端实现复杂，驱动要求实现Smart Client，缓存slots mapping信息并及时更新，提高了开发难度。目前仅JedisCluster相对成熟，异常处理还不完善，比如常见的“max redirect exception”
   2. 节点会因为某些原因发生阻塞（阻塞时间大于 cluster-node-timeout）被判断下线，这种failover是没有必要的
   3. 数据通过异步复制，不保证数据的强一致性
   4. slave充当“冷备”，不能缓解读压力
   5. 批量操作限制，目前只支持具有相同slot值的key执行批量操作，对mset、mget、sunion等操作支持不友好
   6. key事务操作支持有线，只支持多key在同一节点的事务操作，多key分布不同节点时无法使用事务功能
   7. 不支持多数据库空间，单机redis可以支持16个db，集群模式下只能使用一个，即db 0
   
   Redis Cluster模式不建议使用pipeline和multi-keys操作，减少max redirect产生的场景。
   
   **Twitter-Twemproxy**
   
   基本原理是：Redis客户端把请求发送到Twemproxy，Twemproxy根据路由规则发送到正确的Redis实例，最后Twemproxy把结果汇集返回给客户端。
   
   优点：
   1. 客户端像连接Redis实例一样连接Twemproxy，不需要改任何的代码逻辑。
   2. 支持无效Redis实例的自动删除。
   3. Twemproxy与Redis实例保持连接，减少了客户端与Redis实例的连接数。
   
   缺点：
   1. 由于Redis客户端的每个请求都经过Twemproxy代理才能到达Redis服务器，这个过程中会产生性能损失。
   2. 没有友好的监控管理后台界面，不利于运维监控。
   3. 最大的问题是Twemproxy无法平滑地增加Redis实例。对于运维人员来说，当因为业务需要增加Redis实例时工作量非常大。
   
   **Codis 豌豆荚**
   
   支持平滑增加Redis实例的Redis代理软件，其基于Go和C语言开发，开源。
   
   Codis包含下面4个部分：
   1. Codis Proxy：Redis客户端连接到Redis实例的代理，实现了Redis的协议，Redis客户端连接到Codis Proxy进行各种操作。
   Codis Proxy是无状态的，可以用Keepalived等负载均衡软件部署多个Codis Proxy实现高可用。
   
   2. CodisRedis：Codis项目维护的Redis分支，添加了slot和原子的数据迁移命令。Codis上层的 Codis Proxy和Codisconfig只有与这个版本的Redis通信才能正常运行。
   
   3. Codisconfig：Codis管理工具。可以执行添加删除CodisRedis节点、添加删除Codis Proxy、数据迁移等操作。
   另外，Codisconfig自带了HTTP server，里面集成了一个管理界面，方便运维人员观察Codis集群的状态和进行相关的操作，极大提高了运维的方便性，弥补了Twemproxy的缺点。
   
   4. ZooKeeper：Codis依赖于ZooKeeper存储数据路由表的信息和Codis Proxy节点的元信息。另外，Codisconfig发起的命令都会通过ZooKeeper同步到CodisProxy的节点。
   
   优点：
   1. 提供平滑的增加redis实例的解决方案
   2. 集成了管理端界面，方便运维
   3. 支持动态数据迁移
   
   缺点：
   1. 需使用codis维护的codis-redis分支
   2. 最大支持1024个节点