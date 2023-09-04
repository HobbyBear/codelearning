package bitmap

type BitMap struct {
	flags []byte
}

func New(max int64) *BitMap {
	flagLen := max/8 + 1
	return &BitMap{flags: make([]byte, flagLen)}
}

// 关键是计算index 在位图上的位置

func (b *BitMap) Set(index int64) {
	arrIndex := index / 8
	offset := index % 8
	// 将offset位置设置为1,或运算,0 | 1 = 1  1|1= 1, 0|0 =0, 1的| 将原值设置为1 ，0的| 不改变原值
	b.flags[arrIndex] = b.flags[arrIndex] | (0x1 << offset)
}

func (b *BitMap) Clean(index int64) {
	arrIndex := index / 8
	offset := index % 8
	// 0 & 1 = 0 ,0 & 0 = 0, 1&1 =1  1的& 不会改变原来的值， 0的& 将原值变为0
	b.flags[arrIndex] = b.flags[arrIndex] & ^(0x1 << offset)
}

func (b *BitMap) Exits(index int64) bool {
	arrIndex := index / 8
	offset := index % 8
	res := b.flags[arrIndex] & (0x1 << offset)
	if res == 0 {
		return false
	}
	return true
}
