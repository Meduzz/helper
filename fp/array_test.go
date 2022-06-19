package fp

import (
	"fmt"
	"strings"
	"testing"
)

func Test_Map(t *testing.T) {
	subject := Of(1, 2, 3, 4, 5)

	result := subject.Map(func(it int) K {
		return fmt.Sprintf("%d", it)
	})

	if len(result) != 5 {
		t.Errorf("the length does not match")
	}

	if result[0] != "1" {
		t.Errorf("first value was not string one")
	}
}

func Test_Filter(t *testing.T) {
	subject := Of(1, 2, 3, 4, 5)

	result := subject.Filter(func(it int) bool {
		return it%2 == 1
	})

	if len(result) != 3 {
		t.Errorf("the length does not match, was %d", len(result))
	}

	if result[0] != 1 {
		t.Errorf("the first element was not what the test expected, was %d", result[0])
	}
}

func Test_FlatMap(t *testing.T) {
	subject := Of("hello", "world")

	result := subject.FlatMap(func(it string) []K {
		res := make([]K, 0)
		split := strings.Split(it, "")

		for _, val := range split {
			res = append(res, K(val))
		}

		return res
	})

	if len(result) != 10 {
		t.Errorf("the length does not match the expected, was: %d", len(result))
	}
}

func Test_Fold(t *testing.T) {
	subject := Of(1, 2, 3, 4, 5)

	result := subject.Fold(0, func(it int, agg K) K {
		ffs, ok := agg.(int)

		if !ok {
			ffs = 0
		}

		return ffs + it
	})

	if result != 15 {
		t.Errorf("the result was not 15, but: %d", result)
	}
}

func Test_FreestandingMap(t *testing.T) {
	subject := Of(1, 2, 3, 4, 5)

	result := Map(subject, func(val int) int {
		return val * val
	})

	if len(result) != 5 {
		t.Errorf("length of result was not 5, but: %d", len(result))
	}

	if result[0] != 1 {
		t.Errorf("first result was not 1, but: %d", result[0])
	}

	if result[4] != 25 {
		t.Errorf("last result was not 25, but: %d", result[4])
	}
}

func Test_FreestandingFlatMap(t *testing.T) {
	subject := Of("hello", "world")

	result := FlatMap(subject, func(it string) []string {
		return strings.Split(it, "")
	})

	if len(result) != 10 {
		t.Errorf("the length does not match the expected, was: %d", len(result))
	}
}

func Test_FreestandingFold(t *testing.T) {
	subject := Of(1, 2, 3, 4, 5)

	result := Fold(subject, 0, func(it int, agg int) int {
		return agg + it
	})

	if result != 15 {
		t.Errorf("result was not 15, but: %d", result)
	}
}

func Test_ForEach(t *testing.T) {
	subject := Of(1, 1, 1, 1, 1)
	count := 0

	subject.ForEach(func(i int) {
		count++
	})

	if count != 5 {
		t.Errorf("count was not 5, but: %d", count)
	}
}

func Test_Concat(t *testing.T) {
	subject1 := Of(1, 2, 3, 4, 5)
	subject2 := Of(6, 7, 8, 9, 10)

	result := subject1.Concat(subject2)

	if len(result) != 10 {
		t.Errorf("length of result was not 10, but: %d", len(result))
	}
}
