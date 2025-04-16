package cmd

import (
	"fmt"
	"math/rand"
	"github.com/spf13/cobra"
)

var generateCmd = &cobra.Command{
	Use: "generate",
	Short: "generate password",
	Long: "generate password with custom options",

	Run: generatePassword,
}

func init(){
	rootCmd.AddCommand(generateCmd)
	generateCmd.Flags().IntP("length", "l", 4, "Length of the password")
}

func generatePassword(cmd *cobra.Command, args []string){
	length, _ := cmd.Flags().GetInt("length")
	password := ""
	letters := "abcdefghijklmnopqrstuvwxyz"
	// para o numero de length vai buscar uma letra do alfabeto
	for i := 0; i<length; i++ {
		randomIndex := rand.Intn(len(letters))
		password += string(letters[randomIndex])
	}
	fmt.Println(password)
}