package adapter_pattern

import "testing"

func TestPhone_OutputSound(t *testing.T) {
	lightningEarPhone := &LightningEarphoneAdapter{}
	phone := SetLightningEarphone(lightningEarPhone)
	phone.OutputSound()
}
