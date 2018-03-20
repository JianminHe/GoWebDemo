package main

import (
	"github.com/garyburd/redigo/redis"
	"time"
	"WebSample/util"
	"fmt"
	"net/http"
	"github.com/julienschmidt/httprouter"
	"net/url"
)


var (
	RedisClient *redis.Pool
	REDIS_HOST string = "127.0.0.1:6379"
	REDIS_DB int = 0
	REDIS_PASSWORD = "123456"
	WEB_IP string = "127.0.0.1:8080"
)

func init() {

	RedisClient = &redis.Pool{
		// 从配置文件获取maxidle以及maxactive，取不到则用后面的默认值
		MaxIdle : 1,
		MaxActive:  10,
		IdleTimeout: 180 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", REDIS_HOST)
			if err != nil {
				return nil, err
			}
			if _, err := c.Do("AUTH", REDIS_PASSWORD); err != nil {
				c.Close()
				return nil, err
			}
			// 选择db
			c.Do("SELECT", REDIS_DB)
			return c, nil
		},
	}

}

//    /getRedisKey/{key}  请求格式
// 例如   访问 http://127.0.0.1/getRedisKey/wang ,  返回 redis中的key = wang的值
func getRedisKey(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	key := p.ByName("key")
	val, err :=util.GetKey(RedisClient.Get(), key)
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}
	fmt.Fprintln(w, val)
}

// x-www-form-urlencoded  表单格式
func setRedisKey(w http.ResponseWriter, r *http.Request , p httprouter.Params) {

	r.ParseForm()
	key :=  r.PostFormValue("key")
	value := r.PostFormValue("value")
	//fmt.Fprintln(w, r.Form)
	if key == "" {
		fmt.Fprintln(w, "Redis Key is empty")
		return
	}

	if value == "" {
		fmt.Fprintln(w, "Redis value is empty ")
		return
	}

	ok, err := util.AddKey(RedisClient.Get(), key, value)

	if ok == false  && err != nil{
		fmt.Fprintln(w,err)
		return
	}

	if ok == false && err == nil {
		fmt.Fprintln(w, "插入键值已经在服务器中存在.")
	}
	fmt.Fprintf(w,"Add %s=%s to redis server sucess", key, value)
}

//  /getS/key=xxx 请求格式
// 例如  访问http://127.0.0.1/getS/key=wang, 返回redis中的key = wang的值
func getS(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	queryForm, err := url.ParseQuery(r.URL.RawQuery)

	if err != nil {
		fmt.Fprintln(w, err)
		return
	}
	if len(queryForm["key"]) <0 {
		fmt.Fprintln(w, "key参数为空")
		return
	}
	val, err :=util.GetKey(RedisClient.Get(),  queryForm["key"][0])
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}
	fmt.Fprintln(w, val)

}

func main() {

	mux := httprouter.New()
	mux.GET("/getRedisKey/:key", getRedisKey)

	mux.GET("/get", getS)

	mux.POST("/put", setRedisKey)

	server := http.Server{
		Addr:WEB_IP,
		Handler:mux,
	}
	server.ListenAndServe()



}
