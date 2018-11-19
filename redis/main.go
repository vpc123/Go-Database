package main 
import ( 
"time" 
"github.com/kataras/iris" 
"github.com/kataras/iris/context" 
"github.com/kataras/iris/sessions" 
"github.com/kataras/iris/sessions/sessiondb/redis" 
"github.com/kataras/iris/sessions/sessiondb/redis/service" 
) 
func main() { 
	db := redis.New(service.Config{ 
	Network: "tcp", 
	Addr: "127.0.0.1:6379", 
	Password: "12589", 
	Database: "", 
	MaxIdle:  0, 
	MaxActive:  10, 
	IdleTimeout: service.DefaultRedisIdleTimeout, 
	Prefix: ""}) 

    db.Async(true)
    iris.RegisterOnInterrupt(func() {
        db.Close()
    })


    sess := sessions.New(sessions.Config{Cookie: "sessionscookieid", Expires: 45 * time.Minute})

    sess.UseDatabase(db)

    app := iris.New()

    app.Get("/", func(ctx context.Context) {
    	ctx.Writef("You should navigate to the /set, /get, /delete, /clear,/destroy instead")
    })
    app.Get("/set", func(ctx context.Context) {
        s := sess.Start(ctx)
        s.Set("name", "vpc")
            ctx.Writef("All ok session setted to: %s", s.GetString("name"))
    })

    app.Get("/get", func(ctx context.Context) {
    name := sess.Start(ctx).GetString("name")

        ctx.Writef("The name on the /set was: %s", name)
    })

    app.Get("/delete", func(ctx context.Context) {
        // delete a specific key
        sess.Start(ctx).Delete("name")
    })

    app.Get("/clear", func(ctx context.Context) {
        // removes all entries
        sess.Start(ctx).Clear()
    })

    app.Get("/destroy", func(ctx context.Context) {
    	 sess.Destroy(ctx)
     })
    app.Get("/update", func(ctx context.Context) {
        // updates expire date with a new date
        sess.ShiftExpiration(ctx)
    })

        app.Run(iris.Addr(":8080"))
}