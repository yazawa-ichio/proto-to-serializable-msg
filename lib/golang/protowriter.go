package protopack

type protoWriter struct {
	*msgWriter
}

func (w *protoWriter) Length() int {
	return w.buf.Len()
}

func (w *protoWriter) WriteTag(tag uint32) error {
	return w.WriteUint32(tag)
}

func (w *protoWriter) WriteMessage(val Message) error {
	if val == nil {
		return w.WriteNil()
	}
	return val.Write(w)
}
