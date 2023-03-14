package h2tp

import (
	"fmt"

	"github.com/zzztttkkk/faceless/utils"
	"github.com/zzztttkkk/faceless/utils/slices"
)

type KvPair struct {
	idx int
	key []byte
	val []byte
	ok  bool
}

type KvPairs struct {
	ary       []KvPair
	deadIdxes []int
}

func (pairs *KvPairs) add(key []byte, val []byte) {
	if !slices.IsEmpty(pairs.deadIdxes) {
		lastIdx := slices.Pop(&pairs.deadIdxes)
		pair := &(pairs.ary[lastIdx])
		pair.key = append(pair.key, key...)
		pair.val = append(pair.val, val...)
		pair.ok = true
		return
	}

	if pairs.ary == nil {
		pairs.ary = make([]KvPair, 0, _MaxPairs)
	}

	ac := len(pairs.ary) + 1
	pairs.ary = pairs.ary[0:ac]
	pair := &(pairs.ary[ac-1])
	pair.idx = ac - 1
	pair.key = append(pair.key, key...)
	pair.val = append(pair.val, val...)
	pair.ok = true
}

type Headers struct {
	pairs KvPairs
	mmap  map[string][][]byte
}

const _MaxPairs = 16

func (h *Headers) addPairToMap(pair *KvPair) {
	key := utils.B2s(pair.key)
	h.mmap[key] = append(h.mmap[key], pair.val)
}

func (h *Headers) initMmap() {
	h.mmap = make(map[string][][]byte, _MaxPairs<<1)
	for i := 0; i < len(h.pairs.ary); i++ {
		pair := &(h.pairs.ary[i])
		if !pair.ok {
			continue
		}
		h.addPairToMap(pair)
	}
}

func (h *Headers) Add(key []byte, val []byte) {
	if len(key) < 1 || len(val) < 1 {
		panic(fmt.Errorf("empty key/value"))
	}

	if len(h.pairs.ary) < _MaxPairs {
		h.pairs.add(key, val)

		if len(h.pairs.ary) >= _MaxPairs {
			h.initMmap()
		}
		return
	}

	ks := string(key)
	tmp := make([]byte, len(val))
	copy(tmp, val)
	h.mmap[ks] = append(h.mmap[ks], tmp)
}

func (h *Headers) Del(key []byte) {
	if h.mmap != nil {
		delete(h.mmap, utils.B2s(key))
		return
	}
}
