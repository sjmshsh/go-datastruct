package skiplist

import (
	"fmt"
	"math"
	"math/rand"
)

const (
	// MAX_LEVEL 最高层数
	MAX_LEVEL = 16
)

type skipListNode struct {
	// 跳表保存的值
	v interface{}
	// 用于排序的分值
	score int
	// 层高
	level int
	// 每层前进指针
	forwards []*skipListNode
}

// 新建跳表节点
func newSkipListNode(v interface{}, score, level int) *skipListNode {
	return &skipListNode{v: v, score: score, forwards: make([]*skipListNode, level, level), level: level}
}

// SkipList 跳表结构体
type SkipList struct {
	// 跳表头节点
	head *skipListNode
	// 跳表当前的层数
	level int
	// 跳表长度
	length int
}

// NewSkipList 实例化跳表对象
func NewSkipList() *SkipList {
	//头结点，便于操作
	head := newSkipListNode(0, math.MinInt32, MAX_LEVEL)
	return &SkipList{head, 1, 0}
}

// Length 获取跳表长度
func (sl *SkipList) Length() int {
	return sl.length
}

// Level 获取跳表层级
func (sl *SkipList) Level() int {
	return sl.level
}

func (sl *SkipList) String() string {
	return fmt.Sprintf("level:%+v, length:%+v", sl.level, sl.length)
}

func (sl *SkipList) Insert(v interface{}, score int) int {
	if v == nil {
		return 1
	}

	// 查找插入的位置
	cur := sl.head
	// 记录每层的路径
	update := [MAX_LEVEL]*skipListNode{}
	// 跳表的最大高度
	i := MAX_LEVEL - 1

	for ; i >= 0; i-- {
		for nil != cur.forwards[i] {
			if cur.forwards[i].v == v {
				// 代表跳表当中有这个元素，不需要重复进行插入
				return 2
			}
			if cur.forwards[i].score > score {
				update[i] = cur
				break
			}
			cur = cur.forwards[i]
		}
		if cur.forwards[i] == nil {
			update[i] = cur
		}
	}

	// 通过随机算法获取该节点层数
	level := 1
	for i := 1; i < MAX_LEVEL; i++ {
		if rand.Int31()%7 == 1 {
			level++
		}
	}

	// 创建一个新的跳表节点
	newNode := newSkipListNode(v, score, level)

	// 原有节点连接
	for i := 0; i <= level-1; i++ {
		next := update[i].forwards[i]
		update[i].forwards[i] = newNode
		newNode.forwards[i] = next
	}

	// 如果当前节点的层数大于之前跳表的层数
	// 更新当前跳表的层数
	if level > sl.level {
		sl.level = level
	}

	// 更新跳表的长度
	sl.length++

	return 0
}

// Find 查找
func (sl *SkipList) Find(v interface{}, score int) *skipListNode {
	if nil == v || sl.length == 0 {
		return nil
	}
	cur := sl.head
	// 从最高层的索引开始遍历
	for i := sl.level - 1; i >= 0; i-- {
		for cur.forwards[i] != nil {
			// 如果该层不为空，就需要进入单链表的比较了
			if cur.forwards[i].score == score && cur.forwards[i].v == v {
				return cur.forwards[i]
			} else if cur.forwards[i].score > score {
				break
			}
			cur = cur.forwards[i]
		}
	}
	return nil
}

// Delete 删除节点
func (sl *SkipList) Delete(v interface{}, score int) int {
	if nil == v {
		return 1
	}

	// 查找前驱节点
	cur := sl.head
	// 记录前驱节点
	updated := [MAX_LEVEL]*skipListNode{}
	for i := sl.level - 1; i >= 0; i-- {
		updated[i] = sl.head
		for nil != cur.forwards[i] {
			if cur.forwards[i].score == score && cur.forwards[i].v == v {
				updated[i] = cur
				break
			}
			cur = cur.forwards[i]
		}
	}

	// 开始删除
	cur = updated[0].forwards[0]
	for i := cur.level - 1; i >= 0; i-- {
		if updated[i] == sl.head && cur.forwards[i] == nil {
			sl.level = i
		}

		if nil == updated[i].forwards[i] {
			updated[i].forwards[i] = nil
		} else {
			updated[i].forwards[i] = updated[i].forwards[i].forwards[i]
		}
	}

	sl.length--

	return 0
}
