package service

type PackageService interface {
	HandleOrder(order int) map[int]int
	SetPackageSizes(newSizes []int)
}
