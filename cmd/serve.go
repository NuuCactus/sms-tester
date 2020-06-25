package cmd

import (
	"github.com/nuucactus/sms-tester/pkg/serve"
	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "SMS Tester service used for development and performance testing",
	Long:  `SMS Tester service used for development and performance testing when external systems are not available`,
	Run:   serve.RunServe(),
}

func init() {
	rootCmd.AddCommand(serveCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	serveCmd.Flags().StringP("metrics-api-url", "", "http://0.0.0.0:80", "Serve Metrics API")
	serveCmd.Flags().StringP("rest-api-url", "", "http://0.0.0.0:8080", "Serve REST API")

}
