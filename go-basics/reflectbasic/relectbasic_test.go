package reflectbasic

import "testing"

func TestPrintStructInfo(t *testing.T) {
	p := Person{
		name:  "John",
		age:   30,
		famel: true,
	}
	PrintStructInfo(p)
}

func TestCallMethod(t *testing.T) {
	useCallMethod()
}

func TestPrintType(t *testing.T) {
	p := Person{
		name: "Yaoshuai",
	}
	PrintType(p)
}

func TestAddPlugins(t *testing.T) {
	addAndRemovePlugs()
}

func TestBasicRoleToken(t *testing.T) {
	doReflectTypeNewMain()
}
