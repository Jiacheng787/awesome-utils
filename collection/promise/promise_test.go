package promise

import (
	"fmt"
	"testing"
)

func TestSyncCallback(t *testing.T) {
	p := New(func(resolve func(T), reject func(error)) {
		resolve(2333)
	})
	p.Then(func(res T) T {
		if res == 2333 {
			t.Logf("sync callback success")
		} else {
			t.Fatalf("sync callback failed, expected 2333, got %v", res)
		}
		return nil
	}, func(err error) T {
		t.Fatalf("sync callback failed")
		return nil
	})
}

func TestSyncError(t *testing.T) {
	p := New(func(resolve func(T), reject func(error)) {
		reject(fmt.Errorf("some error"))
	})
	p.Then(func(res T) T {
		t.Fatalf("sync error failed")
		return nil
	}, func(err error) T {
		t.Logf("sync error success")
		return nil
	})
}

func TestSyncMultiCallback(t *testing.T) {
	p := New(func(resolve func(T), reject func(error)) {
		resolve(2333)
	})

	count := 0

	p.Then(func(res T) T {
		count++
		return nil
	}, func(err error) T {
		return nil
	})

	p.Then(func(res T) T {
		count++
		return nil
	}, func(err error) T {
		return nil
	})

	if count == 2 {
		t.Logf("sync multi callback success")
	} else {
		t.Fatalf("sync multi callback failed, expected 2, got %v", count)
	}
}

func TestSyncChainCallback(t *testing.T) {
	p := New(func(resolve func(T), reject func(error)) {
		resolve(2333)
	})

	count := 0

	p.Then(func(res T) T {
		count++
		return nil
	}, func(err error) T {
		return nil
	}).Then(func(res T) T {
		count++
		return nil
	}, func(err error) T {
		return nil
	})

	if count == 2 {
		t.Logf("sync chain callback success")
	} else {
		t.Fatalf("sync chain callback failed, expected 2, got %v", count)
	}
}
