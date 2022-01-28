package cmd

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/spf13/cobra"

	"github.com/michaelenger/kurikaeshi/lib"
)

// hiraganaCmd represents the hiragana command
var hiraganaCmd = &cobra.Command{
	Use:   "hiragana",
	Short: "Test hiragana memorisation",
	Run: func(cmd *cobra.Command, args []string) {
		rand.Seed(time.Now().UnixNano())

		words, err := lib.LoadHiragana()
		if err != nil {
			fmt.Errorf("ERROR: %v", err)
			os.Exit(1)
		}

		lib.Challenge(words)
	},
}

func init() {
	rootCmd.AddCommand(hiraganaCmd)
}
