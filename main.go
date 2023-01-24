package main

import (
	"fmt"

	"github.com/daniial79/data-structures-GO/dataStructures/hashTable"
)

func main() {
	mht := hashTable.GenHashTable()

	fmt.Printf("%+v\n", mht)

	mht.Set(hashTable.Pair{
		Key:   "nails",
		Value: 1000,
	})

	mht.Set(hashTable.Pair{
		Key:   "bolts",
		Value: 750,
	})

	mht.Set(hashTable.Pair{
		Key:   "pins",
		Value: 500,
	})

	fmt.Printf("%+v\n", mht)

	fmt.Printf("%#v\n", mht.Keys())

}
