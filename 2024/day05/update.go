package main

//planned update of pages in manual
type update struct {
	//flag if the update is valid for further processing
	valid bool
	//pages in update represented as a map of page and it's place(index) in update -> [pageNr]index
	indexedPages map[int]int
	//raw page numbers
	pages []int
}

func newEmptyUpdate() update {
	return update{
		valid:        true,
		indexedPages: make(map[int]int),
		pages:        make([]int, 0),
	}
}

func (u update) getMiddlePageNumber() int {
	if len(u.pages)%2 == 0 {
		panic("what is the middle?!")
	}
	middle := (len(u.pages) / 2)
	return u.pages[middle]
}
