package setting

import (
	"github.com/go-ini/ini"
	"github.com/sirupsen/logrus"
)

type HarborInfo struct {
	Url      string
	UserName string
	PassWord string
}

type KubernetesConfig struct {
	ConfigPath string
}

var SourceHarborSetting = &HarborInfo{}
var DestinationHarborSetting = &HarborInfo{}
var KubernetesConfigSetting = &KubernetesConfig{}

func SetUp() {
	cfg, err := ini.Load("conf/harbor.ini")
	if err != nil {
		logrus.Error("Fail to parse 'conf/app.ini': %v", err)
	}
	err = cfg.Section("SourceHarbor").MapTo(SourceHarborSetting)
	if err != nil {
		logrus.Error("Cfg.MapTo sourceHarborSetting err: %v", err)
	}
	err = cfg.Section("DestinationHarbor").MapTo(DestinationHarborSetting)
	if err != nil {
		logrus.Error("Cfg.MapTo destinationHarborSetting err: %v", err)
	}

	err = cfg.Section("KubernetesConfig").MapTo(KubernetesConfigSetting)
	if err != nil {
		logrus.Error("Cfg.MapTo KubernetesConfigSetting err: %v", err)
	}
}
