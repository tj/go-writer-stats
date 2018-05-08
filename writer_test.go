package writer

import (
	"io/ioutil"
	"testing"

	"github.com/tj/assert"
)

func TestWrites(t *testing.T) {
	w := New(ioutil.Discard)

	assert.Equal(t, uint64(0), w.Writes())
	w.Write([]byte("hello"))
	assert.Equal(t, uint64(1), w.Writes())
	w.Write([]byte("hello"))
	assert.Equal(t, uint64(2), w.Writes())
}

func TestBytes(t *testing.T) {
	w := New(ioutil.Discard)

	assert.Equal(t, uint64(0), w.Bytes())
	w.Write([]byte("hello"))
	assert.Equal(t, uint64(5), w.Bytes())
	w.Write([]byte("hello"))
	assert.Equal(t, uint64(10), w.Bytes())
}
