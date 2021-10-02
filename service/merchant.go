package service

import (
	"context"
	"gitee.com/cristiane/micro-mall-users-cron/model/args"
	"gitee.com/cristiane/micro-mall-users-cron/repository"
	"gitee.com/cristiane/micro-mall-users-cron/vars"
	"gitee.com/kelvins-io/common/json"
	"gitee.com/kelvins-io/kelvins"
	"github.com/google/uuid"
)

var (
	merchantSearchPageSize = 50
	merchantSearchPageNum  = 1
)

func MerchantInfoSearchSync() {
	count := 0
	for {
		if count > 2 {
			break
		}
		count++
		merchantInfoSearchSyncOne(merchantSearchPageSize, merchantSearchPageNum)
		merchantSearchPageNum++
	}
}

const sqlSelectMerchantInfo = "*"

func merchantInfoSearchSyncOne(pageSize, pageNum int) {
	ctx := context.TODO()
	merchantList, err := repository.ListMerchantInfo(sqlSelectMerchantInfo, pageSize, pageNum)
	if err != nil {
		kelvins.ErrLogger.Errorf(ctx, "ListMerchantInfo err: %v", err)
		return
	}
	if len(merchantList) == 0 {
		return
	}
	uidList := make([]int, len(merchantList))
	for i := 0; i < len(merchantList); i++ {
		uidList[i] = merchantList[i].Uid
	}
	userInfoList, err := repository.FindUserInfo("user_name", uidList)
	if err != nil {
		kelvins.ErrLogger.Errorf(ctx, "FindUserInfo err: %v, uids: %v", err, json.MarshalToStringNoError(uidList))
		return
	}
	if len(userInfoList) == 0 {
		return
	}
	uidToUserName := map[int]string{}
	for i := 0; i < len(userInfoList); i++ {
		uidToUserName[userInfoList[i].Id] = userInfoList[i].UserName
	}
	for i := 0; i < len(merchantList); i++ {
		userName, _ := uidToUserName[merchantList[i].Uid]
		info := &args.MerchantInfoSearch{
			Uid:          int64(merchantList[i].Uid),
			UserName:     userName,
			MerchantCode: merchantList[i].MerchantCode,
			RegisterAddr: merchantList[i].RegisterAddr,
			HealthCardNo: merchantList[i].HealthCardNo,
			TaxCardNo:    merchantList[i].TaxCardNo,
		}
		merchantsMaterialSearchNotice(info)
	}
}

func merchantsMaterialSearchNotice(info *args.MerchantInfoSearch) {
	var ctx = context.TODO()
	userInfoMsg := args.CommonBusinessMsg{
		Type:    args.MerchantsMaterialInfoNoticeType,
		Tag:     "商户申请信息-搜索同步",
		UUID:    uuid.New().String(),
		Content: json.MarshalToStringNoError(info),
	}
	vars.QueueServerUserInfoSearchPusher.PushMessage(ctx, &userInfoMsg)
}
