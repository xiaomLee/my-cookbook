# 链表

## 基本技巧

链表相关的核心点

- null/nil 异常处理
- dummy node 哑巴节点
- 快慢指针
- 插入一个节点到排序链表
- 从一个链表中移除一个节点
- 翻转链表
- 合并两个链表
- 找到链表的中间节点

## 常见题型

### 反转

[206.反转链表](https://leetcode.cn/problems/reverse-linked-list/)
```go
func reverseList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    node := reverseList(head.Next)
    head.Next.Next = head
    head.Next = nil
    return node
}
```

[92.反转链表II](https://leetcode.cn/problems/reverse-linked-list-ii/)
```go
func reverseBetween(head *ListNode, left int, right int) *ListNode {
    if left == 1 {
        return reverseN(head, right)
    }
    head.Next = reverseBetween(head.Next, left-1, right-1)
    return head
}

var tail *ListNode
func reverseN(head *ListNode, n int) *ListNode {
    if n == 1 {
        tail = head.Next
        return head
    }
    p := reverseN(head.Next,n-1)
    head.Next.Next = head
    head.Next = tail
    return p
}


func reverseN(head *ListNode, n int) *ListNode {
    // 1->2->3->4->5   n=3  pre=nil  cur=1
    var pre *ListNode
    cur := head
    for ; n>0; n-- {
        cur.Next, pre, cur = pre, cur, cur.Next
    }
    // 1<-2<-3 4->5  n=0  pre=3 cur=4 head=1
    
    head.Next = cur
    // 5<-4<=1<-2<-3
    return pre
}

```

[25.K个一组翻转链表](https://leetcode.cn/problems/reverse-nodes-in-k-group/)
```go
func reerseKGroup(head *ListNode, k int) *ListNode {
    if head == nil {
        return head
    }
    // 判断是否有k个元素
    a, b := head, head
    for i:=0; i<k; i++ {
        if b == nil {
            return head
        }
        b = b.Next
    }
    newHead := reverseBetween(a, b)
    a.Next = reverseKGroup(b, k)
    return newHead
}

func reverseBetween(a, b *ListNode) *ListNode {
    var pre *ListNode
    cur := a
    for cur != b {
        cur.Next, pre, cur = pre, 
    }
    return pre
}
```

**小结**
- 遍历的方式需使用一个pre变量记录前一个节点
- 递归需要关注当前节点需要做啥
- 善用子函数抽象拆解问题 reverseN(head, n) reverseBetween(a, b *ListNode)

### 合并 分割 相加

[21.合并两个有序链表](https://leetcode.cn/problems/merge-two-sorted-lists/)
```go
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    // 思路 1 递归解法 mergeTwoLists(list1.Next, list2) or mergeTwoLists(list1, list2.Next)
    // 思路 2 遍历list1 list2: 初始化dummy节点， p p1 p2分别指向三个链表
    dummy := &ListNode{}
    p, p1, p2 := dummy, list1, list2

    for p1 != nil && p2 != nil {
        if p1.Val < p2.Val {
            p.Next = p1
            p1 = p1.Next
        }else {
            p.Next = p2
            p2 = p2.Next
        }
        p = p.Next
    }
    if p1 != nil {
        p.Next = p1
    }
    if p2 != nil {
        p.Next = p2
    }
    return dummy.Next
}
```

[23.合并K个升序链表](https://leetcode.cn/problems/merge-k-sorted-lists/)
```go
func mergeKLists(lists []*ListNode) *ListNode {
    // 思路 二分递归合并 转换成 mergeTwoLists
    if len(lists) <=0 {
        return nil
    }
    if len(lists) == 1 {
        return lists[0]
    }
    mid := len(lists)/2
    left := mergeKLists(lists[:mid])
    right := mergeKLists(lists[mid:])
    return mergeTwoLists(left, right)
}

func mergeTwoLists(list1, list2 *ListNode) *ListNode {
    if list1 == nil {
        return list2
    }else if list2 == nil {
        return list1
    }else if list1.Val < list2.Val {
        list1.Next = mergeTwoLists(list1.Next, list2)
        return list1
    }else {
        list2.Next = mergeTwoLists(list1, list2.Next)
        return list2
    }
}
```

[86.分隔链表](https://leetcode.cn/problems/partition-list/)
```go
func partition(head *ListNode, x int) *ListNode {
    if head == nil {
        return nil
    }
    // 思路 合并两个链表的逆操作, 
    // 定义两个链表 small large；之后将large 接与small之后
    dummy1, dummy2 := &ListNode{}, &ListNode{}
    small, large :=dummy1, dummy2
    for head != nil {
        if head.Val < x {
            small.Next = head
            small = small.Next
        }else {
            large.Next = head
            large = large.Next
        }
        head = head.Next
    }
    // 防止有环 leetcode 校验结果时有环会oom
    large.Next = nil
    small.Next = dummy2.Next
    return dummy1.Next
}
```


### 有环、公共节点

[环形链表](https://leetcode.cn/problems/linked-list-cycle/)
```go
func hasCycle(head *ListNode) bool {
    if head == nil {
        return false
    }
    // 快慢指针 快指针两倍速 若再次相遇，说明有环
    slow, fast := head, head
    for fast.Next != nil && fast.Next.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
        if slow == fast {
            return true
        }
    }
    return false
}
```

[环形链表II](https://leetcode.cn/problems/linked-list-cycle-ii/)
```go
func detectCycle(head *ListNode) *ListNode {
    if head == nil {
        return nil
    }
    // 快慢指针 快指针两倍速 若再次相遇，说明有环
    slow, fast := head, head
    for fast.Next != nil && fast.Next.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
        if slow == fast {
            break
        }
    }
    // nocycle
    if fast.Next == nil || fast.Next.Next == nil {
        return nil
    }
    // 此时slow回到链表头, 双指针同步前进， 再次相遇点即是入环点
    slow = head
    for slow != fast {
        slow = slow.Next
        fast = fast.Next
    }
    return slow
}
```

[相交链表](https://leetcode.cn/problems/intersection-of-two-linked-lists/)
```go
func getIntersectionNode(headA, headB *ListNode) *ListNode {
    // 思路1 
    // 将b接在a的后面， 从a出发判断是否有环 以及环的起点
    // 思路2
    // 分别从a-> b->a 遍历链表 若中途相遇即是相交点
    if headA == nil || headB == nil {
        return nil
    }
    a, b := headA, headB
    for a!=b {
        if a == nil {
            a=headB
        }else {
            a = a.Next
        }
        if b == nil {
            b = headA
        }else {
            b = b.Next
        }
    }
    return a
}
```


### 查找 删除、插入

[链表的中间结点](https://leetcode.cn/problems/middle-of-the-linked-list/)
```go
func middleNode(head *ListNode) *ListNode {
    if head == nil {
        return nil
    }
    // 快慢指针 快指针两倍速
    slow, fast := head, head
    for fast.Next!=nil && fast.Next.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
    }
    return slow
}
```

[删除链表的倒数第 N 个结点](https://leetcode.cn/problems/remove-nth-node-from-end-of-list/)
```go
func removeNthFromEnd(head *ListNode, n int) *ListNode {
    if head == nil {
        return nil
    }
    // 快慢指针
    // 注意点 头结点也可能被删 故初始化dummy
    dummy := &ListNode{}
    dummy.Next = head
    slow, fast := dummy, dummy

    // fast 先走n+1步 保证删除的时候是删除slow.Next
    for i:=0; i<=n && fast != nil; i++ {
        fast = fast.Next
    }
    for fast != nil {
        slow = slow.Next
        fast = fast.Next
    }
    // delete slow.Next
    slow.Next = slow.Next.Next
    return dummy.Next
}
```

[旋转链表](https://leetcode.cn/problems/rotate-list/)
```go
func rotateRight(head *ListNode, k int) *ListNode {
    n := 0
    p := head
    for p!=nil {
        p = p.Next
        n++ 
    }
    if n == 0 || k%n == 0{
        return head
    }
    p = head
    for i:=0; i<(n-1-k%n); i++ {
        p = p.Next
    }
    newHead := p.Next
    p.Next = nil

    p = newHead
    for p.Next != nil {
        p = p.Next
    }
    p.Next = head
    return newHead
}
```

[删除排序链表中的重复元素](https://leetcode.cn/problems/remove-duplicates-from-sorted-list/)

[删除排序链表中的重复元素II](https://leetcode.cn/problems/remove-duplicates-from-sorted-list-ii/)
```go
func deleteDuplicates(head *ListNode) *ListNode {
    if head == nil {
        return nil
    }
    // 思路1 hash+两次遍历
    // 思路2 指针 cur.Next == cur.Next.Next delete cur.Next = cur.Next.Next
    dummy := &ListNode{}
    dummy.Next = head

    cur := dummy
    for cur.Next != nil && cur.Next.Next !=nil {
        if cur.Next.Val == cur.Next.Next.Val {
            reaptVal := cur.Next.Val
            // delete reapt
            for cur.Next != nil && cur.Next.Val == reaptVal {
                cur.Next = cur.Next.Next
            }
        }else {
            cur = cur.Next
        }
    }
    return dummy.Next
}
```

### 排序

[排序链表](https://leetcode.cn/problems/sort-list/)

[对链表进行插入排序](https://leetcode.cn/problems/insertion-sort-list/)

[重排链表](https://leetcode.cn/problems/reorder-list/)

### 其他 

[回文链表](https://leetcode.cn/problems/palindrome-linked-list/)
[奇偶链表](https://leetcode.cn/problems/odd-even-linked-list/)


## 常见工程应用

### LRU LFU

### SKIPLIST(跳表)

跳表本质上是一个链表, 它其实是由有序链表发展而来。跳表在链表之上做了一些优化，跳表在有序链表之上加入了若干层用于索引的有序链表。索引链表的结点来源于基础链表，不过只有部分结点会加入索引链表中，并且越高层的链表结点数越少。跳表查询从顶层链表开始查询，碰到比自身大的，往下跳一级，然后逐级展开，直到底层链表。

跳表使用概率均衡技术而不是使用强制性均衡，因此对于插入和删除结点比传统上的平衡树算法更为简洁高效。因此跳表适合增删操作比较频繁，并且对查询性能要求比较高的场景。

原理简化示意图

leveln:
......:
level3: H                   ->                  T
level2: H         ->        4         ->        T        
level1: H    ->   2    ->   4   ->    6    ->   T
level0: H -> 1 -> 2 -> 3 -> 4 -> 5 -> 6 -> 7 -> T

- level0是基础链表，所有节点都会被添加进来
- level1~leveln是索引链表，图例中每层是上一层的1/2
- 如节点在第i层出现，则小于i的每一层都会出现该节点
- 实际使用中，level层高是通过随机函数来确定的，会规范一个最大值

[跳表原理及Go实现](https://segmentfault.com/a/1190000041645807)

### HASH冲突 链地址法