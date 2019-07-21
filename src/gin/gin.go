package gin

import (
	"log"
	"strconv"
	"time"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-xorm/xorm"
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

//Main is an entry
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

	app.GET("/insert", func(ctx *gin.Context) {
		user := &User{Username: "test user", Password: "123456", CreatedAt: time.Now(), UpdatedAt: time.Now()}
		_, err := orm.Insert(user)
		if err != nil {
			ctx.String(http.StatusExpectationFailed, "inert user error: %v", err)
		} else {
			ctx.String(http.StatusOK, "user inserted: %#v", user)
		}
	})

	app.GET("/get/:id", func(ctx *gin.Context) {
		idStr := ctx.Param("id")
		log.Printf("Get user ID:%v", idStr)
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			ctx.String(500, "invaild user id: %v", idStr)
			return
		}
		user := User{ID: id}
		if ok, _ := orm.Get(&user); ok {
			ctx.String(200, "user found: %#v", user)
		} else {
			ctx.String(500, "user not found, id: %v", user.ID)
		}
	})

	app.Run(":8080")
}
