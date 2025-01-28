
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

3. 编写完成后使用Makefile进行生成golang代码，在项目根目录或者protobuf目录下执行make gen，两者等价


## 项目启动
每个服务下创建makefile文件，进入到服务文件夹下，make dev执行

## 项目依赖
0. golang
1. kitex
2. proto，以及proto的golang插件
3. mysql
4. redis
5. etcd
6. makefile



因为不同服务之间会有自定义的配置，所以服务需要组合AppServiceConfig使用，详情查看example服务:
![alt text](./imgs/image.png)

## 其他
一定一定先看一遍example服务！！