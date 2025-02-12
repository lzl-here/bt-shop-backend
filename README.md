
## 项目目录
``` shell
.
├── Makefile
├── README.md
├── apps                    # 存放服务代码
│   ├── example                 # 示例服务
│   ├── gateway                 # 网关，负责和前端交互
│   ├── goods                   # 商品服务
│   ├── order                   # 订单服务
│   ├── pay                     # 支付服务
│   └── user                    # 用户服务
├── deploy                      # 启动项目部署相关
├── go.mod  
├── go.sum
├── imgs                    # 存放图片
├── kitex_gen               # proto生成的golang代码
├── pkg                     # 公共代码
│   ├── config                  # 服务配置相关
│   ├── middleware              # 中间件
│   ├── model                   # model
│   └── utils                   # 工具，比如幂等防重
├── protobuf                # protobuf相关
│   ├── api                     # proto文件，用于生成网关的api代码，需要暴露在网关外和前端交互的接口需要在这冗余写一份
│   └── proto                   # proto文件，用于生成服务内部调用的grpc代码
└── todo.txt                # todo list
```



## 迭代方式

1. 在项目根目录下的protobuf中的proto下编写proto文件

2. 每个服务对应一个文件夹，详情查看example，一个文件对应server: XXX_server.proto，其他的一个文件对应一组相关的接口

3. 编写proto文件完成后在protobuf文件夹下make gen 和 make api 进行生成golang代码


## 项目依赖
0. golang
1. kitex 代码生成插件(cloudwego官网下载)
2. hertz代码生成插件(cloudwego官网下载)
3. proto，以及proto的golang插件
4. mysql
5. redis
6. etcd
7. makefile

## 项目启动
0. 安装好上述依赖项，去服务下的.env.production文件中修改对应的配置（服务端口最好别改，前端网关写死的localhost:8081）
1. 先运行根目录下的deploy下的sql
2. 运行根目录下的make deploy建立es索引
3. 每个服务下创建makefile文件，进入到服务文件夹下，make dev启动每个服务


## 其他

### 如果需要新增配置：
因为不同服务之间会有自定义的配置，所以服务需要组合AppServiceConfig使用，详情查看example服务:
![alt text](./imgs/image.png)

### 日志
日志输出到了cmd的log的app.log文件中了，不会输出到控制台