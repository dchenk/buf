// Package buf provides a minimalistic buffer utility especially suited for small amounts of data.
package buf

// A Buf is a simple buffer that is most efficient for writing either short amounts or data (under 300 bytes) or for just
// a few writes to the buffer (not more than 3). For big buffers, or for the implemented io.Writer or io.Reader interfaces
// or more features, the bytes.Buffer may come in more useful.
type Buf struct {
	b []byte
}

// Bytes returns the buffered data as a slice, which is guaranteed to refer to the internal buffer of Buf only until the next
// modification of the buffer. The returned slice should not be written to if Buf will still be used for more writing.
func (b *Buf) Bytes() []byte { return b.b }

// String returns the buffered data as a string.
func (b *Buf) String() string { return string(b.b) }

// Write appends the given bytes to the buffer.
func (b *Buf) Write(bts []byte) {
	n := b.grow(len(bts)) // Evaluate, grow slice before copying.
	copy(b.b[n:], bts)
}

// WriteString appends the given string to the buffer.
func (b *Buf) WriteString(str string) {
	n := b.grow(len(str))
	copy(b.b[n:], str)
}

// WriteByte appends the given byte to the buffer.
func (b *Buf) WriteByte(bt byte) {
	n := b.grow(1)
	b.b[n] = bt
}

// grow ensures that the internal buffer has at least n more bytes free for writing and returns the index at
// which more can be written to the buffer.
func (b *Buf) grow(n int) (indx int) {

	indx = len(b.b)
	need := indx + n // The minimum total length needed for the write.

	if need <= cap(b.b) { // No growing is necessary.
		b.b = b.b[:need] // Since we never append to the internal buffer (but only copy), we must re-slice as needed.
		return
	}

	// Allocate a new slice for the buffer that will be bigger than 64 bytes.
	nb := make([]byte, need, need+64) // Expect there to be some more writing but not much.
	copy(nb, b.b)
	b.b = nb

	return

}

// Len returns the number of bytes in the buffer; b.Len() == len(b.Bytes()).
func (b *Buf) Len() int { return len(b.b) }

// Reset resets the buffer to be empty but keeps the underlying storage for use by future writes.
func (b *Buf) Reset() { b.b = b.b[:0] }
