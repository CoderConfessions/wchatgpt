package utils

import (
	"hash/fnv"
	"sync"
)

var locks []sync.Mutex

func InitLocks() {
	locks = make([]sync.Mutex, 1000, 1000)
}

func Lock(key string) {
	h := fnv.New64()
	h.Write([]byte(key))
	locks[h.Sum64()%1000].Lock()
}

func UnLock(key string) {
	h := fnv.New64()
	h.Write([]byte(key))
	locks[h.Sum64()%1000].Unlock()
}
