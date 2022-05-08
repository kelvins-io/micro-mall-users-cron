package vars

type EmailConfigSettingS struct {
	User     string
	Password string
	Host     string
	Port     string
}

type UserInfoSearchSyncTaskSettingS struct {
	Cron          string
	SingleSyncNum int
}
