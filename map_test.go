// map_test.go. Solution by Ilya Pokolev

package scripts
import "testing"
import "sync"
import "sync/atomic"
import "math/rand"

type Map struct {
	data map[int]int
	mu   sync.Mutex
	naccesses atomic.Int64
	nadditions atomic.Int64
}

func NewMap() *Map {
	return &Map{ make(map[int]int ), sync.Mutex{}, atomic.Int64{}, atomic.Int64{} }
}

func (m *Map) update(key int, existing int, val int) bool {
	m.mu.Lock()
	defer m.mu.Unlock()
	v, found := m.data[key]
	if !found {
		if existing != -1 { return false }
		m.nadditions.Add(1)
	} else {
		if existing != v { return false }
		m.naccesses.Add(1)
	}
	m.data[key] = val
	return true
}

func (m *Map) get(key int) (val int, found bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	val, found = m.data[key]
	return
}

func (m *Map) accesses() int64 { return m.naccesses.Load() }
func (m *Map) additions() int64 { return m.nadditions.Load() }

func (m *Map) len() int {
	m.mu.Lock()
	defer m.mu.Unlock()
	return len(m.data)
}

func do(m *Map, cnt int) {
	var indexes = make([]int, cnt)
	for i := 0; i < cnt; i++ { indexes[i] = i + 1 }
	// shuffle indexes
	for i := 0; i < cnt / 2; i++ {
		var j = rand.Intn(cnt)
		var k = rand.Intn(cnt)
		for k == j { k = rand.Intn(cnt) }
		indexes[j], indexes[k] = indexes[k], indexes[j]
	}
	for _, i := range indexes {
		var updated = false
		for !updated {
			val, found := m.get(i)
			if !found { val = -1 }
			updated = m.update(i, val, val + 1)
		}
	}
}

func Test(t *testing.T) {
	var m = NewMap()
	var n = 2009 // 3rd of January 2009 Bitcoin network started to work, genesis block were generated
	var wg sync.WaitGroup
	go func() { do(m, n); wg.Done() }()
	wg.Add(1)
	go func() { do(m, n); wg.Done() }()
	wg.Add(1)
	go func() { do(m, n); wg.Done() }()
	wg.Add(1)
	go func() { do(m, n); wg.Done() }()
	wg.Add(1)
	wg.Wait()
	if m.len() != n {
		t.Fatal("incorrect size", m.len())
	}
	for i := 1; i <= n; i++ {
		if v, found := m.get(i); !found || v != 3 {
			t.Fatal("not found or incorrect value", v, "for key", i)
		}
	}
	if m.additions() != int64(n) {t.Fatal("invalid additions count", m.additions()) }
	if m.accesses() != int64(n * 3) {t.Fatal("invalid accesses count", m.accesses()) }
}
