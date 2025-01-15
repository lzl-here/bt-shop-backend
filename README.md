
## 项目目录
``` shell
.
├── Makefile.      # 全局构建脚本，通过proto生成代码
├── apps           # 项目文件夹， 服务 + 网关
│   ├── example. 
│   ├── gateway
│   ├── goods
│   ├── order
│   ├── pay
│   └── user
├── deploy         # 部署相关，后续sql放这，并且加入一键部署所有服务 
├── go.mod
├── go.sum
├── kitex_gen      # 通过proto生成的代码 
├── pkg            # 公共包
├── protobuf       # proto文件
└── todo.txt       
```



## 迭代方式

1. 假如需要新增服务，在项目根目录下的protobuf中的proto下编写proto文件

2. 每个服务对应一个文件夹，详情查看example，一个文件对应server: XXX_server.proto，其他的一个文件对应一个接口

3. 编写完成后使用Makefile进行生成golang代码： make gen


## 项目启动
每个服务下有一个makefile文件，进入到服务文件夹下，make dev执行

导入proto生成代码：
``` golang
import (
	pgen "github.com/lzl-here/bt-shop-backend/kitex_gen/example"
)
```


example 服务的 vscode 的 debug启动: 
``` json
{
    "version": "0.2.0",
    "configurations": [
        {
            "name": "example",
            "type": "go",
            "request": "launch",
            "mode": "auto",
            "program": "${workspaceFolder}/cmd/main.go",
            "args": [
                "--cfgFile", "${workspaceFolder}/.env.production",
            ]
        }
    ]
}
```

