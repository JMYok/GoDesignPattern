package factory

import (
	"errors"
	"strings"
)

// 将v0代码中涉及 parser 创建的部分逻辑剥离出来，抽象成 createParser() 函数
func (s *RuleConfigSource) createParser(configFormat string) (IRuleConfigParser, error) {
	var parser IRuleConfigParser

	switch strings.ToLower(configFormat) {
	case "json":
		parser = &JsonRuleConfigParser{}
	case "xml":
		parser = &XmlRuleConfigParser{}
	case "yaml":
		parser = &YamlRuleConfigParser{}
	case "properties":
		parser = &PropertiesRuleConfigParser{}
	default:
		return nil, errors.New("Rule config file format is not supported: " + configFormat)
	}
	return parser, nil
}

func (s *RuleConfigSource) Load2(ruleConfigFilePath string) (*RuleConfig, error) {
	ruleConfigFileExtension := getFileExtension(ruleConfigFilePath)
	parser, err := s.createParser(ruleConfigFileExtension)
	if err != nil {
		return nil, err
	}
	configText := ""
	// 从ruleConfigFilePath文件中读取配置文本到configText中

	ruleConfig, err := parser.Parse(configText)
	// 省略其他代码
	return ruleConfig, nil
}
