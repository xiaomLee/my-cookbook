# 数组

## 常用技巧

- 排序
- 双指针
- 滑动窗口
- 二分搜索
- 前缀和
- 差分数组
- 花式遍历

## 常见题型

### 排序

**快速排序**

快排算法框架类似二叉树的前序遍历, partition函数选出数组的分割位置pos, 在该位置前都比nums[pos]小, 之后都大于等nums[pos]。

quickSort(nums []int, start, end int) {
    if start>=end {
        return
    }
    pos := partition(nums, start, end)
    quickSort(nums, start, pos-1)
    quickSort(nums, pos+1, end)
}

```go
func QuickSort(nums []int) {
    sort(nums, 0, len(nums)-1)
}

func sort(nums []int, start, end int){
    if start >= end {
        return
    }
    pos := partition(nums, start, end)
    sort(nums, start, pos-1)
    sort(nums, pos+1, end)
}

func partition(nums, start, end int) {
    // 选取基准值pivot, 此处以nums[end]为基准值进行比较
    pos := end
    pivot := nums[pos]
    // 双指针i, j 初始化为start; j开始遍历, 当nums[j]< pivot时, 交换i，j && i++
    // 遍历结束，交换i， pos； 返回i
    i, j := start, start
    for j<pos {
        if nums[j] < pivot {
            nums[i], nums[j] = nums[j], nums[i]
            i++
        }
        j++
    }
    // 遍历结束，交换i， pos； 返回i
    nums[i], nums[pos] = nums[pos], nums[i]
    return i
}
```

技巧，有时数组会部分有序或整体倒序，导致算法复杂度很高，此时可以先进行shuffle洗牌操作，重新打乱数组顺序。

