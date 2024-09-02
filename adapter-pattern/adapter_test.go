package adapter_pattern

import "testing"

func TestPhone_OutputSound(t *testing.T) {
	phone := &Phone{}
	lightningEarPhone := &LightningEarphoneAdapter{}
	phone.SetLightningEarphone(lightningEarPhone)
	phone.OutputSound()
}
