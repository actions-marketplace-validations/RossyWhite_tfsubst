package cmd

import (
	"context"
	"fmt"
	"github.com/fujiwara/tfstate-lookup/tfstate"
	"github.com/spf13/cobra"
	"os"
	"text/template"
)

var (
	stateLoc   string
	inputFile  string
	outputFile string
	funcName   string
	tfsubstCmd = &cobra.Command{
		Use: "tfsubst",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := context.Background()
			funcMap, err := tfstate.FuncMapWithName(ctx, funcName, stateLoc)
			if err != nil {
				return err
			}

			if inputFile == "" {
				inputFile = os.Stdin.Name()
			}

			b, err := os.ReadFile(inputFile)
			if err != nil {
				return err
			}

			var out *os.File
			if outputFile == "" {
				out = os.Stdout
			} else {
				out, err = os.Create(outputFile)
				if err != nil {
					return err
				}
			}

			t, err := template.New("file").Funcs(funcMap).Parse(string(b))
			if err = t.Execute(out, nil); err != nil {
				return err
			}
			return nil
		},
	}
)

func Execute() {
	err := tfsubstCmd.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	tfsubstCmd.Flags().StringVarP(&stateLoc, "state", "s", "", "tfstate file path or URL")
	tfsubstCmd.Flags().StringVarP(&inputFile, "input", "i", "", "specify file input, otherwise use stdin")
	tfsubstCmd.Flags().StringVarP(&outputFile, "output", "o", "", "specify file output, otherwise use stdout")
	tfsubstCmd.Flags().StringVar(&funcName, "func-name", "tfstate", "func name to use in template")

	_ = tfsubstCmd.MarkFlagRequired("state")
}
