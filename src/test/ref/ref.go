package ref

import (
	"fmt"
	"reflect"
)

type anything interface{}
type u struct {
	Name string  `f:"xxx"`
	Pwd  []int64 `don't print out`
}

func (u1 *u) Hello(msg string) string {
	return fmt.Sprintln("u.Name:", u1.Name, " ", msg)
}

func Main() {
	list := []anything{
		&u{Name: "my",
			Pwd: []int64{1, 2, 3, 4}}, 2, 3.4, "5"}
	for idx, e := range list {
		fmt.Print("index: ", idx, ": ")
		switch v := e.(type) {
		case *u:
			// type conert
			u2 := (*u)(v)
			tu := reflect.ValueOf(v)
			u3 := tu.Interface().(*u)
			fmt.Println(u2.Pwd, ", ", u3.Name)

			// reflect: e should be a pointer type
			v1 := reflect.ValueOf(e)
			el := v1.Elem()
			// get a field by name and set the value
			name := el.FieldByName("Name")
			name.SetString("t1")

			// get all field
			t1 := reflect.TypeOf(e)
			tt := t1.Elem()
			fmt.Println("number of field:", tt.NumField())
			for i := 0; i < el.NumField(); i++ {
				f := tt.Field(i)
				val := el.Field(i)
				fmt.Printf("%v(%v,tag`%v`)=%v\n", f.Name, f.Type, f.Tag, val.Interface())
			}

			// get all
			fmt.Println("number of method:", v1.NumMethod())
			for i := 0; i < v1.NumMethod(); i++ {
				m := t1.Method(i)
				mm := v1.Method(i)
				fmt.Printf("%v(%v)\n", m.Name, m.Type)
				msg := reflect.ValueOf("No message")
				result := mm.Call([]reflect.Value{msg})
				fmt.Println("call method and get result:", result[0])
			}

		case int:
			t := reflect.TypeOf(v)
			fmt.Printf("%v is a %v\n", e, t)
		case float64:
			fmt.Printf("%v is a float64\n", v)
		case string:
			fmt.Printf("%v is a string\n", v)
		default:
			fmt.Printf("%v is an unknow\n", v)
		}

	}
}
