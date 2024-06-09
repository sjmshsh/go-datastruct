package heap

import (
	sort2 "go-datastruct/sort"
)

type Interface interface {
	sort2.Interface
	Push(x any)
	Pop() any
}

func Push(h Interface, x interface{}) {
	h.Push(x)        // 向数据集添加一个元素
	up(h, h.Len()-1) // 从下向上堆化
}

// 从下向上堆化内容
func up(h Interface, j int) {
	// h表示堆, j代表要堆化的元素的index
	for {
		// 定义j的父index
		i := (j - 1) / 2
		// 如果两个元素相等(此时j=i=-1, 即已经对比到根节点了) 或者父元素小于当前元素
		if i == j || !h.Less(j, i) {
			break // 堆化完成
		}
		// 交换父元素和当前元素
		h.Swap(i, j)
		// index 变为父元素的 index
		j = i
	}
}

// Pop 返回堆顶的元素, 并删除它
func Pop(h Interface) interface{} {
	n := h.Len() - 1
	// 交换堆顶和最后一个元素
	h.Swap(0, n)
	// 从上到下优化
	down(h, 0, n)
	// 弹出最后一个元素
	return h.Pop()
}

// n代表堆长度
func down(h Interface, i0, n int) bool {
	i := i0 // 堆定 index
	for {
		j1 := 2*i + 1          // 左孩子 index
		if j1 >= n || j1 < 0 { // j1 < 0 after int overflow
			break
		}
		j := j1 // j = 左孩子
		if j2 := j1 + 1; j2 < n && h.Less(j2, j1) {
			// j2 = 右孩子; j 小于堆长度 && 右孩子小于左孩子
			j = j2 // j = 2*i + 2 = 右孩子
		}

		// 上面的代表是从左右孩子选出较小的那个, 将index赋值给j
		if !h.Less(j, i) {
			// 如果堆顶小于j, 堆化结束
			break
		}

		h.Swap(i, j)
		i = j
	}
	// 返回元素是否有移动
	// 此处是一个特殊设计, 用来判断向下堆化是否真的有操作
	// 当删除中间的元素的时候, 如果向下堆化没有操作的话, 就需要向上堆化
	return i > i0
}

// Remove 删除堆中指定元素, 不一定是堆顶
func Remove(h Interface, i int) interface{} {
	n := h.Len() - 1
	// 如果不是堆顶
	if n != i {
		// 交换删除元素 和 最后一个元素
		h.Swap(i, n)
		if !down(h, i, n) {
			up(h, i)
		}
	}
	return h.Pop()
}

// Fix 当某一个元素的值有变化的时候, 用来重新堆化
func Fix(h Interface, i int) {
	// i 是值被改变的 index
	if !down(h, i, h.Len()) { // 从上到下堆化
		up(h, i) // 如果没有成功就从下到上堆化
	}
}
