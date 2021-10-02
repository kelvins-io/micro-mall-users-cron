package startup

import (
	"gitee.com/cristiane/micro-mall-users-cron/service"
	"gitee.com/cristiane/micro-mall-users-cron/vars"
	"gitee.com/kelvins-io/kelvins"
)

func GenCronJobs() []*kelvins.CronJob {
	tasks := make([]*kelvins.CronJob, 0)
	if vars.UserInfoSearchSyncTaskSetting != nil {
		if vars.UserInfoSearchSyncTaskSetting.Cron != "" {
			tasks = append(tasks, &kelvins.CronJob{
				Name: "用户-商户信息同步搜索",
				Spec: vars.UserInfoSearchSyncTaskSetting.Cron,
				Job:  service.UserInfoSearchSync,
			})
		}
	}

	return tasks
}
