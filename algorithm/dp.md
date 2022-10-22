dp.md

# 动态规划

## 使用场景

满足两个条件

- 满足以下条件之一 
  - 求最值
  - 求是否可行
  - 求可行个数
- 满足不能排序&交换

## 解题思路

1. **状态 State**
   - 明确当前dp数组的定义，一般问题中有多少个会影响结果的变量就定义多少维的数组：dp[state1][state2][...]
2. 状态转移
   - 写出当前状态如何由子状态递推而来，若无法推导，重新思考dp数组的定义
3. base case
   - 初始化最开始的极限状态

## 常见题型

### 矩阵类型 遍历&递推

1. [最小路径和](https://leetcode.cn/problems/minimum-path-sum/)
```go
// [[1,3,1],[1,5,1],[4,2,1]]
func minPathSum(grid [][]int) int {
    // dp[i][j] 表示从(0, 0)出发到达(i, j)时的最小路径和
    // 状态转移: (i, j) 的位置只能由(i-1, j) (i, j-1)而来， 问题转换为分别到达(i-1, j) (i, j-1)的最小路径和, 注意i, j遇到边界的处理
    // dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
    // base case: dp[0][0] = grid[0][0]
    // return dp[m-1][n-1]

}
```

2. [三角形最小路径和](https://leetcode.cn/problems/triangle/)
```go
func minimumTotal(triangle [][]int) int {

    // 思路1
    // dp[i][j] 表示从(i ,j)出发， 到达最底层时的最小路径和
    // 状态转移：当前值+相邻下一层的最小值
    // dp[i][j] = min(dp[i+1][j], dp[i+1][j+1]) + triangle[i][j]
    // base case: 最后一层dp[i][j] = triangle[i][j]
    // return dp[0][0]



    // 思路2
    // dp[i][j] 表示从(0, 0)出发到达 (i, j)的最小路径和
    // 状态转移 当前值+上一层相邻的最小路径和
    // dp[i][j] = min(dp[i-1][j], dp[i-1][j-1]) + triangle[i][j]
    // base case dp[0][0] = triangle[0][0]
    // return max(dp[len(triangle)-1])
}
```

3. [不同路径数I](https://leetcode.cn/problems/unique-paths/)
```go
func uniquePaths(m int, n int) int {
    // dp[i][j] 表示从(0, 0)出发 到达(i, j)共有多少种路径
    // 状态转移 (i, j)的位置只能由(i-1, j) (i, j-1)而来， 问题转换为分别到达(i-1, j) (i, j-1)的路径数量相加, 注意i, j遇到边界的处理
    // dp[i][j] = dp[i-1][j]+dp[i][j-1]
    // base case dp[0][0] = 1
    // return dp[m-1][n-1]

}

```

4. [不同路径数II](https://leetcode.cn/problems/unique-paths-ii/)
```go
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
    // dp[i][j] 表示从(0, 0)出发 到达(i, j)共有多少种路径
    // 状态转移 (i, j)的位置只能由(i-1, j) (i, j-1)而来， 问题转换为分别到达(i-1, j) (i, j-1)的路径数量相加, 当grid[i][j]==1 时， dp[i][j]=0; 注意i, j遇到边界的处理
    // dp[i][j]= grid[i][j]==1? 0:dp[i-1][j]+dp[i][j-1]
    // base case dp[0][0] = 1 if grid[0][0] == 1 return 0
    // return dp[m-1][n-1]

}

```


### 跳跃&爬楼梯

1. [爬楼梯](https://leetcode.cn/problems/climbing-stairs/)
```go
func climbStairs(n int) int {
    // dp[i] 表示到达第i阶有多少种方法
    // 状态转移 dp[i] = dp[i-1] + dp[i-2]
    // base case 第0阶表示地面 dp[0] = 1 dp[1] = 1
    // return dp[n]
}
```

2. [跳跃游戏](https://leetcode.cn/problems/jump-game/)
```go
func canJump(nums []int) bool {

    // 思路1
    // dp[i] 表示从0出发 能否到达第i个下标
    // 状态转移 dp[i] 依赖于所有 <i 的点当中是否有可以到达i的下标，即 dp[j] && nums[j]>=i-j 
    // base case dp[0][0]=true
    // return dp[len(nums)-1]

    // 思路2
    // dp[i] 表示从出发， 能否到达终点
    // 状态转移 dp[i] 依赖于下标i所能到达的下一批点位中是否有能到达终点的下标， 即  for j<=nums[i]{ dp[i] = dp[i] || dp[i+j] }
    // base case dp[len(nums)-1] = true
    // return dp[0]
}
```

3. [跳跃游戏2](https://leetcode.cn/problems/jump-game-ii/) 
```go
func jump(nums []int) int {
    // dp[i] 表示从nums[i] 到最后一个位置的最少次数
    // 状态转移 dp[i] = min(dp[i+0], dp[i+1], ..., dp[i+nums[i]]) + 1
    // base case dp[n-1] = 0
}
```

**小结**
跳跃爬楼梯问题通常用一维dp表来存储状态。

定义基本为：从初始点到达当前点 或 从当前点到达终点

状态推导思路：从之前几个状态进行推导 mix max


### 子序列&子串问题

1. [连续子数组最大和](https://leetcode.cn/problems/maximum-subarray/)
```go
func maxSubArray(nums []int) int {
    // dp[i] 表示以nums[i]为结尾的最大子连续数组和
    // 状态转移 如果dp[i-1]>0 dp[i]=dp[i-1]+nums[i] 否则 dp[i] = nums[i]
    // base case dp[0]=nums[0]
    // return max(dp[:])
    
    dp := make([]int, len(nums))
    dp[0] = nums[0]
    max := dp[0]
    for i:=1; i<len(nums); i++ {
        dp[i] = nums[i]
        if dp[i-1] > 0 {
            dp[i] = dp[i-1] + nums[i]
        }
        if dp[i]>max {
            max = dp[i]
        }
    }
    return max
}
```

2. [分割回文串](https://leetcode.cn/problems/palindrome-partitioning/) 回溯算法

3. [分割回文串2](https://leetcode.cn/problems/palindrome-partitioning-ii/) 求最小分割次数
```go
func minCut(s string) int {
    // dp[i] 表示将s[:i]分割成回文子串的最小分割次数
    // 状态转移 对s[:i]进行遍历， 若s[j:i]是回文串 dp[i] = min(dp[i], dp[j] + 1)
    // base case dp[0] = -1 dp[1] = 0
    // return dp[len(s)]
}
```

4. [最长回文子串](https://leetcode.cn/problems/longest-palindromic-substring/)
```go
func longestPalindrome(s string) string {

    // 思路1 动规
    // dp[i][j] 表示 s[i:j]是否是回文串
    // 状态转移 dp[i][j] 依赖于 s[i] == s[j] && dp[i+1][j-1]
    // base case dp[0][0] = true dp[i][i] = true
    // return range dp && max(i-j)


    // 思路2 中心扩散法
    // 回文串存在奇数偶数之分， 分别以i,  (i, i+1)为中心寻找最长回文子串
    // max(res, max(res1, res2))
}

// 扩散寻找最长回文串
func findPalindrome(s string, i, j int) string{
    for i>=0 && j<len(s) {
        if s[i] != s[j] {
            break
        }
        i--
        j++
    }
    return s[i+1:j]
}
```

5. [最长括号子串](https://www.nowcoder.com/practice/45fd68024a4c4e97a8d6c45fc61dc6ad?tpId=295&tqId=715&ru=/exam/oj&qru=/ta/format-top101/question-ranking&sourceUrl=%2Fexam%2Foj)

6. [单词拆分](https://leetcode.cn/problems/word-break/)
```go
func wordBreak(s string, wordDict []string) bool {
    // dp[i] 表示s[:i]能否由wordDict拼接出
    // 状态转移 遍历s[:i] dp[i] = (s[:i] in wordDict) || (dp[j] && s[j:i] in wordDict)
    // base case dp[0] = false
    // return dp[len(s)]
}
```

7. [最长递增子序列](https://leetcode.cn/problems/longest-increasing-subsequence/)
```go
func lengthOfLIS(nums []int) int {
    // dp[i] 表示以nums[i]为结尾的最长子序列长度
    // 状态转移 遍历nums[:i] if nums[j]<nums[i] dp[i] = max(dp[i], dp[j]+1)
    // base case dp[0] = 0 dp[1] = 1
    // return max(dp[:])
}
```

**小结**
动规子序列/子串问题通常用一维dp表来存储状态。

定义基本为：一维数组dp[i], 表示将当前点作为结尾的结果，然后对s[i] 或者s[:i]做复合题意的判断


### 两序列比对问题


1. [最长公共子序列](https://leetcode.cn/problems/longest-common-subsequence/)
```go
func longestCommonSubsequence(text1 string, text2 string) int {
    // dp[i][j] 表示text1[:i] text2[:j] 的最长公共子序列长度
    // 状态转移 
    // text1[i] == text2[j] dp[i][j] = dp[i-1][j-1] + 1
    // text1[i] != text2[j] dp[i][j] = max(dp[i-1][j], dp[i][j-1])
    // base case dp[0][0] = 0 dp[i][0] = 0 dp[0][j] = 0
    // return dp[len(text1)][len(text2)]
}

```

2. [最长公共子串](https://www.nowcoder.com/practice/f33f5adc55f444baa0e0ca87ad8a6aac?tpId=295&tqId=991150&ru=/exam/oj&qru=/ta/format-top101/question-ranking&sourceUrl=%2Fexam%2Foj)
```go
func longestCommonSubstring(str1 string, str2 string) string {
    // dp[i][j] 表示以 str1[i] str2[j]结尾的最长公共子串
    // 状态转移 
    // str1[i] == str2[j] dp[i][j] = dp[i-1][j-1] + str1[i:i+1]
    // str1[i] != str2[j] dp[i][j] = ""
    // base case  dp[i][0] = "" dp[0][j] = "" dp[0][0] = str1[0]==str2[0]?str1[0:1]:""
    // return max(dp[i][j])
}
```

3. [编辑距离](https://leetcode.cn/problems/edit-distance/)
```go
func minDistance(word1 string, word2 string) int {
    
    // 思路1
    // dp[i][j] 表示将 word1[:i] 变成 word2[:j] 所需的最少操作
    // 状态转移 
    // 1. word1[i-1] == word2[j-1] 无需操作 dp[i][j] = dp[i-1][j-1]
    // 2. word1[i-1] != word2[j-1] 插入、删除、替换 取最小 dp[i][j] = min(dp[i][j-1], dp[i-1][j], dp[i-1][j-1]) + 1
    // base case dp[0][0] = 0  dp[i][0] = i dp[0][j] = j 
    // return dp[len(word1)][len(word2)]

    // 思路2
    // max(len(str1), len(str2)) - LCS
    // LCS算法参考最长公共子序列
}
```

4. [正则表达式](https://leetcode.cn/problems/regular-expression-matching/)
```

```

5. [字符串匹配 KMP算法](https://leetcode.cn/problems/find-the-index-of-the-first-occurrence-in-a-string/)

**小结**
两个序列的动规问题，需定义一个二维dp表， 所代表的意义一般是 str1[:i] str2[:j] 的最值
状态转移一般 dp[i][j] = min/max(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]) 此类型的方程
需注意dp的长度通常会Len(str1)+1


### 零钱兑换

1. [零钱兑换](https://leetcode.cn/problems/coin-change/)
```go
func coinChange(coins []int, amount int) int {
    // dp[i] 表示凑出 i金额最少的硬币数
    // 状态转移 针对金额i 循环coins  dp[i] = min(dp[i], dp[i-coins[j]]+1)
    // base case dp[0] = 0 dp[i] = amount+1
    // return dp[amount]==amount+1 ? -1:dp[amount]
}
```

2. [零钱兑换II](https://leetcode.cn/problems/coin-change-2/)
```go
func change(amount int, coins []int) int {
    // dp[i] 表示 凑出 i 金额的组合数
    // 状态转移 为凑出i金额 可先凑出 ’i - 某一种面额的硬币’ 的金额；类比爬楼梯，要到达i  先到达 i-step
    // 循环coins dp[i] += dp[i-coins[j]]
    // base case dp[0][0] = 1
    // return dp[amount]
}
```

**小结**

- 零钱兑换问题需定义一维dp数组dp[i]，表示凑出i金额时的所求解
- 状态转移：为凑出i金额 可先凑出 ’i - 某一种面额的硬币’ 的金额，dp[i] = min(dp[i], dp[i-coins[j]]+1)
- base case dp[0] = 0 dp[i] = amount+1
- return dp[amount]

### 背包问题

1. [背包问题](https://www.lintcode.com/problem/92/)
```go
func BackPack(m int, a []int) int {
    // dp[i]


}
```

2. [背包问题II](https://www.lintcode.com/problem/125/)
```

```

3. [分割等和子集](https://leetcode.cn/problems/partition-equal-subset-sum

**小结**

- 背包问题需定义二维数组dp[i][w], 表示对于前i个物品，且前背包容量为w时 所求问题的结果——一般是最值
- 状态转移：对于第i个物品，有两种选择——装入、不装入，对应的状态分别是dp[i-1][w]  value[i] + dp[i-1][w-wt[i]]
- base case dp[0][...] = dp[...][0] = 0
- return dp[n][w]

### 股票问题

1. [股票卖卖的最佳时机](https://leetcode.cn/problems/best-time-to-buy-and-sell-stock/)
```go
func maxProfit(prices []int) int {
    // dp[i][j] 表示第天i天 分别在持仓、未持仓状态下的收益 j = {0, 1}
    // 状态转移 
    // 第i天未持仓: 前一天也未持仓、或前一天持仓当天卖出 dp[i][0] = max(dp[i-1][0], dp[i-1][1] + prices[i])
    // 第i天持仓： 前一天也持仓、或前一天未持仓今天买入 dp[i][1] = max(dp[i-1][1], 0-prices[i])  ps: 因为只买卖一次， 所以此处是0-prices[i]
    // base case dp[0][0] = 0 dp[0][1] = -prices[0]
    // return dp[len(prices)-1][0]

    dp:=make([][]int, len(prices))
    dp[0] = make([]int, 2)
    dp[0][0] = 0
    dp[0][1] = -prices[0]

    var max func(a, b int) int
    max = func(a, b int) int {
        if a>b {
            return a
        }
        return b
    }

    for i:=1; i<len(prices); i++ {
        dp[i] = make([]int, 2)
        dp[i][0] = max(dp[i-1][0], dp[i-1][1] + prices[i])
        dp[i][1] = max(dp[i-1][1], 0-prices[i])
    }
    return dp[len(prices)-1][0]


    // 思路2
    // 因为只进行一次买卖， 只需找到相差最大的波峰与波谷即可
    // 用一个变量记录 prices[:i]之前的最小价格 minPrice, profit = prices[i] - minPrice
    // 循环prices， return max(profit)
}
```

2. [股票卖卖的最佳时机2](https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-ii/)
```go
func maxProfit(prices []int) int {
    // dp[i][j] 表示第天i天 分别在持仓、未持仓状态下的收益 j = {0, 1}
    // 状态转移 
    // 第i天未持仓: 前一天也未持仓、或前一天持仓当天卖出 dp[i][0] = max(dp[i-1][0], dp[i-1][1] + prices[i])
    // 第i天持仓： 前一天也持仓、或前一天未持仓今天买入 dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i]) 
    // base case dp[0][0] = 0 dp[0][1] = -prices[0]
    // return dp[len(prices)-1][0]

    dp:=make([][]int, len(prices))
    dp[0] = make([]int, 2)
    dp[0][0] = 0
    dp[0][1] = -prices[0]

    var max func(a, b int) int
    max = func(a, b int) int {
        if a>b {
            return a
        }
        return b
    }

    for i:=1; i<len(prices); i++ {
        dp[i] = make([]int, 2)
        dp[i][0] = max(dp[i-1][0], dp[i-1][1] + prices[i])
        dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
    }
    return dp[len(prices)-1][0]


    // 贪心算法
    // 因为不限买卖次数， 即每一天的价格差都可获利
    // if prices[i+1] > prices[i] { profit += prices[i+1] - prices[i] }
    // return profit
}
```

3. [股票卖卖的最佳时机3](https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-iii/) 交易次数限制 2次
```go
func maxProfit(prices []int) int {
    // dp[i][k][j] 表示第i天 最大k次交易限制 持仓/未持仓时的收益; K只在买入的时候-1, 初始值为最大买入次数
    // 状态转移
    // 1. 第i天 最大k次交易 未持仓：前一天也未持仓今天不操作(dp[i-1][k][0]), 或 前一天持仓今天卖出(dp[i-1][k][1]+price[i]) 
    // 2. 第i天 最大k次交易 持仓：前一天未持仓今天买入(dp[i-1][k-1][0]-prices[i]), 或  前一天持仓今天不操作(dp[i-1][k][1])
    // base case dp[0][k][0] = 0 dp[0][k][1] = -prices[0] dp[i][0][0] = 0 dp[i][0][1] = math.MinInt32
    // return dp[len(preices)-1][max_k][0]

    dp := make([][][]int, len(prices))
    K := 2
    var max func(a, b int) int
    max = func(a, b int) int {
        if a>b {
            return a
        }
        return b
    }

    for i:=0; i<len(preices); i++ {
        if dp[i] == nil {
            dp[i] = make([][]int, K+1)
        }
        for k:=K; k>=0; k-- {
            if dp[i][k] == nil {
                dp[i][k] = make([]int, 2)
            }
            if i==0 {
                dp[i][k][0] = 0
                dp[i][k][1] = -prices[i]
                continue
            }
            if k == 0 {
                dp[i][0][0] = 0
                dp[i][0][1] = math.MinInt32
            }
            dp[i][k][0] = max(dp[i-1][k][0], dp[i-1][k][1]+prices[i])
            dp[i][k][1] = max(dp[i-1][k-1][0]-prices[i], dp[i-1][k][1])
        }
    }

    return dp[len(prices)-1][K][0]
}
```

4. [股票卖卖的最佳时机4](https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-iv/) 交易次数限制 k次
```go
func maxProfit(K int, prices []int) int {
    // 同上 3
}
```

5. [股票买卖的最佳时机5](https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-with-cooldown/) 交易冷冻期
```go
func maxProfit(prices []int) int {
    cooldown := 2
    // dp[i][j] 表示在第i天分别持仓 不持仓状态下的最大收益
    // 状态转移 
    // dp[i][0] 第i天不持仓: i-1天也不持仓 或 i-1天持仓今天卖出； max(dp[i-1][0], dp[i-1][1]+prices[i])
    // dp[i][1] 第i天持仓：今天买入且冷冻期前不持仓 或 i-1天持仓； max(dp[i-cooldown][0]-prices[i], dp[i-1][1]) 
    // base case dp[0][0] = 0 dp[0][1] = -prices[0]
    // return dp[len(prices)-1][0]

    dp:=make([][]int, len(prices))
    var max func(a, b int) int
    max = func(a, b int) int {
        if a>b {
            return a
        }
        return b
    }

    for i:=0; i<len(prices); i++ {
        if dp[i] == nil {
            dp[i] = make([]int, 2)
        }
        if i == 0 {
            dp[i][0] = 0
            dp[i][1] = -prices[i]
            continue
        }
        if i-cooldown < 0 {
            // base case 2
            dp[i][0] = max(dp[i-1][0], dp[i-1][1] + prices[i]);
            // i - 2 小于 0 时根据状态转移方程推出对应 base case
            dp[i][1] = max(dp[i-1][1], -prices[i]);
            continue
        }
        dp[i][0] =  max(dp[i-1][0], dp[i-1][1]+prices[i])
        dp[i][1] = max(dp[i-cooldown][0]-prices[i], dp[i-1][1])
    }
    return dp[len(prices)-1][0]
}
```
 
6. [股票买卖的最佳时机6](https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-with-transaction-fee/) 手续费
```go
func maxProfit(prices []int, fee int) int {
    // 不限交易次数， 同2， 注意收益值减去fee
}
```

**小结**

股票类问题定义一个二维 或三维的dp数组(适用于有交易次数限制的题)， 例如 dp[i][j] i表示天数 j取值0,1分别代表是否持仓； 
然后根据当天的持仓状态写出状态转移方程
基本形如
- dp[i][0] 第i天不持仓: i-1天也不持仓 或 i-1天持仓今天卖出； max(dp[i-1][0], dp[i-1][1]+prices[i])
- dp[i][1] 第i天持仓：今天买入且前一天不持仓 或 i-1天持仓； max(dp[i-1][0]-prices[i], dp[i-1][1])

针对只限一次交易 或不限交易次数的题， 亦有其他取巧方法， 比如不限交易可用贪心算法(每天的价格差累计即可)

对于含冷冻期 或手续费的题状态转移方程稍作变更即可

若限制交易次数 则需定义三维数组 dp[i][k][j] 表示在第i天 在最大交易次数k的限制下 分别持仓不持仓的收益

状态转移
1. 第i天 最大k次交易 未持仓：前一天也未持仓今天不操作(dp[i-1][k][0]), 或 前一天持仓今天卖出(dp[i-1][k][1]+price[i]) 
2. 第i天 最大k次交易 持仓：前一天未持仓今天买入(dp[i-1][k-1][0]-prices[i]), 或  前一天持仓今天不操作(dp[i-1][k][1])
3. base case dp[0][k][0] = 0 dp[0][k][1] = -prices[0] dp[i][0][0] = 0 dp[i][0][1] = math.MinInt32


### 打家劫舍

1. [打家劫舍1](https://leetcode.cn/problems/house-robber/) 数组链表
2. [打家劫舍2](https://leetcode.cn/problems/house-robber-ii/) 环形链表
3. [打家劫舍3](https://leetcode.cn/problems/house-robber-iii/) 树形结构

**小结**

- 定义一维数组dp[i] 表示到第 i 家房屋，所能偷的最大金额
- 状态转移：对当前房屋i有两种选择——偷、不偷，最大收益分别对应 dp[i-2] + nums[i], dp[i-1]; max(偷， 不偷)
- base case dp[0] = 0 dp[1] = nums[0]
- 对于环形链表，分两种base case：偷第一家不偷最后一家(dp[1]=nums[0]), 不偷第一家(dp[1]=0)