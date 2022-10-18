# 面试问题准备

## 项目实践

1. 项目介绍

智能边缘平台、云边协同框架研发项目。

背景：公司各部门间存在众多形态各异的产品，各产品线各自管理自身研发体系、存在重复造轮子，资源浪费较严重的情况。
经各部门统一梳理，各产品线产品虽形态各异，但基本功能都大有雷同，故各部门领导牵头有此项目。

团队组成：多部门合作(tob企业业务研发部、tob平台技术中心-算法引擎、研究院算法sdk/模型团队、tog安防团队)

目标：公司级别统一的云边端开发框架，将云边端所需的设备管理、任务调度、算法引擎、算力监控、数据回流、云边协同、远程升级等能力提供抽象聚合实现。实际产品业务方在使用时可按需裁剪框架能力，做最小化的业务入侵。

个人核心工作：作为产品线业务研发身份参与，抽象实现业务侧所需的设备管理、任务调度、数据回流等模块。

2. 项目架构讲解

![云边端架构设计](./%E4%BA%91%E8%BE%B9%E7%AB%AF%E5%BC%80%E5%8F%91%E6%A1%86%E6%9E%B6.jpg)

说明：
- 整个框架主要分3大块，分别是云侧、边缘侧、以及端
- 云侧服务部署在公有云或私有云，负责所有边、端设备的管理、任务配置管理调度、OTA升级，另外提供openapi供其他业务系统集成
- 边缘侧负责局域网子设备管理，任务调度执行，充当算力中心的角色
- 端设备主要指IP摄像机等，主要提供视频流
- 边缘侧分为service、engine、infra三层抽象；所有组件支持拔插替换，可集群部署
- service负责提供云边交互、子设备管理、算法仓管理、任务调度、算力监控、事件输出等功能的实现
- engine层主要负责产品授权、视频统一接入、编解码、封装算法sdk、模型检测、事件输出
- infra层主要提供服务运行环境、各类基础组件
- 云侧与边侧主要通过mqtt实现影子协议进行信息交互，异步、声明式配置；图片、视频等数据使用aliyun-oss进行存储管理；OTA通过https协议实现、支持断点重传等
- 边侧内部服务通过grpc/ecosystem、消息队列(kafka/mqtt)交互，边侧与子设备通过tcp/ip、rtsp、gb2818、onvif、mqtt等协议进行接入
- 设备管理的实现，主要包括设备注册、心跳保持；子设备接入、视频流预览；云边信息同步
- 任务管理主要包括云边任务策略同步、算力监控、任务调度等功能
- 事件管理主要包括算法检测事件的再包装以及回流、图片帧信息回流、短视频截取

3. 主要技术栈
- etcd做为任务等关键数据的存储，支持监听、分布式
- sqlite主要用于云边同步时的影子消息存储，嵌入式、资源占用少
- pg主要用于事件数据的存储、对json数据的查找支持友好；此处曾实现过引入es做存储、支持搜索事件属性的方案，组件较重，后舍弃
- cassandra用于人脸特征数据的存储
- minio/ceph主要用于图片、短视频等二进制数据的对象存储
- kafka/mqtt用于服务间的消息传递

4. 项目挑战难点、如何解决、如何改进
- 项目整体涉及模块众多、设计相对复杂、跨部门合作
- 视频、图像处理领域之前没有接触、需要补充很多领域知识

- 多进行积极有效的沟通、遇到无法独自解决的问题尽早报风险、寻求领导、同事的帮助
- 业余多花时间补充相关领域知识、多看代码、代码是最好的老师
- 善用搜索、内部外部，你所遇到的大部分问题别人基本都遇到过，借助前人的力量解决问题

## 算法

1. [链表](./algorithm/list.md)

技术要点总结
- dummy node 哑巴节点
- 快慢指针

解题思路：
- 有环判断：快慢指针，快指针两倍速再次相遇即有环；此时slow回到链表头, 双指针同步前进， 再次相遇点即是入环点
- 倒数第N个：快慢指针，快指针先走N步，之后同步行进，fast==nil，slow即是倒数第N
- 找中点：快慢指针
- 回文链表：快慢指针，注意奇偶
- 奇偶重排：快慢指针
- 插入删除合并：使用dummy节点作为head
- 翻转：有遍历与递归两种方式；遍历要点——使用一个pre变量记录前一个node；递归要点——当前节点的next.next指向自己，cur.Next置空，递归下一个节点


