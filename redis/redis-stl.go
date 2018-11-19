package main 
// tls


import (
	"fmt"
	"time"
	"github.com/garyburd/redigo/redis"

)

func main() {
	c,err := redis.Dial("tcp","127.0.0.1:6379")

	if err != nil {
		fmt.Println("Connect to redis error", err)
		return
	}
	c.Do("AUTH", "12589")
	defer c.Close()

	_,err = c.Do("SET","myKey", "superWang","EX","5")
	if err != nil {
        fmt.Println("redis set failed:", err)
    }else{
    	fmt.Println("sucess!")
    }

    username,err:=redis.String(c.Do("GET","myKey"))

    if err != nil{
    	fmt.Println("redis get failed:",err)

    }else{
    	fmt.Println("Get username:\n",username)
    }

    time.Sleep(8*time.Second)
    username2,err:=redis.String(c.Do("GET","myKey"))

    if err != nil{
    	fmt.Println("redis get failed:",err)

    }else{
    	fmt.Println("Get username2:\n",username2)
    }
}