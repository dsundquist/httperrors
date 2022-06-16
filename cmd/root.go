/*
Copyright © 2022 Dean Sundquist dean@sundquist.net

*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gotestserver",
	Short: "A webserver that'll return headers, or can generate 5xx errors",
	Long: `
	This is a webserver that was particularily designed for troubleshooting websites behind cloudflared
	
	For more information about cloudflared please see:
	https://github.com/cloudflare/cloudflared.
	`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.httperrors.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	serveCmd.Flags().String("cert", "./server.crt", "Specify the server certificate for HTTPS")
	serveCmd.Flags().String("clientcert", "./client.crt", "Specify the client certificate for mTLS")
	serveCmd.Flags().BoolP("debug", "d", false, "Set the logging level to debug.")
	serveCmd.Flags().String("key", "./server.key", "Specify the server key for HTTPS")
	serveCmd.Flags().BoolP("mtls", "m", false, "Run a mTLS server. Requires: ./server.crt, ./server.key, and ./client.crt")
	serveCmd.Flags().IntP("port", "p", 80, "Specify which port for the webserver to run on.")
	serveCmd.Flags().BoolP("tls", "s", false, "Run a HTTPS server. Requires: ./server.crt and ./server.key")

	clientCmd.Flags().BoolP("insecure", "k", false, "Ingore Self Signed or bad certificates.")
	clientCmd.Flags().StringP("location", "l", "http://localhost:80", "The address or IP address that we're connecting to.")
}
