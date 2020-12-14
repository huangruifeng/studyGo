package reflecttest

import (
	"fmt"
	"reflect"
	"testing"
)

//reflect.TypeOf
//reflect.ValueOf 也能获得类型
// kind 判断类型

func checkType(v interface{}) {
	t := reflect.TypeOf(v)
	switch t.Kind() {
	case reflect.Float32, reflect.Float64:
		fmt.Println("Float")
	case reflect.Int, reflect.Int32, reflect.Int64:
		fmt.Println("Integer")
	case reflect.String:
		fmt.Println("string")
	default:
		fmt.Println("Unknown", t)
	}
}

func TestBasicType(t *testing.T) {
	var f float64 = 64
	checkType(f)
}

func TestTypeAndValue(t *testing.T) {
	var f int64 = 100
	t.Log(reflect.TypeOf(f), reflect.ValueOf(f))
	t.Log(reflect.ValueOf(f).Type())
}

//reflect.ValueOf(*e).FieldByName("")
//reflect.ValueOf(e).MethodByName("").Call([]reflect.Value{reflect.ValueOf(1)})
type Employee struct {
	EmployeeID string
	Name       string `format:"normal"`
	Age        int32
}

func (e *Employee) UpdateAge(newval int32) {
	e.Age = newval
}

func TestInvokeByName(t *testing.T) {
	e := &Employee{"0", "Bob", 20}
	t.Logf("Name: value(%[1]v),Type(%[1]T)", reflect.ValueOf(*e).FieldByName("Name"))

	if nameField, ok := reflect.TypeOf(*e).FieldByName("Name"); !ok {
		t.Error("Failed to get Name!")
	} else {
		t.Log("Tag:format", nameField.Tag.Get("format"))
	}

	reflect.ValueOf(e).MethodByName("UpdateAge").Call([]reflect.Value{reflect.ValueOf(int32(32))})

	t.Log("Update Age:", e)
}
