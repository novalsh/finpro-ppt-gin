package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"finpro/config"
	"finpro/controllers"
	"finpro/middleware"
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
	todoController     controllers.TodoController     = controllers.NewTodoController(todoService)
	categoryController controllers.CategoryController = controllers.NewCategoryController(categoryService, jwtService)
	userController     controllers.UserController     = controllers.NewUserController(userService, jwtService)
	stepController     controllers.StepController     = controllers.NewStepController(stepService)
)

func main() {
	defer config.CloseDatabaseConnection(db)
	r := gin.Default()

	authRoutes := r.Group("auth")
	{
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/register", authController.Register)
	}

	authRoutes = r.Group("user", middleware.AuthorizeJWT(jwtService))
	{
		authRoutes.PUT("/", userController.UpdateUser)
		authRoutes.GET("/:id", userController.ProfileUser)

		autRoutes := r.Group("todos", middleware.AuthorizeJWT(jwtService))
		{
			autRoutes.GET("/", todoController.FindAllTodo)
			autRoutes.GET("/:id", todoController.FindTodoById)
			autRoutes.POST("/", todoController.InsertTodo)
			autRoutes.PUT("/", todoController.UpdateTodo)
			autRoutes.DELETE("/:id", todoController.DeleteTodo)
		}

		authRoutes = r.Group("categories", middleware.AuthorizeJWT(jwtService))
		{
			authRoutes.GET("/", categoryController.AllCategory)
			authRoutes.GET("/:id", categoryController.FindCategoryById)
			authRoutes.POST("/", categoryController.InsertCategory)
			authRoutes.PUT("/", categoryController.UpdateCategory)
			authRoutes.DELETE("/:id", categoryController.DeleteCategory)
		}

		authRoutes = r.Group("steps", middleware.AuthorizeJWT(jwtService))
		{
			authRoutes.GET("/", stepController.FindAllStep)
			authRoutes.GET("/:id", stepController.FindStepById)
			authRoutes.POST("/", stepController.InsertStep)
			authRoutes.PUT("/", stepController.UpdateStep)
			authRoutes.DELETE("/:id", stepController.DeleteStep)
		}

		r.Run()
	}
}
