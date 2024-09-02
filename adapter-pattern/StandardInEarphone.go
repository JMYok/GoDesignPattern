package adapter_pattern

import "fmt"

// 标准插头耳机中的 入耳式耳机
type StandardInEarphone struct {
}

func (earphone *StandardInEarphone) PlayByStandardInterface() {
	fmt.Println("入耳式-标准接口耳机 播放声音")
}
