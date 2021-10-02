package util

import (
	"context"
	"gitee.com/kelvins-io/kelvins/config/setting"
	"gitee.com/kelvins-io/kelvins/util/client_conn"
	"gitee.com/kelvins-io/kelvins/util/middleware"
	"google.golang.org/grpc"
)

func GetGrpcClient(ctx context.Context, serverName string) (*grpc.ClientConn, error) {
	client, err := client_conn.NewConnClient(serverName)
	if err != nil {
		return nil, err
	}
	// 调用其它RPC填写OAUTH信息
	// 应该根据serverName去加载对应的token（如果需要的话），这里只是为了方便将所有服务的token设为一样
	conf := &setting.RPCAuthSettingS{
		Token:             "c9VW6ForlmzdeDkZE2i8", // 如果对方服务不需要授权，则token置空
		TransportSecurity: false,
	}
	opts := middleware.GetRPCAuthDialOptions(conf)
	return client.GetConn(ctx, opts...)
}

func GetHttpEndpoints(ctx context.Context, serverName string) ([]string, error) {
	client, err := client_conn.NewConnClient(serverName)
	if err != nil {
		return nil, err
	}
	return client.GetEndpoints(ctx)
}
