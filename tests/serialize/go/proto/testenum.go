// tests/proto/test.proto

package proto

// TestEnum  comment
type TestEnum int32

const (
	// TestEnum_TestNone  TEST_NONE 0 comment
	TestEnum_TestNone TestEnum = 0
	// TestEnum_TestValue  TestValue 0 comment
	// 2Line
	TestEnum_TestValue TestEnum = 1
	TestEnum_SnakeTestValue TestEnum = 2
)

// String TestEnum to string
func (x TestEnum) String() string {
	switch x {
	case TestEnum_TestNone:
		return "TestNone"
	case TestEnum_TestValue:
		return "TestValue"
	case TestEnum_SnakeTestValue:
		return "SnakeTestValue"
	default:
		return "Unknown"
	}
}

// ParseTestEnum string to TestEnum
func ParseTestEnum(val string) (TestEnum , bool) {
	switch val {
	case "TestNone":
		return TestEnum_TestNone, true
	case "TestValue":
		return TestEnum_TestValue, true
	case "SnakeTestValue":
		return TestEnum_SnakeTestValue, true
	default:
		return TestEnum(0), false
	}
}
