package ranges

import "math/big"

type IntRange[T any] struct {
	start, end int
	indices    []T
}
type BigIntRange struct {
	start, end *big.Int
	indices    []*big.Int
}
type Range[T any] interface {
	Len() int
	ForEach(f func(T))
}

// big int ranges

var bigOne = big.NewInt(1)

func NewBigIntRange(start, end *big.Int) *BigIntRange {
	result := make([]*big.Int, end.Sub(end, start).Int64())
	for i := new(big.Int).Set(start); i.Cmp(end) < 0; i.Add(i, bigOne) {
		result[new(big.Int).Sub(i, start).Int64()] = new(big.Int).Set(i)
	}
	return &BigIntRange{start, end, result}
}

func (r BigIntRange) Len() int {
	return len(r.indices)
}

func (r BigIntRange) ForEach(f func(*big.Int)) {
	for _, i := range r.indices {
		f(i)
	}
}

//int ranges

func NewIntRange(start, end int) IntRange[int] {
	result := make([]int, end-start)
	for i := start; i != end; i++ {
		result[i-start] = i
	}
	return IntRange[int]{start, end, result}
}

func NewInt8Range(start, end int) IntRange[int8] {
	result := make([]int8, end-start)
	for i := start; i != end; i++ {
		result[i-start] = int8(i)
	}
	return IntRange[int8]{start, end, result}
}

func NewInt16Range(start, end int) IntRange[int16] {
	result := make([]int16, end-start)
	for i := start; i != end; i++ {
		result[i-start] = int16(i)
	}
	return IntRange[int16]{start, end, result}
}

func NewInt32Range(start, end int) IntRange[int32] {
	result := make([]int32, end-start)
	for i := start; i != end; i++ {
		result[i-start] = int32(i)
	}
	return IntRange[int32]{start, end, result}
}

func NewInt64Range(start, end int) IntRange[int64] {
	result := make([]int64, end-start)
	for i := start; i != end; i++ {
		result[i-start] = int64(i)
	}
	return IntRange[int64]{start, end, result}
}

// Uint ranges

func NewUintRange(start, end int) IntRange[uint] {
	result := make([]uint, end-start)
	for i := start; i != end; i++ {
		result[i-start] = uint(i)
	}
	return IntRange[uint]{start, end, result}
}

func NewUint8Range(start, end, step int) IntRange[uint8] {
	result := make([]uint8, end-start)
	for i := start; i != end; i++ {
		result[i-start] = uint8(i)
	}
	return IntRange[uint8]{start, end, result}
}

func NewUint16Range(start, end int) IntRange[uint16] {
	result := make([]uint16, end-start)
	for i := start; i != end; i++ {
		result[i-start] = uint16(i)
	}
	return IntRange[uint16]{start, end, result}
}

func NewUint32Range(start, end int) IntRange[uint32] {
	result := make([]uint32, end-start)
	for i := start; i != end; i++ {
		result[i-start] = uint32(i)
	}
	return IntRange[uint32]{start, end, result}
}

func NewUint64Range(start, end int) IntRange[uint64] {
	result := make([]uint64, end-start)
	for i := start; i != end; i++ {
		result[i-start] = uint64(i)
	}
	return IntRange[uint64]{start, end, result}
}

func (r IntRange[T]) Len() int {
	return len(r.indices)
}

func (r IntRange[T]) ForEach(f func(T)) {
	for _, i := range r.indices {
		f(i)
	}
}
