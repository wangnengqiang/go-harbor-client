package main

import (
	"github.com/sirupsen/logrus"
	"github.com/wangnengqiang/go-harbor-client/pkg/setting"
)

func main() {
	setting.SetUp()
	logrus.Info(setting.SourceHarborSetting.Url)
}
