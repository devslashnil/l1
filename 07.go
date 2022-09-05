// Реализовать конкурентную запись данных в map.

package task

import "sync"

type Map[K comparable, V any] struct {
	m  map[K]V
	mu sync.RWMutex // для блокировки на запись методов
}

// Store блокируется на запись, благодаря этому возможна лишь одна единовременная запись
func (m *Map[K, V]) Store(key K, value V) {
	m.mu.Lock()
	m.m[key] = value
	m.mu.Unlock()
}
