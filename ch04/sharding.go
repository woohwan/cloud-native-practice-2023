package ch04

import (
	"crypto/sha1"
	"sync"
)

type Shard struct {
	sync.RWMutex                        // sync.RWMutex를 구성
	m            map[string]interface{} // m은 샤드의 데이터를 갖고 있습니다.
}

type ShardedMap []*Shard // ShardedMap은 *Shard 슬라이스 타입

func NewShardedMap(nshards int) ShardedMap {
	shards := make([]*Shard, nshards) // *Shards slice를 초기화

	for i := 0; i < nshards; i++ {
		shard := make(map[string]interface{})
		shards[i] = &Shard{m: shard}
	}
	return shards
}

func (m ShardedMap) getShardIndex(key string) int {
	checksum := sha1.Sum([]byte(key)) // crypto/sha1 패키지의 Sum method 사용
	hash := int(checksum[17])         //해시로 사용할 임의의 바이트를 고릅니다.
	return hash % len(m)              // 인덱스를 얻기 위해 len(m)으로 나머지를 구합니다.
}

func (m ShardedMap) getShard(key string) *Shard {
	index := m.getShardIndex(key)
	return m[index]
}

// Delete
func (m ShardedMap) Delete(key string) {
	shard := m.getShard(key)
	shard.Lock()
	defer shard.Unlock()

	delete(shard.m, key)
}

// Get retrieve
func (m ShardedMap) Get(key string) interface{} {
	shard := m.getShard(key)
	shard.RLock()
	defer shard.RUnlock()
	return shard.m[key]
}

func (m ShardedMap) Set(key string, value interface{}) {
	shard := m.getShard(key)
	shard.Lock()
	defer shard.Unlock()

	shard.m[key] = value
}

// keys
func (m ShardedMap) Keys() []string {
	keys := make([]string, 0)
	mutex := sync.Mutex{}

	wg := sync.WaitGroup{}
	wg.Add(len(m))

	for _, shard := range m {
		go func(s *Shard) {
			s.RLock()
			for key := range s.m {
				mutex.Lock()
				keys = append(keys, key)
				mutex.Unlock()
			}
			s.RUnlock()
			wg.Done()
		}(shard)
	}
	wg.Wait() // Block until all reads are done

	return keys
}
