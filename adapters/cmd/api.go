/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"github.com/spf13/cobra"
	server "github.com/wy7-source/iti-challenge/adapters/api"
	"github.com/wy7-source/iti-challenge/application/services"
)

var port string

// apiCmd represents the api command
var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "Api is the way to interact with application.",
	Long:  `Api enable the '/validate' endpoint, that can receive a json with password to be validated.`,
	Run: func(cmd *cobra.Command, args []string) {
		passwordService := services.NewPasswordService()
		server := server.MakeServer(port, passwordService)
		server.Serve()
	},
}

func init() {
	rootCmd.AddCommand(apiCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// apiCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	apiCmd.Flags().StringVarP(&port, "port", "p", ":9000", "Custom port to be used on api")
}
