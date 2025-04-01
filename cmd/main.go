package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"golang-clean-architecture/api/http"
	noteDao "golang-clean-architecture/internal/app/ports/out/dao"
	"golang-clean-architecture/internal/infrastructure/transport/http/handlers"
	"golang-clean-architecture/pkg/config_loader"
	"log"
)

func main() {
	config := config_loader.LoadConfig()
	fmt.Printf("%s!\n", config.Server.Name)

	dao, db := noteDao.Create(config)
	defer db.Close()

	router, apiInterface := gin.Default(), handlers.NewNoteBackendApi(dao)

	http.Load(router)
	http.RegisterHandlers(router, apiInterface)

	log.Println(fmt.Sprintf("Starting server on : %s", config.Server.Port))
	if err := router.Run(":" + config.Server.Port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
