package sort

type Interface interface {
	Len() int

	Less(i, j int) bool // 返回index i 是否小于 index j

	Swap(i, j int) // 交换i和j的值
}
