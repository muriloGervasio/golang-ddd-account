package interfaces

import "github.com/gin-gonic/gin"

type Server struct {
}

func NewServer() *Server {
	return &Server{}
}

func (server *Server) Run() {
	r := gin.Default()

	for _, route := range Routes {
		switch route.Method {
		case GET:
			r.GET(route.Path, route.Handler)
		case POST:
			r.POST(route.Path, route.Handler)
		case PUT:
			r.PUT(route.Path, route.Handler)
		case DELETE:
			r.DELETE(route.Path, route.Handler)
		case PATCH:
			r.PATCH(route.Path, route.Handler)
		}
	}

	r.Run()
}
