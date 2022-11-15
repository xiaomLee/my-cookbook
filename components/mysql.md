
# MySQL

## 架构设计

整体mysql架构分为连接层、服务层、引擎层、存储层。

服务层的主要模块有如下：
1. 连接器。核心功能是监听客户端连接、权限认证、连接保持、响应请求。
2. SQL接口层。crud等操作的handler，根据查询请求做不同逻辑分发。此处会查询server层的全局缓存(mysql8.0版本已移除)，若存在所需数据，直接返回；反之继续往下。
3. 解析器。SQL词法分析，即将一个请求任务分成几个步骤进行操作。
4. 优化器。将解析器拆解的逻辑步骤生成查询路径树，然后选择出最优路径。
5. 执行器。一步步执行优化器选择的路径，实际就是不断调用存储引擎层的api进行数据存取。


## 存储引擎

### InnoDB

#### 内存缓冲池与三大日志

- 内存缓冲池

    InnoDB中有一个非常重要的内存组件——缓冲池（Buffer Pool），这里会缓存很多数据便于查询时，直接从读取缓存数据，而不需要访问磁盘。
    在执行更新操作时，如对“id=007”数据进行更新，会先查询BufferPool中是否存在，不存在的话，就从磁盘中加载到缓冲池中来，然后还要对这行记录加独占锁。

- undo日志

    所有写操作都会记录一个undolog用于历史数据回滚、以及mvcc多版本控制下快照读。

- redo日志

    在更新操作时会先更新buffer pool，再由异步线程执行刷盘，防止期间掉电宕机等，引入redolog，保证数据不丢失。
    写redolog是也是先写redolog buffer，可通过innodb_flush_log_at_trx_commit=1来控制刷盘策略。
    - commit=0时，事务提交成功，redo buffer不会写入redo log
    - commit=1时，只要事务提交成功，redo buffer一定会写入redo log（推荐）
    - commit=2时，事务提交成功，redo buffer先写入os cache，然后过段时间才刷入redo log

- binlog日志

    - redolog是有引擎层实现的一个偏物理性质的日志，InnoDB独有。而binlog是由server层的实现，记录的是对表中某行数据做了什么操作，且修改后的值是多少。binlog是主从同步的基础，可作为流量的重放。
    - binlog日志的落盘是在上面redo日志落盘以后才会去执行的，而且落盘策略也是可以选择直接刷盘还是先刷到OS cache中，这个配置项取决于sync_binlog。sync_binlog 这个参数设置成 1 的时候，表示每次事务的 binlog 都持久化到磁盘，这样可以保证 MySQL 异常重启之后 binlog 不丢失（一般都推荐这个值）。
    - binlog配合redolog实现两阶段提交，保证数据不丢失。

- IO线程随机刷盘

    当事务提交完毕以后，所有的日志文件都已经更新到最新了，此时系统不惧任何宕机断电行为了，后台的IO线程会随机把内存中的那个更新后的脏数据刷入到磁盘文件中，此时内存和磁盘中的数据已经保持一致。

**两阶段提交**
- 写操作在更新完buffer pool之后，会写入redolog，此时redolog 为prepare阶段
- 当执行commit之后，才会进行binlog的写入
- binlog写入成功后，将binlog文件名、offset等信息回写到redolog对应的记录上，同时对该记录打上commit标记
- 两阶段提交完成
- 目的：解决redolog binlog可能会存在的数据一致性问题；两个日志模块所充当的功能角色不一样，Redolog是为事务服务的，binlong是为数据归档、数据同步等功能服务的，在进行写操作时需要保证二者数据一致

**更新数据流程**
1. 加载磁盘文件到buffer Pool中；
2. 更新数据之前，写入旧数据到undo日志，便于回退；
3. 更新内存中的buffer pool数据；
4. 将更新部分的redo log写入到redo log buffer中；
5. redo日志刷入磁盘
6. binlog日志刷入磁盘
7. 将binlog文件和位置写入到redo日志文件中，并写入commit。
8. 后台的IO线程某个时间随机将buffer pool中的脏数据同步到磁盘文件。


#### 事务

数据库中的事务是指对数据库执行一批操作，这些操作最终要么全部执行成功，要么全部失败，不会存在部分成功的情况。

