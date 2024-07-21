package repository

type MemDB struct {
	Sizes []int
}

func NewMemDB(sizes []int) *MemDB {
	return &MemDB{
		Sizes: sizes,
	}
}

func (db *MemDB) SetPackageSizes(newSizes []int) {
	db.Sizes = newSizes
}

func (db *MemDB) GetAllPackageSizes() []int {
	return db.Sizes
}
