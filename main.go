package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"mmt.com/lolbank/handlers"
	"mmt.com/lolbank/repos"
	"mmt.com/lolbank/services"
	"net/http"
)

func main() {
	router := gin.New()
	router.Use(gin.Recovery())
	router.RedirectTrailingSlash = false
	router.Use(gin.Logger())

	repo := repos.NewInMemAccountRepo()
	accSvc := services.NewAccountService(repo)
	accHandler := handlers.NewAccountHandler(accSvc)

	api := router.Group("/api/v1")
	{
		account := api.Group("accounts")
		{
			account.POST("", accHandler.CreateAccount)
			account.PUT(":accountId/deposit", accHandler.DepositToAccount)
			account.PUT(":accountId/withdraw", accHandler.WithdrawMoney)
			account.GET(":accountId", accHandler.GetAccount)
		}
	}

	srv := &http.Server{
		Addr:    ":8000",
		Handler: router,
	}

	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("listen error: %s\n", err)
	}
}
