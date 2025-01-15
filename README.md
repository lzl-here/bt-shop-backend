
## 项目目录
``` shell
.
├── Makefile     全局脚本，根据protobuf生成代码
├── README.md   
├── gateway      网关代码
├── kitex_gen.   kitex生成的代码
├── pkg          公共代码
├── protobuf     protobuf文件
└── services     存放各个服务的业务代码
    ├── example
    ├── goods
    ├── order
    ├── pay
    └── user
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

