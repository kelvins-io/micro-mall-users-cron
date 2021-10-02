package repository

import (
	"gitee.com/cristiane/micro-mall-users-cron/model/mysql"
	"gitee.com/kelvins-io/kelvins"
)

func ListUserInfo(sqlSelect string, pageSize, pageNum int) (result []mysql.User, err error) {
	result = make([]mysql.User, 0)
	err = kelvins.XORM_DBEngine.Table(mysql.TableUser).Select(sqlSelect).
		Desc("update_time").
		Limit(pageSize, (pageNum-1)*pageSize).
		Find(&result)
	return result, err
}

func FindUserInfo(sqlSelect string, uidList []int) (result []mysql.User, err error) {
	result = make([]mysql.User, 0)
	err = kelvins.XORM_DBEngine.Table(mysql.TableUser).Select(sqlSelect).
		In("id", uidList).
		Find(&result)
	return result, err
}

func ListMerchantInfo(sqlSelect string, pageSize, pageNum int) ([]*mysql.Merchant, error) {
	var result = make([]*mysql.Merchant, 0)
	err := kelvins.XORM_DBEngine.Table(mysql.TableMerchantInfo).Select(sqlSelect).
		Desc("update_time").
		Limit(pageSize, (pageNum-1)*pageSize).
		Find(&result)
	return result, err
}
