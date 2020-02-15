// sample/proto/Forum.proto

package proto

import (
	protopack "github.com/yazawa-ichio/proto-to-serializable-msg/lib/golang"
)

type Forum_ForumInfo struct {
	Users map[int32]*Forum_User
}

type Forum_ForumInfo_ForumNestInfo struct {
	UserPost map[int32]*Forum_PostData
}

// Pack Serialize Message
func (m *Forum_ForumInfo_ForumNestInfo) Pack() ([]byte, error) {
	return protopack.Pack(m)
}

// Write Serialize Message
func (m *Forum_ForumInfo_ForumNestInfo) Write(w protopack.Writer) error {
	// Write Map Length
	err := w.WriteMapHeader(1)
	if err != nil {
		return err
	}

	// Write user_post
	err = w.WriteTag(1)
	if err != nil {
		return err
	}
	if m.UserPost == nil {
		err = w.WriteNil()
		if err != nil {
			return err
		}
	} else {
		mapLen := len(m.UserPost)
		err = w.WriteMapHeader(mapLen)
		if err != nil {
			return err
		}
		for mapUserPostKey, mapUserPostValue := range m.UserPost {
			err = w.WriteInt32(mapUserPostKey)
			if err != nil {
				return err
			}
			err = w.WriteMessage(mapUserPostValue)
			if err != nil {
				return err
			}
		}
	}
	return err
}

// Unpack Serialize Message
func (m *Forum_ForumInfo_ForumNestInfo) Unpack(buf []byte) error {
	return protopack.Unpack(m, buf)
}

// Read Deserialize Message
func (m *Forum_ForumInfo_ForumNestInfo) Read(r protopack.Reader) error {
	// Read Map Length
	len, err := r.ReadMapHeader()
	if err != nil {
		return err
	}

	for i := uint(0); i < len; i++ {

		// Read Tag
		tag, err := r.ReadTag()
		if err != nil {
			return err
		}

		switch tag {
		case 1: // Read user_post
			isMapNil, err := r.NextFormatIsNull()
			if err != nil {
				return err
			}
			if isMapNil {
				r.ReadNil()
				m.UserPost = nil
				continue
			}
			mapUserPostLen, err := r.ReadMapHeader()
			if err != nil {
				return err
			}
			m.UserPost = make(map[int32]*Forum_PostData, mapUserPostLen)
			for mapIndex := uint(0); mapIndex < mapUserPostLen; mapIndex++ {
				var mapUserPostKey int32
				var mapUserPostValue *Forum_PostData
				mapUserPostKey, err = r.ReadInt32()
				if err != nil {
					return err
				}
				isNil, err := r.NextFormatIsNull()
				if err != nil {
					return err
				}
				if isNil {
					mapUserPostValue = nil
					err = r.ReadNil()
				} else {
					mapUserPostValue = &Forum_PostData{}
					err = mapUserPostValue.Read(r)
				}
				if err != nil {
					return err
				}
				m.UserPost[mapUserPostKey] = mapUserPostValue
			}
			break
		default:
			err = r.Skip()
			if err != nil {
				return err
			}
			break
		}
	}
	return err
}

// Pack Serialize Message
func (m *Forum_ForumInfo) Pack() ([]byte, error) {
	return protopack.Pack(m)
}

// Write Serialize Message
func (m *Forum_ForumInfo) Write(w protopack.Writer) error {
	// Write Map Length
	err := w.WriteMapHeader(1)
	if err != nil {
		return err
	}

	// Write users
	err = w.WriteTag(1)
	if err != nil {
		return err
	}
	if m.Users == nil {
		err = w.WriteNil()
		if err != nil {
			return err
		}
	} else {
		mapLen := len(m.Users)
		err = w.WriteMapHeader(mapLen)
		if err != nil {
			return err
		}
		for mapUsersKey, mapUsersValue := range m.Users {
			err = w.WriteInt32(mapUsersKey)
			if err != nil {
				return err
			}
			err = w.WriteMessage(mapUsersValue)
			if err != nil {
				return err
			}
		}
	}
	return err
}

// Unpack Serialize Message
func (m *Forum_ForumInfo) Unpack(buf []byte) error {
	return protopack.Unpack(m, buf)
}

// Read Deserialize Message
func (m *Forum_ForumInfo) Read(r protopack.Reader) error {
	// Read Map Length
	len, err := r.ReadMapHeader()
	if err != nil {
		return err
	}

	for i := uint(0); i < len; i++ {

		// Read Tag
		tag, err := r.ReadTag()
		if err != nil {
			return err
		}

		switch tag {
		case 1: // Read users
			isMapNil, err := r.NextFormatIsNull()
			if err != nil {
				return err
			}
			if isMapNil {
				r.ReadNil()
				m.Users = nil
				continue
			}
			mapUsersLen, err := r.ReadMapHeader()
			if err != nil {
				return err
			}
			m.Users = make(map[int32]*Forum_User, mapUsersLen)
			for mapIndex := uint(0); mapIndex < mapUsersLen; mapIndex++ {
				var mapUsersKey int32
				var mapUsersValue *Forum_User
				mapUsersKey, err = r.ReadInt32()
				if err != nil {
					return err
				}
				isNil, err := r.NextFormatIsNull()
				if err != nil {
					return err
				}
				if isNil {
					mapUsersValue = nil
					err = r.ReadNil()
				} else {
					mapUsersValue = &Forum_User{}
					err = mapUsersValue.Read(r)
				}
				if err != nil {
					return err
				}
				m.Users[mapUsersKey] = mapUsersValue
			}
			break
		default:
			err = r.Skip()
			if err != nil {
				return err
			}
			break
		}
	}
	return err
}
