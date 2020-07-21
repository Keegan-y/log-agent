package main

import (
	pb "ankr.com/log-agent/proto"
	"context"
	"encoding/json"
	"fmt"
	"google.golang.org/grpc"
)

func main() {
	//字符串
	str_obj := `{
  "sync_info": {
      "latest_block_hash": "1414FC7313F85D7451831D029295F729F4CA3D07F18B21621BC18BF16392FBCD",
      "latest_app_hash": "198B627EF0CDE9144CFBC566986C56CAEA75F12A6D66A7345B32059D2B9E9760",
      "latest_block_height": "598995",
      "latest_block_time": "2020-01-29T17:01:11.450658236Z",
      "catching_up": true
    }
}`

	var d map[string]interface{}
	//将字符串反解析为字典
	json.Unmarshal([]byte(str_obj),&d)
	//fmt.Printf("%v\n",d)
	//打印字典里的值
	//for key,value := range d{
	//	fmt.Printf("key:%v,value:%v\n",key,value)
	//}
	//key:sync_info,value是个字典，再次[]取值
	//fmt.Printf("latest_block_height:%#v",d["sync_info"].(map[string]interface{})["latest_block_height"])
	// 连接服务器
	conn, err := grpc.Dial(":8972", grpc.WithInsecure())
	if err != nil {
		fmt.Printf("faild to connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewAppinfoClient(conn)
	// 调用服务端的SayHello
    Name := d["sync_info"].(map[string]interface{})["latest_block_height"].(string)//数据类型问题问刘奎
	r, err := c.BlockHeight(context.Background(), &pb.BlockHeightRequest{Name:Name})//字符串对象
	if err != nil {
		fmt.Printf("could not greet: %v", err)
	}
	fmt.Printf("latest_block_height: %s\n", r.Message)
}
