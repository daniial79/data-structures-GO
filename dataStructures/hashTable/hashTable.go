package hashTable

type pair struct {
	key   string
	value int
}

type HashTable struct {
	Buckets [7][]pair
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
