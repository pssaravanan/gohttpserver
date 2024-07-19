/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/spf13/cobra"
)

var port int

var path string

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Starts the HTTP server",
	Long:  `Starts the HTTP Server`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("starting server port:", 8080)
		mux := http.NewServeMux()
		mux.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
			_, err := os.Executable()
			if err != nil {
				panic(err)
			}
			dir, err := os.Getwd()
			if err != nil {
				return
			}
			fmt.Println(dir)
			fmt.Println("Incoming request", r.URL.Path)
			filepath := fmt.Sprintf("%s/%s/%s", dir, path, "test.html")
			fmt.Println("Path", filepath)
			data, _ := os.ReadFile(filepath)
			fmt.Println("File read", r.URL.Path)
			// cwd, _ := os.Getwd()
			// fmt.Println(string(err.Error()))
			fmt.Println(string(data))
			// fmt.Println(cwd)
			fmt.Fprintf(w, "%s", string(data))
		})
		mux.HandleFunc("/s.css", func(w http.ResponseWriter, r *http.Request) {
			fmt.Println("Incoming request", r.URL.Path)
			data, _ := os.ReadFile("./testroot/s.css")
			fmt.Println("File read", r.URL.Path)
			// cwd, _ := os.Getwd()
			// fmt.Println(string(err.Error()))
			fmt.Println(string(data))
			// fmt.Println(cwd)
			fmt.Fprintf(w, "%s", string(data))
		})
		s := &http.Server{
			Addr:           ":8080",
			Handler:        mux,
			ReadTimeout:    10 * time.Second,
			WriteTimeout:   10 * time.Second,
			MaxHeaderBytes: 1 << 20,
		}
		s.ListenAndServe()

	},
}

func init() {
	rootCmd.AddCommand(runCmd)

	runCmd.Flags().IntVar(&port, "port", 3000, "Port")
	runCmd.Flags().StringVar(&path, "path", ".", "Path")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// runCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// runCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
