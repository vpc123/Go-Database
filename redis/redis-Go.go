package main
import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func main() {
	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println(err)
		return
	}
	c.Do("AUTH", "12589")
	c.Do("SET", "a", 134)
	a, err := redis.Int(c.Do("GET", "a"))
	fmt.Println(a)
	defer c.Close()
}