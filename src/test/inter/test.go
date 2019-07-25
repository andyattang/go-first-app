package inter

import (
	"fmt"
)

type say interface {
	hello()
}

type user struct {
	name string
}

func (u user) hello() {
	fmt.Println(u.name, ", hello user")
}

type test int

// type admin user
type admin struct {
	user
	email string
}

// func (u admin) hello() {
// 	fmt.Println(u.name, ", hello admin")
// }

func Main() {
	say := admin{user: user{name: "andy"}, email: "andy@test.com"}
	say.hello()
	a := admin(say)
	fmt.Println(a.email, a.name)

	// u := user(say)
	// fmt.Println(u.email, u.name)
}