**事务ACID原则**
- 原子性(Atomicity) 不可分割的统一整体，要不全成功，要不全失败。
- 一致性(Consistency) 在事务执行开始至执行结束，不同客户端读取的数据都是一致的。
- 隔离性(Isolation) 一个事务的执行不能被其他事务干扰，并发执行的各个事务时独立的，无法取到其他事务中正在修改的对象。
- 持久性(Durability) 事务一旦提交，对数据的更改即是持久的，即得保证数据已经持久化落盘。

**mysql中的事务**
- 隐式事务：mysql中事务默认是隐式事务，执行写操作的时候，数据库自动开启事务、提交或回滚事务。是否开启隐式事务是由变量autocommit控制的。
```shell
mysql> show variables like 'autocommit';
+---------------+-------+
| Variable_name | Value |
+---------------+-------+
| autocommit    | ON    |
+---------------+-------+
1 row in set, 1 warning (0.00 sec)
```
- 显示事务：由开发者自己控制，通过start transaction / rollback / commit 等命令控制
- savepoint关键字：只rollback事务当中的某些部分。
```shell
mysql> start transaction;
Query OK, 0 rows affected (0.00 sec)

mysql> insert into test1 values (1);
Query OK, 1 row affected (0.00 sec)

mysql> savepoint part1;//设置一个保存点
Query OK, 0 rows affected (0.00 sec)

mysql> insert into test1 values (2);
Query OK, 1 row affected (0.00 sec)

mysql> rollback to part1;//将savepint = part1的语句到当前语句之间所有的操作回滚
Query OK, 0 rows affected (0.00 sec)

mysql> commit;//提交事务
Query OK, 0 rows affected (0.00 sec)

mysql> select * from test1;
+------+
| a    |
+------+
|    1 |
+------+
1 row in set (0.00 sec)
```
- 只读事务：mysql支持开启只读事务，start transaction read only; 只读事务内不可进行写操作。
```shell
mysql> commit;
Query OK, 0 rows affected (0.00 sec)

mysql> start transaction read only;
Query OK, 0 rows affected (0.00 sec)

mysql> select * from test1;
+------+
| a    |
+------+
|    1 |
|    1 |
+------+
2 rows in set (0.00 sec)

mysql> delete from test1;
ERROR 1792 (25006): Cannot execute statement in a READ ONLY transaction.
mysql> commit;
Query OK, 0 rows affected (0.00 sec)

mysql> select * from test1;
+------+
| a    |
+------+
|    1 |
|    1 |
+------+
2 rows in set (0.00 sec)
```

**事务中常见的问题**
- 脏读：一个事务在执行的过程中读取到了其他事务还没有提交的数据。
- 读已提交：一个事务操作过程中可以读取到其他事务已经提交的数据。
- 可重复读：一个事务操作中对于一个读取操作不管多少次，读取到的结果都是一样的。
- 幻读：事务中后面的操作（插入号码X）需要上面的读取操作（查询号码X的记录）提供支持，但读取操作却不能支持下面的操作时产生的错误，就像发生了幻觉一样。只会在可重复读的情况下会出现幻读。
```shell
事务A操作如下：
1、打开事务
2、查询号码为X的记录，不存在
3、插入号码为X的数据，插入报错（为什么会报错，先向下看）
4、查询号码为X的记录，发现还是不存在（由于是可重复读，所以读取记录X还是不存在的）

事物B操作：在事务A第2步操作时插入了一条X的记录，所以会导致A中第3步插入报错（违反了唯一约束）
上面操作对A来说就像发生了幻觉一样，明明查询X（A中第二步、第四步）不存在，但却无法插入成功
```

**事务隔离级别**
- 隔离级别分为4种：读未提交(READ-UNCOMMITTED)、读已提交(READ-COMMITTED)、可重复读(REPEATABLE-READ)、串行(SERIALIZABLE)
- 隔离级别越来越强，会导致数据库的并发性也越来越低
- 查看隔离级别：show variables like 'transaction_isolation';
- 设置隔离级别：修改mysql中的my.init文件 transaction-isolation 配置项，之后重启服务

