package reflectbasic

import (
	"fmt"
	"reflect"
	"unsafe"
)

func PrintTypeAndValue(any interface{}) {
	t := reflect.TypeOf(any)
	v := reflect.ValueOf(any)

	fmt.Println("type:", t.String())
	fmt.Printf("value:%v\n", v)
}

type Person struct {
	name  string
	age   int
	famel bool
}

func ChangePersonAge(age int) {
	p := Person{
		name: "John",
		age:  30,
	}
	v := reflect.ValueOf(&p)
	ele := v.Elem().FieldByName("age")
	if ele.CanSet() {
		ele.SetInt(int64(age))
		fmt.Println("age:", ele.Int())
	} else {
		ptr := unsafe.Pointer(ele.UnsafeAddr())
		*(*int)(ptr) = age
		fmt.Println("can not set age:", ele.Int())
	}
	fmt.Printf("modified p:%v\n", p)
	TestChangePrivateNameField()
}

func TestChangePrivateNameField() {
	p := Person{
		name: "John",
		age:  30,
	}
	v := reflect.ValueOf(&p)
	field := v.Elem().FieldByName("name")
	if field.CanSet() {
		field.SetString("Jane")
		fmt.Println("name:", field.String())
	} else {
		ptr := unsafe.Pointer(field.UnsafeAddr())
		*(*string)(ptr) = "YaoShuai"
		fmt.Println("can not set name:", field.String())
	}
	fmt.Printf("modified p:%v\n", p)
	TestChangePrivateFamelField()
}

func TestChangePrivateFamelField() {
	p := Person{
		name:  "yaoshuai",
		age:   40,
		famel: false,
	}
	pValue := reflect.ValueOf(&p)
	field := pValue.Elem().FieldByName("famel")
	if field.CanSet() {
		field.SetBool(true)
	} else {

		unsafePtr := unsafe.Pointer(field.UnsafeAddr())
		*(*bool)(unsafePtr) = true
		fmt.Println("cannot Set famle :", field.String())
	}
	fmt.Printf("modified p famel:%v\n", p)

	fn := func() {
		fmt.Print("Im a function!")
	}
	fnType := reflect.TypeOf(&fn)
	fmt.Println("fnName :", fnType.Name())
}

func PrintStructInfo(obj any) {
	// objType := reflect.TypeOf(&obj)
	// objValue := reflect.ValueOf(&obj)

	// fmt.Println(objType.Field())
}
