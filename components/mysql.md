
# mysql

## 参考

[mysql知识体系详解](https://pdai.tech/md/db/sql-mysql/sql-mysql-overview.html)

1. myisam innodb的区别

    
    InnoDB支持事物，而MyISAM不支持事物
    InnoDB支持行级锁，而MyISAM支持表级锁
    InnoDB支持MVCC, 而MyISAM不支持
    InnoDB支持外键，而MyISAM不支持
    InnoDB不支持全文索引，而MyISAM支持
    索引结构不一样，myisam表可以不存在主键索引
    
    
   索引
   
    MyISAM的索引方式索引和数据存放是分开的，非聚集”的，所以也叫做非聚集索引。
    MyISAM中索引检索的算法为首先按照B+Tree搜索算法搜索索引，如果指定的Key存在，则取出其data域的值，
    data域里存放的是数据的行号，MyISAM会按照数据插入的顺序分配行号，从0开始，然后按照数据插入的顺序存储在磁盘上。
    如果行是定长的，可以从表的开头跳过相应的字节找到需要的行，变长有其他策略。然后以data域的值计算出地址，读取相应数据记录。
    
    innodb的索引分为一级索引和二级索引，一级索引即主键索引，使用的是聚簇索引(索引与数据存放在一起)，二级索引非聚簇,底层数据结构都是b+tree.
    聚簇索引的每一个叶子节点都包含了主键值、事务ID、用于事务和MVCC的回滚指针以及所有的剩余列，即所有数据。
    二级索引的叶子节点存放的是主键的值，这种策略的缺点是二级索引需要两次索引查找，第一次在二级索引中查找主键，第二次在聚簇索引中通过主键查找需要的数据行。
    
    聚簇索引的优点
    可以把相关数据存储在一起，减少数据查询时的磁盘I/O
    数据访问更快，因为聚簇索引就是表，索引和数据保存在一个B+Tree中
    使用索引覆盖的查询时可以直接使用页节点中的主键值
    
    缺点
    插入慢，严重依赖插入顺序。主键建议用自增ID，保证插入顺序行。

[参考](https://juejin.im/post/6844903701480472590)
    
    
2. 事务隔离级别

    
    读未提交(RU)
    读已提交(RC)
    可重复读(RR)
    串行
    
    innodb默认可重复读
    

3. 事务的实现原理

    
    

4. redo undo binlog的作用以及区别

    
    redolog
    用于重做保证数据完整性，防止内存中的数据未及时刷新到磁盘
    
    undolog
    用于数据回滚，对未提交的数据进行回滚操作
    
    binlog
    binlog是server层实现的,追加的方式写入的。记录数据变更的操作，有stament row 以及 mutil三中模式
    
    每当修改数据的时候，会记录redolog，同时记录一条相反的undo，此时记录状态是prepared，在记录binlog。
    commit的时候状态变成commit.

5. 性能调优

    
    explain出来的各种item的意义
    
    select_type：表示查询中每个select子句的类型
    type：表示MySQL在表中找到所需行的方式，又称“访问类型”
    possible_keys：查询涉及到的字段上若存在索引，则该索引将被列出，但不一定被查询使用
    key：实际使用的索引，若没有使用索引，显示为NULL
    key_len：索引中使用的字节数
    ref：表示上述表的连接匹配条件，即哪些列或常量被用于查找索引列上的值
    Extra：包含不适合在其他列中显示但十分重要的额外信息
    
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

6. innodb行锁 乐观锁和悲观锁

    乐观锁
    
    假定数据在读取期间不会被其他进程修改，每次提交时检查。innodb没有相应实现。
    实现方式：给表增加version or timestamp 列，每次提交时加上where id=xxx and version=xxx
    update col=xxx, version=version+1 where id=xxx and version=xxx
    
    悲观锁
    
    假定数据在读取期间会被其他进程修改，所以在操作之前提前加锁，使得只有自身或指定类型的进程才能操作当前数据。
    分为共享锁、排它锁，innodb引擎有相应实现。
    
    共享锁
    
    又称读锁，用户可以并发的读。会阻塞排它锁，数据使用共享锁后，update、insert、delete语句执行时会自动加排它锁进行阻塞等待。
    实现方式：select * from t where id=1 lock in share mode
    
    排他锁
    
    写锁，会阻塞其他的读锁与写锁。
    实现方式：select * from t where id=1 for update
    
    行锁/表锁
    
    指数据被锁定的方式，写锁、读锁都涉及。当查询所需要的数据不存在索引，或未能使用到索引，则会变成表锁。
    
    死锁排查
    
    查询是否锁表：show open tables where IN_use>0;
    查询进程：show processlist;
    查看事务：select * from INFORMATION_SCHEMA.INNODB_TRX;
    查看当前锁定的事务：select * from INFORMATION_SCHEMA.INNODB_LOCKS;
    查看当前等锁的事务：select * from INFORMATION_SCHEMA.INNODB_LOCK_WAITS;
    

[参考](https://segmentfault.com/a/1190000015815061)

7. drop、delete与truncate的区别
    
    
    1、delete和truncate只删除表的数据不删除表的结构
    
    
    2、速度,一般来说: drop> truncate >delete
    
    
    3、delete语句是dml,这个操作会放到rollback segement中,事务提交之后才生效;
    
    
    4、如果有相应的trigger,执行的时候将被触发. truncate,drop是ddl, 操作立即生效,原数据不放到rollback segment中,不能回滚. 操作不触发trigger.
    
    ps:
    1、不再需要一张表的时候，用drop
    
    2、想删除部分数据行时候，用delete，并且带上where子句
    
    3、保留表而删除所有数据的时候用truncate
