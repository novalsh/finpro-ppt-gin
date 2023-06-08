package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"finpro/config"
	"finpro/controllers"
	"finpro/repository"
	"finpro/service"
)

var (
	db                 *gorm.DB                       = config.SetupDatabaseConnection()
	userRepository     repository.UserRepository      = repository.NewUserRepository(db)
	categoryRepository repository.CategoryRepository  = repository.NewCategoryRepository(db)
	stepRepository     repository.StepRepository      = repository.NewStepRepository(db)
	todoRepostiory     repository.TodoRepository      = repository.NewTodoRepository(db)
	jwtService         service.JWTService             = service.NewJWTService()
	authService        service.AuthService            = service.NewAuthService(userRepository)
	userService        service.UserService            = service.NewUserService(userRepository)
	categoryService    service.CategoryService        = service.NewCategoryService(categoryRepository)
	stepService        service.StepService            = service.NewStepService(stepRepository)
	todoService        service.TodoService            = service.NewTodoService(todoRepostiory)
	authController     controllers.AuthController     = controllers.NewAuthController(authService, jwtService)
	todoController     controllers.TodoController     = controllers.NewTodoController(todoService, jwtService)
	categoryController controllers.CategoryController = controllers.NewCategoryController(categoryService, jwtService)
)

func main() {
	defer config.CloseDatabaseConnection(db)
	r := gin.Default()

	authRoutes := r.Group("users")
	{
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/register", authController.Register)
	}

	autRoutes := r.Group("todos")
	{
		autRoutes.GET("/", todoController.AllTodo)
		autRoutes.GET("/:id", todoController.FindTodoById)
		autRoutes.POST("/", todoController.InsertTodo)
		autRoutes.PUT("/", todoController.UpdateTodo)
		autRoutes.DELETE("/:id", todoController.DeleteTodo)
	}

	authRoutes = r.Group("categories")
	{
		authRoutes.GET("/", categoryController.AllCategory)
		authRoutes.GET("/:id", categoryController.FindCategoryById)
		authRoutes.POST("/", categoryController.InsertCategory)
		authRoutes.PUT("/", categoryController.UpdateCategory)
		authRoutes.DELETE("/:id", categoryController.DeleteCategory)
	}

	r.Run()
}
