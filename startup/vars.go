package startup

import (
	"gitee.com/cristiane/micro-mall-users-cron/model/args"
	"gitee.com/cristiane/micro-mall-users-cron/vars"
	"gitee.com/kelvins-io/kelvins"
	"gitee.com/kelvins-io/kelvins/setup"
	"gitee.com/kelvins-io/kelvins/util/queue_helper"
)

// SetupVars 加载变量
func SetupVars() error {
	var err error
	if vars.QueueAMQPSettingUserInfoSearchNotice != nil && vars.QueueAMQPSettingUserInfoSearchNotice.Broker != "" {
		vars.QueueServerUserInfoSearch, err = setup.NewAMQPQueue(vars.QueueAMQPSettingUserInfoSearchNotice, nil)
		if err != nil {
			return err
		}
		vars.QueueServerUserInfoSearchPusher, err = queue_helper.NewPublishService(vars.QueueServerUserInfoSearch, &queue_helper.PushMsgTag{
			DeliveryTag:    args.UserInfoSearchNotice,
			DeliveryErrTag: args.UserInfoSearchNoticeErr,
			RetryCount:     vars.QueueAMQPSettingUserInfoSearchNotice.TaskRetryCount,
			RetryTimeout:   vars.QueueAMQPSettingUserInfoSearchNotice.TaskRetryTimeout,
		}, kelvins.BusinessLogger)
		if err != nil {
			return err
		}
	}
	return err
}

func StopFunc() error {
	var err error
	// 当kelvins收到退出信号时将会调用
	// 本应用的资源回收通常在这里进行，kelvins自动加载的资源回收由框架进行
	return err
}
