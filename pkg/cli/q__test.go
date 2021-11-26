package cli

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"k8s.io/klog"
)

func Test11(t *testing.T) {
	logBuff := bytes.Buffer{}
	klog.CopyStandardLogTo("WARNING")
	// klog.InitFlags(nil)
	// flag.Set("logtostderr", "false")
	// flag.Parse()
	klog.SetOutput(&logBuff)
	defer func() {
		klog.Flush()
		if t.Failed() {
			t.Error("failed-11")
			t.Error(logBuff.String())
		}
	}()
	tt := assert.New(t)
	klog.Error("qwe")
	rv := e11()
	tt.True(rv)
	tt.False(rv)
}

func Test12(t *testing.T) {
	logBuff := bytes.Buffer{}
	// klog.CopyStandardLogTo("INFO")
	klog.SetOutput(&logBuff)
	defer func() {
		if t.Failed() {
			t.Error("failed-11")
			t.Error(logBuff.String())
		}
	}()
	tt := assert.New(t)
	rv := e22()
	tt.True(rv)
	// tt.False(rv)
}
