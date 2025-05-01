package collections

import (
	"sort"
)

type Counter[T comparable] struct {
	data map[T]int
}

type Pair[T comparable] struct {
	Item  T
	Count int
}

func NewCounter[T comparable](input ...any) *Counter[T] {
	c := &Counter[T]{data: make(map[T]int)}
	if len(input) == 0 {
		return c
	}
	switch v := input[0].(type) {
	case []T:
		for _, item := range v {
			c.data[item]++
		}
	case map[T]int:
		for key, val := range v {
			c.data[key] = val
		}
	}
	return c
}

func (c *Counter[T]) Add(item T) {
	c.data[item]++
}

func (c *Counter[T]) AddN(item T, n int) {
	c.data[item] += n
}

func (c *Counter[T]) Get(item T) int {
	return c.data[item]
}

func (c *Counter[T]) Set(item T, n int) {
	c.data[item] = n
}

func (c *Counter[T]) Delete(item T) {
	delete(c.data, item)
}

func (c *Counter[T]) Clear() {
	c.data = make(map[T]int)
}

func (c *Counter[T]) Items() map[T]int {
	return c.data
}

func (c *Counter[T]) Total() int {
	sum := 0
	for _, count := range c.data {
		sum += count
	}
	return sum
}

func (c *Counter[T]) Elements() []T {
	var elems []T
	for item, count := range c.data {
		if count > 0 {
			for i := 0; i < count; i++ {
				elems = append(elems, item)
			}
		}
	}
	return elems
}

func (c *Counter[T]) MostCommon(n int) []Pair[T] {
	pairs := make([]Pair[T], 0, len(c.data))
	for item, count := range c.data {
		pairs = append(pairs, Pair[T]{item, count})
	}
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].Count > pairs[j].Count
	})
	if n > 0 && n < len(pairs) {
		return pairs[:n]
	}
	return pairs
}

func (c *Counter[T]) Update(input any) {
	switch v := input.(type) {
	case []T:
		for _, item := range v {
			c.data[item]++
		}
	case map[T]int:
		for k, v2 := range v {
			c.data[k] += v2
		}
	case *Counter[T]:
		for k, v2 := range v.data {
			c.data[k] += v2
		}
	}
}

func (c *Counter[T]) Subtract(input any) {
	switch v := input.(type) {
	case []T:
		for _, item := range v {
			c.data[item]--
		}
	case map[T]int:
		for k, v2 := range v {
			c.data[k] -= v2
		}
	case *Counter[T]:
		for k, v2 := range v.data {
			c.data[k] -= v2
		}
	}
}

func (c *Counter[T]) Copy() *Counter[T] {
	newC := NewCounter[T]()
	for k, v := range c.data {
		newC.data[k] = v
	}
	return newC
}

func (c *Counter[T]) AddCounter(other *Counter[T]) *Counter[T] {
	result := c.Copy()
	for k, v := range other.data {
		result.data[k] += v
	}
	return result
}

func (c *Counter[T]) SubCounter(other *Counter[T]) *Counter[T] {
	result := NewCounter[T]()
	for k, v := range c.data {
		diff := v - other.data[k]
		if diff > 0 {
			result.data[k] = diff
		}
	}
	return result
}

func (c *Counter[T]) AndCounter(other *Counter[T]) *Counter[T] {
	result := NewCounter[T]()
	for k, v := range c.data {
		if ov, ok := other.data[k]; ok {
			min := v
			if ov < v {
				min = ov
			}
			if min > 0 {
				result.data[k] = min
			}
		}
	}
	return result
}

func (c *Counter[T]) OrCounter(other *Counter[T]) *Counter[T] {
	result := NewCounter[T]()
	for k, v := range c.data {
		result.data[k] = v
	}
	for k, v := range other.data {
		if cv, ok := result.data[k]; !ok || v > cv {
			result.data[k] = v
		}
	}
	return result
}

func (c *Counter[T]) Equals(other *Counter[T]) bool {
	allKeys := make(map[T]bool)
	for k := range c.data {
		allKeys[k] = true
	}
	for k := range other.data {
		allKeys[k] = true
	}
	for k := range allKeys {
		if c.Get(k) != other.Get(k) {
			return false
		}
	}
	return true
}

func (c *Counter[T]) IsSubsetOf(other *Counter[T]) bool {
	for k, v := range c.data {
		if v > other.Get(k) {
			return false
		}
	}
	return true
}

func (c *Counter[T]) IsSupersetOf(other *Counter[T]) bool {
	return other.IsSubsetOf(c)
}

func (c *Counter[T]) Positive() *Counter[T] {
	result := NewCounter[T]()
	for k, v := range c.data {
		if v > 0 {
			result.data[k] = v
		}
	}
	return result
}

func (c *Counter[T]) Negative() *Counter[T] {
	result := NewCounter[T]()
	for k, v := range c.data {
		if v < 0 {
			result.data[k] = -v
		}
	}
	return result
}
