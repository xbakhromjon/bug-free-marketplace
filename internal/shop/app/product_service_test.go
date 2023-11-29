package app

import (
	"errors"
	"fmt"
	"golang-project-template/internal/shop/adapter/mock"
	"golang-project-template/internal/shop/domain"
	"reflect"
	"testing"
)

func TestGetOneProduct(t *testing.T) {
	underTest := productService{repository: mock.NewMockProductAdapter(nil)}
	t.Run("with correct id", func(t *testing.T) {
		id := 1
		got, _ := underTest.GetOne(id)
		want := domain.Product{
			Id:     id,
			Name:   "T-shirt",
			Price:  100,
			ShopId: 1,
		}
		if !reflect.DeepEqual(want, *got) {
			t.Errorf("want %+v but got %+v", want, got)
		}
	})

	t.Run("with invalid id", func(t *testing.T) {
		id := 2
		_, err := underTest.GetOne(id)
		want := domain.ErrProductNotFound{Err: fmt.Sprintf("Product not found with %d", id)}

		if err == nil {
			t.Errorf("expected %T, got nil", want)
			return
		}

		var got *domain.ErrProductNotFound
		if !errors.As(err, &got) {
			t.Errorf("want %T, got %T", want, err)
		}
	})
}
