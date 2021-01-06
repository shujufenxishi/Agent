package main
 
import (
 "context"
 "fmt"
 "github.com/micro/go-micro/v2"
 "Micro/proto"
)
 
func main() {
 
 // 创建一个服务
 service := micro.NewService(micro.Name("greeter.client"))
 // 初始化
 service.Init()
 // 创建一个微服务的客户端
 greeter := pb.NewGreeterService("greeter", service.Client())
 // 调用微服务
 rsp, err := greeter.Hello(context.TODO(), &pb.Request{Name: "Hello Micro"})
 if err != nil {
  fmt.Println(err)
 }
 
 fmt.Println(rsp.Msg)
}
