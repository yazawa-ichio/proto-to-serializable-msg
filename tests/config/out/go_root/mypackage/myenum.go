// tests/proto/packagetest/package.proto

package mypackage

type MyEnum int32

const (
	MyEnum_None MyEnum = 0
	MyEnum_Valut MyEnum = 2
)

// String MyEnum to string
func (x MyEnum) String() string {
	switch x {
	case MyEnum_None:
		return "None"
	case MyEnum_Valut:
		return "Valut"
	default:
		return "Unknown"
	}
}

// ParseMyEnum string to MyEnum
func ParseMyEnum(val string) (MyEnum , bool) {
	switch val {
	case "None":
		return MyEnum_None, true
	case "Valut":
		return MyEnum_Valut, true
	default:
		return MyEnum(0), false
	}
}
