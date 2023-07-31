package duckcache

// ByteView holds an immutable of bytes.
type ByteView struct {
	b []byte
}

// Len returns the view's length.
func (v ByteView) Len() int {
	return len(v.b)
}

// ByteSlice returns a copy of the data as a byte slice.
func (v ByteView) ByteSlice() []byte {
	return cloneBytes(v.b)
}

// String returns the data as a string, making a copy if necessary.
func (v ByteView) String() string {
	return string(v.b)
}

// CloneBytes makes a copy of a byte slice, used in ByteSlice func.
func cloneBytes(b []byte) []byte {
	c := make([]byte, len(b))
	copy(c, b)
	return c
}
