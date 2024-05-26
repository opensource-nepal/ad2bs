package main

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/opensource-nepal/go-nepali/nepalitime"
	"github.com/spf13/cobra"
)

var (
	ErrDateArgsMissing = errors.New("date argument is missing. format: YYYY-MM-DD")
	DateInputFormat    = "2006-1-2"
)

func validateArgs(cmd *cobra.Command, args []string) error {
	if len(args) != 1 {
		return ErrDateArgsMissing
	}
	_, err := time.Parse(DateInputFormat, args[0])
	if err != nil {
		return err
	}
	return nil
}

func convertADToBS(cmd *cobra.Command, args []string) {
	enTime, _ := time.Parse(DateInputFormat, args[0])
	npTime, err := nepalitime.FromEnglishTime(enTime)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		return
	}
	npTimeFormat, _ := npTime.Format("%Y-%m-%d")
	fmt.Println(npTimeFormat)
}

func main() {
	var ad2bsCmd = &cobra.Command{
		Use:   "ad2bs",
		Short: "Convert AD date to BS",
		Args:  validateArgs,
		Run:   convertADToBS,
	}

	if err := ad2bsCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