算法应用：
- LRU 最近最少使用
    ```go
    // 实现要点 map + 双向链表
    // map记录所有数据
    // 链表用于按使用顺序进行排序
    type LRUCache struct {
        elems map[string]*ListNode
        count int // 记录当前数量
        cap int // 记录最大容量
        head *ListNode
        tail *ListNode
    }
    type ListNode struct {
        Key string
        Value interface
        Next *ListNode
        Pre *ListNode
    }

    func (c *LRUCache) Set(key string, value interface{}) {
        // 1. 先执行c.get操作
        // 2. 若 res!=nil update elems[key] return 
        // 3. 若 res==nil new ListNode and addToHead；判断若 count>cap delFromTail 
    }

    func (c *LRUCache) Get(key string) interface{} {
        // if elems[key] movetoHead
        // if !elems[key] return nil
    }
    ```

- SkipList 跳表
    ```go
    // leveln:
    // ......:
    // level3: H                   ->                  T
    // level2: H         ->        4         ->        T        
    // level1: H    ->   2    ->   4   ->    6    ->   T
    // level0: H -> 1 -> 2 -> 3 -> 4 -> 5 -> 6 -> 7 -> T

    // level0是基础链表，所有节点都会被添加进来
    // level1~leveln是索引链表，图例中每层是上一层的1/2
    // 如节点在第i层出现，则小于i的每一层都会出现该节点
    // 实际使用中，level层高是通过随机函数来确定的，会规范一个最大值
    const MaxLevel = 32
    type SkipList struct {
        head *SkipListNode  // 链表头
        level int // 当前层高
    }
    type SkipListNode struct {
        key int
        value interface{}
        next [MaxLevel]*SkipListNode
    }
    ```

2. [栈](./algorithm/stack.md)
先进后出的特点

- 括号有效性判断可借助栈
- 模拟计算器
    ```go
    // 模拟计算器的算法要点
    // 1. 借助栈结构：在遇到一个操作符时，将操作符前一个数字入栈，与上次缓存的操作符结合入栈；
    // 2. 前一个操作符时+、-操作时，直接结合入栈；*、/操作时，栈顶出栈，与当前数字运算后，将结果入栈；
    // 3. 缓存记录当前操作符
    // 4. 若表达式存在括号，需借助递归；遇到"("时进入递归，遇到")"时递归返回
    // 5. 最后sum(stack)
    func calculate(s string) int {
        pos := 0
        return helper(s, &pos)
    }

    func helper(s string, p *int) int {
        stack := make([]int, 0)
        
        // 初始化当前数字、操作符
        num := 0
        sign := '+'
        for *p<len(s) {
            c := s[*p]
            *p += 1
            if '0' <= c && c <= '9' {
                num += num*10 + int(c - '0')
            }
            if c == '+' || c == '-' || c == '*' || c == '/' {
                fmt.Printf("curSign:%c pos:%d preSign:%c num:%d \n", c, *p, sign, num)
                switch sign {
                case '+':
                    stack = append(stack, num)
                case '-':
                    stack = append(stack, -num)
                case '*':
                    stack[len(stack)-1] *= num
                case '/':
                    stack[len(stack)-1] /= num
                }

                // 更新sign, num
                sign = rune(c)
                num = 0
            }
            if c == '(' {
                num = helper(s, p)
            }
            if c == ')' {
                break
            }
        }
        if num != 0 {
            switch sign {
                case '+':
                    stack = append(stack, num)
                case '-':
                    stack = append(stack, -num)
                case '*':
                    stack[len(stack)-1] *= num
                case '/':
                    stack[len(stack)-1] /= num
                }
        }

        sum := 0
        for i:=0; i<len(stack); i++ {
            sum += stack[i]
        }
        return sum
    }
    ```

3. [树](./algorithm/tree.md)

技术要点：
- 遍历与递归，几乎所有题都是递归结构，合理设计递归函数返回值
- 前中后序遍历，解题时只需关注当前节点需要做啥操作、操作的位置(前中后)，剩余交给递归结构即可
- dfs算法要点：深度遍历——自底向下、自左向右的路径遍历，关注当前节点的操作，合理利用返回值(分治算法)
- bfs算法要点：层次遍历，维护一个queue，将每层的子节点添加到队尾，层次遍历到queue为空

