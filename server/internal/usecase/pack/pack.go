package pack_usecase

import (
	"sort"
)

// Usecase responsible for packing items.
type Usecase struct {
	PackSizes []uint64
}

// New returns a new Usecase.
func New(sizes []uint64) *Usecase {
	sort.Slice(sizes, func(i, j int) bool {
		return sizes[i] > sizes[j]
	})

	return &Usecase{sizes}
}

// GetPacksNumber returns the number of packs of each size required to.
func (uc *Usecase) GetPacksNumber(items uint64) map[uint64]uint64 {
	packs := make(map[uint64]uint64, len(uc.PackSizes))

	// Divide items into corresponding size packs.
	for _, packSize := range uc.PackSizes {
		if items >= packSize {
			packs[packSize] = items / packSize
			items %= packSize
		}
	}

	// If there are any remaining items, put them in the smallest pack size.
	if items > 0 {
		packs[uc.PackSizes[len(uc.PackSizes)-1]]++
	}

	// Combine any adjacent packs of the same size.
	for i := len(uc.PackSizes) - 1; i > 0; i-- {
		packSize := uc.PackSizes[i]
		if value, ok := packs[packSize]; ok {
			if value > 1 && packSize*2 == uc.PackSizes[i-1] {
				packs[packSize] -= 2
				packs[uc.PackSizes[i-1]]++
			}

			if packs[packSize] == 0 {
				delete(packs, packSize)
			}
		}
	}

	return packs
}
