package easy

//适配器作为中间桥梁

type LightningEarphoneAdapter struct {
	standardEarphone StandardEarphone
}

func (p *LightningEarphoneAdapter) PlayByLightningInterface() {
	p.standardEarphone = new(StandardInEarphone)
	p.standardEarphone.PlayByStandardInterface()
}
