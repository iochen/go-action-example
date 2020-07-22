package example_test

import (
	"runtime"
	"testing"
)

func Test(t *testing.T) {
	t.Log(runtime.Version())
	t.Log(runtime.GOOS, runtime.GOARCH)
}
