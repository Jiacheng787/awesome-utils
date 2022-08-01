package promise

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestSyncCallback(t *testing.T) {
	p := New(func(resolve func(T), reject func(error)) {
		resolve(2333)
	})
	p.Then(func(res T) T {
		if res == 2333 {
			t.Logf("sync callback success")
		} else {
			t.Fatalf("sync callback failed, expected %v, got %v", 2333, res)
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
		t.Fatalf("sync multi callback failed, expected %v, got %v", 2, count)
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
		t.Fatalf("sync chain callback failed, expected %v, got %v", 2, count)
	}
}

func TestAsyncCallback(t *testing.T) {
	wg := new(sync.WaitGroup)
	p := New(func(resolve func(T), reject func(error)) {
		wg.Add(1)
		go func() {
			defer wg.Done()
			time.Sleep(1 * time.Second)
			resolve(2333)
		}()
	})

	p.Then(func(res T) T {
		t.Logf("async callback success")
		return nil
	}, func(err error) T {
		t.Fatalf("async callback failed")
		return nil
	})

	wg.Wait()
}

func TestNestPromise(t *testing.T) {
	p := New(func(resolve func(T), reject func(error)) {
		resolve(2333)
	})
	p.Then(func(res T) T {
		return New(func(resolve func(T), reject func(error)) {
			resolve(666)
		})
	}, func(err error) T {
		return nil
	}).Then(func(res T) T {
		if res == 666 {
			t.Logf("nest promise success")
		} else {
			t.Fatalf("nest promise failed, expected %v, got %v", 666, res)
		}
		return nil
	}, func(err error) T {
		t.Fatalf("nest promise failed")
		return nil
	})
}
