package p

import "sync"

func foo() {
	var _ sync.Pool // want `sync.Pool cannot be used`
}
