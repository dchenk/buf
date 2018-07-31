package buf

import (
	"bytes"
	"testing"
)

func TestBuf_Write1(t *testing.T) {

	buf := new(Buf)

	data := []byte{'a', 'b', 'c'}

	buf.Write(data)

	if got := buf.String(); got != "abc" {
		t.Errorf("got bad result %q", got)
	}

}

func TestBuf_Write2(t *testing.T) {

	buf := new(Buf)

	data := []byte("hello world")

	buf.Write(data)

	buf.WriteString("--some more...")

	want := "hello world--some more..."

	if got := buf.String(); got != want {
		t.Errorf("got bad result %q", got)
	}

}

func TestBuf_Write3(t *testing.T) {

	buf := new(Buf)

	str := "a really long string. a really long string. a really long string. a really long string. a really long string. a really long string. a really long string. a really long string. a really long string."
	str2 := "something under 32 bytes" // This string written to check that re-slicing is done when no growing is necessary.

	buf.WriteString(str)
	buf.WriteString(str2)

	if got := buf.String(); got != str+str2 {
		t.Errorf("got bad result %q", got)
	}

}

func TestBuf_Write4(t *testing.T) {

	buf := new(Buf)
	buf.Write(someData)
	buf.Write(someDataMore)
	buf.WriteString("hello world")
	buf.Write(someDataMore)
	buf.WriteString("and some more string data here")

	want := append(append(append(append(someData, someDataMore...), "hello world"...), someDataMore...), "and some more string data here"...)

	if got := buf.Bytes(); !bytes.Equal(got, want) {
		t.Errorf("got bad result: %s", got)
	}

}

func TestBuf_WriteByte(t *testing.T) {

	buf := new(Buf)

	str := "a really long string. a really long string. a really long string. a really long string. a really long string. a really long string. a really long string. a really long string. a really long string."

	buf.WriteString(str)
	buf.WriteByte('A')

	if got := buf.String(); got != str+"A" {
		t.Errorf("got bad result %q", got)
	}

}

func TestBuf_LenReset(t *testing.T) {

	buf := new(Buf)
	buf.Write(someData)
	l := buf.Len()
	buf.Reset()
	l2 := buf.Len()
	if l != len(someData) || l2 != 0 {
		t.Errorf("getting Len or doing Reset bad: got l=%d; l2=%d", l, l2)
	}

}

func BenchmarkThisBufSmall(b *testing.B) {
	for n := 0; n < b.N; n++ {
		buf := new(Buf)
		buf.WriteString("hello")
		buf.WriteByte(0x61)
	}
}

func BenchmarkThisBuf2Small(b *testing.B) {
	for n := 0; n < b.N; n++ {
		buf := new(Buf2)
		buf.WriteString("hello")
		buf.WriteByte(0x61)
	}
}

func BenchmarkBytesBufSmall(b *testing.B) {
	for n := 0; n < b.N; n++ {
		buf := new(bytes.Buffer)
		buf.WriteString("hello")
		buf.WriteByte(0x61)
	}
}

func BenchmarkConcatSmall(b *testing.B) {
	for n := 0; n < b.N; n++ {
		s := "hello"
		s += string(0x61)
	}
}

var someData = []byte{0x98, 0x77, 0x45, 0x61, 0x98, 0x77, 0x45, 0x61, 0x98, 0x77, 0x45, 0x61, 0x98, 0x77, 0x45, 0x61, 0x45, 0x61, 0x45, 0x61, 0x45, 0x61, 0x45, 0x61, 0x45}

func BenchmarkThisBufMedium(b *testing.B) {
	for n := 0; n < b.N; n++ {
		buf := new(Buf)
		buf.WriteString("hello world")
		buf.WriteByte(0x61)
		buf.Write(someData)
		buf.WriteString("and some more string data here.............")
	}
}

