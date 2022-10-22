# 回溯算法

回溯算法常用于排列组合问题， 穷举所有解，都是暴力递归，算法框架形如dfs深度搜索遍历。
通常用需要定义一个visited数组用来存储已遍历过的值
track数组用于存储当前遍历栈， res数组用于存储满足要求的解；当track数组达到所需要求后， 将track添加到可行解列表中。

```go
// 排列问题 i 从 0 开始, 需要一个记录已经做过选择的数组
func backtrack(nums []int, visited []bool, track []int, res *[][]int) {
    if track 满足要求 {
        ans := make([]int, len(track))
        copy(ans, track)
        *res = append(*res, ans)
        return
    }

    for i:=0; i<len(nums); i++ {
        if visited[i] || 其他过滤条件 {
            continue
        }

        // 做选择 将nums[i] 添加到track栈
        track = append(track, nums[i])
        visited[i] = true

        // 递归下次选择
        backtrack(nums, visited, track, res)

        // 取消选择 
        track = track[:len(track)-1]
        visited = false
    }
}

// 组合问题 i 从给定的初始pos开始
func backtrack(nums []int, pos int, track []int, res *[][]int) {
    if track 满足要求 {
        ans := make([]int, len(track))
        copy(ans, track)
        *res = append(*res, ans)
        return
    }

    for i:=pos; i<len(nums); i++ {
        if 过滤选择 {
            continue
        }

        // 做选择 将nums[i] 添加到track栈
        track = append(track, nums[i])

        // 递归下次选择
        backtrack(nums, i+1, track, res)

        // 取消选择 
        track = track[:len(track)-1]
    }
}
```

## 常见题型

### 排列