[引用](https://cloud.tencent.com/developer/article/1516737?from=article.detail.1521495)

#### MVCC

全称Multi-Version Concurrency Control，即多版本并发控制。

**当前读、快照读**
- 当前读：就是它读取的是记录的最新版本，读取时还要保证其他并发事务不能修改当前记录，会对读取的记录进行加锁。像select lock in share mode(共享锁), select for update ; update, insert ,delete(排他锁)这些操作都是一种当前读
- 快照读：不加锁的select就是快照读，读取的可能是历史版本。串行隔离级别下快照读会退化成当前读。

**MVCC实现原理**
- 隐式字段。每行记录除了我们自定义的字段外，还有数据库隐式定义的DB_TRX_ID,DB_ROLL_PTR,DB_ROW_ID等字段

    - DB_ROW_ID 6byte, 隐含的自增ID（隐藏主键），如果数据表没有主键，InnoDB会自动以DB_ROW_ID产生一个聚簇索引 
    - DB_TRX_ID 6byte, 最近修改(修改/插入)事务ID：记录创建这条记录/最后一次修改该记录的事务ID 
    - DB_ROLL_PTR 7byte, 回滚指针，指向这条记录的上一个版本（存储于rollback segment里）
    - DELETED_BIT 1byte, 记录被更新或删除并不代表真的删除，而是删除flag变了 

- undo日志。记录插入、更新、删除前的原始行信息，用于rollback操作。同一记录行的所有历史ubdolog可根据DB_ROLL_PTR组成链表，最新的旧记录为链表头
- Read View(读视图)。事务进行快照读时候产生的一个读视图，记录维护当系统内当前活跃的事务ID，用于解决RR模式下的可重复读问题，视图核心结构如下。

    - trx_list 未提交事务ID列表，用来维护Read View生成时刻系统正活跃的事务ID 
    - up_limit_id 记录trx_list列表中事务ID最小的ID 
    - low_limit_id ReadView生成时刻系统尚未分配的下一个事务ID，也就是目前已出现过的事务ID的最大值+1

**整体流程**
- 开启事务
- 事务内快照读，比较当前DB_TRX_ID 与 readview.up_limit_id readview.low_limit_id的关系，是否进行undolog的历史版本查询
- 返回快照读结果

**RR解决不可重读的原理**
- 在RR级别下的某个事务的对某条记录的第一次快照读会创建一个快照及Read View, 将当前系统活跃的其他事务记录起来，此后在调用快照读的时候，还是使用的是同一个Read View，所以只要当前事务在其他事务提交更新之前使用过快照读，那么之后的快照读使用的都是同一个Read View，所以对之后的修改不可见； 
- 即RR级别下，快照读生成Read View时，Read View会记录此时所有其他活动事务的快照，这些事务的修改对于当前事务都是不可见的。而早于Read View创建的事务所做的修改均是可见
- RC级别下的，事务中，每次快照读都会新生成一个快照和Read View, 这就是我们在RC级别下的事务中可以看到别的事务提交的更新的原因
- RC隔离级别下，是每个快照读都会生成并获取最新的Read View；而在RR隔离级别下，则是同一个事务中的第一个快照读才会创建Read View, 之后的快照读获取的都是同一个Read View。

#### InnoDB小结
1. MySQL 默认的事务型存储引擎，只有在需要它不支持的特性时，才考虑使用其它存储引擎。
2. 实现了四个标准的隔离级别——读为提交、读已提交、可重复读、串行，默认级别是可重复读(REPEATABLE READ)。在可重复读隔离级别下，通过多版本并发控制(MVCC)+ 间隙锁(Next-Key Locking)防止幻影读。
3. 通过undolog redolog binlog的配合实现事务，以及数据崩溃恢复的能力


### MyISAM

不支持行级锁，只能对整张表加锁，读取时会对需要读到的所有表加共享锁，写入时则对表加排它锁。但在表有读取操作的同时，也可以往表中插入新的记录，这被称为并发插入(CONCURRENT INSERT)。 可以手工或者自动执行检查和修复操作，但是和事务恢复以及崩溃恢复不同，可能导致一些数据丢失，而且修复操作是非常慢的。 如果指定了 DELAY_KEY_WRITE 选项，在每次修改执行完成时，不会立即将修改的索引数据写入磁盘，而是会写到内存中的键缓冲区，只有在清理键缓冲区或者关闭表的时候才会将对应的索引块写入磁盘。这种方式可以极大的提升写入性能，但是在数据库或者主机崩溃时会造成索引损坏，需要执行修复操作。 ¶

### 比较

- 事务: InnoDB 是事务型的，可以使用 Commit 和 Rollback 语句。 
- 并发: MyISAM 只支持表级锁，而 InnoDB 还支持行级锁。 外键: InnoDB 支持外键。 
- 备份: InnoDB 支持在线热备份。 崩溃恢复: MyISAM 崩溃后发生损坏的概率比 InnoDB 高很多，而且恢复的速度也更慢。
-  其它特性: MyISAM 支持压缩表和空间数据索引。
- InnoDB支持事物，而MyISAM不支持事物
- InnoDB支持行级锁，而MyISAM支持表级锁
- InnoDB支持MVCC, 而MyISAM不支持
- InnoDB支持外键，而MyISAM不支持
- InnoDB不支持全文索引，而MyISAM支持
- 索引结构不一样，myisam表可以不存在主键索引

## 索引

### 索引类型

**B树索引**

B树是一个平衡多路查找树，并且所有叶子节点位于同一层，B为Blance，是为磁盘等外存储设备设计的一种平衡查找树。

系统从磁盘读取数据到内存时是以磁盘块（block）为基本单位的，位于同一个磁盘块中的数据会被一次性读取出来，而不是需要什么取什么。

InnoDB存储引擎中有页（Page）的概念，页是其磁盘管理的最小单位。InnoDB存储引擎中默认每个页的大小为16KB，可通过参数innodb_page_size将页的大小设置为4K、8K、16K，在MySQL中可通过如下命令查看页的大小：show variables like 'innodb_page_size';

而系统一个磁盘块的存储空间往往没有这么大，因此InnoDB每次申请磁盘空间时都会是若干地址连续磁盘块来达到页的大小16KB。InnoDB在把磁盘数据读入到磁盘时会以页为基本单位，在查询数据时如果一个页中的每条数据都能有助于定位数据记录的位置，这将会减少磁盘I/O次数，提高查询效率。

数据结构特点：
- 关键字集合分布在整棵树中；
- 任何一个关键字出现且只出现在一个结点中；
- 搜索有可能在非叶子节点结束；
- 其搜索性能等价于在关键字全集内做一次二分查找；

对比传统用来搜索的平衡二叉树，诸如AVL树、红黑树等，B树是一个多叉树，每个节点可以有>2个子节点(通常是pagesize/datasize)，故可以极大的减小树的高度，且一次磁盘IO能读入更多有效数据用于查找。
B+树的层高一般都在2~4层。


**B+树索引**

B+树是在B树基础上的一种优化，使其更适合实现存储索引结构，InnoDB存储引擎就是用B+Tree实现其索引结构。

数据结构特点：
- 非叶子节点只存储键值信息；
- 所有叶子节点之间都有一个链指针；
- 数据记录都存放在叶子节点中；

对比B树的优点：
- B+ 树的层级更少： 相较于 B 树 B+ 每个非叶子节点存储的关键字数更多(一个page页内，因为只存放key 不存放data，故有容纳更多的key)，树的层级更少所以查询数据更快；B+树的一般层高在2~4层，查找某一行数据时最多只要1~3次磁盘IO
- B+ 树查询速度更稳定： B+ 所有关键字数据地址都存在叶子节点上，所以每次查找的次数都相同所以查询速度要比B树更稳定
- B+ 树支持范围查询： 叶子节点的关键字从小到大有序排列，左边结尾数据都会保存右边节点开始数据的指针
- B+ 树天然具备排序功能： B+ 树所有的叶子节点数据构成了一个有序链表，在查询大小区间的数据时候更方便，数据紧密性很高，缓存的命中率也会比B树高
- B+ 树全节点遍历更快： B+ 树遍历整棵树只需要遍历所有的叶子节点即可，而不需要像 B 树一样需要对每一层进行遍历，这有利于数据库做全表扫描。


**物理分类**
- 聚簇索引

    - 存储记录在物理磁盘上市连续的，按照索引排序，故一个表最多有一个聚簇索引。
    - InnoDB通过主键聚集数据，若未定义主键，则以非空唯一索引代替，若还是没有，则以隐式字段DB_ROW_ID做主键。
    - 聚簇索引的B+树索引结构中，所有叶子结点存放的即是整张表的记录，从左至右按主键排序。

 - 非聚簇索引

    - 数据在物理存储上不按索引排序。
    - 叶子节点存储的并非整行数据，而是主键索引。
    - 查询时先通过辅助索引来查询到主键，之后再通过主键ID 查询聚簇索引来找到记录行信息，俗称回表。

**索引优化**
- 独立的列，索引列不能是表达式的一部分
- 联合索引，最左匹配原则
- where条件将选择性最强的列放在前面
- 主键选择具有自增属性的ID，更接近顺序插入，减少叶子节点分裂，磁盘页分裂，以及磁盘碎片的产生


## 性能优化

explain 可以查看整个sql的执行计划，explain出来的各种item的意义如下：
- select_type：表示查询中每个select子句的类型
- type：表示MySQL在表中找到所需行的方式，又称“访问类型”
- possible_keys：查询涉及到的字段上若存在索引，则该索引将被列出，但不一定被查询使用
- key：实际使用的索引，若没有使用索引，显示为NULL
- key_len：索引中使用的字节数
- ref：表示上述表的连接匹配条件，即哪些列或常量被用于查找索引列上的值
- Extra：包含不适合在其他列中显示但十分重要的额外信息

profile的意义以及使用场景
查询到 SQL 会执行多少时间, 并看出 CPU/Memory 使用量, 执行过程中 Systemlock, Table lock 花多少时间等等

慢查询分析工具mysqldumpslow

-s 表示按照何种方式排序
    c 访问次数
    l 锁定时间
    r 返回记录
    t 查询时间
    al 平均锁定时间
    ar 平均返回记录数
    at  平均查询时间
-t 返回前面多少条数据
-g 后边搭配一个正则匹配模式，大小写不敏感

mysqldumpslow -s r -t 10 /var/lib/mysql/695f5026f0f6-slow.log
mysqldumpslow -s t -t 10 /var/lib/mysql/695f5026f0f6-slow.log

## 分库分表

分表有水平切分，垂直切分的思路；水平切分表示按某一维度，比如时间，把表数据切分成结构一样的两个表；
垂直切分表示按列，把不同的列切分成不同的表。
当分表无法满足需求时可进行分库。
分库主要思路有将业务聚合度高的表作为一个单独的库，比如商品数据库、用户数据库；此外对于冷数据也可采用数据归档的方式


## 主从复制

目标：
- 数据分布 (Data distribution )
- 负载平衡(load balancing)
- 数据备份(Backups) ，保证数据安全
- 高可用性和容错行(High availability and failover)
- 实现读写分离，缓解数据库压力

主要主从结构：
- 一主多从
- 多主一从 从mysql5.7开始支持
- 双主复制 互为主备，数据相互同步
- 级联复制

原理及流程：
- 涉及三个线程：主节点log dump thread， 从节点 IO thread，SQL thread
- 当从节点上执行`start slave`命令之后，从节点会创建一个I/O线程用来连接主节点，请求主库中更新的bin-log。I/O线程接收到主节点的blog dump进程发来的更新之后，保存在本地relay-log（中继日志）中， 同时将binlog offset记录在master-info文件中，用于下次请求的参数
- 当从节点连接主节点时，主节点会为该从节点启动一个log dump thread，每个从节点都会维护一个；该线程用于读取主节点的binlog，在读取时会加锁，读取完之后释放锁，同时开始发送binlog
- 当检测到relaylog有新数据时，SQL thread开始启动，解析log，执行命令，同时记录当前relaylog offset到relay-log.info文件中

主从复制模式：
- 异步模式。该模式下，主库的读写操作以及事务跟从库无关，主库值复制写Binlog，binglog同步完全异步执行；从库数据落后于主库，此时主库宕机从库被选为主，存在数据丢失的风险。
- 半同步模式。主库执行完事务提交后不立即饭后，而是等待至少一个从库接收最新binlog，并写入到relaylog之后才进行返回；或者等待提交超时时间结束，降级为异步模式。不能保证从库已经应用binlog，且会造成一定延时，使用于低延时的网络环境。
- 全同步模式。主库提交完事务，等待所有从节点都复制并且执行了该事务之后才返回。

主从复制方式：基于SQL语句的复制（statement-based replication，SBR），基于行的复制（row-based replication，RBR)，混合模式复制（mixed-based replication,MBR)。对应的bin-log文件的格式也有三种：STATEMENT,  ROW,  MIXED。

## 参考

[mysql知识体系详解](https://pdai.tech/md/db/sql-mysql/sql-mysql-overview.html)

[mysql原理](https://www.jianshu.com/p/31770ad88010)

[mysql主从复制实现原理详解](https://blog.nowcoder.net/n/b90c959437734a8583fddeaa6d102e43)

[参考](https://segmentfault.com/a/1190000015815061)