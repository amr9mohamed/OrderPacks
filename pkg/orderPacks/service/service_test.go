package service

import (
	"github.com/go-playground/assert/v2"
	"orderPacks/pkg/orderPacks/repository"
	"reflect"
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

func TestHandleOrder(t *testing.T) {
	t.Run("HandleOrder", func(t *testing.T) {
		r := repository.NewMemDB([]int{250, 500, 1000, 2000, 5000})
		s := NewPacksManager(r)
		type args struct {
			items int
		}
		tests := []struct {
			name   string
			args   args
			expect map[int]int
		}{
			{
				name: "1",
				args: args{
					items: 1,
				},
				expect: map[int]int{250: 1},
			},
			{
				name: "250",
				args: args{
					items: 250,
				},
				expect: map[int]int{250: 1},
			},
			{
				name: "251",
				args: args{
					items: 251,
				},
				expect: map[int]int{500: 1},
			},
			{
				name: "501",
				args: args{
					items: 501,
				},
				expect: map[int]int{500: 1, 250: 1},
			},
			{
				name: "12001",
				args: args{
					items: 12001,
				},
				expect: map[int]int{5000: 2, 2000: 1, 250: 1},
			},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if got := s.HandleOrder(tt.args.items); !reflect.DeepEqual(got, tt.expect) {
					t.Errorf("HandleOrder() returned %v, expected %v", got, tt.expect)
				}
			})
		}
	})
}