func BenchmarkThisBuf2Medium(b *testing.B) {
	for n := 0; n < b.N; n++ {
		buf := new(Buf2)
		buf.WriteString("hello world")
		buf.WriteByte(0x61)
		buf.Write(someData)
		buf.WriteString("and some more string data here.............")
	}
}

func BenchmarkBytesBufMedium(b *testing.B) {
	for n := 0; n < b.N; n++ {
		buf := new(bytes.Buffer)
		buf.WriteString("hello world")
		buf.WriteByte(0x61)
		buf.Write(someData)
		buf.WriteString("and some more string data here.............")
	}
}

func BenchmarkConcatMedium(b *testing.B) {
	for n := 0; n < b.N; n++ {
		s := "hello world"
		s += string(0x61)
		s += string(someData)
		s += "and some more string data here............."
	}
}

var someDataMore = []byte{0x98, 0x77, 0x45, 0x61, 0x98, 0x77, 0x45, 0x61, 0x98, 0x77, 0x45, 0x61, 0x98, 0x77, 0x45, 0x61, 0x45, 0x61, 0x45, 0x61, 0x45, 0x98, 0x77, 0x45,
	0x61, 0x98, 0x77, 0x45, 0x61, 0x98, 0x77, 0x45, 0x61, 0x98, 0x77, 0x45, 0x61, 0x45, 0x61, 0x45, 0x61, 0x45, 0x98, 0x77, 0x45, 0x61, 0x98, 0x77, 0x45, 0x61, 0x98, 0x77,
	0x45, 0x61, 0x98, 0x77, 0x45, 0x61, 0x45, 0x61, 0x45, 0x61, 0x45, 0x98, 0x77, 0x45, 0x61, 0x98, 0x77, 0x45, 0x61, 0x98, 0x77, 0x45, 0x61, 0x98, 0x77, 0x45, 0x61, 0x45,
	0x61, 0x45, 0x61, 0x45, 0x45, 0x61, 0x45, 0x61, 0x45, 0x61, 0x45, 0x98, 0x77, 0x45, 0x61, 0x98, 0x77, 0x45, 0x61, 0x98, 0x77, 0x45, 0x61, 0x98, 0x77, 0x45, 0x61, 0x45,
	0x45, 0x61, 0x98, 0x77, 0x45, 0x61, 0x45, 0x61, 0x45, 0x61, 0x45, 0x98, 0x77, 0x45, 0x61, 0x98, 0x77, 0x45, 0x61, 0x98, 0x77, 0x45, 0x61, 0x98, 0x77, 0x45, 0x61, 0x45,
	0x45, 0x61, 0x98, 0x77, 0x45, 0x61, 0x45, 0x61, 0x45, 0x61, 0x45, 0x98, 0x77, 0x45, 0x61, 0x98, 0x77, 0x45, 0x61, 0x98, 0x77, 0x45, 0x61, 0x98, 0x77, 0x45, 0x61, 0x45,
	0x45, 0x61, 0x98, 0x77, 0x45, 0x61, 0x45, 0x61, 0x45, 0x61, 0x45, 0x98, 0x77, 0x45, 0x61, 0x98, 0x77, 0x45, 0x61, 0x98, 0x77, 0x45, 0x61, 0x98, 0x77, 0x45, 0x61, 0x45}

func BenchmarkThisBufLarge(b *testing.B) {
	for n := 0; n < b.N; n++ {
		buf := new(Buf)
		buf.Write(someData)
		buf.Write(someDataMore)
		buf.WriteString("hello world")
		buf.WriteByte(0x61)
		buf.Write(someDataMore)
		buf.WriteString("and some more string data here")
	}
}

func BenchmarkThisBuf2Large(b *testing.B) {
	for n := 0; n < b.N; n++ {
		buf := new(Buf2)
		buf.Write(someData)
		buf.Write(someDataMore)
		buf.WriteString("hello world")
		buf.WriteByte(0x61)
		buf.Write(someDataMore)
		buf.WriteString("and some more string data here")
	}
}

