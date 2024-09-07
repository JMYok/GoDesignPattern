package adapter_pattern

// Phone 依赖Lightning接口
type Phone struct {
	lightningEarPhone LightningEarPhone
}

func SetLightningEarphone(earPhone LightningEarPhone) *Phone {
	return &Phone{earPhone}
}

func (p *Phone) OutputSound() {
	p.lightningEarPhone.PlayByLightningInterface()
}
