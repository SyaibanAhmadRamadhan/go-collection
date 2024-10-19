package generic

import "sync"

type SafeMap[KEY any, VALUE any] struct {
	m sync.Map
}

func (s *SafeMap[KEY, VALUE]) Store(key KEY, value VALUE) {
	s.m.Store(key, value)
}

func (s *SafeMap[KEY, VALUE]) Load(key KEY) (VALUE, bool) {
	var zero VALUE
	v, ok := s.m.Load(key)
	if !ok {
		return zero, false
	}
	intVal, ok := v.(VALUE)
	return intVal, ok
}

func (s *SafeMap[KEY, VALUE]) Delete(key string) {
	s.m.Delete(key)
}

func (s *SafeMap[KEY, VALUE]) Range(f func(key KEY, value VALUE) bool) {
	s.m.Range(func(k, v any) bool {
		key, okKey := k.(KEY)
		value, okValue := v.(VALUE)
		if okKey && okValue {
			return f(key, value)
		}
		return true
	})
}
