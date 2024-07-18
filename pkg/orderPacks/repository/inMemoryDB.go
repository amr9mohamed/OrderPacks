package repository

import "sort"

type MemDB struct {
	Sizes []int
}

func NewMemDB(sizes []int) *MemDB {
	sort.Sort(sort.Reverse(sort.IntSlice(sizes)))
	return &MemDB{
		Sizes: sizes,
	}
}

func (db *MemDB) SetPackageSizes(newSizes []int) {
	sort.Sort(sort.Reverse(sort.IntSlice(newSizes)))
	db.Sizes = newSizes
}

func (db *MemDB) GetAllPackageSizes() []int {
	return db.Sizes
}
