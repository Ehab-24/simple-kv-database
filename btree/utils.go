package btree

import "log"

func assert(predicate bool) {
	if predicate {
		return
	}
	log.Fatal("assertion failed")
}
