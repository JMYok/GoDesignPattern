package factory

import (
	"errors"
	"strings"
)

// RuleConfig 表示规则配置
type RuleConfig struct {
	// 假设RuleConfig有一些字段
}

// IRuleConfigParser 是解析规则配置的接口
type IRuleConfigParser interface {
	Parse(configText string) (*RuleConfig, error)
}

// JsonRuleConfigParser 是JSON格式的解析器
type JsonRuleConfigParser struct{}

func (p *JsonRuleConfigParser) Parse(configText string) (*RuleConfig, error) {
	// 实现JSON解析逻辑
	return &RuleConfig{}, nil
}

// XmlRuleConfigParser 是XML格式的解析器
type XmlRuleConfigParser struct{}

func (p *XmlRuleConfigParser) Parse(configText string) (*RuleConfig, error) {
	// 实现XML解析逻辑
	return &RuleConfig{}, nil
}

// YamlRuleConfigParser 是YAML格式的解析器
type YamlRuleConfigParser struct{}

func (p *YamlRuleConfigParser) Parse(configText string) (*RuleConfig, error) {
	// 实现YAML解析逻辑
	return &RuleConfig{}, nil
}

// PropertiesRuleConfigParser 是Properties格式的解析器
type PropertiesRuleConfigParser struct{}

func (p *PropertiesRuleConfigParser) Parse(configText string) (*RuleConfig, error) {
	// 实现Properties解析逻辑
	return &RuleConfig{}, nil
}

// RuleConfigSource 负责加载规则配置
type RuleConfigSource struct{}

func (s *RuleConfigSource) Load(ruleConfigFilePath string) (*RuleConfig, error) {
	ruleConfigFileExtension := getFileExtension(ruleConfigFilePath)
	var parser IRuleConfigParser

	switch strings.ToLower(ruleConfigFileExtension) {
	case "json":
		parser = &JsonRuleConfigParser{}
	case "xml":
		parser = &XmlRuleConfigParser{}
	case "yaml":
		parser = &YamlRuleConfigParser{}
	case "properties":
		parser = &PropertiesRuleConfigParser{}
	default:
		return nil, errors.New("Rule config file format is not supported: " + ruleConfigFilePath)
	}

	configText := ""
	// 从ruleConfigFilePath文件中读取配置文本到configText中

	ruleConfig, err := parser.Parse(configText)
	if err != nil {
		return nil, err
	}

	return ruleConfig, nil
}

// getFileExtension 从文件路径中获取文件扩展名
func getFileExtension(filePath string) string {
	// 解析文件名获取扩展名，比如rule.json，返回json
	return "json"
}
