package main

//planned update of pages in manual
type update struct {
	//flag if the update is valid for further processing
	valid bool
	//pages in update represented as a map of page and it's place(index) in update -> [pageNr]index
	indexedPages map[int]int
	//raw page numbers
	pages     []int
	initPages []int
	//if invalid update was fixed
	fixed          bool
	recursionDepth int
}

func newEmptyUpdate() update {
	return update{
		valid:          true,
		fixed:          false,
		indexedPages:   make(map[int]int),
		pages:          make([]int, 0),
		initPages:      make([]int, 0),
		recursionDepth: 0,
	}
}

func (u update) getMiddlePageNumber() int {
	if len(u.pages)%2 == 0 {
		log(u.pages)
		panic("what is the middle?!")
	}
	middle := (len(u.pages) / 2)
	return u.pages[middle]
}

func (u *update) fixPageIndex(pageNr int, newPlace int) {
	switch newPlace >= 0 {
	case true:
		oldPlace := u.indexedPages[pageNr]
		u.indexedPages[pageNr] = newPlace
		for page, index := range u.indexedPages {
			if page == pageNr {
				continue
			}
			if index >= newPlace && index < oldPlace {
				u.indexedPages[page] = index + 1
			}
			// } else if index > oldPlace && index < newPlace {
			// 	u.indexedPages[page] = index - 1
			// }
		}
	case false:
		// u.indexedPages[pageNr] = 0
		// for page, index := range u.indexedPages {
		// 	if page == pageNr {
		// 		continue
		// 	}
		// 	u.indexedPages[page] = index + 1
		// }
	}
}

func (u *update) refreshPages() {
	for page, index := range u.indexedPages {
		u.pages[index] = page
	}
}
