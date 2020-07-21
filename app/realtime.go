package app
import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Ankr-network/dccn-metrics-agent/core/client"
	metrics "github.com/Ankr-network/dccn-metrics-agent/proto"
	"github.com/golang/glog"

	//"github.com/golang/glog"

	//"github.com/golang/glog"
)

type realTime struct {
	metrics.UnimplementedRealtimeMetricsSrvServer
	//      不实现 嵌入式 兼容
	//在metric对象里
}

/*
調用其他函數接口
*/
func (r *realTime) Query(ctx context.Context, req *metrics.RealTimeRequest) (*metrics.RealTimeResponse, error) {
	glog.Info("kuaigao  query")
	if err := req.Validate(); err != nil {
		return nil, err
	}
	//这里的逻辑我猜是？
	//有些东西封装好了，明确对象，改改方法
	namespace := req.GetNamespace()
	k8sClient := client.GetK8sClientCoreV1()
	ns, err := k8sClient.Namespaces().Get(namespace)//
	if err != nil {
		glog.Errorf("uptime query error: %s", err.Error())
		return nil, err
	}

	createTime := ns.CreationTimestamp.Unix()

	res := &metrics.UptimeResponse{
		Data: &metrics.Uptime{
			Timestamp: createTime,
		},
	}

	return res, nil
}

//定义一个获取块高信息的接口
//func GetBlockHeight() () {
//
//}
