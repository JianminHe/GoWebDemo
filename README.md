Go Wes示例代码

学习Go语言一个月以后，编写的Restful + Redis接口用例

Redis连接使用了连接池方式。主要用于演示怎么在redis中添加、查询键值

本程序需要下载2个额外的go包， 执行以下命令
go get github.com/garyburd/redigo/redis   
go get github.com/julienschmidt/httprouter

本程序需要在本机安装Redis服务器，端口号6379，密码123456

程序编译运行，访问地址 http://127.0.0.1:8080

默认提供RESTFUL 接口读写Redis功能

1. 写入redis键值
POST方式 x-www-form-urlencoded  表单格式
http://127.0.0.1：8080/put, 有2个关键参数key, value,赋予对应值即可

2. 获取redis值
GET方式     /getRedisKey/{key}  请求格式
http://127.0.0.1:8080/getRediskey/demoKey
返回redis中键为demokey的值

3. 获取redis值
GET方式 getS/key=xxx 请求格式
http://127.0.0.1/getS/key=wang, 返回redis中的key = wang的值

