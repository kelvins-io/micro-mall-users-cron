package startup

import (
	"gitee.com/cristiane/micro-mall-users-cron/vars"
	"gitee.com/kelvins-io/kelvins/config"
	"gitee.com/kelvins-io/kelvins/config/setting"
)

const (
	SectionEmailConfig       = "email-config"
	UserInfoSearchNotice     = "user-info-search-notice"
	UserInfoSearchNoticeTask = "user-info-search-notice-task"
)

// LoadConfig 加载配置对象映射
func LoadConfig() error {
	// 邮箱
	vars.EmailConfigSetting = new(vars.EmailConfigSettingS)
	config.MapConfig(SectionEmailConfig, vars.EmailConfigSetting)

	// 用户信息入库
	vars.QueueAMQPSettingUserInfoSearchNotice = new(setting.QueueAMQPSettingS)
	config.MapConfig(UserInfoSearchNotice, vars.QueueAMQPSettingUserInfoSearchNotice)

	vars.UserInfoSearchSyncTaskSetting = new(vars.UserInfoSearchSyncTaskSettingS)
	config.MapConfig(UserInfoSearchNoticeTask, vars.UserInfoSearchSyncTaskSetting)
	return nil
}
