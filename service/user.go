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
	pageSize = 50
	pageNum  = 1
)

func UserInfoSearchSync() {
	if vars.UserInfoSearchSyncTaskSetting != nil {
		if vars.UserInfoSearchSyncTaskSetting.SingleSyncNum > 0 {
			pageSize = vars.UserInfoSearchSyncTaskSetting.SingleSyncNum
		}
	}
	count := 0
	for {
		if count > 5 {
			break
		}
		count++
		userInfoSearchSyncOne(pageSize, pageNum)
		pageNum++
	}

	// 顺便同步商户信息
	go MerchantInfoSearchSync()
}

const sqlSelectUserInfo = "*"

func userInfoSearchSyncOne(pageSize, pageNum int) {
	ctx := context.TODO()
	userInfoList, err := repository.ListUserInfo(sqlSelectUserInfo, pageSize, pageNum)
	if err != nil {
		kelvins.ErrLogger.Errorf(ctx, "ListUserInfo err: %v", err)
		return
	}
	if len(userInfoList) == 0 {
		return
	}
	for i := 0; i < len(userInfoList); i++ {
		info := &args.UserInfoSearch{
			UserName:    userInfoList[i].UserName,
			Phone:       userInfoList[i].CountryCode + ":" + userInfoList[i].Phone,
			Email:       userInfoList[i].Email,
			IdCardNo:    userInfoList[i].IdCardNo.String,
			ContactAddr: userInfoList[i].ContactAddr,
		}
		userInfoSearchNotice(info)
	}
}

func userInfoSearchNotice(info *args.UserInfoSearch) {
	var ctx = context.TODO()
	userInfoMsg := args.CommonBusinessMsg{
		Type:    args.UserInfoSearchNoticeType,
		Tag:     "用户信息同步通知",
		UUID:    uuid.New().String(),
		Content: json.MarshalToStringNoError(info),
	}
	vars.QueueServerUserInfoSearchPusher.PushMessage(ctx, &userInfoMsg)
}
