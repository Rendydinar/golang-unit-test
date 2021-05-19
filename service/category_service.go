package service

import (
	"belaja-golang-unit-test/repository"
	"belaja-golang-unit-test/entity"
 	"errors"
)

type CategoryService struct {
	Repository repository.CategoryRepository
}	

// @params service type CategoryService			   
// @function call Get @params id type string
// @return entity.Category or nil
func (service CategoryService) Get(id string) (*entity.Category, error) {
	category := service.Repository.FindById(id)
	if category == nil {
		return nil, errors.New("Category Not Found")
	} else {
		return category, nil
	}
 
}

