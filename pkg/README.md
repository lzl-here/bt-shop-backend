# protocol

安装golang的protobuf代码生成插件
``` golang
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
```

安装kitex，自动生成rpc代码
``` golang
go install github.com/cloudwego/kitex/tool/cmd/kitex@latest
```

# 目录结构说明

## 总览
``` shell 
.
├── Makefile
├── docs    # swagger.json 产物
│   ├── define
│   └── trade
├── grpc    # 存放所有 pb 生成产物
│   ├── define  # 存放所有 pb 枚举生成产物
│   └── trade 
├── proto   # 存放所有 pb 文件
│   ├── define  # 存放所有 pb 枚举
│   │   ├── base_response.proto 
│   │   └── trade   # 根据业务模块区分存放 pb 枚举
│   └── trade   # 根据业务模块区分存放 pb 文件
└── script  # 生成pb.gw所用shell脚本
```

## proto 目录结构说明
``` shell 
.
├── define
│   ├── base_response.proto
│   └── trade
└── trade   # 存放交易/电商模块下相关项目的 pb 文件
    └── kl-pay  # 服务名称 
        └── v1  # 该服务的所有 pb 文件
```

# 新增服务

## 在 protocol 中新增服务

  本示例将展示如何新增一个服务的pb文件到protocol下， 新增服务: kl-material-audio-center (素材音频中心)

1. 在 proto 下新增业务模块 material;
2. 在 proto/material 下新增文件夹: kl-material-audio-center 以及 存放 pb 文件的 v1 文件夹： 
``` shell 
.
├── define
├── material
│   └── kl-material-audio-center
│       └── v1
└── trade
```

3. 在 proto/material/kl-material-audio-center/v1 下新增 ac_server.proto 文件，定义该服务提供的所有API: 
``` protobuf
syntax = "proto3";

package klmaterialaudiocenter.v1;

option go_package = "github.com/lzl-here/mini_video_proto/sdk/protocol/grpc/material/kl-material-audio-center/v1";

import "google/api/annotations.proto";

service MaterialAudioCenter {
  rpc Ping(PingReq) returns (PingRsp) {
    option (google.api.http) = {
      post : "/xe.material.audio.ping/1.0.0"
      body : "*"
    };
  };
}

// ping - 请求。
message PingReq {}

// ping - 响应。
message PingRsp {
  // 返回状态码。
  int32 code = 1;

  // 返回消息。
  string message = 2;
}
```

4. 在 protocol 项目根目录下执行 make go / buf generate 命令，即可生成pb代码产物: 
``` shell 
.
├── grpc
│   ├── material
│   │   └── kl-material-audio-center
│   │       └── v1
│   │           ├── ac_server.pb.go
│   │           ├── ac_server.pb.gw.go
│   │           └── ac_server_grpc.pb.go
├── proto
│   ├── material
│   │   └── kl-material-audio-center
│   │       └── v1
│   │           └── ac_server.proto
... 
```

## 在服务中引入 protocol 
  本示例将展示如何在服务中引入 protocol 中的 pb 产物;

1. 项目中引入 "github.com/lzl-here/mini_video_proto/sdk/protocol/grpc/material/kl-material-audio-center/v1" (即 proto 文件中的 go_package)
``` go 
  import 	acpb "github.com/lzl-here/mini_video_proto/sdk/protocol/grpc/material/kl-material-audio-center/v1"
``` 

2. 更新 go mod;
``` go 
...
	github.com/lzl-here/mini_video_proto/sdk/protocol latest
...
```

## 更新协议
  本示例将展示如何更新协议;
  
1. 对 pb 文件进行更新 (将 PingReq 拆分到 ac_message_ping.proto 文件中): 
``` protobuf
syntax = "proto3";

package klmaterialaudiocenter.v1;

option go_package = "github.com/lzl-here/mini_video_proto/sdk/protocol/grpc/material/kl-material-audio-center/v1";

import "google/api/annotations.proto";
import "material/kl-material-audio-center/v1/ac_message_ping.proto";

service MaterialAudioCenter {
  rpc Ping(PingReq) returns (PingRsp) {
    option (google.api.http) = {
      post : "/xe.material.audio.ping/1.0.0"
      body : "*"
    };
  };
``` 
  对应的目录结构: 
``` shell 
.
└── kl-material-audio-center
    └── v1
        ├── ac_message_ping.proto
        └── ac_server.proto
```

2. 在 protocol 项目根目录下执行 make go / buf generate 命令，生成pb代码产物: 
``` shell 
.
├── grpc
│   ├── material
│   │   └── kl-material-audio-center
│   │       └── v1
│   │           ├── ac_message_ping.pb.go
│   │           ├── ac_server.pb.go
│   │           ├── ac_server.pb.gw.go
│   │           └── ac_server_grpc.pb.go
├── proto
│   ├── material
│   │   └── kl-material-audio-center
│   │       └── v1
│   │           ├── ac_message_ping.proto
│   │           └── ac_server.proto
... 
```

3. 提交本次改动;
4. 在引用该协议的服务中更新 go mod: 
``` go 
	github.com/lzl-here/mini_video_proto/sdk/protocol latest 
``` 
5. 执行 go mod tidy；
6. 成功更新；

  

  

  

  

  



