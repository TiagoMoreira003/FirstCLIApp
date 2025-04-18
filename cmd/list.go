/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list your passwords",
	Long: `list your passwords`,
	Run: listPassword,
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")
	listCmd.Flags().StringP("file", "f", "all", "Specify the file")
}

func listPassword(cmd *cobra.Command, args []string){
	file, _ := cmd.Flags().GetString("file")
	list := []string{}

	if file == "all"{
		files, _ := os.ReadDir("C:/Users/ALEXA/FirstCLIApp")
		for _, f := range files {
			if strings.HasSuffix(f.Name(), ".txt"){
				content, _ := os.ReadFile(f.Name())
				list = append(list, string(content))
			}
		} 
	} else{
		filename := file + ".txt"
		data, _ := os.ReadFile(filename)
		list = append(list, string(data))
	}

	for f,_ := range list{
		fmt.Println(list[f])
	}
}