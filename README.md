# dongpo_proxy
![image](https://img.shields.io/badge/golang-1.13.4-green)
                                                                                                                          
   ```                                                                                                                       
                                                                                                                          

      _                                                           
     | |                                                          
   __| | ___  _ __   __ _ _ __   ___    _ __  _ __ _____  ___   _ 
  / _` |/ _ \| '_ \ / _` | '_ \ / _ \  | '_ \| '__/ _ \ \/ / | | |
 | (_| | (_) | | | | (_| | |_) | (_) | | |_) | | | (_) >  <| |_| |
  \__,_|\___/|_| |_|\__, | .__/ \___/  | .__/|_|  \___/_/\_\\__, |
                     __/ | |           | |                   __/ |
                    |___/|_|           |_|                  |___/ 


```
# 用于网络代理转发

##教程环境:
1. 客户端Windows x64
2. 服务端CentOs 8 x64
## client端
### 编译
  1. `cd ${your_path}/donpo_proxy/client`
  2. `go build Client.go` 生成`Client.exe`，双击即可运行
### 配置
  + gui配置，根据提示输入
  + 文件配置
    + 在`Client.exe`同级目录下,新建`config.json`文件（内容如下）,启动Client.exe
    
    
```json
{
  "local.port": "8100",
  "rpc.ip": "167.179.102.249",
  "rpc.port": "8080"
}
```
        
 |  字段 | 类型 |  含义 |
 |:-----|-----:|:-----:|
 |local.port |  string  |   本地监听端口  |
 |rpc.ip  |  string  |   代理服务器ip  |
 |rpc.port |  string  |   代理服务器端口  |

## Server端
  1. `cd ${your_path}/donpo_proxy/server`
  2. `go build Server.go` 生成`Server`，linux中运行./Server