package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"orderPacks/pkg/orderPacks/service"
	"strconv"
)

type Server struct {
	Service service.PackageService
}

func NewServer(service service.PackageService) Server {
	return Server{
		Service: service,
	}
}

type SetPackageSizesRequest struct {
	Sizes []int `json:"sizes"`
}

func (s *Server) InitServer() *gin.Engine {
	router := gin.Default()
	router.GET("/handleOrder/:orderSize", s.HandleOrder)
	router.POST("/setPackageSizes", s.SetPackageSizes)
	return router
}

func (s *Server) HandleOrder(c *gin.Context) {
	orderSize := c.Param("orderSize")
	size, err := strconv.ParseInt(orderSize, 10, 64)
	if err != nil {
		return
	}

	res := s.Service.HandleOrder(int(size))
	c.IndentedJSON(http.StatusOK, res)
}

func (s *Server) SetPackageSizes(c *gin.Context) {
	req := SetPackageSizesRequest{}
	if err := c.BindJSON(&req); err != nil {
		return
	}

	s.Service.SetPackageSizes(req.Sizes)
	c.IndentedJSON(http.StatusOK, http.NoBody)
}
