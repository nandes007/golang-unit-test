package service

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"learn-golang-unit-test/entity"
	"learn-golang-unit-test/repository"
	"testing"
)

var categoryRepository = &repository.CategoryRepositoryMock{Mock: mock.Mock{}}
var categoryService = CategoryService{Repository: categoryRepository}

func TestCategoryService_GetNotFound(t *testing.T) {
	categoryRepository.Mock.On("FindById", "1").Return(nil)
	category, err := categoryService.Get("1")
	assert.NotNil(t, err)
	assert.Nil(t, category)
}

func TestCategoryService_GetSuccess(t *testing.T) {
	categoryMock := entity.Category{
		Id:   "1",
		Name: "Laptop",
	}
	categoryRepository.Mock.On("FindById", "2").Return(categoryMock)

	category, err := categoryService.Get("2")
	assert.Nil(t, err)
	assert.NotNil(t, category)
	assert.Equal(t, categoryMock.Id, category.Id)
	assert.Equal(t, categoryMock.Name, category.Name)
}
