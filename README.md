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

## 教程环境:
1. 客户端Windows x64
2. 服务端CentOs 8 x64

## 开发环境准备
1. 安装[MSYS2](http://www.msys2.org/)
2.  `MSYS2` shell中安装gcc
>更新`MSYS2`
> 
>`pacman -Sy pacman`
>
> `pacman -Syu`
>
> `pacman -Su`
>
> `pacman -S mingw-w64-i686-toolchain`
>
> `pacman -S mingw-w64-x86_64-toolchain`
>
3. 配置环境变量`%msys2_path%/mingw64/bin`
## client端
### 编译
  0. 添加图标`rsrc -manifest main.exe.manifest -ico client.ico -o main_amd64.syso  -arch amd64`
  1. `cd ${your_path}/donpo_proxy/client`
  2. `go build  -ldflags="-H windowsgui" -o dongpo_client.exe` 生成`dongpo_client.exe`，双击即可运行
### 配置
  + gui配置，根据提示输入
  + 文件配置
    + 在`dongpo_client.exe`同级目录下,新建`dongpo.json`文件（内容如下）,dongpo_client.exe
    
    
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


### 启动pac代理
client没有开发pac功能(后面有时间会做本地pac)，可以手动设置pac，以windows为例，vpn->代理->脚本地址 填写代理pac文件路径即可。

可用代理路径`https://raw.githubusercontent.com/petronny/gfwlist2pac/master/gfwlist.pac`

## Server端
  1. `cd ${your_path}/donpo_proxy/server`
  2. `go build Server.go` 生成`Server`，linux中运行./Server