func BenchmarkBytesBufLarge(b *testing.B) {
	for n := 0; n < b.N; n++ {
		buf := new(bytes.Buffer)
		buf.Write(someData)
		buf.Write(someDataMore)
		buf.WriteString("hello world")
		buf.WriteByte(0x61)
		buf.Write(someDataMore)
		buf.WriteString("and some more string data here")
	}
}

func BenchmarkConcatLarge(b *testing.B) {
	for n := 0; n < b.N; n++ {
		s := string(someData)
		s += string(someDataMore)
		s += "hello world"
		s += string(0x61)
		s += string(someDataMore)
		s += "and some more string data here"
	}
}

func BenchmarkThisBufXtraLarge(b *testing.B) {
	for n := 0; n < b.N; n++ {
		buf := new(Buf)
		buf.Write(someData)
		for i := 0; i < 30; i++ {
			buf.Write(someDataMore)
			buf.WriteString("something")
		}
	}
}

func BenchmarkThisBuf2XtraLarge(b *testing.B) {
	for n := 0; n < b.N; n++ {
		buf := new(Buf2)
		buf.Write(someData)
		for i := 0; i < 30; i++ {
			buf.Write(someDataMore)
			buf.WriteString("something")
		}
	}
}

// Old version of Buf:

// A Buf is a simple buffer that is most efficient for writing either short amounts or data (under 300 bytes) or for just
// a few writes to the buffer (not more than 3). For big buffers, or for the implemented io.Writer or io.Reader interfaces
// or more features, the bytes.Buffer may come in more useful.
type Buf2 struct {
	bootstrap [64]byte // allocate some memory for small buffers (often avoids two allocations)
	b         []byte   // the internal buffer
}

// Bytes returns the buffered data as a slice, which is guaranteed to refer to the internal buffer of Buf only until the next
// modification of the buffer. The returned slice should not be written to if Buf will still be used for more writing.
func (b *Buf2) Bytes() []byte { return b.b }

// String returns the buffered data as a string.
func (b *Buf2) String() string { return string(b.b) }

// Write appends the given bytes to the buffer.
func (b *Buf2) Write(bts []byte) {
	n := b.grow(len(bts)) // Evaluate, grow slice before copying.
	copy(b.b[n:], bts)
}

// WriteString appends the given string to the buffer.
func (b *Buf2) WriteString(str string) {
	n := b.grow(len(str))
	copy(b.b[n:], str)
}

// WriteByte appends the given byte to the buffer.
func (b *Buf2) WriteByte(bt byte) {
	n := b.grow(1)
	b.b[n] = bt
}

// grow ensures that the internal buffer has at least n more bytes free for writing and returns the index at
// which more can be written to the buffer.
func (b *Buf2) grow(n int) (indx int) {

	indx = len(b.b)
	need := indx + n // The minimum total length needed for the write.

	if need <= cap(b.b) { // No growing is necessary.
		b.b = b.b[:need] // Since we never append to the internal buffer (but only copy), we must re-slice as needed.
		return
	}

	// Only use the bootstrap array if the internal buffer slice is either nil or not greater than 64 because if cap(b.b) is bigger
	// than 64 it should be re-usable after a Reset().
	if need <= 64 && cap(b.b) <= 64 {
		b.b = b.bootstrap[:need] // Here we don't know if the internal slice b.b is nil. Just slice from bootstrap just in case.
		return
	}

	// Allocate a new slice for the buffer that will be bigger than 64 bytes.
	nb := make([]byte, need, need+64) // Expect there to be some more writing but not much.
	copy(nb, b.b)
	b.b = nb

	return

}

// Len returns the number of bytes in the buffer; b.Len() == len(b.Bytes()).
func (b *Buf2) Len() int { return len(b.b) }

// Reset resets the buffer to be empty but keeps the underlying storage for use by future writes.
func (b *Buf2) Reset() { b.b = b.b[:0] }
