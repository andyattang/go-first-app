package inter

import (
	"fmt"
)

func Main() {
	// var n notificater
	n := user{name: "andy", email: "test.163.com"}
	notificate(&n)
}

type notificater interface {
	notify()
}

type user struct {
	name  string
	email string
}

func (u user) notify() {
	fmt.Printf("send mail(%s) to %s", u.email, u.name)
}

func notificate(n notificater) {
	n.notify()
}
