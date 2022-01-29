// Package cmd handles all the CLI commands.
package cmd

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/cobra"

	"kurikaeshi/internal/colors"
	"kurikaeshi/internal/data"
)

// Amount of words to test
var wordCount int

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

	wordCountStr := "∞"
	if wordCount != -1 {
		wordCountStr = strconv.Itoa(wordCount)
	}
	fmt.Printf("ようこそ！ You are guessing %s words!\n (an empty input exits the program)\n\n", wordCountStr)

	var guess string
	var output string
	var wordCounter int = 0
	var correctWords int = 0

	for { // there is no escape
		word := words[rand.Intn(len(words))]
		guess = ""
		wordCounter += 1

		fmt.Printf("%v: ", word.Letters)
		fmt.Scanln(&guess)

		if guess == "" {
			wordCounter -= 1
			break // there is some escape
		}

		if guess == data.Sanitize(word.Romaji) {
			output = colors.Green(fmt.Sprintf("%v (%v): %v - %v", word.Letters, word.Kanji, word.Romaji, word.Translation))
			correctWords += 1
		} else {
			output = colors.Red(fmt.Sprintf("%v (%v): %v - %v", word.Letters, word.Kanji, word.Romaji, word.Translation))
		}

		fmt.Println(output)

		if wordCount != -1 && wordCounter == wordCount {
			break // there is actually a lot of escape
		}
	}

	var percentage int = 0
	if wordCounter > 0 {
		percentage = int(float32(correctWords) / float32(wordCounter) * 100)
	}
	fmt.Printf("\nYou got %d/%d words! (%d%%)\n", correctWords, wordCounter, percentage)

	return nil
}

func validateArgs(cmd *cobra.Command, args []string) error {
	if len(args) < 1 {
		return errors.New("requires a syllabary argument")
	}
	if args[0] != "hiragana" && args[0] != "katakana" {
		return fmt.Errorf("invalid syllabary: %s", args[0])
	}

	return nil
}

var rootCmd = &cobra.Command{
	Use:   "kurikaeshi <syllabary>",
	Short: "繰り返し - Learn you Japanese good through repetition",
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
	rootCmd.Flags().IntVarP(&wordCount, "amount", "n", -1, "Amount of words to test")

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
