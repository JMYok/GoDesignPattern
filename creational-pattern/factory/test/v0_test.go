package test

import (
	"fmt"
	"jmy/go-design-pattern/creational-pattern/factory"
	"testing"
)

func TestV0(t *testing.T) {
	source := &factory.RuleConfigSource{}
	config, err := source.Load("./rule.json")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Config loaded:", config)
	}
}
