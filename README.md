# dongpo_proxy
![image](https://img.shields.io/badge/golang-v1.13.4%2B-green)
                                                                                                                          
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
4. 安装go依赖包,参照`go.mod`文件内容中的依赖项
## client端
### 编译
  1. `cd ${your_path}/donpo_proxy/client/windows/view`
  2. 添加图标`rsrc -manifest main.exe.manifest -ico client.ico -o main_amd64.syso  -arch amd64`
  3. `cd ${your_path}/donpo_proxy/client/windows`
  4. `go build  -ldflags="-H windowsgui" -o dongpo_client.exe` 生成`dongpo_client.exe`，双击即可运行
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
不支持本地pac

## Server端
  1. `cd ${your_path}/donpo_proxy/server`
  2. `go build Server.go` 生成`Server`，linux中运行./Server
  > 默认启动`8080`端口,需要修改端口，可以更改`Server.go`
