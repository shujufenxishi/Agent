package main
 
import (
 "context"
 "fmt"
 "github.com/micro/go-micro/v2"
 "Micro/proto"
)
 
type Greeter struct {}
 
func (g *Greeter) Hello(ctx context.Context, req *pb.Request, rsp *pb.Response) error  {
 //把客户端的请求回射给客户端
 rsp.Msg = req.Name
 return nil
}
 
func main() {
 
 // 新创建一个服务，服务名为greeter，服务注册中心会用这个名字来发现服务
 service := micro.NewService(
  micro.Name("greeter"),
 )
 
 // 初始化
 service.Init()
 
 // 注册处理器
 pb.RegisterGreeterHandler(service.Server(), new(Greeter))
 
 // 启动服务运行
 if err := service.Run(); err != nil {
  fmt.Println(err)
  
 }
}
