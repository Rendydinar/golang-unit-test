package repository

import "belaja-golang-unit-test/entity"

type CategoryRepository interface {
	FindById(id string) *entity.Category
}