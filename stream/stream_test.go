package stream

import (
	"reflect"
	"testing"
)

func TestStream(t *testing.T) {
	t.Run("Contains", func(t *testing.T) {
		stream := Of(1, 3, 5, 7, 9)
		hasOne := stream.Contains(1)
		hasTwo := stream.Contains(2)

		if !hasOne || hasTwo {
			t.Errorf("Expected true, false, got %v, %v", hasOne, hasTwo)
		}
	})

	t.Run("ForEach", func(t *testing.T) {
		stream := Of(1, 3, 5, 7, 9)
		var result []int
		stream.ForEach(func(elem int) {
			result = append(result, elem)
		})
		if !reflect.DeepEqual(result, []int{1, 3, 5, 7, 9}) {
			t.Errorf("Expected [1, 3, 5, 7, 9], got %v", result)
		}
	})

	t.Run("RemoveIndex", func(t *testing.T) {
		stream := Of(1, 2, 3, 4, 5).RemoveIndex(2)
		if stream.Contains(3) {
			t.Error("Stream should not contain 3")
		}
	})

	t.Run("Filter", func(t *testing.T) {
		stream := Of(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
		stream.Filter(func(elem int) bool {
			return elem%2 == 0
		}).ForEach(func(elem int) {
			if elem%2 != 0 {
				t.Errorf("Expected even number, got %v", elem)
			}
		})
	})

	t.Run("Count", func(t *testing.T) {
		stream := Of(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
		count := stream.Count(func(elem int) bool {
			return elem%2 == 0
		})
		if count != 5 {
			t.Errorf("Expected 5, got %v", count)
		}
	})

	t.Run("NotEmpty", func(t *testing.T) {
		if Of[any]().NotEmpty() {
			t.Error("Stream should be empty")
		} else if !Of(1).NotEmpty() {
			t.Error("Stream should not be empty")
		}
	})

	t.Run("Last", func(t *testing.T) {
		stream := Of(1, 2, 3, 4, 5)
		last := stream.Last(func(i int) bool {
			return i%2 == 0
		})
		if last != 4 {
			t.Errorf("Expected 4, got %v", last)
		}
	})

	t.Run("Slice", func(t *testing.T) {
		stream := Of(1, 2, 3, 4, 5).Slice([]int{1, 3})
		correct := []int{2, 4}
		if !reflect.DeepEqual(stream.Values(), correct) {
			t.Errorf("Expected %v, got %v", correct, stream.Values())
		}
	})

	t.Run("All", func(t *testing.T) {
		streamOne := Of(1, 1, 1, 1, 1, 1, 1, 1, 1, 1)
		streamTwo := Of(2, 1, 5, 7, 2, 8, 2, 5, 3, 9)

		allOnes := streamOne.All(func(elem int) bool {
			return elem == 1
		})

		if !allOnes {
			t.Error("Expected true, got false")
		}

		allTwos := streamTwo.All(func(elem int) bool {
			return elem == 1
		})

		if allTwos {
			t.Error("Expected true, got false")
		}
	})

	t.Run("Any", func(t *testing.T) {
		stream := Of(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
		a := stream.Any(func(elem int) bool {
			return elem == 0 || elem == 11
		})

		if a {
			t.Error("Expected false, got true")
		}
	})

	t.Run("Find", func(t *testing.T) {
		stream := Of(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
		found := stream.Find(func(elem int) bool {
			return elem == 0
		})
		if found != nil && *found != 0 {
			t.Errorf("Expected 0, got %v", found)
		}
	})

	t.Run("TakeLast", func(t *testing.T) {
		stream := Of(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11).TakeLast(5)
		correct := []int{7, 8, 9, 10, 11}
		if !reflect.DeepEqual(stream.Values(), correct) {
			t.Errorf("Expected %v, got %v", correct, stream.Values())
		}
	})

	t.Run("DropLast", func(t *testing.T) {
		stream := Of(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11).DropLast(5)
		correct := []int{1, 2, 3, 4, 5, 6}
		if !reflect.DeepEqual(stream.Values(), correct) {
			t.Errorf("Expected %v, got %v", correct, stream.Values())
		}
	})

	t.Run("Take", func(t *testing.T) {
		stream := Of(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11).Take(5)
		correct := []int{1, 2, 3, 4, 5}
		if !reflect.DeepEqual(stream.Values(), correct) {
			t.Errorf("Expected %v, got %v", correct, stream.Values())
		}
	})
}
