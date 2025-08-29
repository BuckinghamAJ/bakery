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
		Description: "Our classic sourdough bread is crafted using a traditional fermentation process, resulting in a hearty loaf with a crisp crust and a chewy, tangy interior. Perfect for sandwiches, toast, or enjoying on its own.",
	},
	{
		Id:          2,
		Name:        "Chocolate Croissant",
		Cost:        3.25,
		ImgPath:     "assets/img/chocolate-croissants.jpg",
		Description: "A flaky, buttery croissant filled with rich, smooth chocolate. Baked fresh daily, this pastry is a delightful treat for breakfast or an afternoon snack, combining classic French technique with decadent flavor.",
	},
	{
		Id:          3,
		Name:        "Croissant",
		Cost:        3.00,
		ImgPath:     "assets/img/French-Croissants.jpg",
		Description: "Our traditional French croissant features a golden, crisp exterior and a soft, airy interior. Made with premium butter and expertly layered dough, it's the perfect choice for a light breakfast or to accompany your coffee.",
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

	r.GET("/web", func(c *gin.Context) {
		templ.Handler(HelloForm()).ServeHTTP(c.Writer, c.Request)
	})

	r.GET("/", func(c *gin.Context) {
		templ.Handler(Orders(orders)).ServeHTTP(c.Writer, c.Request)
	})

	r.POST("/hello", func(c *gin.Context) {
		HelloWebHandler(c.Writer, c.Request)
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

func (s *Server) HelloWorldHandler(c *gin.Context) {
	resp := make(map[string]string)
	resp["message"] = "Hello World"

	c.JSON(http.StatusOK, resp)
}

func (s *Server) healthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, s.db.Health())
}

func (s *Server) makeErrorMessage(code int, msg string, c *gin.Context) {
	resp := make(map[string]string)
	resp["error"] = msg

	c.JSON(code, resp)
}
