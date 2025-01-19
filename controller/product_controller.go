package controller

import (
	"go-api/model"
	"go-api/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type prodcutController struct {
	productUseCase usecase.ProductUseCase
}

func NewProductController(usecase usecase.ProductUseCase) prodcutController {
	return prodcutController{
		productUseCase: usecase,
	}
}

func(p *prodcutController) GetProducts(ctx *gin.Context) {
	products := []model.Product{
		{
			ID: 1,
			Name: "Batata-frita",
			Price: 10,
		},
		
	}

	ctx.JSON(http.StatusOK, products)
}