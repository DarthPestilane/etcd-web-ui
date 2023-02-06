package start

import (
	"etcd-web-ui/cmd/start/flags"
	"etcd-web-ui/pkg/controller"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"net/http"
)

const Port = 8259 // http port

var Command = &cobra.Command{
	Use:   "start",
	Short: "start api server",
	RunE: func(cmd *cobra.Command, args []string) error {
		// start web server
		engine := gin.New()
		engine.Use(gin.Logger(), gin.Recovery(), CorsMiddleware())

		// register routes
		{
			engine.NoRoute(controller.NoRoute)
			engine.GET("/options", controller.GetOptions)
			engine.POST("/refresh", controller.Refresh)
			engine.POST("/get", controller.Get)
			engine.PUT("/put", controller.Put)
			engine.POST("/delete", controller.Delete)
		}

		return engine.Run(fmt.Sprintf(":%d", Port))
	},
}

func init() {
	flagSet := Command.Flags()

	flagSet.SortFlags = false
	flagSet.StringVar(&flags.Endpoint, "endpoint", "127.0.0.1:2379", "the endpoint to connect to")
	flagSet.StringVar(&flags.Username, "username", "", "etcd auth username")
	flagSet.StringVar(&flags.Password, "password", "", "etcd auth password")
	flagSet.StringVar(&flags.KeyPrefix, "key-prefix", "/", "etcd key prefix")
}

func CorsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Origin, Origin, Content-Type, Content-Length, Accept-Encoding")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
