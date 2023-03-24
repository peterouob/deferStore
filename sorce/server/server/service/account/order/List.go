package serviceOrder

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"server/server/service/h"
)

type ListRequest struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
}

type ListRequestItem struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}
type ListResponse struct {
	ListItems []*ListRequestItem
}

func List(c *gin.Context) {
	req := ListRequest{}
	c.ShouldBindJSON(&req)
	fmt.Println(req)
	//SQL Query

	//Seed
	resp := &ListResponse{}
	for i := 0; i < 5; i++ {
		resp.ListItems = append(resp.ListItems, &ListRequestItem{
			Id:   fmt.Sprintf("%d", i),
			Name: "name",
		})
	}
	tk := h.GetToken(c)
	c.JSON(http.StatusOK, tk)
}
