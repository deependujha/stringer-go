package stringer

import (
    "fmt"

    "github.com/deependujha/stringer-go/pkg/stringer"
    "github.com/spf13/cobra"
)

var reverseCmd = &cobra.Command{
    Use:   "reverse",
    Aliases: []string{"rev"},
    Short:  "Reverses a string",
    Args:  cobra.ExactArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Passed args are: ", args)
        res := stringer.Reverse(args[0])
        fmt.Println(res)
    },
}

func init() {
    rootCmd.AddCommand(reverseCmd)
}