package args

const (
	UserInfoSearchNotice    = "user_info_search_notice"
	UserInfoSearchNoticeErr = "user_info_search_notice_err"

	UserInfoSearchNoticeType        = 10014
	MerchantsMaterialInfoNoticeType = 10015
)

type UserInfoSearch struct {
	UserName    string `json:"user_name"`
	Phone       string `json:"phone"`
	Email       string `json:"email"`
	IdCardNo    string `json:"id_card_no"`
	ContactAddr string `json:"contact_addr"`
}

type MerchantInfoSearch struct {
	Uid          int64  `json:"uid"`
	UserName     string `json:"user_name"`
	MerchantCode string `json:"merchant_code"`
	RegisterAddr string `json:"register_addr"`
	HealthCardNo string `json:"health_card_no"`
	TaxCardNo    string `json:"tax_card_no"`
}

type CommonBusinessMsg struct {
	Type    int    `json:"type"`
	Tag     string `json:"tag"`
	UUID    string `json:"uuid"`
	Content string `json:"content"`
}
