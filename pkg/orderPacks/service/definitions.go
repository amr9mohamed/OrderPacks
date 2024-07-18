package service

type Service interface {
	HandleOrder(order int) map[int]int
	SetPackageSizes(newSizes []int)
}
