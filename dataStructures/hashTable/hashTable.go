package hashTable

type Pair struct {
	Key   string
	Value int
}

type HashTable struct {
	Buckets [7][]Pair
}

func GenHashTable() *HashTable {
	return &HashTable{}
}

func (h *HashTable) hash(key string) int {
	hash := 0

	for _, charByte := range key {
		hash = (hash + int(charByte)*23) % len(h.Buckets)
	}

	return hash
}

func (h *HashTable) Set(pair Pair) {
	index := h.hash(pair.Key)
	h.Buckets[index] = append(h.Buckets[index], pair)
}

func (h *HashTable) Get(key string) (int, bool) {
	index := h.hash(key)

	for _, pair := range h.Buckets[index] {
		if pair.Key == key {
			return pair.Value, true
		}
	}

	return -1, false
}
