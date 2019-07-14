package gin

import (
	"log"
	"time"

	"github.com/go-xorm/xorm"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

// User should be use to store user information
type User struct {
	ID        int64
	Version   int    `xorm:"version"`
	Username  string `xorm:"varchar(250)"`
	Password  string `xorm:"varchar(500)", json:"password"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func Main() {
	// inital the app
	var app = gin.New()
	app.Use(gin.Logger())

	// inital xorm
	orm, err := xorm.NewEngine("sqlite3", "./test.db")
	if err != nil {
		log.Fatalf("failed to initalized sqlite3: %v", err)
	}
	err = orm.Sync2(new(User))
	if err != nil {
		log.Fatalf("orm failed to initalized User table: %v", err)
	}
	app.Post("/insert", func(ctx *gin.Context) {
		user := &User{Username: "test user", Password: "123456", CreatedAt: time.Now(), UpdatedAt: time.Now()}
		_, err := orm.Insert(user)
		if err != nil {
			ctx.Writef("inert user error: %v", err)
		} else {
			ctx.Writef("user inserted: %#v", user)
		}
	})

	app.Get("/get/:id", func(ctx *gin.Context) {
		user := User{ID: ctx.("id")}
		if ok, _ := orm.Get(&user); ok {
			ctx.Writef("user found: %#v", user)
		} else {
			ctx.Writef("user not found, id: %v", user.ID)
		}
	})

	app.Run(":8080")
}
