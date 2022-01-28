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

// katakanaCmd represents the katakana command
var katakanaCmd = &cobra.Command{
	Use:   "katakana",
	Short: "Test katakana memorisation",
	Run: func(cmd *cobra.Command, args []string) {
		rand.Seed(time.Now().UnixNano())

		words, err := data.LoadKatakana()
		if err != nil {
			fmt.Errorf("ERROR: %v", err)
			os.Exit(1)
		}

		challenge.RunChallenge(words)
	},
}

func init() {
	rootCmd.AddCommand(katakanaCmd)
}
