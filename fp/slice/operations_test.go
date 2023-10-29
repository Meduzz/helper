package slice

import (
	"testing"
)

var subject = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

func TestPartition(t *testing.T) {
	result := Partition(subject, 2)

	if len(result) != 5 {
		t.Errorf("result of partition was not 5 but %d", len(result))
	}
}

func TestLengthOfSkip(t *testing.T) {
	result := Skip(subject, 6)

	if len(result) != 4 {
		t.Errorf("len of result was not 4 but %d", len(result))
	}
}

func TestLengthOfTake(t *testing.T) {
	result := Take(subject, 6)

	if len(result) != 6 {
		t.Errorf("len of result was not 6 but %d", len(result))
	}
}

func TestSharding(t *testing.T) {
	result := Shard(subject, func(i int) int64 {
		return int64(i % 2)
	})

	if len(result) > 2 {
		t.Errorf("len of result was not 2 but %d", len(result))
	}
}

func TestFilter(t *testing.T) {
	result := Filter(subject, func(i int) bool {
		return i%2 == 0
	})

	if len(result) != 5 {
		t.Errorf("result was not 5 but %d", len(result))
	}
}

func TestHeadTailBounds(t *testing.T) {
	subject := make([]string, 0)
	partial := make([]string, 0)
	partial = append(partial, "first")

	headResult := Head(subject)

	if headResult != "" {
		t.Error("headResult was not empty")
	}

	tailResult := Tail(subject)

	if len(tailResult) > 0 {
		t.Error("tailResult was not empty")
	}

	tailResult = Tail(partial)

	if len(tailResult) > 0 {
		t.Error("tailResult of partial was not empty")
	}

	takeResult := Take(subject, 1)

	if len(takeResult) > 0 {
		t.Error("takeResult was not empty")
	}

	takeResult = Take(partial, 2)

	if len(takeResult) != 1 {
		t.Errorf("takeResult of partial was not 1 but %d", len(takeResult))
	}

	skipResult := Skip(subject, 1)

	if len(skipResult) > 0 {
		t.Error("skipResult was not empty")
	}

	skipResult = Skip(partial, 1)

	if len(skipResult) > 0 {
		t.Error("skipResult of partial was not empty")
	}
}
