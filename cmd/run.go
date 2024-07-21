/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

var port int

var path string

var currdir string

var cache *Cache

func handleFunc(w http.ResponseWriter, r *http.Request) {
	filepath := fmt.Sprintf("%s/%s/%s", currdir, path, r.URL.Path)
	data, ok := cache.fetch(filepath)
	fmt.Println(ok)
	if !ok {
		content, err := os.ReadFile(filepath)
		if err != nil {
			if strings.Contains(err.Error(), "no such file or directory") {
				w.WriteHeader(404)
				fmt.Fprintf(w, "%s", "Not found")
				return
			} else {
				panic(err)
			}
		}
		data = string(content)
		cache.persist(filepath, data)
		fmt.Println("persist in cache")
	} else {
		fmt.Println("fetch from cache")
	}

	fmt.Println("sending output")
	fmt.Fprintf(w, "%s", data)
}

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Starts the HTTP server",
	Long:  `Starts the HTTP Server`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("starting server port:", port)
		fmt.Println("serving directory:", path)
		cache = NewCache()
		mux := http.NewServeMux()
		mux.HandleFunc("/", handleFunc)
		s := &http.Server{
			Addr:           fmt.Sprintf(":%d", port),
			Handler:        mux,
			ReadTimeout:    10 * time.Second,
			WriteTimeout:   10 * time.Second,
			MaxHeaderBytes: 1 << 20,
		}
		s.ListenAndServe()

	},
}

func init() {
	var err error
	currdir, err = os.Getwd()
	if err != nil {
		panic(err)
	}
	fmt.Println("CWD:", currdir)
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
