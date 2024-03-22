package cmd

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/opensource-nepal/go-nepali/dateConverter"
	"github.com/spf13/cobra"
)

var errorInvalidDateString = errors.New("requires exactly one argument in fromat YYYY-MM-DD")

const DateFormat = "2006-1-2"

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ad2bs",
	Short: "convert AD date to BS",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errorInvalidDateString
		}
		_, err := time.Parse(DateFormat, args[0])
		if err != nil {
			return err
		}
		return nil
	},
	// Preconditions:
	// 1. arg[0] is a valid AD date in YYYY-MM-DD format
	Run: func(_ *cobra.Command, args []string) {
		ad, _ := time.Parse(DateFormat, args[0])
		bs, err := dateConverter.EnglishToNepali(ad.Year(), int(ad.Month()), ad.Day())
		if err != nil {
			panic(err)
		}
		fmt.Printf("%d-%02d-%02d\n", bs[0], bs[1], bs[2])
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
