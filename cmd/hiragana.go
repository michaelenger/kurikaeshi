package cmd

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/spf13/cobra"

	"kurikaeshi/internal/challenge"
	"kurikaeshi/internal/data"
)

// hiraganaCmd represents the hiragana command
var hiraganaCmd = &cobra.Command{
	Use:   "hiragana",
	Short: "Test hiragana memorisation",
	Run: func(cmd *cobra.Command, args []string) {
		rand.Seed(time.Now().UnixNano())

		words, err := data.LoadHiragana()
		if err != nil {
			fmt.Errorf("ERROR: %v", err)
			os.Exit(1)
		}

		challenge.RunChallenge(words)
	},
}

func init() {
	rootCmd.AddCommand(hiraganaCmd)
}
