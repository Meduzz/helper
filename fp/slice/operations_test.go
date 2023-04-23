package slice

import "testing"

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
