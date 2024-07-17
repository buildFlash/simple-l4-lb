package main

import (
	"fmt"
	"hash/fnv"
	"sort"
	"sync"
)

type BalancingStrategy interface {
	Init([]*Backend)
	GetNextBackend(req IncomingReq) *Backend
	RegisterBackend(backend *Backend)
	PrintTopology()
}

type RRBalancingStrategy struct {
	Backends []*Backend
	Index    int
}

func NewRRBalancingStrategy(backends []*Backend) *RRBalancingStrategy {
	strategy := new(RRBalancingStrategy)
	strategy.Init(backends)
	return strategy
}

func (s *RRBalancingStrategy) Init(backends []*Backend) {
	s.Index = 0
	s.Backends = backends
}

func (s *RRBalancingStrategy) GetNextBackend(_ IncomingReq) *Backend {
	s.Index = (s.Index + 1) % len(s.Backends)
	return s.Backends[s.Index]
}

func (s *RRBalancingStrategy) RegisterBackend(backend *Backend) {
	s.Backends = append(s.Backends, backend)
}

func (s *RRBalancingStrategy) PrintTopology() {
	for index, backend := range s.Backends {
		fmt.Println(fmt.Sprintf("       [%d] %s", index, backend))
	}
}

type ConsistentHashingStrategy struct {
	Backends     map[uint32]*Backend
	backingArray []uint32
	vNodes       int
	mu           sync.RWMutex
}

func NewConsistentHashingStrategy(backends []*Backend, vNodes int) *ConsistentHashingStrategy {
	strategy := new(ConsistentHashingStrategy)
	strategy.vNodes = vNodes
	strategy.Init(backends)
	return strategy
}

func (s *ConsistentHashingStrategy) Init(backends []*Backend) {
	s.Backends = make(map[uint32]*Backend)
	s.backingArray = make([]uint32, 0)
	s.mu = sync.RWMutex{}
	for idx := range backends {
		s.RegisterBackend(backends[idx])
	}
}

func (s *ConsistentHashingStrategy) GetNextBackend(req IncomingReq) *Backend {
	s.mu.RLock()
	defer s.mu.RUnlock()
	if len(s.Backends) == 0 {
		return nil
	}
	hashKey := s.getHashKey(req.reqId)
	idx := s.searchKey(hashKey)
	return s.Backends[s.backingArray[idx]]
}

func (s *ConsistentHashingStrategy) RegisterBackend(backend *Backend) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	for idx := range s.vNodes {
		hashKey := s.getHashKey(fmt.Sprintf("%s:%d:%d", backend.Host, idx, backend.Port))
		s.Backends[hashKey] = backend
		s.backingArray = append(s.backingArray, hashKey)
	}
	sort.Slice(s.backingArray, func(i, j int) bool {
		return s.backingArray[i] < s.backingArray[j]
	})
}

func (s *ConsistentHashingStrategy) PrintTopology() {
	s.mu.RLock()
	defer s.mu.RUnlock()

	for idx, key := range s.backingArray {
		fmt.Println(fmt.Sprintf("       [%d] %s", idx, s.Backends[key]))
	}
}

func (s *ConsistentHashingStrategy) getHashKey(key string) uint32 {
	h := fnv.New32a()
	_, err := h.Write([]byte(key))
	if err != nil {
		panic(err)
	}
	return h.Sum32() % 10007
}

func (s *ConsistentHashingStrategy) searchKey(key uint32) int {
	idx := sort.Search(len(s.backingArray), func(i int) bool {
		return s.backingArray[i] >= key
	})
	if idx == len(s.backingArray) {
		return 0
	}
	return idx
}
