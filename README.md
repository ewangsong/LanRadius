### 关于
- 懒投资RADIUS是一个开源的Radius服务软件。

- 懒投资RADIUS支持标准RADIUS协议（RFC 2865, RFC 2866），提供完整的AAA实现。支持灵活的策略管理，支持各种主流接入设备并轻松扩展，具备丰富的计费策略支持。

- 编写参考：ToughRADIUS、SoftRadius。

### 安装

---
#### 1、编译安装
 编译需要go环境需自己安装下载

- go get github.com/astaxie/beego
- go get github.com/spf13/cobra
- go get -u github.com/go-sql-driver/mysql   

```
在go源目录下创建文件夹用来git clone 源代码
mkdir radius
cd radius
git clone https://github.com/ewangsong/LanRadius.git
cd Lanradius
go build -o lanradius main.go
cp -r LanRadius /opt/lanradius
cd /opt/lanradius
ln -s /otp/lanradius /usr/bin/
```
#### 2、快捷安装

```
cd /opt
git clone https://github.com/ewangsong/LanRadius.git
mv LanRadius lanradius
cd /opt/lanradius
ln -s /otp/lanradius /usr/bin/
```

###  命令行参数

```
admin			Start Lanradius             //启动后台管理程序
help        	Help about any command         //显示帮助选项
radiusct    	Start Lanradius                 //启动radius认证程序
stop        	Stop Lanradius                  //停止所有lanradius程序
version     	Print the version number of Lanradius   //显示版本
``` 

### 配置文件详解

```
appname = radiusweb             //app名称   
httpport = 8080                 //后台web端口
runmode = dev                   //变成模式
dbtype = "mysql"                //数据库类型
dbinfo = "root:123456@tcp(192.168.220.138:3306)/test?charset=utf8&loc=Local"                             //数据库配置
radiusport="1812"               //radius服务端口
```
### 日志文件

```
/var/log/lanraiuds下
```
### 待做
- QA
- Readis
- OTP二次认证
