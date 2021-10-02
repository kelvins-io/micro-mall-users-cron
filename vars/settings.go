package vars

type EmailConfigSettingS struct {
	User     string
	Password string
	Host     string
	Port     string
}

type EmailNoticeSettingS struct {
	Receivers []string
}

type UserInfoSearchSyncTaskSettingS struct {
	Cron string
}
