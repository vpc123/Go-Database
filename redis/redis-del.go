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
	_,err = c.Do("SET","myKey","superLiu")
	if err != nil {
		fmt.Println("redis set faild:",err)
	}else{
		fmt.Println("sucess!")
	}

	is_key_exit, err := redis.Bool(c.Do("EXISTS", "myKey"))
	if err !=nil{
		fmt.Println("error:",err)
	}else{
		fmt.Println("exists or not:", is_key_exit)
	}

	_,err=c.Do("DEL","myKey")
	if err !=nil{
		fmt.Println("redis delete failed:",err)
	}else{
		fmt.Println("delete is sucess.")
	}

	username,err := redis.String(c.Do("GET","myKey"))
	if err !=nil{
		fmt.Println("redis get failed:",err)

	}else{
		fmt.Println("get mykey:",username)
	}
}