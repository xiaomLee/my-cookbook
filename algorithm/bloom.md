# 布隆过滤器

布隆过滤器通常用于判断一个元素是否存在于集合中，熟悉的应用场景有：解决缓存穿透、爬虫url去重。
相比传统的Hash-Table模式有占用内存少的优势，但对于元素是否存在则有一定的误判率，即：可以 100% 判定元素一定不存在，无法 100% 判断元素一定存在。
原因详见下述原理部分。

## 原理

布隆过滤器首先初始化一个长度为 m 的比特向量，然后针对元素进行 k 次哈希，每次哈希的结果都对应到某个比特位，然后将该位置为 1 (若已然是 1 则跳过)，总共会产生 k 个比特位。
查询时使用同样的哈希顺序查找对应的比特位，判断是否为 0，若有一个为 0， 则元素必定不存在。
通过上述描述可知：某个比特位可能会被其他元素共用并置为 1，故无法准确判断一个元素一定存在集合中。

[原理参考](https://sunyunqiang.com/blog/bloom_filter/)

## 定量分析

[原理参考](https://sunyunqiang.com/blog/bloom_filter/)

## 工程实践

https://github.com/riobard/go-bloom

```go
// Package bloom implements Bloom Filter using double hashing
package bloom

import (
	"math"
    "hash/fnv"
)

// Filter is a generic Bloom Filter
type Filter interface {
	Add([]byte)       // add an entry to the filter
	Test([]byte) bool // test if an entry is in the filter
	Size() int        // size of the filter in bytes
	Reset()           // reset the filter to initial state
}

// Classic Bloom Filter
type classicFilter struct {
	b []byte
	k int
	h func([]byte) (uint64, uint64)
}

// New creates a classic Bloom Filter that is optimal for n entries and false positive rate of p.
// h is a double hash that takes an entry and returns two different hashes.
func New(n int, p float64, h func([]byte) (uint64, uint64)) Filter {
	k := -math.Log(p) * math.Log2E   // number of hashes
	m := float64(n) * k * math.Log2E // number of bits
	return &classicFilter{b: make([]byte, int(m/8)), k: int(k), h: h}
}

func (f *classicFilter) getOffset(x, y uint64, i int) uint64 {
	return (x + uint64(i)*y) % (8 * uint64(len(f.b)))
}

// Add 
// 函数向布隆过滤器中添加数据, 只使用了一个哈希函数, 
// 该哈希函数返回两个不同的哈希值, 其通过将第二个哈希值按递增倍数进行偏移 (我们 将两个哈希值记为 v1 和 v2, 
// 第一次将 v1 对应的比特置为 1, 第二次将 v1 + 1 * v2 对应的比特置位 1, 第三次将 v1 + 2 * v2 对应的比特置为 1, 以此类推) 来对比特向量赋值 (或运算)
func (f *classicFilter) Add(b []byte) {
	x, y := f.h(b)
	for i := 0; i < f.k; i++ {
		offset := f.getOffset(x, y, i)
		f.b[offset/8] |= 1 << (offset % 8)
	}
}

// Test 
// 函数与 Add 函数的逻辑基本相同, 差别在于 Test() 函数在定位到比特位之后不是将比特位赋值为 1, 而是检查比特位是否等于 0
// 若比特位为 0 说明元素一定不在集合中 (与运算)
func (f *classicFilter) Test(b []byte) bool {
	x, y := f.h(b)
	for i := 0; i < f.k; i++ {
		offset := f.getOffset(x, y, i)
		if f.b[offset/8]&(1<<(offset%8)) == 0 {
			return false
		}
	}
	return true
}

func (f *classicFilter) Size() int { return len(f.b) }

func (f *classicFilter) Reset() {
	for i := range f.b {
		f.b[i] = 0
	}
}

// simply use Double FNV here as our Bloom Filter hash
func DoubleFNV(b []byte) (uint64, uint64) {
	hx := fnv.New64()
	hx.Write(b)
	x := hx.Sum64()
	hy := fnv.New64a()
	hy.Write(b)
	y := hy.Sum64()
	return x, y
}
```