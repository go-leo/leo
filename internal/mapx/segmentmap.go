package mapx

import "golang.org/x/exp/maps"

const maxShards = 65536

type segment[K comparable, V any] struct {
	m Map[K, V]
}

type SegmentMap[K comparable, V any] struct {
	hasher   Hash[K]
	segments []*segment[K, V]
	shards   int
}

func NewSegmentMap[K comparable, V any](initSize int, shards int, hasher Hash[K]) *SegmentMap[K, V] {
	segments := make([]*segment[K, V], 0, shards)
	for i := 0; i < shards; i++ {
		stdMap := NewStdMap[K, V](initSize)
		seg := &segment[K, V]{m: stdMap}
		segments = append(segments, seg)
	}
	if shards > maxShards {
		shards = maxShards
	}
	return &SegmentMap[K, V]{hasher: hasher, segments: segments, shards: shards}
}

func (m *SegmentMap[K, V]) getSegment(key K) *segment[K, V] {
	return m.segments[int(m.hasher.Sum(key))%m.shards]
}

func (m *SegmentMap[K, V]) Load(key K) (value V, ok bool) {
	return m.getSegment(key).m.Load(key)
}

func (m *SegmentMap[K, V]) Store(key K, value V) {
	m.getSegment(key).m.Store(key, value)
}

func (m *SegmentMap[K, V]) LoadOrStore(key K, value V) (actual V, loaded bool) {
	return m.getSegment(key).m.LoadOrStore(key, value)
}

func (m *SegmentMap[K, V]) LoadAndDelete(key K) (value V, loaded bool) {
	return m.getSegment(key).m.LoadAndDelete(key)
}

func (m *SegmentMap[K, V]) Delete(key K) {
	m.getSegment(key).m.Delete(key)
}

func (m *SegmentMap[K, V]) DeleteFunc(del func(K, V) bool) {
	for _, seg := range m.segments {
		seg.m.DeleteFunc(del)
	}
}

func (m *SegmentMap[K, V]) Map() map[K]V {
	var res map[K]V
	for _, seg := range m.segments {
		maps.Copy(res, seg.m.Map())
	}
	return res
}

func (m *SegmentMap[K, V]) Len() int {
	var length int
	for _, seg := range m.segments {
		length += seg.m.Len()
	}
	return length
}

func (m *SegmentMap[K, V]) Keys() []K {
	var res []K
	for _, seg := range m.segments {
		res = append(res, seg.m.Keys()...)
	}
	return res
}

func (m *SegmentMap[K, V]) Values() []V {
	var res []V
	for _, seg := range m.segments {
		res = append(res, seg.m.Values()...)
	}
	return res
}

func (m *SegmentMap[K, V]) Clear() {
	for _, seg := range m.segments {
		seg.m.Clear()
	}
}

func (m *SegmentMap[K, V]) Clone() Map[K, V] {
	segments := make([]*segment[K, V], m.shards)
	for i, seg := range m.segments {
		segments[i] = &segment[K, V]{
			m: seg.m.Clone(),
		}
	}
	return &SegmentMap[K, V]{
		hasher:   m.hasher,
		segments: segments,
		shards:   m.shards,
	}
}
