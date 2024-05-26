package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/opensource-nepal/go-nepali/nepalitime"
	"github.com/spf13/cobra"
)

var (
	ErrDateArgsMissing = errors.New("date argument is missing. format: YYYY-MM-DD")
	DateOutputFormat   = "2006-01-02"
)

func validateArgs(cmd *cobra.Command, args []string) error {
	if len(args) != 1 {
		return ErrDateArgsMissing
	}
	_, err := nepalitime.Parse(args[0], "%Y-%m-%d")
	if err != nil {
		return err
	}
	return nil
}

func convertBSToAD(cmd *cobra.Command, args []string) {
	npTime, _ := nepalitime.Parse(args[0], "%Y-%m-%d")
	enTime := npTime.GetEnglishTime()
	enTimeFormat := enTime.Format(DateOutputFormat)
	fmt.Println(enTimeFormat)
}

func main() {
	var bs2adCmd = &cobra.Command{
		Use:   "bs2ad",
		Short: "Convert BS to AD",
		Args:  validateArgs,
		Run:   convertBSToAD,
	}

	if err := bs2adCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
