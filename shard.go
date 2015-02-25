package gorp

import (
	"fmt"
	"hash/crc32"
)

type Shard interface {
	FindForKey(key interface{}) int
}

type HashShard struct {
	ShardNum int
}

func (s *HashShard) FindForKey(key interface{}) int {
	h := HashValue(key)
	return int(h % uint64(s.ShardNum))
}

func HashValue(value interface{}) uint64 {
	switch val := value.(type) {
	case int:
		return uint64(val)
	case uint64:
		return uint64(val)
	case int64:
		return uint64(val)
	case string:
		return uint64(crc32.ChecksumIEEE([]byte(val)))
	case []byte:
		return uint64(crc32.ChecksumIEEE(val))
	}
	panic(fmt.Sprintf("Unexpected key variable type %T", value))
}
