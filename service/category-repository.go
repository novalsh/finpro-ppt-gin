package service

import (
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
}

type categoryService struct {
	categoryRepository repository.CategoryRepository
}

func NewCategory(cateRepo repository.CategoryRepository) CategoryService {
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
	res := service.categoryRepository.InsertCategory(category.CategoryName) // Menggunakan category.CategoryName sebagai argumen
	return res
}

func (service *categoryService) UpdateCategory(b dto.CategoryUpdateDto) models.Category {
	category := models.Category{}
	err := smapping.FillStruct(&category, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := service.categoryRepository.UpdateCategory(b.CategoryId, category.CategoryName) // Menggunakan ID dan nama kategori yang sesuai
	return res
}

func (service *categoryService) DeleteCategory(b models.Category) {
	categoryID := b.CategoryId                            // Mengambil ID kategori dari variabel b
	service.categoryRepository.DeleteCategory(categoryID) // Menghapus kategori berdasarkan ID
}

func (service *categoryService) FindAllCategory() []models.Category {
	return service.categoryRepository.FindAllCategory() // Mengambil semua kategori
}

func (service *categoryService) FindCategoryById(categoryID uint64) models.Category {
	return service.categoryRepository.FindCategoryById(categoryID) // Mengambil kategori berdasarkan ID
}
