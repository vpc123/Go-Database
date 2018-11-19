package main 


import (
	"fmt"
	"github.com/garyburd/redigo/redis"

)

func main() {
	c,err := redis.Dial("tcp","127.0.0.1:6379")

	if err != nil {
		fmt.Println("Connect to redis error", err)
		returna
	}
	c.Do("AUTH", "12589")
	defer c.Close()

	_,err = c.Do("SET","username", "liuyusheng")
	if err != nil {
        fmt.Println("redis set failed:", err)
    }else{
    	fmt.Println("sucess!")
    }

    username,err :=redis.String(c.Do("GET","username"))

    if err != nil{
    	fmt.Println("redis get failed:",err)

    }else{
    	fmt.Println("Get username: %v \n",username)
    }

}