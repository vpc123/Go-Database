package main 

import (
	"fmt"
	"encoding/json"
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

	key:="profile1"
	imap:=map[string]string{"username":"liu","phonenumber": "yusheng"}
	value,_:=json.Marshal(imap)
	n,err :=c.Do("SETNX",key,value)
	if err !=nil{
		fmt.Println(err)
	}
	// 设置过期时间为24小时  
	n, _ := rs.Do("EXPIRE", key, 24*3600)  
	if n == int64(1){
		fmt.Println("success")
	}


	var imapGet map[string]string
	valueGet,err:=redis.Bytes(c.Do("GET",key))
	if err !=nil{
		fmt.Println(err)
	}
	errShal:=json.Unmarshal(valueGet,&imapGet)
	if errShal!=nil{
		fmt.Println(err)
	}
	fmt.Println(imapGet["username"])
	fmt.Println(imapGet["phonenumber"])

}