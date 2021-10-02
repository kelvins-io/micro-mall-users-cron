package code

import "gitee.com/kelvins-io/common/errcode"

const (
	Success     = 29999999
	ErrorServer = 29999998
)

var ErrMap = make(map[int]string)

func init() {
	dict := map[int]string{
		Success:     "OK",
		ErrorServer: "服务器错误",
	}
	errcode.RegisterErrMsgDict(dict)
	for key, _ := range dict {
		ErrMap[key] = dict[key]
	}
}
