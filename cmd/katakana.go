package cmd

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/spf13/cobra"

	"github.com/michaelenger/kurikaeshi/lib"
)

// katakanaCmd represents the katakana command
var katakanaCmd = &cobra.Command{
	Use:   "katakana",
	Short: "Test katakana memorisation",
	Run: func(cmd *cobra.Command, args []string) {
		rand.Seed(time.Now().UnixNano())

		katakana, err := lib.LoadKatakana()
		if err != nil {
			fmt.Errorf("ERROR: %v", err)
			os.Exit(1)
		}

		index := rand.Intn(len(katakana))
		fmt.Printf("%v", katakana[index])
	},
}

func init() {
	rootCmd.AddCommand(katakanaCmd)
}
