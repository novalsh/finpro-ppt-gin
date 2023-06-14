package controllers

import (
	"fmt"
	"math"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"finpro/dto"
	"finpro/service"
)

type TodoController interface {
	InsertTodo(ctx *gin.Context)
	UpdateTodo(ctx *gin.Context)
	DeleteTodo(ctx *gin.Context)
	FindAllTodo(ctx *gin.Context)
	FindTodoById(ctx *gin.Context)
	KMeans(ctx *gin.Context)
}

type FilteredTodo struct {
	Name     string  `json:"name"`
	Deadline string  `json:"deadline"`
	Level    float64 `json:"level"`
}

type Cluster struct {
	Centroid FilteredTodo
	Items    []FilteredTodo
}

type HasilCluster struct {
	Name     string  `json:"name"`
	Deadline string  `json:"deadline"`
	Level    float64 `json:"level"`
}

type todoController struct {
	todoService service.TodoService
}

func NewTodoController(todoServ service.TodoService) TodoController {
	return &todoController{
		todoService: todoServ,
	}
}

func (controller *todoController) InsertTodo(ctx *gin.Context) {
	var todoCreateDto dto.TodoCreateDto
	err := ctx.ShouldBindJSON(&todoCreateDto)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Panggil service untuk menyimpan todo
	createdTodo := controller.todoService.InsertTodo(todoCreateDto)

	ctx.JSON(http.StatusOK, createdTodo)
}

func (controller *todoController) UpdateTodo(ctx *gin.Context) {
	var todoUpdateDto dto.TodoUpdateDto
	err := ctx.ShouldBindJSON(&todoUpdateDto)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Panggil service untuk mengupdate todo
	updatedTodo := controller.todoService.UpdateTodo(todoUpdateDto)

	ctx.JSON(http.StatusOK, updatedTodo)
}

func (controller *todoController) DeleteTodo(ctx *gin.Context) {
	// Ambil ID todo dari parameter URL
	TodoID := ctx.Param("id")

	// Konversi TodoID menjadi uint64
	todoIDUint, err := strconv.ParseUint(TodoID, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid todo ID"})
		return
	}

	// Panggil service untuk menghapus todo
	err = controller.todoService.DeleteTodoById(todoIDUint)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Todo deleted successfully"})
}

func (controller *todoController) FindAllTodo(ctx *gin.Context) {
	// Panggil service untuk mendapatkan daftar semua todo
	todos := controller.todoService.FindAllTodo()

	ctx.JSON(http.StatusOK, todos)
}

func (controller *todoController) FindTodoById(ctx *gin.Context) {
	// Ambil ID todo dari parameter URL
	todoId := ctx.Param("id")

	TodoId, err := strconv.ParseUint(todoId, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Panggil service untuk mendapatkan todo berdasarkan ID
	todo := controller.todoService.FindTodoById(TodoId)

	ctx.JSON(http.StatusOK, gin.H{"todo": todo})
}

func (controller *todoController) KMeans(ctx *gin.Context) {
	todos := controller.todoService.KMeans()

	filteredTodos := make([]FilteredTodo, 0)

	for _, todo := range todos {
		filteredTodo := FilteredTodo{
			Name:     todo.Name,
			Deadline: todo.Deadline,
			Level:    todo.Level,
		}

		filteredTodos = append(filteredTodos, filteredTodo)
	}

	k := 2

	centroids := initializeCentroids(filteredTodos, k)

	for {
		clusters := clusterItems(filteredTodos, centroids)

		newCentroids := calculateNewCentroids(clusters)

		if centroidsConverged(centroids, newCentroids) {
			break
		}
		centroids = newCentroids
	}

	clusters := clusterItems(filteredTodos, centroids)
	hasilClusters := make(map[string][]map[string]interface{})

	for i, cluster := range clusters {
		clusterItems := make([]map[string]interface{}, 0)
		fmt.Printf("Cluster %d:\n", i+1)
		for _, item := range cluster.Items {
			fmt.Printf("- %s (Deadline: %s, Level: %.2f)\n", item.Name, item.Deadline, item.Level)

			clusterItem := map[string]interface{}{
				"name":     item.Name,
				"deadline": item.Deadline,
				"level":    item.Level,
			}
			clusterItems = append(clusterItems, clusterItem)
		}
		hasilClusters[fmt.Sprintf("Cluster %d", i+1)] = clusterItems
	}

	ctx.JSON(http.StatusOK, hasilClusters)
}

func initializeCentroids(items []FilteredTodo, k int) []FilteredTodo {
	centroids := make([]FilteredTodo, k)
	for i := 0; i < k; i++ {
		centroids[i] = items[i]
	}
	return centroids

}

func clusterItems(items []FilteredTodo, centroids []FilteredTodo) []Cluster {
	clusters := make([]Cluster, len(centroids))
	for _, item := range items {
		minDistance := math.MaxFloat64
		clusterIndex := 0
		for i, centroid := range centroids {
			distance := euclideanDistance(item, centroid)
			if distance < minDistance {
				minDistance = distance
				clusterIndex = i
			}
		}
		clusters[clusterIndex].Items = append(clusters[clusterIndex].Items, item)
	}
	return clusters
}

func euclideanDistance(item1, item2 FilteredTodo) float64 {
	return math.Sqrt(math.Pow(item1.Level-item2.Level, 2))
}

func calculateNewCentroids(clusters []Cluster) []FilteredTodo {
	newCentroids := make([]FilteredTodo, len(clusters))
	for i, cluster := range clusters {
		var sumLevel float64
		for _, item := range cluster.Items {
			sumLevel += item.Level
		}
		centroid := FilteredTodo{
			Name:     fmt.Sprintf("Centroid %d", i+1),
			Deadline: cluster.Centroid.Deadline,
			Level:    sumLevel / float64(len(cluster.Items)),
		}
		newCentroids[i] = centroid
	}
	return newCentroids
}

func centroidsConverged(centroids, newCentroids []FilteredTodo) bool {
	for i := range centroids {
		if centroids[i].Level != newCentroids[i].Level {
			return false
		}
	}
	return true
}
