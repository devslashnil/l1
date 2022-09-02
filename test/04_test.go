package test

import (
	. "l1"
	"testing"
)

func TestWorkerPool(t *testing.T) {
	ReadDataFromMainThread(4)
}
