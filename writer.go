// Wrap an io.Writer for metrics.
package writer

import "sync/atomic"
import "io"

// Writer implements an io.Writer which wraps another
// in order to provide write and byte-size metrics.
type Writer struct {
	w      io.Writer
	writes uint64
	bytes  uint64
}

// New writer wrapping `w`.
func New(w io.Writer) *Writer {
	return &Writer{w: w}
}

// Write implements io.Writer.
func (w *Writer) Write(b []byte) (int, error) {
	atomic.AddUint64(&w.writes, 1)
	atomic.AddUint64(&w.bytes, uint64(len(b)))
	return w.w.Write(b)
}

// Writes returns the total number of writes.
func (w *Writer) Writes() uint64 {
	return atomic.LoadUint64(&w.writes)
}

// Bytes returns the number of total bytes written.
func (w *Writer) Bytes() uint64 {
	return atomic.LoadUint64(&w.bytes)
}
