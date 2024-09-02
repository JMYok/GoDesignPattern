package adapter_pattern

import "fmt"

type StandardEarphone struct {
}

func (p *StandardEarphone) PlayByStandardInterface() {
	fmt.Println("播放声音")
}
