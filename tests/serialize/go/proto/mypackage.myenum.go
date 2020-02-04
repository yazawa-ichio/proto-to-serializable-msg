// tests/proto/packagetest/package.proto

package proto

type MyPackage_MyEnum int32

const (
	MyPackage_MyEnum_None MyPackage_MyEnum = 0
	MyPackage_MyEnum_Valut MyPackage_MyEnum = 2
)

// String MyPackage_MyEnum to string
func (x MyPackage_MyEnum) String() string {
	switch x {
	case MyPackage_MyEnum_None:
		return "None"
	case MyPackage_MyEnum_Valut:
		return "Valut"
	default:
		return "Unknown"
	}
}

// ParseMyPackage_MyEnum string to MyPackage_MyEnum
func ParseMyPackage_MyEnum(val string) (MyPackage_MyEnum , bool) {
	switch val {
	case "None":
		return MyPackage_MyEnum_None, true
	case "Valut":
		return MyPackage_MyEnum_Valut, true
	default:
		return MyPackage_MyEnum(0), false
	}
}
