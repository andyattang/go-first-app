package iris

import (
	"time"

	"github.com/go-xorm/xorm"
	"github.com/kataras/iris"
	_ "github.com/mattn/go-sqlite3"
)

// User should be use to store user information
type User struct {
	ID        int64
	Version   int    `xorm: "version"`
	Username  string `xorm: "varchar(250)"`
	Password  string `xorm: "varchar(500)"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func Main() {
	var app = iris.New()
	app.Logger().SetLevel("debug")

	orm, err := xorm.NewEngine("sqlite3", "./test.db")
	if err != nil {
		app.Logger().Fatalf("failed to initalized sqlite3: %v", err)
	}
	err = orm.Sync2(new(User))
	if err != nil {
		app.Logger().Fatalf("orm failed to initalized User table: %v", err)
	}
	app.Get("/insert", func(ctx iris.Context) {
		user := &User{Username: "test user", Password: "123456", CreatedAt: time.Now(), UpdatedAt: time.Now()}
		_, err := orm.Insert(user)
		if err != nil {
			ctx.Writef("inert user error: %v", err)
		} else {
			ctx.Writef("user inserted: %#v", user)
		}
	})

	app.Get("/get/{id:int64}", func(ctx iris.Context) {
		var id = ctx.Params().GetInt64Default("id", 0)
		app.Logger().Infof("Get id:%v", id)
		user := User{ID: id}
		if ok, _ := orm.Get(&user); ok {
			ctx.Writef("user found: %#v", user)
		} else {
			ctx.Writef("user not found, id: %v", user.ID)
		}
	})

	app.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed))
}
