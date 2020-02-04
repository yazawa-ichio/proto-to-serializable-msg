// sample/proto/Forum.proto

package proto

// Forum_Roll ロールです
type Forum_Roll int32

const (
	// Forum_Roll_Guest ゲスト
	Forum_Roll_Guest Forum_Roll = 0
	// Forum_Roll_Master マスター
	Forum_Roll_Master Forum_Roll = 1
)

// String Forum_Roll to string
func (x Forum_Roll) String() string {
	switch x {
	case Forum_Roll_Guest:
		return "Guest"
	case Forum_Roll_Master:
		return "Master"
	default:
		return "Unknown"
	}
}

// ParseForum_Roll string to Forum_Roll
func ParseForum_Roll(val string) (Forum_Roll , bool) {
	switch val {
	case "Guest":
		return Forum_Roll_Guest, true
	case "Master":
		return Forum_Roll_Master, true
	default:
		return Forum_Roll(0), false
	}
}