题型总结：
- 遍历路径
- 翻转、最大深度、最大序列和：分治 递归
- 平衡二叉树：分别判断 并计算左右子树是否平衡，以及高度，且 |left-right| > 1 
- 最近公共祖先：分治、left!=nil && right!=nil return root, else return left == nil ? right:left
- 层次遍历、z型遍历、完整性校验：层次遍历解题
- 序列化反序列化：只有前序 后续 层序遍历可同时实现序列化、反序列化；整体代码结构与对应的遍历类似
- 二叉搜索树：利用有序的特性，减少复杂度
- 前缀树：多叉树结构，将每个字符作为路径存储，利用额外isEnd字段表示root到当前节点的路径是否表示完整单词
- btree/b+tree/avltree

4. [数组](./algorithm/array.md)

技术要点：
- 排序
- 双指针
- 滑动窗口
- 二分查找
- 前缀和
- 差分
- 花式遍历

题型总结：
- [排序](./algorithm/array.md#排序)：快排、归并、堆排序；堆排序可用于构建优先队列，解决topk问题
- [双指针](./algorithm/array.md#双指针)：回文串、两树和、接雨水、翻转
- [滑动窗口](./algorithm/array.md#滑动窗口)：最长无重复子串、最小覆盖子串、字母异位词、strstr、DNA序列
- [二分查找](./algorithm/array.md#二分查找)：模板要点 left + 1 < right 之后分别判断arr[left] arr[right]
- [差分数组](./algorithm/array.md#差分数组)：航班预定、拼车问题
- [前缀和](./algorithm/array.md#前缀和)

5. 字符串

技术要点：
- 回文串。中心扩散法，区分奇偶。

6. [回溯算法](./algorithm/backtrack.md)

技术要点：
- 穷举所有解，暴力递归。
- 排列问题，下次递归 从 0 开始, 需要一个记录已经做过选择的数组。函数签名 func backtrack(nums []int, visited []bool, track []int, res *[][]int)；针对有重复数字问题，先排序，然后把有序数字当做一个整体看待，即对i>0 && nums[i] == nums[i-1] && && !visited[i-1] 需做跳过的判断
- 组合问题，下次递归 从 i+1开始， 需要一个pos记录当前递归索引。函数签名 func backtrack(nums []int, pos int, track []int, res *[][]int)；针对有重复数字问题，先排序，对 i>pos && nums[i] == nums[i-1] 需做跳过的判断
- 子集问题，与组合类似

题型汇总：
- [排列问题](./algorithm/backtrack.md#排列)
- [组合](./algorithm/backtrack.md#组合)
- [子集](./algorithm/backtrack.md#子集)
- [n皇后](./algorithm/backtrack.md#其他)

7. [动态规划](./algorithm/dp.md)

技术要点：
- 问题基本求最值、是否可行、可行解个数，且不支持排序、交换等操作
- 通常定义一维或二维dp，明确状态转移——即当前问题可以由子问题递推，无法递推 重新明确dp定义；明确base case

题型汇总：
- [矩阵类问题](./algorithm/dp.md#矩阵类型-遍历递推)：通常是求从起始到终点，路径总数 最小路径和之类
    ```go
    // 定义二维数组dp[i][j]。表示从(0,0)出发到达(i,j)时的情况；或者表示从(i,j)出发到达终点时的情况
    // dp[i][j] 可由 dp[i-1][j] || dp[i][j-1]推导而来。表示若要到达(i, j), 必须从(i-1, j) 或 (i, j-1)出发
    // base case dp[0][...]  dp[...][j] 

    ```
- [跳跃&爬楼梯](./algorithm/dp.md#跳跃爬楼梯)：给出一个数组序列，从起点到达终点的路径数，最小路径和等
    ```go
    // 定义一维数组dp[i]。表示从 0 出发到达 i 时的结果；或者表示从 i 出发到达终点时的结果；两种方案都可
    // dp[i] 由 dp[i-1] dp[i-2] dp[i-nums[i]] 推导出
    // base case dp[0]=0 
    ```
- [子串子序列](./algorithm/dp.md#子序列子串问题)：一维数组dp[i], 表示将当前点作为结尾的结果，然后对s[i] 或者s[:i]做复合题意的判断
- [两序列比对](./algorithm/dp.md#两序列比对问题)：LCS、编辑距离、正则表达式、KMP
    ```go
    // 两个序列的动规问题，需定义一个二维dp表， 所代表的意义一般是 str1[:i] str2[:j] 的最值
    // 状态转移一般 dp[i][j] = min/max(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]) 此类型的方程
    // 需注意dp的长度通常会Len(str1)+1
    ```
- [零钱兑换](./algorithm/dp.md#零钱兑换)：
    ```go
    // 零钱兑换问题需定义一维dp数组dp[i]，表示凑出i金额时的所求解
    // 状态转移：为凑出i金额 可先凑出 ’i - 某一种面额的硬币’ 的金额，dp[i] = min(dp[i], dp[i-coins[j]]+1)
    // base case dp[0] = 0 dp[i] = amount+1
    // return dp[amount]
    ```
- [背包问题](./algorithm/dp.md#背包问题)：背包问题，求最值
    ```go
    // 背包问题需定义二维数组dp[i][w], 表示对于前i个物品，且前背包容量为w时 所求问题的结果——一般是最值
    // 状态转移：对于第i个物品，有两种选择——装入、不装入，对应的状态分别是dp[i-1][w]  value[i] + dp[i-1][w-wt[i]]
    // base case dp[0][...] = dp[...][0] = 0
    // return dp[n][w]
    ```
- [股票问题](./algorithm/dp.md#股票问题)：所有股票类问题
    ```go
    
    // 1. 无交易次数限制
    // 定义二维数组dp[i][j]，i代表第i天，j代表是否持有(取值0, 1)
    // 状态转移
    // dp[i][0] 第i天不持仓: i-1天也不持仓 或 i-1天持仓今天卖出； max(dp[i-1][0], dp[i-1][1]+prices[i])
    // dp[i][1] 第i天持仓：今天买入且前一天不持仓 或 i-1天持仓； max(dp[i-1][0]-prices[i], dp[i-1][1])
    // base case dp[0][0] = 0

    // 2. 有交易次数限制
    // 定义三维数组dp[i][k][j]，i代表第i天，k代表交易次数限制, j代表是否持仓
    // 含义：第i天还剩k次交易次数时，持仓不持仓分别所获的收益
    // 状态转移
    // dp[i][k][0] 第i天 最大k次交易 未持仓：前一天也未持仓今天不操作(dp[i-1][k][0]), 或 前一天持仓今天卖出(dp[i-1][k][1]+price[i]) 
    // dp[i][k][0] 第i天 最大k次交易 持仓：前一天未持仓今天买入(dp[i-1][k-1][0]-prices[i]), 或  前一天持仓今天不操作(dp[i-1][k][1])
    // base case dp[0][k][0] = 0 dp[0][k][1] = -prices[0] dp[i][0][0] = 0 dp[i][0][1] = math.MinInt32


    // 手续费、冷冻期不影响dp数组定义，在含有冷冻期时关注状态转移的i下标
    ```
- [打家劫舍](./algorithm/dp.md#打家劫舍)：
    ```go
    // 定义一维数组dp[i] 表示到第 i 家房屋，所能偷的最大金额
    // 状态转移：对当前房屋i有两种选择——偷、不偷，最大收益分别对应 dp[i-2] + nums[i], dp[i-1]; max(偷， 不偷)
    // base case dp[0] = 0 dp[1] = nums[0]
    // 对于环形链表，分两种base case：偷第一家不偷最后一家(dp[1]=nums[0]), 不偷第一家(dp[1]=0)
    ```

8. 数学运算

- 洗牌算法：保证每个元素被选取的概率都是1/n
    ```go
    func shuffle(nums []int) {
        for i:=0; i<len(nums); i++ {
            // 生成一个 [i, n-1] 区间内的随机数
            j := i + rand.Intn(n-i)
            // swap nums[i] nums[j]
            nums[i], nums[j] = nums[j], nums[i]
        }
    }
    ```
- 水塘抽样算法：一次遍历对一个序列随机抽取k个元素；
    ```go
    // 算法要点：当遇到第 i 个元素时，应该有 k/i 的概率选择该元素，1 - k/i 的概率保持原有的选择
    func getRandomK(nums []int, k int) []int {
        res := make([]int, k)
        
        // 首先默认选择前k个，当索引大于k时，对原有选择进行概率替换
        for i:=0; i<len(nums); i++ {
            if i<k {
                res[i] = nums[i]
                continue
            }
            // 生成【0, i]之间的随机数
            j := rand.Intn(i+1)
            // 这个整数小于 k 的概率就是 k/i
            if j<k {
                res[j] = nums[i]
            }
        }
        return res
    }
    ```
- 位操作：对二进制数据进行操作
    ```go
    // 1. n & (n-1) 消除数字n的二进制中的最后一个1。 
    // 可用于计算汉明权重：n = n & (n-1) 不断循环 至n==0
    // 判断n是否是2的指数：return (n & (n-1)) == 0 
    
    // 2. 一个数和它本身做异或运算结果为 0，即 a ^ a = 0；一个数和 0 做异或运算的结果为它本身，即 a ^ 0 = a。
    // 寻找只出现一次的数字：遍历 res ^= nums[i]
    // 寻找缺失的数字：遍历 res ^= i ^ nums[i]
    ```
- 阶乘末尾零的个数：
    ```go
    // 两数相乘末尾有0，那必然是可以分解出2*5这样一对因子。0~n中，2的数量远远大于5，故只要计算有多少个5的因子数就行了
    func tailZeroes(n int) int {
        res := 0
        divisor := 5
        for divisor <= n {
            res += n/divisor
            divisor *= 5
        }
        return res
    }
    ```

9. 其他

常见题型总结：
- [岛屿问题](./algorithm/backtrack.md#岛屿)：遍历矩阵，碰到陆地('1')之后，使用dfs(grid[i][j])进行淹没处理；dfs内部需对上下左右四个方向递归淹没。子岛屿问题先将grid2中的岛屿不存在Grid1中的进行淹没，然后统计岛屿数量。飞地问题/封闭岛屿问题， 先将四周边界的岛屿进行淹没，然后按题目需求进行求解


## 计算机基础

1. [进程、线程、协程的概念与区别](./linux/process-thread-coroutine.md)

经典的冯诺依曼结构把计算机系统抽象成 CPU + 存储器 + IO，那么计算机资源无非就两种：计算资源、存储资源

CPU只负责指令的计算，不负责存储分配。

[进程](./linux/process-thread-coroutine.md#进程)
- 进程是为存储资源的分配服务的，它是操作系统分配存储资源的的最小单位
- 进程拥有独立的虚拟地址空间，是个逻辑内存，它是实际物理内存分配的具体实现；虚拟地址空间最终通过页表进行映射
- 虚拟地址空间分内核空间、用户空间，在32位操作系统下可实现的内存映射为4G，内核与用户分别占1G、3G
- 内核空间存储内核代码、数据，以及内核堆栈——用于提供内核线程运行所需的资源，内核堆栈里保存着当前进程有关的信息(PCB结构、页表、文件描述符)；内核空间所有进程共享
- 用户空间存储用户代码、数据，以及用户代码执行所需的堆栈空间；用户空间每个进程独有
- 子进程从父类fork时，采用写时复制(COW)技术加快创建过程；实际就是直接把父进程的页面先直接COPY，故能映射到相同的物理内存，当需要对某块内存进行更改时，再将该内存块进行复制更新，同时更新页表结构
- 进程间通信只能基于信号、信号量、socket等方式，本质都需要借助内核来实现

[线程](./linux/process-thread-coroutine.md#线程)
- 线程是为计算资源的分配服务的，计算资源的分配通过内核调度器实现，而内核调度器调度的实体是(Kernal Scheduling Entry， KSE)，与用户线程一一对应，常称为内核线程
- 单核单CPU时代，一个进程可以近似的认为就是一个线程——进程在创建的时候，默认会创建thread0，用于对应内核调度实体
- 线程要基于进程创建，多线程共用进程的所有存储资源，但每个线程拥有独立的栈空间、以及自身的数据结构；栈空间用于函数调用，自身数据结构用于保存线程自身信息以及切换时的现场，用于上下文恢复
- 因为共享进程内存，故同一进程下的多个线程间可基于共享内存进行通信

[协程](./linux/process-thread-coroutine.md#协程)
- 协程可以看做是编程语言在用户态提供的线程，一种粒度更细的资源调度单元，进一步压榨cpu
- 协程的调度理想状态下不涉及内核态的切换，可以近似抽象的认为线程在占用一个CPU时间片内，调度执行多个不同的用户逻辑块
- 协程就是这些逻辑块的封装，操作系统不实现，需由编程语言自身运行时或者共享库来实现。C语系下有Coroutine, Goroutine
- 从协程的实现原理来看，协程适用于IO密集型场景——即在每次遇到IO阻塞时就去调度下一个用户代码块(协程)，最大限度压榨cpu，同时尽量减少线程切换
- 协程是用户态的实现，有自身的运行栈、以及状态信息，总体来说运行时的上下文信息大小比线程小很多 go里面初始栈大小为2kb Linux线程栈大小通常在2~4m

[切换调度](./linux/process-thread-coroutine.md#多任务切换)

进程的切换实际是基于线程的，通过内核调度器实现。切换流程如下
- 发生中断或系统调用，进入内核态执行进程上下文保存：寄存器信息入内核栈
- 内核栈的保存：将进程相关的内核栈信息保存在内核空间与进程对应的PCB结构中(包括寄存器、Task、页表、内核栈、文件描述符等)
- 加载将要被调度的进程信息到内核栈，切换页表
- PC寄存器更新为目标进程的内核栈
- 执行目标进程的内核栈，恢复寄存器状态，切换回用户态进程

- 流程总结：t1用户栈 -> t1内核栈 -> 执行现场保存、t2现场恢复 -> t2内核栈 -> t2用户栈 
- 主要涉及的资源有进程运行时寄存器状态信息、PCB结构、文件描述符、页表的保存与恢复
- 线程的切换在保存现场与恢复现场时略有不同，主要就是若是同一个进程内的线程，省略了页表文件描述符之类的切换，提高效率


2. [tcp/ip](./tcp%26ip/tcp%26ip.md)

**三次握手**
```go

client             server
         |  
         |
         | SYN=1 seq=J   |               
SYN_SENT | ----------->  | 
```


3. http2实现原理、以及与http的区别

## go

1. [go调度器实现](./language/golang.md#调度器)

go调度器是由goruntime实现的一种基于用户态线程而进行的调度操作，调度的对象是用户的程序代码块，也即goroutine。
整体调度逻辑的实现是一个调度器实例Sched + MGP三个模型抽象

**调度模型——MGP模型**
- M：系统线程的抽象。持有一个系统线程，绑定到P后，可调度执行goroutine抽象
- G：Goroutine抽象。用户代码块的抽象，反向持有m，同时存储有自身的栈信息，另外持有一个sched结构，存储调度相关的信息——现场保存、现场恢复
- P：处理器的抽象Process。持有本地运行队列、mcache、绑定的m；核心目标是通过资源本地化来解决协程调度时的资源竞争问题

- sched：调度器实例。组织运行MGP三大组件的管理者，持有空闲的P队列、空闲的M队列、全局的G队列、以及其他一些全局的缓存资源

**调度器启动&循环**

运行时启动：通过汇编实现，_rt0_amd64 -> runtime·rt0_go -> runtime·schedinit -> runtime·newproc -> runtime·mstart -> runtime·main -> main.main

- 在runtime·rt0_go中初始化了g0 m0, 分别表示系统栈、主线程；g0使用系统栈Linux下默认8M，不扩容，g0绑定m0，主要用于执行调度任务，以及其他一些需要依赖系统栈的函数执行

- runtime·schedinit
    主要执行了 stackinit -> mallocinit -> mcommoninit -> gcinit -> procresize；
    stackinit: goroutine执行栈初始化，实际分配在进程的堆内存，被go运行时分段锁定，只能用于goroutine执行栈；
    mallocinit: 内存分配器初始化，管理进程的堆内存，无法操作goroutine栈空间；
    mcommoninit: allm的相关初始化；
    gcinit: 垃圾回收器初始化；
    procresize: allp的初始化，根据runtime.GOMAXPROCS进行调整，绑定p0到m0；

- runtime·newproc
    主goroutine的初始化，使用g0将runtime·main封装为一个G，加入p0运行队列；
    runtime·main为主程序的入口，里面执行了sysmon、forcegchelper等函数的启动，最后执行main.main进入用户程序
    sysmon直接在单独的一个系统线程里启动，不参与调度器调度

- runtime·mstart
    mstart -> mstart1 -> schedule -> execute -> gogo -> ... -> goexit -> goexit1 -> goexit0 -> schedule 启动调度循环；
    mstart用于运行时初始化、以及每个新建的M启动；
    进入mstart1之后执行信号相关的初始化，之后执行schedule，并且正常情况下永不返回；当使用了LockOsThread，协程退出后，会回到mstart退出调度循环，执行mexit，退出系统线程
    在程序启动时，系统内的第一个goroutine——runtime·main被调度执行，之后进入到实际用户编码程序，执行相关代码，运行时启动完成。

**用户Goroutine创建及调度**

每次调用go时，都会：
- 调用runtime·newproc进入系统栈创建G，然后加入下当前P队列，或者全局队列
- 若此时存在空闲的P，即表示当前系统资源并未达最大使用值，创建M，并绑定到空闲P；若无空闲P至此退出当前逻辑，等待新建的G被调度执行即可
- M实际会创建并绑定到系统线程，执行mstart进入调度循环，寻找能执行的G
- 寻找顺序：当前P的本地队列、全局队列、网络定时器、偷其他P的等待队列

协程阻塞切换：
- 进入系统调用时，M、P、G都将进入syscall的状态，等待系统调用返回(比如网络IO)
- sysmon会根据P处于P_syscall状态的时间，P的待运行队列，当前系统的繁忙程度的情况判断是否retake当前P
- 被retake的P将解绑与M的关系，寻找自旋状态的M、或者创建新的M来调度执行P的剩余待运行G队列
- G从系统调用中返回时，会优先加入到原有的P队列，若原P队列已满，则加入到全局队列，等待再次被调度


2. go内存分配机制

使用的是空闲链表分配器，将内存按不同大小的内存分成多个等级，每个等级用链表进行管理分配。
整体实现借鉴TCMalloc，使用多级缓存将内存大小分类，对不同类别实施不同分配策略。

Go内存分配器核心主要包含一下几个组件：内存管理单元、线程缓存、中心缓存、页堆。

内存管理单元：runtime.mspan
- 双向链表结构，同时记录当前管理单元的内存起始地址、页数(页大小为8kb)
- allocBits 和 gcmarkBits — 分别用于标记内存的占用和回收情况
- allocCache通过位存储记录可被使用的内存、快速查找
- spanClass 跨度类，标识该管理的单元内存块大小，共有68个跨度类(8b~32kb)，ID=0的跨度类为特殊类型，用于管理>32kb的大内存
- spanClass 是uint8类型整数，前7位用于表示类别，最后一位表示是否包含指针

线程缓存：runtime.mcache
- 被P所持有的线程缓存，负责当前处理器所处理的任务的内存分配
- alloc数组，存储有136个内存管理单元(mspan)——共有68类，根据是否包含指针*2 = 136
- 在p初始化时，mcache被初始化为emptymspan，当需要申请内存时，从上一级mcentrel中获取内存
- tiny tinyoffset local_tinyallocs 一组用于微对象(<16b && 非指针)的内存分配器，内存与mcache结构连续

中心缓存：runtime.mcentral
- 分别维护有两个mspan集合，分别存储包含空闲、非空闲对象的列表
- 通过cacheSpan方法获取新的内存管理单元mspan，mcahce需申请内存时调用该方法
- 通过grow方法，从mheap中获取新的mspan

页堆：mheap
- 内存分配的入口，运行时初始化一个全局变量runtime.mheap_用于管理所有在堆上初始化的对象
- runtime.mheap_在schedinit -> mallocinit中被初始化，主要逻辑是初始化物理内存对应的页表大小
- mheap管理有central数组，长度为136的mcentral中心缓存结构体
- allspans存储维护所有内存管理单元
- arenas存储所有从虚拟地址空间申请的heapArena， 64位OS中每个arena管理64MB的内存空间，64MB按pagasize切分成相同大小的块被mspan管理
- mheap.alloc方法进行内存分配，返回mspan，在系统栈(g0)中被执行
- 实际会执行到pageCache.alloc， 无内存时调用mheap.grow向系统申请

总结：
- 通过全局的mheap_变量维护堆内存的申请与分配，以及GC
- 组件关系如下：mheap -> mcentral -> mcache
- 微小对象申请内存顺序：先通过本地P持有的mcache进行申请，不够时根据spanClass调用mheap_.central[spanClass]mcentral中心缓存进行申请，中心缓存再调用mheap.alloc方法申请，最终调用pageAlloc.alloc向内存申请
- 大对象直接调用mheap.alloc进行内存申请


3. go垃圾回收原理

go实现的是并发垃圾收集器，可与用户程序并行执行。分为标记和清扫两个主要阶段。标记使用的三色算法，为了减少STW的时间，引入混合写屏障实现三色不变性。

三色标记算法流程：
- 从灰色对象的集合中选择一个灰色对象并将其标记成黑色
- 将黑色对象指向的所有对象标记成灰色，保证该对象及该对象所引用的对象都不会被回收
- 重复上述过程，直至集合中不存在灰色对象

并发标记时引入屏障技术实现三色不变性：
- 强三色不变性 — 黑色对象不会指向白色对象，只会指向灰色对象或者黑色对象
- 弱三色不变性 — 黑色对象指向的白色对象必须包含一条从灰色对象经由多个白色对象的可达路径

go混合写屏障实现：
- 插入写屏障与删除写屏障组合使用
- 将被覆盖的指针和新指针都标记成灰色
- 将所有新创建的对象直接标记为黑色
- 开启屏障后，会影响所有的赋值操作从而影响用户程序。简单来说就是原先一条指令可完成的赋值，现在可能需要变成十几条指令

垃圾收集器的各个阶段：
- 清理终止阶段。完成上一阶段的清扫工作，检查是否满足垃圾收集启动条件；后台启动处理标记任务的goroutine、启动前准备、STW
- 标记阶段。
    状态切换到_GCmark、开启写屏障、用户程序协助（Mutator Assists）并将根对象入队；
    恢复应用程序(StarrTheWorld)，标记协程、以及用户程序(gcmalloc)开始并发标记内存中的对象；
    扫描根对象，包括所有 Goroutine 的栈、全局对象以及不在堆中的运行时数据结构，扫描 Goroutine 栈期间会暂停当前处理器；
    全部扫描完成后(终止算法)，调佣gcMarkDone进入标记终止阶段
- 标记终止阶段。暂停程序、将状态切换至 _GCmarktermination 并关闭辅助标记的用户程序；
- 清理阶段。
    状态切换至_GCoff，开始清理阶段，初始化清理状态并关闭写屏障；
    恢复用户程序，所有新创建的对象会标记成白色；
    后台并发清理所有的内存管理单元，当 Goroutine 申请新的内存管理单元时就会触发清理；

所有的垃圾收集启动入口都是gcStart，共用三个触发时机(所有引用runtime.gcTrigger结构体的地方都是触发点)，如下：
- runtime.sysmon 和 runtime.forcegchelper — 后台运行定时检查和垃圾收集；
- runtime.GC — 用户程序手动触发垃圾收集；
- runtime.mallocgc — 申请内存时根据堆大小触发垃圾收集；


4. go栈内存管理

goroutine栈内存设计：
- 用runtime.stack来表示，lo/hi指针分别对应寄存器的BP/SP
- 由runtime.stackpool、runtime.stackLarge两个全局变量进行管理，分别对应<32KB >32KB的内存分配
- 栈内存的实际管理单元同样是mspan，也即goroutine栈内存实际是分配在堆上
- 线程缓存mcache会持有一个stackcache，用于就近的本地g栈空间分配
- stackpool、stackLarge都无法满足要求时，会调用 runtime.stackcacherefill 从堆上获取新的内存

栈扩容：
- 编译器会为大部分函数调用插入runtime.morestack的运行时检查
- 当检测到需要扩容时，调用runtime.newstack创建新的栈来进行扩容
- 将旧栈中的所有内容复制到新栈中；
- 将指向旧栈对应变量的指针重新指向新栈；
- 销毁并回收旧栈的内存空间；

栈缩容：
- 运行时会在栈内存使用不足1/4时进行缩容，缩容调用runtime.shrinkstack
- 每次缩容都是原先的大小的一半
- 若缩容后大小小于2kB，停止缩容
- 缩容同样涉及数据拷贝、指针修改


5. slice底层实现、扩容机制

runttime.SliceHeader{
    Data uintptr    // 指向底层数组开头
    Len int         // 当前元素数量
    Cap int         // 最大可容纳的元素数量、也即底层数组的长度
}

cap<1024 若期望容量是当前容量的2倍以上，则使用期望容量；否则使用2倍
cap>1024 每次扩容25%直至满足期望容量

## redis

1. redis使用场景


2. redis作为缓存时的常见问题


3. 5大数据类型的底层实现


4. 持久化机制


5. 高可用方案


## mysql

1. 常用引擎


2. ACID事务一致性


## etcd

1. 使用场景


2. 实现原理


3. raft原理


4. 底层数据存储模型


## kafka

1. 使用场景


2. 使用过程中的问题


3. 实现原理


4. 消息队列选型对比


## elaticsearch

1. 使用场景


2. 组件架构原理


## k8s

1. 整体架构


2. 实现原理、各组件如何交互