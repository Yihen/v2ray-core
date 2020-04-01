// +build !confonly

package dispatcher

import (
	"v2ray.com/core/common"
	"v2ray.com/core/common/buf"
)

type SizeStatWriter struct {
	Writer buf.Writer
}

func (w *SizeStatWriter) WriteMultiBuffer(mb buf.MultiBuffer) error {
	return w.Writer.WriteMultiBuffer(mb)
}

func (w *SizeStatWriter) Close() error {
	return common.Close(w.Writer)
}

func (w *SizeStatWriter) Interrupt() {
	common.Interrupt(w.Writer)
}
