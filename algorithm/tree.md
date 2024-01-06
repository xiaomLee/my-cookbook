# 树

<!-- vscode-markdown-toc -->
* 1. [二叉树(Binary Tree)](#BinaryTree)
	* 1.1. [二叉树遍历](#)
	* 1.2. [深度遍历(DFS)/层次遍历(BFS)](#DFSBFS)
	* 1.3. [常见题型](#-1)
* 2. [二叉搜索树BST(Binary Search Tree)](#BSTBinarySearchTree)
* 3. [其他常见树形结构及应用](#-1)
	* 3.1. [前缀树](#-1)
	* 3.2. [B-树(Balance Tree/Bayer Tree)](#B-BalanceTreeBayerTree)
	* 3.3. [B+树(B树Plus)](#BBPlus)
	* 3.4. [AVL树(Balance Binary Search Tree)](#AVLBalanceBinarySearchTree)
	* 3.5. [红黑树(Red Black Tree)](#RedBlackTree)
* 4. [参考](#-1)

<!-- vscode-markdown-toc-config
	numbering=true
	autoSave=true
	/vscode-markdown-toc-config -->
<!-- /vscode-markdown-toc -->

##  1. <a name='BinaryTree'></a>二叉树(Binary Tree)

**核心解题思维**
- 遍历：是否可以通过遍历一遍二叉树得到答案？如果可以，用一个 traverse 函数配合外部变量来实现；
- 递归：原问题是否可以通过子问题推导出？如果可以，定义一个递归函数，充分利用函数返回值，入参基本是root.left root.right

明确二叉树每个节点需要做什么， 需要在什么时候做(前中后序)

###  1.1. <a name=''></a>二叉树遍历

**前序遍历**：**先访问根节点**，再前序遍历左子树，再前序遍历右子树
**中序遍历**：先中序遍历左子树，**再访问根节点**，再中序遍历右子树
**后序遍历**：先后序遍历左子树，再后序遍历右子树，**再访问根节点**

```go
func traverse(root *TreeNode) {
    if root == nil {
        return
    }
    // 前序位置
    traverse(root.Left)
    // 中序位置
    traverse(root.Right)
    // 后序位置
}
```

###  1.2. <a name='DFSBFS'></a>深度遍历(DFS)/层次遍历(BFS)

**DFS-自顶向下**
```go
func dfs(root *TreeNode, result *[]interface{}) {
    if root == nil {
        return
    }
    // 前序位置 此时前中后序都可以，按编码习惯通常放于前序位置
    *result = append(*result, root)
    dfs(root.Left, result)
    dfs(root.Right, result)
}
```

**DFS-自底向上 又称分治法**
```go
func divide(root *TreeNode) []interface{} {
    if root == nil {
        return nil
    }

    // 分治
    left := divide(root.Left)
    right := divide(root.Right)

    // merge
   return merge(left, right)
}
```
快排、归并排序 都是借助分治的思想。 

快排是在前序位置计算出partitionIndex 然后分别递归排序左右两边 quickSort(nums, lo, partitionIndex-1) quickSort(nums, partitionIndex+1, hi) 

归并排序是计算mid=(start+end)/2, 然后分别对左右递归排序 left:=mergeSort(nums, start, mid)  right:=mergeSort(nums, mid+1, end), 最后merge两个有序数组 merge(left, right)；left, right已经有序

**BFS层序遍历**
```go
func levelTraverse(root *TreeNode) [][]interface {
    result := make([][]interface{}, 0)
    if root == nil {
        return nil
    }
    
    queue := make([]*TreeNode, 0)
    queue = append(queue, root)
    for len(queue) > 0 {
        l := len(queue)
        level := make([][]interface{}, 0)
        for i:=0; i<l; i++ {
            // 队头出队
            root := queue[0]
            queue = queue[1:]
            level = append(level, root)
            if root.Left != nil {
                queue = append(queue, root.Left)
            }
            if root.Right !=nil {
                queue = append(queue, root.Right)
            }
        }
        result = append(result, level)
    }

    return result
}
```

###  1.3. <a name='-1'></a>常见题型

**遍历、分治**

[二叉树的前序遍历](https://leetcode.cn/problems/binary-tree-preorder-traversal/)
[二叉树的中序遍历](https://leetcode.cn/problems/binary-tree-inorder-traversal/)
[二叉树的后序遍历](https://leetcode.cn/problems/binary-tree-postorder-traversal/)

[二叉树的所有路径](https://leetcode.cn/problems/binary-tree-paths/)
```go
func traverse(root *TreeNode, path string, res *[]string) {
    if root == nil {
        return
    }

    path += fmt.Sprintf("->%d", root.Val)
    if root.Left == nil && root.Right == nil {
        // 到达叶子节点
        *res = append(*res, path[2:])
        return
    }
    traverse(root.Left, path, res)
    traverse(root.Right, path, res) 
}
```

[二叉树展开为链表](https://leetcode.cn/problems/flatten-binary-tree-to-linked-list)
```go
func flatten(root *TreeNode)  {
    if root == nil {
        return
    }
    // 拉平左右子树
    flatten(root.Left)
    flatten(root.Right)

    left := root.Left
    right := root.Right

    root.Left = nil
    root.Right = left

    // 将链表指针移到末尾
    p:=root
    for p.Right != nil {
        p = p.Right
    }
    p.Right = right
}
```

[翻转二叉树](https://leetcode.cn/problems/invert-binary-tree/)
```go
func invertTree(root *TreeNode) *TreeNode {
    if root == nil {
        return nil
    }
    left := invertTree(root.Left)
    right := invertTree(root.Right)
    root.Right = left
    root.Left = right
    return root
}
```

[合并二叉树](https://leetcode.cn/problems/merge-two-binary-trees/)
```go
func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
    if root1 == nil {
        return root2
    }
    if root2 == nil {
        return root1
    }

    // 递归子问题：分别将左右子树进行合并
    left := mergeTrees(root1.Left, root2.Left)
    right := mergeTrees(root1.Right, root2.Right)

    // 依赖递归返回 后序处理
    root := &TreeNode{
        Val: root1.Val + root2.Val,
    }
    root.Left = left
    root.Right = right
    return root
}
```

[二叉树的最大深度](https://leetcode.cn/problems/maximum-depth-of-binary-tree/)
```go
func maxDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }
    left := maxDepth(root.Left)
    right := maxDepth(root.Right)
    
    if left>right {
        return left+1
    }
    return right+1
}
```

[平衡二叉树](https://leetcode.cn/problems/balanced-binary-tree/)
```go
func isBalanced(root *TreeNode) bool {
    return maxDepth(root)>=0
}

func maxDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }
    left := maxDepth(root.Left)
    right := maxDepth(root.Right)
    
    // 非平衡 结束递归
    if left == -1 || left-right>1 || right==-1 || right-left>1 {
        return -1
    }

    if left>right {
        return left+1
    }
    return right+1
}
```

[二叉树的直径](https://leetcode.cn/problems/diameter-of-binary-tree/)
```go
func diameterOfBinaryTree(root *TreeNode) int {
    res := 0
    maxDepth(root, &res)
    return res
}

func maxDepth(root *TreeNode, res *int) int {
    if root == nil {
        return 0
    }
    left := maxDepth(root.Left, res)
    right := maxDepth(root.Right, res)
    
    if *res<left+right {
        *res = left+right
    }
    if left>right {
        return left+1
    }
    return right+1
}
```

[完全二叉树的节点个数](https://leetcode.cn/problems/count-complete-tree-nodes/)
```go
func countNodes(root *TreeNode) int {
    // 思路 
    // 假设以当前节点为根节点，它是一个满二叉树，则 count=math.Pow(2, hight)
    // 否则 递归左子树右子树分别计算个数
    // 沿最左侧和最右侧分别计算高度, 如果相等则是满二叉树

    l,r := root, root
    hl, hr := 0, 0
    
    for l!=nil {
        l=l.Left
        hl++
    }
    for r!=nil {
        r=r.Right
        hr++
    }
    if hl == hr {
        // 满二叉树
        return int(math.Pow(2, float64(hl)))-1
    }
    return 1 + countNodes(root.Left) + countNodes(root.Right)
}
```

[二叉树中的最大路径和](https://leetcode.cn/problems/binary-tree-maximum-path-sum/)
```go
func maxPathSum(root *TreeNode) int {
    // 思路 计算以root节点为起点的最大序列和 与初始值比较 取较大
    // 遍历 分别计算以左右子树的最大序列和 与初始值比较 取较大

    // 定义 单边路径最大和函数 slidMax(root, res) int

    res := -1001
    slidMax(root, &res)
    return res
}

func slidMax(root *TreeNode, res *int) int {
    if root == nil {
        return 0
    }
    left := max(0, slidMax(root.Left, res))
    right := max(0, slidMax(root.Right, res))

    if *res < root.Val + left + right {
        *res = root.Val + left + right 
    }
    return root.Val + max(left, right)
}
```

[路径总和II](https://leetcode.cn/problems/path-sum-ii/)
```go
func pathSum(root *TreeNode, targetSum int) [][]int {
    var (
        sum = 0
        path =make([]int, 0)
        res = make([][]int, 0)
    )
    traverse(root, targetSum, sum, path, &res)
    return res
}

func traverse(root *TreeNode, targetSum, sum int, path []int, res *[][]int) {
    if root == nil {
        return
    }

    path = append(path, root.Val)
    sum += root.Val
    if root.Left == nil && root.Right == nil {
        if sum == targetSum {
            ans := make([]int, len(path))
            copy(ans, path)
            *res = append(*res, ans)
        }
        return
    }
    traverse(root.Left, targetSum, sum, path, res)
    traverse(root.Right, targetSum, sum, path, res)
}
```

[二叉树的最近公共祖先](https://leetcode.cn/problems/lowest-common-ancestor-of-a-binary-tree/)
```go
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
    // 判断当前节点是否等于p 或 q 是的话直接返回， 反之递归左右子树
    // 针对递归返回做后序处理
    // 1. p,q 分别在左右两边 left!=nil && right!=nil return root
    // 2. p,q 在左右某一边 或 都不在(left==nil && right==nil)  return left != nil ? left:right
    if root == nil {
        return root
    }
    if root == p || root == q {
        return root
    }
    left := lowestCommonAncestor(root, p, q)
    right := lowestCommonAncestor(root, p, q)
    
    if left!=nil && right!=nil {
        return root
    }
    if left != nil {
        return left
    }
    return right
}
```

**层次遍历**

[二叉树的层序遍历](https://leetcode-cn.com/problems/binary-tree-level-order-traversal/)
```go
func levelOrderBottom(root *TreeNode) [][]int {
    if root == nil {
        return nil
    }

    res := make([][]int, 0)
    queue := make([]*TreeNode, 0)
    queue = append(queue, root)

    for len(queue) > 0 {
        count := len(queue)
        level := make([]int, 0)
        for i:=0; i<count; i++ {
            root = queue[i]
            level = append(level, root.Val)
            if root.Left != nil {
                queue = append(queue, root.Left)
            }
            if root.Right != nil {
                queue = append(queue, root.Right)
            }
        }
        queue = queue[count:]
        if len(level) > 0 {
            res = append(res, level)
        }
    }
    // reverse(res)
    return res
}
```

[二叉树的层序遍历-ii](https://leetcode-cn.com/problems/binary-tree-level-order-traversal-ii/)
```go
// 遍历同上 遍历完成后翻转result
func reverse(list [][]int) {
    for i, j:=0, len(list)-1; i<j; i, j=i+1, j-1 {
        list[i], list[j] = list[j], list[i]
    }
}
```

[二叉树的z字形遍历](https://leetcode-cn.com/problems/binary-tree-zigzag-level-order-traversal/)
```go
// 层序遍历
// 将每层结果append到result时 做标志位flag的判断是否翻转 同时标志位置反
if len(level) > 0 {
    if flag {
        reverse(level)
    }
    flag = !flag
    res = append(res, level)
}
```

[二叉树的完全性检验](https://leetcode.cn/problems/check-completeness-of-a-binary-tree/)
```go
func isCompleteTree(root *TreeNode) bool {
    // 思路 bfs
    // 1. 对每个节点检验 root.Left == nil && root.Right!=nil return false
    // 2. 一旦出现叶子节点或不完整节点(root.Left == nil && root.Right==nil) || (root.Left!=nil && root.Right==nil)
    // 则 后续不能再出现完整节点
    if root == nil {
        return true
    }
    queue := make([]*TreeNode, 0)
    queue = append(queue, root)
    end := false

    for len(queue)>0 {
        l := len(queue)
        for i:=0; i<l; i++ {
            root = queue[0]
            queue = queue[1:]
            // check valid
            if root.Left == nil && root.Right != nil {
                return false
            }
            if end && (root.Left != nil || root.Right != nil) {
                return false
            }

            // add queue
            if root.Left != nil {
                queue = append(queue, root.Left)
            }
            if root.Right != nil {
                queue = append(queue, root.Right)
            }
            // 是否到达最后一层
            if (root.Left == nil && root.Right==nil) || (root.Left!=nil && root.Right==nil) {
                end = true
            }
        }
    }
    return true
}
```

**序列化、反序列化**

[二叉树的序列化与反序列化](https://leetcode.cn/problems/serialize-and-deserialize-binary-tree/)
```go
// 思路
// 只有前序遍历、后序遍历、层序遍历可以同时实现序列化与反序列化
// 原理
// 前序遍历的结果 第一个元素 即是root节点，之后根据左子树右子树顺序递归剩余字符串即可
// 后序遍历的结果 最后一个元素是root节点，之后根据右子树左子树顺序递归剩余字符串
// 层序遍历的结果 第一个节点是root节点， 之后紧跟的分别是left right; 下一层的每两个元素分别是上一层每一个元素的子节点
// 中序遍历因为无法确定root节点的位置，故无法实现序列化

// 解题技巧
// 反序列化时， 可先将data字符串转换成 []*TreeNode, 即将每个元素预先构造成树, # 转成 nil；然后根据不同序列化的方式 将树串联起来

// 有如下树结构
//       1
// 2           3
//     4   5       6

// 前序遍历序列化反序列化
// 序列化结果： 1, 2, #, 4, #, #, 3, 5, #, #, 6, #, #
func serialize(root *TreeNode) string {
    if root == nil {
        return "#"
    }
    left := this.serialize(root.Left)
    right := this.serialize(root.Right)
    data := fmt.Sprintf("%d,%s,%s", root.Val, left, right)
    return data
}
// 反序列化，用一个pos索引指针记录当前反序列化位置
func deserialize(data string, pos *int) *TreeNode {
    if *pos >= len(data) {
        return nil
    }
    // 获取当前节点， 即root节点
    var first string
    index := strings.Index(data[*pos:], ",")
    if index == -1 {
        first = data[*pos:]
    }else {
        first = data[*pos:*pos+index]
    }
    // 指针后移
    *pos = *pos + len(first) +1

    if first == "#" {
        // leaf node
        return nil
    }
    // covert val
    val, err:=strconv.ParseInt(first, 10, 64)
    if err!=nil {
        panic(err)
    }
    root := &TreeNode{Val:int(val)}
    
    // 获取当前节点左右子树
    root.Left = deserialize(data, pos) // pos 在递归中不断后移
    root.Right = deserialize(data, pos)
    return root
}

// 后续遍历 序列化反序列化
func serialize(root *TreeNode) string {
    if root == nil {
        return "#"
    }
    left := this.serialize(root.Left)

    right := this.serialize(root.Right)
    data := fmt.Sprintf("%s,%s,%d",left, right, root.Val)
    return data
}
// 反序列化 最后一个元素是root 之后递归右子树 左子树; pos指针不断前移
func deserialize(data string, pos *int) *TreeNode{}


// 层序遍历
// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
    if root == nil {
        return ""
    }
    res := ""
    queue := make([]*TreeNode, 0)
    queue = append(queue, root)
    for len(queue) > 0 {
        count := len(queue)
        for i:=0; i<count; i++ {
            node := queue[i]
            if node == nil {
                res += fmt.Sprintf(",%s", "#")
                continue
            }else if len(res) == 0 {
                res += fmt.Sprintf("%d", node.Val)
            }else {
                res += fmt.Sprintf(",%d", node.Val)
            }
            queue = append(queue, node.Left)
            queue = append(queue, node.Right)
        }
        queue = queue[count:]
    }
    return res
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {   
    // fmt.Println("data:", data) 
    if len(data) == 0 {
        return nil
    }

    // get the first node
    var getFirstNode func(data string, pos *int) *TreeNode 
    getFirstNode = func(data string, pos *int) *TreeNode {
        if *pos > len(data) {
            return nil
        }
        var first string
        index := strings.Index(data[*pos:], ",")
        if index == -1 {
            first = data[*pos:]
        }else {
            first = data[*pos:*pos+index]
        }
        *pos = *pos + len(first) +1

        if first == "#" {
            return nil
        }
        val, err:=strconv.ParseInt(first, 10, 64)
        if err !=nil {
            fmt.Printf("panic data:%s pos:%d index:%d first:%s err:%s\n", data, *pos, index, first, err)
            //panic(err)
            return nil
        }
        return &TreeNode{Val:int(val)}
    }

    var (
        queue = make([]*TreeNode, 0)
        pos int
    )
    root := getFirstNode(data, &pos)

    if root == nil {
        return nil
    }
    queue = append(queue, root)

    for pos<len(data) {
        node := queue[0]
        queue = queue[1:]
        
        // 获取左右节点
        node.Left = getFirstNode(data, &pos)
        if node.Left != nil {
            queue = append(queue, node.Left)
        }
        node.Right = getFirstNode(data, &pos)
        if node.Right != nil {
            queue = append(queue, node.Right)
        }
    }
    return root
}

```

[前序遍历构造二叉搜索树](https://leetcode.cn/problems/construct-binary-search-tree-from-preorder-traversal)
[从中序与后序遍历序列构造二叉树](https://leetcode.cn/problems/construct-binary-tree-from-inorder-and-postorder-traversal/)
[最大二叉树](https://leetcode.cn/problems/maximum-binary-tree/)
```go
// 思路 与快排类似：partition找到maxIndex 构建root, 之后递归左右两边 root.Left = construct(nums[:maxIndex-1]) root.Right = construct(maxIndex+1)
```
[最大二叉树II](https://leetcode.cn/problems/maximum-binary-tree-ii/)



##  2. <a name='BSTBinarySearchTree'></a>二叉搜索树BST(Binary Search Tree)

1. 对于 BST 的每一个节点 node，左子树节点的值都比 node 的值要小，右子树节点的值都比 node 的值大。
2. 对于 BST 的每一个节点 node，它的左侧子树和右侧子树都是 BST。

特性：中序遍历有序(升序) 解题基本都是中序遍历框架

[二叉搜索树中第K小的元素](https://leetcode.cn/problems/kth-smallest-element-in-a-bst/)
```go
func kthSmallest(root *TreeNode, k int) int {
    // 思路 中序遍历 当 rank==k 时  return root.Val
    var rank, res int
    traverse(root, k, &rank, &res)
    return res
}

func traverse(root *TreeNode, k int, rank, res *int) {
    if root==nil {
        return
    }
    traverse(root.Left, k, rank, res)
    *rank +=1
    if *rank == k {
        *res = root.Val
        return
    }
    traverse(root.Right, k, rank, res)
}
```

[把二叉搜索树转换为累加树](https://leetcode.cn/problems/convert-bst-to-greater-tree/)
```go
func convertBST(root *TreeNode) *TreeNode {
    // 思路 中续遍历 降序 计算累加和 并更新当前节点
    // 定义 traverse(root, sum) 以他为根的累加和 从右子树开始遍历
    // 将当前节点值变为 root.Val += sum
    var sum int
    traverse(root, &sum)
    return root
}

func traverse(root *TreeNode, sum *int) {
    if root == nil {
        return
    }
    
    traverse(root.Right, sum)
    *sum += root.Val
    root.Val = *sum
    traverse(root.Left, sum)
}
```

[删除二叉搜索树中的节点](https://leetcode.cn/problems/delete-node-in-a-bst/)
```go
func deleteNode(root *TreeNode, key int) *TreeNode {
    // 思路 比较 root.Val vs key 判断递归左 or 右子树
    // 若 root.Val == key if root.Right !=nil {将root.Left 挂到root.Right 的最左叶子结点} else {root = root.Left}
    // root.Val < key 递归 右子树
    // root.Val > key 递归左子树

    if root == nil {
        return nil
    }

    if root.Val == key {
        if root.Right != nil {
            p:=root.Right
            for p.Left != nil {
                p = p.Left
            }
            p.Left = root.Left
            root = root.Right
        }else {
            root = root.Left
        }
        return root
    }
    if root.Val > key {
        root.Left = deleteNode(root.Left, key)
    }else{
        root.Right = deleteNode(root.Right, key)
    }
    return root
}
```

[二叉搜索树中的搜索](https://leetcode.cn/problems/search-in-a-binary-search-tree/)

[二叉搜索树中的插入操作](https://leetcode.cn/problems/insert-into-a-binary-search-tree/)
```go
func insertIntoBST(root *TreeNode, val int) *TreeNode {
    // 思路 不断向左右递归 找到最后一个叶子节点，进行插入
    // base case root == nil return &TreeNode{Val:val}
    
    if root == nil {
        return &TreeNode{Val:val}
    }
    if root.Val < val {
        root.Right = insertIntoBST(root.Right, val)
    }else {
        root.Left = insertIntoBST(root.Left, val)
    }
    return root
}
```

[验证二叉搜索树](https://leetcode.cn/problems/validate-binary-search-tree/)
```go
func isValidBST(root *TreeNode) bool {
    // 思路 中序遍历有序 pre变量记录遍历的上一个值
    pre := math.MinInt64
    valid := true
    traverse(root, &pre, &valid)
    return valid
}
```

[不同的二叉搜索树](https://leetcode.cn/problems/unique-binary-search-trees/)
```go
func numTrees(n int) int {
    // 思路 递归函数count(lo, hi int) int 计算 lo~hi直接有多少种
    // base case lo >= hi return 1
    // 添加备忘录
    memo := make([][]int, n+1)
    return count(1, n, memo)
}
```

[不同的二叉搜索树II](https://leetcode.cn/problems/unique-binary-search-trees-ii/)
```go
func generateTrees(n int) []*TreeNode {
    // 思路 定义递归函数
    // generate(lo, hi int) []*TreeNode
    // 循环lo~hi 固定root， 然后分别计算lefts, rights
    // 排列组合lefts, rights
    // 优化 备忘录
    return generate(1, n)
}
```


##  3. <a name='-1'></a>其他常见树形结构及应用

###  3.1. <a name='-1'></a>前缀树

Trie 树又叫字典树、前缀树、单词查找树，是一种二叉树衍生出来的高级数据结构(多叉树)，主要应用场景是处理字符串前缀相关的操作， 例如自动补全 拼写检查等。

基本结构及方法
```go
type Trie struct {
    // 数组大小26表示支持的key字符串为a~z 亦可根据实际情况自定义
    // 数组下标对应ASCII码 当前定义下标转换 index = char - 'a'
    children [26]*Trie  
    // 表示从root->当前节点的路径是否表示一个完整的key
    isEnd bool
    // 在 isEnd为TRUE时，key 对应的value值 [option]
    val interface{}
}

func Constructor() Trie {
    return Trie{}
}
// 插入字符串
func (this *Trie) Insert(word string)  {
    cur := this
    for _, ch := range word {
        index := ch - 'a'
        if cur.children[index] == nil {
            // 不存在当前字符，新增插入
            cur.children[index] = &Trie{}
        }
        cur = cur.children[index]
    }
    cur.isEnd = true
}
// 前缀匹配
func (this *Trie) SearchWithPrefix(word string) *Trie {
    cur := this
    for _, ch := range word {
        index := ch - 'a'
        if cur.children[index] == nil {
            // 不存在当前字符
            return nil
        }
        cur = cur.children[index]
    }
    return cur
} 
// 查找
func (this *Trie) Search(word string) bool {
    node := this.SearchWithPrefix(word)
    return node != nil && node.isEnd
}
// 前缀查找
func (this *Trie) StartsWith(prefix string) bool {
    return node != nil
}
```

[参考](https://mp.weixin.qq.com/s/hGrTUmM1zusPZZ0nA9aaNw)

[实现前缀树(Trie)](https://leetcode.cn/problems/implement-trie-prefix-tree/)

[添加与搜索单词-数据结构设计](https://leetcode.cn/problems/design-add-and-search-words-data-structure/)

[单词替换](https://leetcode.cn/problems/replace-words/)

[键值映射](https://leetcode.cn/problems/map-sum-pairs/)


###  3.2. <a name='B-BalanceTreeBayerTree'></a>B-树(Balance Tree/Bayer Tree)

[B树图文详解](https://segmentfault.com/a/1190000038749020)


###  3.3. <a name='BBPlus'></a>B+树(B树Plus)

###  3.4. <a name='AVLBalanceBinarySearchTree'></a>AVL树(Balance Binary Search Tree)

[AVL树](https://baike.baidu.com/item/AVL%E6%A0%91/10986648)

AVL树得名于它的发明者G. M. Adelson-Velsky和E. M. Landis，他们在1962年的论文《An algorithm for the organization of information》中发表了它。

###  3.5. <a name='RedBlackTree'></a>红黑树(Red Black Tree)

[红黑树](https://baike.baidu.com/item/%E7%BA%A2%E9%BB%91%E6%A0%91/2413209)


##  4. <a name='-1'></a>参考

[东哥带你刷二叉树（纲领篇）](https://labuladong.github.io/algo/2/21/36/)