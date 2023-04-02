package utils

import (
	"hash/fnv"
	"sync"
)

var locks []sync.Mutex
var size uint64

func InitLocks() {
	size = 1000
	locks = make([]sync.Mutex, size, size)
}

func Lock(key string) {
	h := fnv.New64()
	h.Write([]byte(key))
	locks[h.Sum64()%size].Lock()
}

func UnLock(key string) {
	h := fnv.New64()
	h.Write([]byte(key))
	locks[h.Sum64()%size].Unlock()
}
