package service

import (
	"fmt"
	"log"

	"github.com/mashingan/smapping"

	"finpro/dto"
	"finpro/models"
	"finpro/repository"
)

type CategoryService interface {
	InsertCategory(b dto.CategoryCreateDto) models.Category
	UpdateCategory(b dto.CategoryUpdateDto) models.Category
	DeleteCategory(b models.Category)
	FindAllCategory() []models.Category
	FindCategoryById(categoryID uint64) models.Category
	IsAllowedToEdit(CategoryId string, Id uint64) bool
}

type categoryService struct {
	categoryRepository repository.CategoryRepository
}

func NewCategoryService(cateRepo repository.CategoryRepository) CategoryService {
	return &categoryService{
		categoryRepository: cateRepo,
	}
}

func (service *categoryService) InsertCategory(b dto.CategoryCreateDto) models.Category {
	category := models.Category{}
	err := smapping.FillStruct(&category, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := service.categoryRepository.InsertCategory(category)
	return res
}

func (service *categoryService) UpdateCategory(b dto.CategoryUpdateDto) models.Category {
	category := models.Category{}
	category.Id = b.Id
	category.Name = b.Name
	res := service.categoryRepository.UpdateCategory(category)
	return res
}

func (service *categoryService) DeleteCategory(b models.Category) {
	service.categoryRepository.DeleteCategory(b)
}

func (service *categoryService) FindAllCategory() []models.Category {
	return service.categoryRepository.FindAllCategory() // Mengambil semua kategori
}

func (service *categoryService) FindCategoryById(categoryID uint64) models.Category {
	return service.categoryRepository.FindCategoryById(categoryID) // Mengambil kategori berdasarkan ID
}

func (service *categoryService) IsAllowedToEdit(Id string, CategoryId uint64) bool {
	b := service.categoryRepository.FindCategoryById(CategoryId)
	id := fmt.Sprintf("%v", b.Id)
	return Id == id
}
