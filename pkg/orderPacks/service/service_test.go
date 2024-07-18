package service

import (
	"github.com/go-playground/assert/v2"
	"orderPacks/pkg/orderPacks/repository"
	"testing"
)

func TestSetPackageSizes(t *testing.T) {
	t.Run("SetPackageSizes", func(t *testing.T) {
		// Arrange
		r := repository.NewMemDB([]int{100})
		s := NewPacksManager(r)

		// Act
		s.SetPackageSizes([]int{30, 10, 20})

		// Assert
		assert.Equal(t, r.GetAllPackageSizes(), []int{30, 20, 10})
	})
}