1. [全排列](https://leetcode.cn/problems/permutations/) 不含重复数字
```go
func backtrack(nums []int, visited []bool, track []int, res *[][]int) {
    if len(track) == len(nums) {
        ans := make([]int, len(track))
        copy(ans, track)
        *res = append(*res, ans)
        return
    }
    for i:=0; i<len(nums); i++ {
        if visited[i]  {
            continue
        }
        track = append(track, nums[i])
        visited[i] = true
        backtrack(nums, visited, track, res)
        track = track[:len(track)-1]
        visited[i] = false
    }
}
```

2. [全排列II](https://leetcode.cn/problems/permutations-ii/) 含重复数字
```go
// 含重复数字， 先排序， 固定所有数字的相对顺序
func permuteUnique(nums []int) [][]int {
    sort.Slice(nums, func(i, j int) bool{return nums[i]<nums[j]})
    
    var res = make([][]int, 0)
    var backtrack func
    backtrack = func(nums []int, visited []bool, track []int) {
        if len(track) == len(nums) {
            ans := make([]int, track)
            copy(ans, track)
            res = append(res, ans)
            return
        }
        for i:=0; i<len(nums); i++ {
            if visited[i]  {
                continue
            }
            // 跳过重复数字; 将重复的数字当做一个整体
            if i>0 && nums[i-1] == nums[i] && !visited[i-1] {
                continue
            }
            track = append(track, nums[i])
            visited[i] = true
            backtrack(nums, visited, track, res)
            track = track[:len(track)-1]
            visited[i] = false
        }
    }

    var (
        visited = make([]bool, len(nums))
        track = make([]int, 0)
    )
    backtrack(nums, visited, track)
    return res
}
```

### 组合

1. [组合](https://leetcode.cn/problems/combinations/)
```go
func backtrack(nums []int, k int, pos int, track []int, res *[][]int) {
    if pos == len(nums)-k+1 {
        return
    }
    if len(track) == k {
        ans := make([]int, len(track))
        copy(ans, track)
        *res = append(*res, ans)
        return
    }
    if len(track) > k {
        return
    }

    for i:=pos; i<len(nums); i++ {
        track = append(track, nums[i])
        backtrack(nums, k, i+1, track, res)
        track = track[:len(track)-1]
    }
}
```

2. [组合总和](https://leetcode.cn/problems/combination-sum/) 可重复选
```go
func backtrack(nums []int, target int, pos int, sum int, track []int, res *[][]int) {
    if sum == target {
        ans := make([]int, len(track))
        copy(ans, track)
        *res = append(*res, ans)
        return
    }

    // nums 是正整数
    if sum > target {
        return 
    }

    for i:=pos; i<len(nums); i++ {
        sum += nums[i]
        track = append(track, nums[i])
        // 可重复选 pos=i
        backtrack(nums, target, i, sum, track, res)
        sum-=nums[i]
        track = track[:len(track)-1]
    }
}
```

3. [组合总和II](https://leetcode.cn/problems/combination-sum-ii/) 重复只可用一次
```go
// 先排序 sort.Slice(nums, func(i, j)bool {nums[i]<nums[j]})
func backtrack(nums []int, target int, pos int, sum int, track []int, res *[][]int) {
    if sum == target {
        ans := make([]int, len(track))
        copy(ans, track)
        *res = append(*res, ans)
        return
    }

    // nums 是正整数
    if sum > target {
        return 
    }

    for i:=pos; i<len(nums); i++ {
        if i>pos && nums[i] == nums[i-1] {
            continue
        }
        sum += nums[i]
        track = append(track, nums[i])
        // 可重复选 pos=i
        backtrack(nums, target, i, sum, track, res)
        sum-=nums[i]
        track = track[:len(track)-1]
    }
}
```

4. [组合总和III](https://leetcode.cn/problems/combination-sum-iii/)
```go
func combinationSum3(k int, n int) [][]int {
    var (
        track = make([]int, 0)
        res = make([][]int, 0)
        sum = 0
    )
    backtrack(1, n, k, sum, track, &res)
    return res
}

func backtrack(i int, target int, k int, sum int, track []int, res *[][]int) {
    if sum == target && len(track) == k {
        ans := make([]int, k)
        copy(ans, track)
        *res = append(*res, ans)
        return
    }
    if sum>target && len(track)>k {
        return
    }

    for ; i<=9; i++ {
        sum += i
        track = append(track, i)
        backtrack(i+1, target, k, sum, track, res)
        track = track[:len(track)-1]
        sum -= i
    }
}
```

### 子集

1. [子集](https://leetcode.cn/problems/subsets/)
```go
func backtrack(nums []int, pos int, track []int, res *[][]int) {
    ans := make([]int, len(track))
    copy(ans, track)
    *res = append(*res, ans)
    
    for i:=pos; i<len(nums); i++ {
        track = append(track, nums[i])
        backtrack(nums, i+1, track, res)
        track = track[:len(track)-1]
    }
}
```

2. [子集II](https://leetcode.cn/problems/subsets-ii/) 重复数字
```go
// 先排序
func backtrack(nums []int, pos int, track []int, res *[][]int) {
    ans := make([]int, len(track))
    copy(ans, track)
    *res = append(*res, ans)
    
    for i:=pos; i<len(nums); i++ {
        if i>pos && nums[i] == nums[i-1] {
            continue
        }
        track = append(track, nums[i])
        backtrack(nums, i+1, track, res)
        track = track[:len(track)-1]
    }
}
```

**小结**

- 排列问题backtrack里层循环是需从0开始, 且需传入一个visited数组；针对有重复数字问题，先排序，然后把有序数字当做一个整体看待，即对
i>0 && nums[i] == nums[i-1] && && !visited[i-1] 需做跳过的判断
- 组合/子集问题backtrack里层循环初始值i+1；针对有重复数字问题，先排序，对 i>pos && nums[i] == nums[i-1] 需做跳过的判断

### 岛屿

1. [岛屿数量](https://leetcode.cn/problems/number-of-islands/)
```go
func numIslands(grid [][]byte) int {
    res := 0
    for i:=0; i<len(grid); i++ {
        for j:=0; j<len(grid[i]); j++ {
            if grid[i][j] == '1' {
                res++
                // 淹没当前岛屿
                dfs(grid, i, j)
            }
        }
    }
    return res
}

func dfs(grid [][]byte, i, j int) {
    // 到达边界
    if i<0 || i>=len(grid) || j>=len(grid[i]) || j<0 {
        return
    }
    // 已经是水
    if grid[i][j] == '0' {
        return
    }

    // 淹没
    grid[i][j] = '0'
    // 淹没相邻陆地
    dfs(grid, i+1, j)
    dfs(grid, i-1, j) 
    dfs(grid, i, j+1) 
    dfs(grid, i, j-1)
}

```

2. [统计子岛屿](https://leetcode.cn/problems/count-sub-islands/)
```go
func countSubIslands(grid1 [][]int, grid2 [][]int) int {
    // 思路 将grid2中的岛屿不存在grid1中的淹掉， 然后统计grid2岛屿的数量
    
    for i:=0; i<len(grid2); i++ {
        for j:=0; j<len(grid2[i]); j++ {
            if grid2[i][j] == 1 && grid1[i][j] == 0 {
                // 淹没当前岛屿
                dfs(grid2, i, j)
            }
        }
    }

    res := 0
    for i:=0; i<len(grid2); i++ {
        for j:=0; j<len(grid2[i]); j++ {
            if grid2[i][j] == 1 {
                res++
                dfs(grid2, i, j)
            }
        }
    }
    return res
}
```

3. [统计封闭岛屿的数目](https://leetcode.cn/problems/number-of-closed-islands/)
```go
func closedIsland(grid [][]int) int {
    // 思路 把靠边的岛屿先淹没， 然后统计岛屿数量
    m := len(grid)
    n := len(grid[0])

    for i:=0; i<n; i++ {
        dfs(grid, 0, i)
        dfs(grid, m-1, i)
    }
    for i:=0; i<m; i++ {
        dfs(grid, i, 0)
        dfs(grid, i, n-1)
    }

    res := 0
    for i:=0; i<m; i++ {
        for j:=0; j<n; j++ {
            if grid[i][j] == 0 {
                res++
                dfs(grid, i, j)
            }
        }
    }
    return res
}
```

4. [飞地的数量](https://leetcode.cn/problems/number-of-enclaves/)
```go
func numEnclaves(grid [][]int) int {
    // 思路 对四周边界的岛屿进行淹没， 然后统计1的数量
    m := len(grid)
    n := len(grid[0])

    for i:=0; i<n; i++ {
        dfs(grid, 0, i)
        dfs(grid, m-1, i)
    }
    for i:=0; i<m; i++ {
        dfs(grid, i, 0)
        dfs(grid, i, n-1)
    }

    res := 0
    for i:=0; i<m; i++ {
        for j:=0; j<n; j++ {
            if grid[i][j] == 1 {
                res++
            }
        }
    }
    return res
}
```

5. [岛屿的最大面积](https://leetcode.cn/problems/max-area-of-island/)

**小结**

- 岛屿问题基本思想：if grid[i][j]=='1' res++ 然后用dfs算法对相邻陆地进行淹没 
   - left：dfs(grid, i, j-1)
   - right: dfs(grid, i, j+1)
   - top: dfs(grid, i-1, j)
   - bootn: dfs(grid, i+1, j)
- 子岛屿问题先将grid2中的岛屿不存在Grid1中的进行淹没，然后统计岛屿数量
- 飞地问题/封闭岛屿问题， 先将四周边界的岛屿进行淹没，然后按题目需求进行求解


### 其他 

1. [N皇后](https://leetcode.cn/problems/n-queens/)

2. [数独](https://leetcode.cn/problems/sudoku-solver/)

3. [括号生成](https://leetcode.cn/problems/generate-parentheses/)


## 参考阅读

[回溯算法解题套路框架](https://labuladong.github.io/algo/1/8/)
[回溯算法秒杀所有排列/组合/子集问题](https://labuladong.github.io/algo/1/9/)
[一文秒杀所有岛屿题目](https://labuladong.github.io/algo/4/31/108/)
[回溯算法最佳实践：解数独](https://labuladong.github.io/algo/4/31/109/)


