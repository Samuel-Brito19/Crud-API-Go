package controller

import (
	"go-api/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type prodcutController struct {

}

func NewProductController() prodcutController {
	return prodcutController{}
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