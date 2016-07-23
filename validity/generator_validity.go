package validity

import (
	"errors"
	"path/filepath"

	"CodeFly/global"
)

// GenerateCommandInfo 命令信息结构
type GenerateCommandInfo struct {
	Lang   string
	Input  string
	Output string
}

// CheckGenerateCommandInfoValidity 检查输入命令合法性
func (gci *GenerateCommandInfo) CheckGenerateCommandInfoValidity() error {

	if gci.Lang == "" {
		return errors.New("语言名称为空")
	}
	switch gci.Lang {
	case global.Swift:
		break
	default:
		return errors.New("未被支持的语言")
	}

	if gci.Input == "" {
		return errors.New("thrift文件路径为空")
	}
	p, err := filepath.Abs(gci.Input)
	if err != nil {
		return errors.New("thrift文件路径错误")
	}
	gci.Input = p

	if gci.Output == "" {
		return errors.New("输出文件夹路径为空")
	}
	p, err = filepath.Abs(gci.Output)
	if err != nil {
		return errors.New("输出文件路径错误")
	}
	gci.Output = p

	return nil
}