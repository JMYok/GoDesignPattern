package adapter_pattern

// Phone 依赖Lightning接口
type Phone struct {
	lightningEarPhone LightningEarPhone
}

func (p *Phone) SetLightningEarphone(earPhone LightningEarPhone) {
	p.lightningEarPhone = earPhone
}

func (p *Phone) OutputSound() {
	p.lightningEarPhone.PlayByLightningInterface()
}
