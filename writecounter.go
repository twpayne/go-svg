package svg

import "io"

type writeCounter struct {
	w            io.Writer
	bytesWritten int
}

func (wc *writeCounter) Write(p []byte) (int, error) {
	n, err := wc.w.Write(p)
	wc.bytesWritten += n
	return n, err
}
