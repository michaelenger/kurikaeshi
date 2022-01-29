// Package cmd handles all the CLI commands.
package cmd

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/spf13/cobra"

	"kurikaeshi/internal/challenge"
	"kurikaeshi/internal/data"
)

func runCommand(cmd *cobra.Command, args []string) error {
	rand.Seed(time.Now().UnixNano())

	var words []data.Word
	var err error

	if args[0] == "hiragana" {
		words, err = data.LoadHiragana()
	} else if args[0] == "katakana" {
		words, err = data.LoadKatakana()
	}

	if err != nil {
		return err
	}

	challenge.RunChallenge(words)

	return nil
}

func validateArgs(cmd *cobra.Command, args []string) error {
	if len(args) < 1 {
		return errors.New("requires a syllabary argument\n")
	}
	if args[0] != "hiragana" && args[0] != "katakana" {
		return fmt.Errorf("invalid syllabary: %s\n", args[0])
	}

	return nil
}

var rootCmd = &cobra.Command{
	Use:   "kurikaeshi [syllabary]",
	Short: "Learn you Japanese good through repetition",
	Args:  validateArgs,
	RunE:  runCommand,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	usageLines := strings.Split(rootCmd.UsageTemplate(), "\n")

	newUsageLines := append(usageLines[:12], "\nAvailable syllabaries:\n  hiragana\n  katakana\n")
	newUsageLines = append(newUsageLines, usageLines[13:]...)

	usageTemplate := strings.Join(newUsageLines, "\n")

	rootCmd.SetUsageTemplate(usageTemplate)
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
