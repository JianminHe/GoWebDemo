package util

import (
	"github.com/garyburd/redigo/redis"
)

func AddKey(rc redis.Conn, key string, value string)  (bool, error){
	defer rc.Close()
	n, err := rc.Do("SETNX", key, value)
	if n == int64(1) {
		return true, nil
	} else {
		return false, err
	}
}

func SetExpire(rc redis.Conn, key string, value int64) (bool, error) {
	defer rc.Close()
	n, err := rc.Do("EXPIRE", key, value)
	if n == int64(1) {
		return true, nil
	} else {
		return false, err
	}
}

func GetKey(rc redis.Conn, key string) (string, error) {
	defer rc.Close()
	val, err := redis.String(rc.Do("GET", key))
	if err != nil {
		return "", err
	}
	return val, nil
}

