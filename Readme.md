# writer

Wrap an io.Writer for metrics.

## Usage

#### type Writer

```go
type Writer struct {}
```

Writer implements an io.Writer which wraps another in order to provide write and
byte-size metrics.

#### func  New

```go
func New(w io.Writer) *Writer
```
New writer wrapping `w`.

#### func (*Writer) Bytes

```go
func (w *Writer) Bytes() uint64
```
Bytes returns the number of total bytes written.

#### func (*Writer) Write

```go
func (w *Writer) Write(b []byte) (int, error)
```
Write implements io.Writer.

#### func (*Writer) Writes

```go
func (w *Writer) Writes() uint64
```
Writes returns the total number of writes.
