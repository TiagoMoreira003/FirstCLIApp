package cmd

import (
	"fmt"
	"math/rand"
	"os"

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
	generateCmd.Flags().BoolP("special-characters", "c", false, "Special characters (true/false)" )
	generateCmd.Flags().BoolP("upper", "u", false, "Use uppercase letters")
	generateCmd.Flags().StringP("file", "f", "password", "Specify the filse destination")
}

func generatePassword(cmd *cobra.Command, args []string){
	length, _ := cmd.Flags().GetInt("length")
	specialChar, _ := cmd.Flags().GetBool("special-characters")
	uppercase, _ := cmd.Flags().GetBool("upper")
	file, _ := cmd.Flags().GetString("file") 
	
	password := ""
	filename := file + ".txt"

	var letters string

	lower := "abcdefghijklmnopqrstuvwxyz"
	upper := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	special := "!@#$%&*-_"

	if uppercase && specialChar {
		letters = lower + upper + special
	} else if uppercase && !specialChar {
		letters = lower + upper
	} else if !uppercase && specialChar {
		letters = lower + special
	} else {
		letters = lower
}
	// para o numero de length vai buscar uma letra do alfabeto
	for i := 0; i<length; i++ {
		randomIndex := rand.Intn(len(letters))
		password += string(letters[randomIndex])
	}

	os.WriteFile(filename, []byte(password), 0644)

	fmt.Println(password)
}