package queue

import (
	"container/list"
	"testing"
)

func verifyLength(t *testing.T, q *Queue, len int) {
	if q.length != len {
		t.Errorf("Queue length: result-%d, expected-%d", q.length, len)
	}

	if l := q.end - q.start + 1; l != len {
		t.Errorf("Queue length: result-%d, expected-%d", l, len)
	}
}

func verifyEmpty(t *testing.T, q *Queue) {
	if q.length != 0 {
		t.Errorf("Queue empty: result-nonempty, expected-empty")
	}
}

func verifyErrNil(t *testing.T, err error) {
	if err != nil {
		t.Errorf("Error check: result-not nil, expected-nil")
	}
}

func verifyErrEqual(t *testing.T, err error, compare error) {
	if err != compare {
		t.Errorf("Error Equal: result-false , expected-true")
	}
}

func TestNew(t *testing.T) {
	q := NewQueue()
	verifyEmpty(t, q)
}

func TestAdd(t *testing.T) {
	q := NewQueue()
	for i := 1; i <= 10; i++ {
		q.Add(i)
		verifyLength(t, q, i)
	}
}

func TestPoll(t *testing.T) {
	q := NewQueue()
	q.Add(1)
	poll := q.Poll()
	if poll != 1 {
		t.Errorf("Queue poll: result-%d, expect-%d", poll, 1)
	}
	verifyEmpty(t, q)

	for i := 2; i <= 5; i++ {
		q.Add(i)
	}

	for i := 2; i <= 5; i++ {
		poll = q.Poll()
		if poll != i {
			t.Errorf("Queue poll: result-%d, expect-%d", poll, i)
		}
		verifyLength(t, q, 5-i)
	}

	poll = q.Poll()

	if poll != nil {
		t.Errorf("Queue poll: result-%d, expect-%v", poll, nil)
	}
}

func TestRemove(t *testing.T) {
	q := NewQueue()
	q.Add(1)
	remove, err := q.Remove()
	verifyErrNil(t, err)
	if remove != 1 {
		t.Errorf("Queue poll: result-%d, expect-%d", remove, 1)
	}
	verifyEmpty(t, q)

	for i := 2; i <= 5; i++ {
		q.Add(i)
	}

	for i := 2; i <= 5; i++ {
		remove, err = q.Remove()
		verifyErrNil(t, err)
		if remove != i {
			t.Errorf("Queue poll: result-%d, expect-%d", remove, i)
		}
		verifyLength(t, q, 5-i)
	}

	remove, err = q.Remove()

	verifyErrEqual(t, err, ErrNoElement)
}

func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		queue := NewQueue()
		for n := 1; n <= 1024; n++ {
			queue.Add(n)
		}
	}
}

func BenchmarkAddByList(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var list list.List
		for n := 1; n <= 1024; n++ {
			list.PushBack(i)
		}
	}
}