[参考](https://github.com/labuladong/fucking-algorithm/blob/master/%E7%AE%97%E6%B3%95%E6%80%9D%E7%BB%B4%E7%B3%BB%E5%88%97/%E6%B4%97%E7%89%8C%E7%AE%97%E6%B3%95.md)
```go
// Knuth 洗牌算法
// golang 自带有 rand.Shuffle(n, func(i, j))
func shuffle(nums []int) {
    n := len(nums)
    for i:=0; i<n; i++ {
        j := rand.Intn(n-i) + i
        nums[i], num[j] = nums[j], nums[i]
    }
}
```

**归并排序**

类似二叉树的后续遍历

```go
func mergeSort(nums []int) []int {
    if len(nums) <= 1 {
        return
    }
    mid := len(nums)/2
    left := mergeSort(nums[:mid])
    right := mergeSort(nums[mid:])
    return merge(left, right)
}
func merge(nums1, nums2 []int) []int {
    res := make([]int, 0)
    i, j := 0, 0
    for i<len(nums1) && j<len(nums2) {
        if nums1[i] < nums2[j] {
            res = append(res, nums1[i])
            i++
        }else {
            res = append(res, nums2[j])
            j++
        }
    }
    if i<len(nums1) {
        res = append(res, nums1[i:]...)
    }
    if j<len(nums2) {
        res = append(res, nums2[j:]...)
    }
    return res
}
```

**堆排序**

排序步骤：
1. 将数组构建成大顶堆(小顶堆)——用数组表示的完全二叉树结构——nums[0]是根节点，堆中最大(最小)
2. 堆顶出队，与数组最后一位交换(将最大值放于数组最后) swap(nums, 0, len(nums)-1)
3. 剩余元素nums[:len(nums)-2]继续构建堆结构——将nums[0]下沉到合适的位置 sink(nums, end)
3. 循环2，至所有数据排序

构建堆：
1. 首先构建大顶堆, nums[0] 为树的根节点，构建完全二叉树，则存在如下下标关系 left = 2*i+1 right = 2*i+2
2. 从最后一个非叶子节点开始进行上浮操作，最后一个非叶子节点为 len(nums)/2-1
3. swim(nums, i) 以nums[i]为根， 分别与left right进行比较 nums[2*i+1] nums[2*i+2]
4. 循环至i==0

```go
func heapSort(nums []int) {
    //构建大顶堆
    for i:=len(nums)/2-1; i>=0; i-- {
        swim(nums, i)
    }
    // 交换 排序
    for i:=len(nums)-1; i>=0; i-- {
        // swap 0, i
        nums[0], nums[i] = nums[i], nums[0]
        // sink
        sink(nums, i)
    }
    return nums
}

func swim(nums []int, root int) {
    for {
        l := 2*root + 1
        r := 2*root +2
        idx := root
        if l<len(nums) && nums[l] > nums[idx] {
            idx = l
        }
        if r<len(nums) && nums[r] > nums[idx] {
            idx = r
        }
        // 不需要上浮
        if idx == root {
            break
        }
        // 交换
        nums[root], nums[idx] = nums[idx], nums[root]
        root = idx
    }
}

func sink(nums []int, end int) {
    root := 0
    for {
        l := 2*root + 1
        r := 2*root +2
        idx := root
        if l<end && nums[l] > nums[idx] {
            idx = l
        }
        if r<end && nums[r] > nums[idx] {
            idx = r
        }
        // 不需要下沉
        if idx == root {
            break
        }
        // 交换
        nums[root], nums[idx] = nums[idx], nums[root]
        root = idx
    }
}
```

### 双指针

[两数之和 II - 输入有序数组](https://leetcode.cn/problems/two-sum-ii-input-array-is-sorted/)
```go
// 对有重复数字的 采用hash表
func twoSum(numbers []int, target int) []int {
    // 双指针 
    i, j := 0, len(numbers)-1
    for i<j {
        sum := numbers[i] + numbers[j]
        if sum == target {
            return []int{i+1, j+1}
        }else if sum < target {
            i++
        }else {
            j--
        }
    }
    return nil
}
```

[删除有序数组中的重复项](https://leetcode.cn/problems/remove-duplicates-from-sorted-array/)
```go
func removeDuplicates(nums []int) int {
    // 双指针
    // i, j num[i]!=nums[j]时 将j移到i+1的位置， i++
    i, j := 0, 0
    for j<len(nums) {
        if nums[i]!=nums[j] {
            nums[i+1], nums[j] = nums[j], nums[i+1]
            i++
        }
        j++
    }
    // 返回长度 下标+1
    return i+1
}
```

[移除元素](https://leetcode.cn/problems/remove-element/)
```go
// 同上
```

[移动零](https://leetcode.cn/problems/move-zeroes/)
```go
// 同上
```

[反转字符串](https://leetcode.cn/problems/reverse-string/)

[最长回文子串](https://leetcode.cn/problems/longest-palindromic-substring/)
```go
// 中心扩散法
func longestPalindrome(s string) string {
    // 以i, i+1为中心向两边扩散
    n := len(s)
    ans := ""
    for i:=0; i<n; i++ {
        res1 := findPalindrome(s, i, i)
        res2 := findPalindrome(s, i, i+1)
        if len(res1) > len(ans) {
            ans = res1
        }
        if len(res2) > len(ans) {
            ans = res2
        }
    }
    return ans
}

// 以i, j为中心的最长回文串
func findPalindrome(s string, i, j int) string {
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

[删除排序链表中的重复元素](https://leetcode.cn/problems/remove-duplicates-from-sorted-list/)

[盛最多水的容器](https://leetcode.cn/problems/container-with-most-water/description/)

[接雨水](https://leetcode.cn/problems/trapping-rain-water/description/)
```go
// 算法思路
// 将数组看做一个木桶 arr[0] arr[len(arr)-1]看做木桶的两个边，不断收缩
// 木桶原理：水量以短边为准
func trap(height []int) int {
    left, right := 0, len(height)-1
    maxLeft, maxRight := height[0], height[right]
    res := 0
    for left < right {
        if height[left]>maxLeft {
            maxLeft = height[left]
        }
        if height[right]>maxRight {
            maxRight = height[right]
        }
        if maxLeft<maxRight {
            res += maxLeft-height[left]
            left++
        }else {
            res += maxRight-height[right]
            right--
        }
    }
    return res
}
```

### 滑动窗口

解题模板
```go
func slideWindow(s string) int {
    // 1. 初始化窗口 left, right := 0, 0 win := make(map[byte]int) 窗口内所有值始终只出现一次
    // 2. 更新窗口 右指针后移 
    // 3. 判断窗口 左指针收缩
    // 4. 统计结果
    left, right := 0, 0
    win := make(map[byte]int)

    for right < len(s) {
        // 更新窗口 右指针后移
        c := s[right]
        win[c] += 1
        right++

        // 收缩左指针 更新窗口
        for win[c] > 1 {
            win[s[left]]--
            left++
        }
        // t统计结果
    }

}
```

[3.无重复字符的最长子串](https://leetcode.cn/problems/longest-substring-without-repeating-characters/)
```go
func lengthOfLongestSubstring(s string) int {
    // 滑动窗口 + map
    // 1. 初始化窗口 left, right := 0, 0 win := make(map[byte]int) 窗口内所有值始终只出现一次
    // 2. 更新窗口 右指针后移 
    // 3. 判断窗口 左指针收缩
    // 4. 统计结果
    left, right := 0, 0
    win := make(map[byte]int)
    res := 0

    for right < len(s) {
        // 更新窗口 右指针后移
        c := s[right]
        win[c] += 1
        right++

        // 收缩左指针 更新窗口
        for win[c] > 1 {
            win[s[left]]--
            left++
        }
        // 统计结果
        if right - left > res {
            res = right-left
        }
    }
    return res 
}
```

[438.找到字符串中所有字母异位词](https://leetcode.cn/problems/find-all-anagrams-in-a-string/)
```go
func findAnagrams(s string, p string) []int {
    // 滑动窗口
    left, right := 0, 0
    win := make(map[byte]int)
    res := make([]int, 0)

    match := make(map[byte]int)
    for i:=0; i<len(p); i++ {
        match[p[i]] +=1
    }

    for right < len(s) {
        // 更新窗口 指针后移
        win[s[right]] += 1
        right++

        // 收缩窗口
        for right-left>len(p) {
            win[s[left]]--
            left++
        }
        // fmt.Printf("right:%d left:%d win:%v substr:%s \n", right, left, win, s[left:right])

        if right-left == len(p) && valid(win, match) {
            res = append(res, left)
        }
    }
    return res
}

func valid(win map[byte]int, match map[byte]int) bool {
    for key, val := range match {
        if num, ok := win[key]; !ok || num != val {
            return false
        }
    }
    return true
}
```

[字符串的排列](https://leetcode.cn/problems/permutation-in-string/)
```go
// 同上
```

[最小覆盖子串](https://leetcode.cn/problems/minimum-window-substring/)

[重复的DNA序列](https://leetcode.cn/problems/repeated-dna-sequences/)

[实现 strStr()](https://leetcode.cn/problems/implement-strstr/)


### 二分查找

给一个**有序数组**和目标值，找第一次/最后一次/任何一次出现的索引，如果没有出现返回-1

模板四点要素

- 1、初始化：start=0、end=len-1
- 2、循环退出条件：start + 1 < end
- 3、比较中点和目标值：A[mid] ==、 <、> target
- 4、判断最后两个元素是否符合：A[start]、A[end] ? target

```go
// 二分搜索最常用模板
func search(nums []int, target int) int {
    // 1、初始化start、end
    start := 0
    end := len(nums) - 1
    // 2、处理for循环
    for start+1 < end {
        mid := start + (end-start)/2
        // 3、比较a[mid]和target值
        if nums[mid] == target {
            end = mid
        } else if nums[mid] < target {
            start = mid
        } else if nums[mid] > target {
            end = mid
        }
    }
    // 4、最后剩下两个元素，手动判断
    if nums[start] == target {
        return start
    }
    if nums[end] == target {
        return end
    }
    return -1
}
```

[34.在排序数组中查找元素的第一个和最后一个位置](https://leetcode.cn/problems/find-first-and-last-position-of-element-in-sorted-array/)
```go
func searchRange(nums []int, target int) []int {
    // 二分查找
    // 1. 左边界：right不断左移收缩
    // 2. 右边界：left不断右移收缩
    if len(nums) == 0 {
        return []int{-1, -1}
    }

    res := []int{-1, -1}
    
    // 寻找左边界
    left, right := 0, len(nums)-1
    for left+1<right {
        mid := left + (right-left)/2
        if nums[mid] >= target {
            right = mid
        }else {
            left = mid
        }
    }
    // 判断最后两个元素nums[left] nums[right]
    if nums[right] == target {
        res[0] = right
    }
    if nums[left] == target {
        res[0] = left
    }
    if res[0] == -1 {
        return res
    }

    // 寻找右边界
    left, right = 0, len(nums)-1
    for left+1<right {
        mid := left + (right-left)/2
        if nums[mid] > target {
            right = mid
        }else {
            left = mid
        }
    }
    // 判断最后两个元素nums[left] nums[right]
    if nums[left] == target {
        res[1] = left
    }
    if nums[right] == target {
        res[1] = right
    }
    return res
}
```

[二分查找](https://leetcode.cn/problems/binary-search/)
```go
func search(nums []int, target int) int {
    if len(nums) == 0 {
        return -1
    }
    left, right := 0, len(nums)-1

    for left+1<right {
        mid := left + (right-left)/2
        if nums[mid] == target {
            return mid
        }else if nums[mid]>target {
            right = mid
        }else {
            left = mid
        }
    }
    if nums[left] == target {
        return left
    }else if nums[right] == target {
        return right
    }
    return -1
}
```

[search-insert-position](https://leetcode-cn.com/problems/search-insert-position/)
```go
func searchInsert(nums []int, target int) int {
    // 二分查找 左边界 第一个 >= target 的位置 
    if len(nums) == 0 {
        return 0
    }
    left, right := 0, len(nums)-1

    for left+1 < right {
        mid := left + (right - left)/2
        if nums[mid] >= target {
            right = mid
        }else {
            left = mid
        }
    }
    if nums[left] >= target {
        return left
    }else if nums[right] >= target {
        return right
    }else if nums[right]<target {
        return right +1
    }
    return 0
}
```

[search-a-2d-matrix](https://leetcode-cn.com/problems/search-a-2d-matrix/)
```go
func searchMatrix(matrix [][]int, target int) bool {
    // 思路 二分查找 row=mid/n col=mid%n

    m := len(matrix)
    n := len(matrix[0])

    left, right := 0, m*n-1

    for left+1<right {
        mid := left + (right-left)/2
        row := mid/n
        col := mid%n
        if matrix[row][col] == target {
            return true
        }else if matrix[row][col] > target {
            right = mid
        }else {
            left = mid
        }
    }

    if matrix[left/n][left%n] == target || matrix[right/n][right%n] == target {
        return true
    }
    return false
}
```

[first-bad-version](https://leetcode-cn.com/problems/first-bad-version/)
```go
func firstBadVersion(n int) int {
    // 二分查找 左边界
    left, right := 1, n
    for left +1 < right {
        mid := left + (right-left)/2
        if isBadVersion(mid) {
            right = mid
        }else {
            left = mid
        }
    }
    if isBadVersion(left) {
        return left
    }
    if isBadVersion(right) {
        return right
    }
    return -1
}
```

[find-minimum-in-rotated-sorted-array](https://leetcode-cn.com/problems/find-minimum-in-rotated-sorted-array/)

[find-minimum-in-rotated-sorted-array-ii](https://leetcode-cn.com/problems/find-minimum-in-rotated-sorted-array-ii/)

[search-in-rotated-sorted-array](https://leetcode-cn.com/problems/search-in-rotated-sorted-array/)

[search-in-rotated-sorted-array-ii](https://leetcode-cn.com/problems/search-in-rotated-sorted-array-ii/)
```go
func search(nums []int, target int) bool {
    // 二分搜索
    start, end := 0, len(nums) -1
    for start+1 < end {
        // 去除重复元素
        for start < end && nums[start] == nums[start+1] {
            start++
        }
        for start < end && nums[end] == nums[end-1] {
            end--
        }
        mid := start + (end-start)/2
        if nums[mid] == target {
            return true
        }

        // 判断在左上升区间 还是右上升区间
        if nums[start] < nums[mid] {
            // 左上升区间
            if nums[start] <= target && nums[mid] >= target {
                end = mid
            }else {
                start = mid
            }
        }else if nums[end] > nums[mid] {
            // 右上升区间
            if nums[end] >= target && nums[mid] <= target {
                start = mid
            }else {
                end = mid
            }
        }
    }
    // 判断最后剩余的 start end
    if nums[start] == target || nums[end] == target {
        return true
    }
    return false
}
```

[在排序数组中查找数字 I](https://leetcode.cn/problems/zai-pai-xu-shu-zu-zhong-cha-zhao-shu-zi-lcof/)

[在 D 天内送达包裹的能力](https://leetcode.cn/problems/capacity-to-ship-packages-within-d-days/)

[分割数组的最大值](https://leetcode.cn/problems/split-array-largest-sum/)

[爱吃香蕉的珂珂](https://leetcode.cn/problems/koko-eating-bananas/)

### 前缀和

[区域和检索 - 数组不可变](https://leetcode.cn/problems/range-sum-query-immutable/)
```go
type NumArray struct {
    data []int
    preSum []int
}

func Constructor(nums []int) NumArray {
    preSum := make([]int, len(nums)+1)
    for i := 0; i < len(nums); i++ {
        preSum[i+1] = preSum[i] + nums[i]
    }
    //fmt.Println(preSum)
    return NumArray{
        data: nums,
        preSum: preSum,
    }
}

func (this *NumArray) SumRange(left int, right int) int {
    return this.preSum[right+1] - this.preSum[left]
}
 ```

[二维区域和检索 - 矩阵不可变](https://leetcode.cn/problems/range-sum-query-2d-immutable/)
```go
type NumMatrix struct {
    preSum [][]int
}


func Constructor(matrix [][]int) NumMatrix {
    preSum := make([][]int, len(matrix)+1)
    preSum[0] = make([]int, len(matrix[0])+1)
    for i := 0; i<len(matrix); i++ {
        if preSum[i+1] == nil {
            preSum[i+1] = make([]int, len(matrix[0])+1)
        }
        for j:=0; j< len(matrix[0]); j++ {
            preSum[i+1][j+1] = preSum[i+1][j] + preSum[i][j+1] - preSum[i][j] + matrix[i][j]
        }
    }
    return NumMatrix{preSum: preSum}
}


func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
    preSum := this.preSum
    return preSum[row2+1][col2+1] - preSum[row2+1][col1] - preSum[row1][col2+1] + preSum[row1][col1]
}
 ```

[二维子矩阵的和](https://leetcode.cn/problems/O4NDxx/)

### 差分数组

[拼车](https://leetcode.cn/problems/car-pooling/)

[航班预订统计](https://leetcode.cn/problems/corporate-flight-bookings/)


### 数组遍历

[翻转字符串里的单词](https://leetcode.cn/problems/reverse-words-in-a-string/)
```go
// 多次翻转
// 首先整个字符串翻转
// 之后 每个单词再次翻转
```

[旋转图像](https://leetcode.cn/problems/rotate-image/)
```go
func rotate(matrix [][]int)  {
    // 1. 对角线交换，完成行列转换
    // 2. 每行翻转，完成旋转

    // 1 2 3
    // 4 5 6
    // 7 8 9

    // 1 4 7
    // 2 5 8
    // 3 6 9

    // 7 4 1
    // 8 5 2
    // 9 6 3

    for i:=0; i<len(matrix); i++ {
        for j:=i; j<len(matrix); j++ {
            matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
        }
    }
    for i :=0; i<len(matrix); i++ {
       reverse(matrix[i])
    }
}

func reverse(nums []int) {
    i, j :=0, len(nums)-1
    for i<j {
        nums[i], nums[j] = nums[j], nums[i]
        i++
        j--
    }
}
```

[螺旋矩阵](https://leetcode.cn/problems/spiral-matrix/)


[螺旋矩阵 II](https://leetcode.cn/problems/spiral-matrix-ii/)

[顺时针打印矩阵](https://leetcode.cn/problems/shun-shi-zhen-da-yin-ju-zhen-lcof/)


