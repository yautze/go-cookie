package cmd

import (
	"go-cookie/router"

	"github.com/kataras/iris/v12"
	"github.com/spf13/cobra"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use: "server",
	Run: func(cmd *cobra.Command, args []string) {
		start()
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
	//serverCmd.Flags().StringVarP(&config.C.Info.Port, "port", "p", "1314", "service port")
}

// start -
func start() {
	app := iris.Default()

	router.SetRouter(app)

	app.Run(iris.Addr(":3000"), iris.WithoutInterruptHandler)
}
