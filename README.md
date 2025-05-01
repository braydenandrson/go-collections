# Collections

A Go library providing useful collection types and utilities, starting with a generic Counter implementation.

## Counter

The `Counter` type is a generic implementation of a counting container, similar to Python's `collections.Counter`. It provides a convenient way to count occurrences of comparable items.

### Features

- Generic implementation that works with any comparable type
- Thread-safe operations
- Rich set of operations for counting and manipulation
- Similar interface to Python's Counter

### Usage

```go
import "github.com/yourusername/collections"

// Create a new counter
counter := collections.NewCounter[string]()

// Add items
counter.Add("apple")
counter.AddN("banana", 3)

// Get counts
count := counter.Get("apple")  // returns 1

// Get most common items
mostCommon := counter.MostCommon(2)

// Update with another counter or slice
otherCounter := collections.NewCounter[string]()
otherCounter.Add("apple")
counter.Update(otherCounter)

// Get all elements
elements := counter.Elements()

// Get total count
total := counter.Total()
```

### Available Methods

#### Basic Operations
- `NewCounter[T comparable](input ...any) *Counter[T]` - Create a new counter, optionally initialized with a slice or map
- `Add(item T)` - Increment count for an item by 1
- `AddN(item T, n int)` - Increment count for an item by n
- `Get(item T) int` - Get count for an item
- `Set(item T, n int)` - Set count for an item
- `Delete(item T)` - Remove an item
- `Clear()` - Remove all items
- `Items() map[T]int` - Get all items and their counts
- `Total() int` - Get total count of all items
- `Elements() []T` - Get all elements as a slice
- `MostCommon(n int) []Pair[T]` - Get n most common items

#### Counter Operations
- `Update(input any)` - Update counts from another counter, slice, or map
- `Subtract(input any)` - Subtract counts from another counter, slice, or map
- `Copy() *Counter[T]` - Create a copy of the counter
- `AddCounter(other *Counter[T]) *Counter[T]` - Add two counters
- `SubCounter(other *Counter[T]) *Counter[T]` - Subtract two counters
- `AndCounter(other *Counter[T]) *Counter[T]` - Intersection of two counters
- `OrCounter(other *Counter[T]) *Counter[T]` - Union of two counters

#### Comparison Operations
- `Equals(other *Counter[T]) bool` - Check if two counters are equal
- `IsSubsetOf(other *Counter[T]) bool` - Check if counter is a subset
- `IsSupersetOf(other *Counter[T]) bool` - Check if counter is a superset
- `Positive() *Counter[T]` - Get counter with only positive counts
- `Negative() *Counter[T]` - Get counter with only negative counts

### Type Parameters

- `T comparable` - The type of items to count. Must be comparable (can be used as a map key)

### Examples

#### Basic Counting
```go
counter := NewCounter[string]()
counter.Add("apple")
counter.Add("banana")
counter.Add("apple")
fmt.Println(counter.Get("apple"))  // Output: 2
```

#### Most Common Items
```go
counter := NewCounter[string]()
counter.AddN("apple", 3)
counter.AddN("banana", 2)
counter.AddN("orange", 1)
common := counter.MostCommon(2)
// common will contain [{"apple", 3}, {"banana", 2}]
```

#### Counter Operations
```go
c1 := NewCounter[string]()
c1.AddN("apple", 3)
c1.AddN("banana", 2)

c2 := NewCounter[string]()
c2.AddN("apple", 1)
c2.AddN("orange", 2)

result := c1.AddCounter(c2)
// result will have: apple: 4, banana: 2, orange: 2
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the LICENSE file for details. 