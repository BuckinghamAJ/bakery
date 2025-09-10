package server

import (
	"net/http"
	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"buckingham_bakery/cmd/web"
	"io/fs"

	. "buckingham_bakery/cmd/web/templates"
	. "buckingham_bakery/internal/dto"

	"github.com/a-h/templ"
)

var orders = []FoodOrder{
	{
		Id:          1,
		Name:        "Sourdough Bread",
		Cost:        4.50,
		ImgPath:     "assets/img/Sourdough-Bread.jpg",
		Description: "Classic sourdough bread crafted using a traditional fermentation process.",
	},
	{
		Id:          2,
		Name:        "Chocolate Croissant",
		Cost:        3.25,
		ImgPath:     "assets/img/chocolate-croissants.jpg",
		Description: "A flaky, buttery croissant filled with rich, smooth chocolate.",
	},
	{
		Id:          3,
		Name:        "Croissant",
		Cost:        3.00,
		ImgPath:     "assets/img/French-Croissants.jpg",
		Description: "Our traditional French croissant features a golden, crisp exterior and a soft, airy interior.",
	},
}

func (s *Server) RegisterRoutes() http.Handler {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, // Add your frontend URL
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: true, // Enable cookies/auth
	}))

	r.GET("/health", s.healthHandler)

	staticFiles, _ := fs.Sub(web.Files, "assets")
	r.StaticFS("/assets", http.FS(staticFiles))

	r.GET("/", func(c *gin.Context) {
		// clearCartFromSession(c.Writer, c.Request)

		cartOrders, _ := getCartFromSession(c.Request)

		templ.Handler(Orders(orders, cartOrders)).ServeHTTP(c.Writer, c.Request)
	})

	r.GET("/sidecart", func(c *gin.Context) {
		GetSideCartOrders(c.Writer, c.Request)
	})

	r.PUT("/cart/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			// TODO: Probably will need to change to some alert message on the front end
			s.makeErrorMessage(400, "Incorrect order id provided", c)
		}

		PutCartWebHandler(id, c.Writer, c.Request)
	})

	return r
}

func (s *Server) healthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, s.db.Health())
}

func (s *Server) makeErrorMessage(code int, msg string, c *gin.Context) {
	resp := make(map[string]string)
	resp["error"] = msg

	c.JSON(code, resp)
}
