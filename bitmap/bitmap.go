package bitmap

type BitMap []byte

// New 这里的长度要加一，因为如果长度不是8的整数，那么最后一个字节也需要占用
// 如果传入100作为参数，那么位图的长度就是100位，能够表示0到99这100个值
// 每个字节可以表示8位，所以实际分配的内存大小就是100/8 + 1 = 13个字节
func New(length uint) BitMap {
	return make([]byte, length/8+1)
}

func (b BitMap) Set(value uint) {
	// 计算出我的值在哪一个字节的位置
	byteIndex := value / 8
	// 如果超出了位图的范围
	if byteIndex >= uint(len(b)) {
		return
	}
	// 计算出我的值在字节当中的哪一个位
	bitIndex := value % 8
	[]byte(b)[byteIndex] |= 1 << bitIndex
}

func (b BitMap) Get(value uint) bool {
	byteIndex := value / 8
	if byteIndex >= uint(len(b)) {
		return false
	}
	bitIndex := value % 8
	return []byte(b)[byteIndex]&(1<<bitIndex) != 0
}
