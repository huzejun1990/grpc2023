# grpc2023


golang开发框架-高性能RPC框架gRPC
2023年12月28日00:20:22
grpc 用于微服务
1.鉴权问题
2.grpc 数据传递，类似 http header
3.拦截器
4.客户端负载均衡（如果服务已经部署为负载均衡，那么无需客户端负载均衡）
5.服务的健康检查
6.数据传输的方式（一元请求或流式请求）
7.服务之间的认证问题
8.服务限流的问题，服务接口限流
9.服务的熔断，通过判断发生错误的次数，对服务做降级
10.日志追踪

protobuf
1.编译器
2.安装编译器插件
3.生成代码
4.将代码放到项目中编译

优势：
1.传输会更快，以二进制的方式传输


2023-12-28 20:04:20

helloworld

helloworld

protoc --go_out=plugins=grpc:. ./helloworld/*.proto


protoc.exe -I . --go_out=. --go_opt=paths=source_relative ./helloworld/proto/helloworld.proto

protoc.exe -I . --go-grpc_out=helloworld/proto --go-grpc_opt=paths=source_relative ./helloworld/proto/helloworld.proto

protoc.exe -I . --go-grpc_out=. --go-grpc_opt=paths=source_relative ./helloworld/proto/helloworld.proto
