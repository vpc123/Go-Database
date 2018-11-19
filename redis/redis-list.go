package main 

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)
func main() {
	c,err :=redis.Dial("tcp","127.0.0.1:6379")
	if err !=nil{
		fmt.Println("Connect to redis error",err)
		return
	}
	c.Do("AUTH", "12589")
	defer c.Close()
	_,err=c.Do("lpush","runoobkey","redis")
	if err!=nil{
		fmt.Println("redis set failed:",err)
	}

	_,err=c.Do("lpush","runoobkey","mongodb")
	if err!=nil{
		fmt.Println("redis set failed:",err)
	}
	_,err=c.Do("lpush","runoobkey","mysql")
	if err!=nil{
		fmt.Println("redis set failed:",err)
	}
	fmt.Println("1 to 0 1 to 0 1 to 0")
	values,_:=redis.Values(c.Do("lrange","runoobkey","0","100"))
	for _,v :=range values{
		fmt.Println(string(v.([]byte)))
	}
}