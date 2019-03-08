// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "file-server",
	Short: "A file server for web files",
	Long:  `A static file server. Serves html, javascript, css, icons, fonts and image files`,
	RunE: func(cmd *cobra.Command, args []string) error {
		port, err := cmd.Flags().GetString("port")
		if err != nil {
			return err
		}
		dir, err := cmd.Flags().GetString("dir")
		if err != nil {
			return err
		}
		log.Printf("Http file server started on port: %s\n", port)
		return http.ListenAndServe(port, http.FileServer(http.Dir(dir)))
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringP("port", "p", ":9090", "Port to serve files")
	rootCmd.Flags().StringP("dir", "d", "", "Directory to serve files")
	rootCmd.Flags().BoolP("tls", "tls", false, "Encrypt communication with tls")
}
