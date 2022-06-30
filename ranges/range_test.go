package ranges

import (
	"reflect"
	"testing"
)

func TestRanges(t *testing.T) {

	t.Run("IntRange", func(t *testing.T) {
		var looped []int
		NewIntRange(0, 10).ForEach(func(i int) {
			looped = append(looped, i)
		})
		if !reflect.DeepEqual(looped, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}) {
			t.Errorf("Expected [0, 1, 2, 3, 4, 5, 6, 7, 8, 9], got %v", looped)
		}
	})

	t.Run("Int16Range", func(t *testing.T) {
		var looped []int16
		NewInt16Range(0, 10).ForEach(func(i int16) {
			looped = append(looped, i)
		})
		if !reflect.DeepEqual(looped, []int16{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}) {
			t.Errorf("Expected [0, 1, 2, 3, 4, 5, 6, 7, 8, 9], got %v", looped)
		}
	})

	t.Run("Int32Range", func(t *testing.T) {
		var looped []int32
		NewInt32Range(0, 10).ForEach(func(i int32) {
			looped = append(looped, i)
		})
		if !reflect.DeepEqual(looped, []int32{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}) {
			t.Errorf("Expected [0, 1, 2, 3, 4, 5, 6, 7, 8, 9], got %v", looped)
		}
	})

	t.Run("Int8Range", func(t *testing.T) {
		var looped []int8
		NewInt8Range(0, 10).ForEach(func(i int8) {
			looped = append(looped, i)
		})
		if !reflect.DeepEqual(looped, []int8{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}) {
			t.Errorf("Expected [0, 1, 2, 3, 4, 5, 6, 7, 8, 9], got %v", looped)
		}
	})

	t.Run("Int64Range", func(t *testing.T) {
		var looped []int64
		NewInt64Range(0, 10).ForEach(func(i int64) {
			looped = append(looped, i)
		})
		if !reflect.DeepEqual(looped, []int64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}) {
			t.Errorf("Expected [0, 1, 2, 3, 4, 5, 6, 7, 8, 9], got %v", looped)
		}
	})

	t.Run("UintRange", func(t *testing.T) {
		var looped []uint
		NewUintRange(0, 10).ForEach(func(i uint) {
			looped = append(looped, i)
		})
		if !reflect.DeepEqual(looped, []uint{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}) {
			t.Errorf("Expected [0, 1, 2, 3, 4, 5, 6, 7, 8, 9], got %v", looped)
		}
	})

	t.Run("Uint8Range", func(t *testing.T) {
		var looped []uint8
		NewUint8Range(0, 10).ForEach(func(i uint8) {
			looped = append(looped, i)
		})
		if !reflect.DeepEqual(looped, []uint8{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}) {
			t.Errorf("Expected [0, 1, 2, 3, 4, 5, 6, 7, 8, 9], got %v", looped)
		}
	})

	t.Run("Uint16Range", func(t *testing.T) {
		var looped []uint16
		NewUint16Range(0, 10).ForEach(func(i uint16) {
			looped = append(looped, i)
		})
		if !reflect.DeepEqual(looped, []uint16{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}) {
			t.Errorf("Expected [0, 1, 2, 3, 4, 5, 6, 7, 8, 9], got %v", looped)
		}
	})

	t.Run("Uint32Range", func(t *testing.T) {
		var looped []uint32
		NewUint32Range(0, 10).ForEach(func(i uint32) {
			looped = append(looped, i)
		})
		if !reflect.DeepEqual(looped, []uint32{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}) {
			t.Errorf("Expected [0, 1, 2, 3, 4, 5, 6, 7, 8, 9], got %v", looped)
		}
	})

	t.Run("Uint64Range", func(t *testing.T) {
		var looped []uint64
		NewUint64Range(0, 10).ForEach(func(i uint64) {
			looped = append(looped, i)
		})
		if !reflect.DeepEqual(looped, []uint64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}) {
			t.Errorf("Expected [0, 1, 2, 3, 4, 5, 6, 7, 8, 9], got %v", looped)
		}
	})

}
