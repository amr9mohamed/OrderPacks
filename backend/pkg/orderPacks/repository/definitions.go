package repository

// PackageRepository is an interface that need to be implemented by the repo/storage layer
type PackageRepository interface {
	SetPackageSizes([]int)
	GetAllPackageSizes() []int
}
