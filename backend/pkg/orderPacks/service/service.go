package service

import (
	"orderPacks/pkg/orderPacks/repository"
	"sort"
)

type PacksManager struct {
	Repo repository.PackageRepository
}

func NewPacksManager(repo repository.PackageRepository) *PacksManager {
	return &PacksManager{
		Repo: repo,
	}
}

// SetPackageSizes method to update the package sizes based on new requests
func (p *PacksManager) SetPackageSizes(newSizes []int) {
	p.Repo.SetPackageSizes(newSizes)
}

// HandleOrder method to calculate packs needed to fulfill requested order
func (p *PacksManager) HandleOrder(orderSize int) map[int]int {
	output := make(map[int]int)

	packageSizes := p.Repo.GetAllPackageSizes()
	sort.Sort(sort.Reverse(sort.IntSlice(packageSizes)))
	smallestSize := packageSizes[len(packageSizes)-1]

	for _, ps := range packageSizes {
		curPacks := orderSize / ps
		if curPacks > 0 {
			output[ps] = curPacks
		}
		orderSize = orderSize % ps

		if ps-orderSize < smallestSize {
			output[ps]++
			orderSize = 0
		}
	}

	return output
}
