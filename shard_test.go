package gorp

import (
	"strconv"
	"testing"
)

func TestHashShard(t *testing.T) {

	shard := &HashShard{ShardNum: 10}
	t.Logf(strconv.Itoa(shard.FindForKey("123")))

}
