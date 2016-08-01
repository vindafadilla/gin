package main
import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// export "GIN_MODE=release"

func main() {

	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()

	// This handler will match /user/john but will not match neither /user/ or /user
	router.GET("/user/vinda", func(c *gin.Context) {
		name := c.Param("vinda")
		c.String(http.StatusOK, "Hello %s", name)
	})

	// However, this one will match /user/john/ and also /user/john/send
	// If no other routers match /user/john, it will redirect to /user/john/
	router.GET("/user/vinda/send", func(c *gin.Context) {
		name := c.Param("vinda")
		action := c.Param("send")
		message := name + " is " + action
		c.String(http.StatusOK, message)
	})

	router.Run(":8080")
}