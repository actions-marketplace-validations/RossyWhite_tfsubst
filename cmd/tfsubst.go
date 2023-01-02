package cmd

import (
	"context"
	"fmt"
	"github.com/fujiwara/tfstate-lookup/tfstate"
	"github.com/spf13/cobra"
	"io"
	"os"
	"text/template"
)

func Execute() {
	err := tfsubstCmd.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var (
	stateLoc   string
	inputFile  string
	outputFile string
	funcName   string
	tfsubstCmd = &cobra.Command{
		Use: "tfsubst",
		RunE: func(cmd *cobra.Command, args []string) error {
			c := &tfsubst{}
			ctx := context.Background()
			var in, out *os.File

			if inputFile == "" {
				in = os.Stdin
			} else {
				f, err := os.Open(inputFile)
				if err != nil {
					return err
				}
				defer func() { _ = f.Close() }()
				in = f
			}

			if outputFile == "" {
				out = os.Stdout
			} else {
				f, err := os.Create(outputFile)
				if err != nil {
					return err
				}
				defer func() { _ = f.Close() }()
				out = f
			}

			return c.execute(ctx, stateLoc, in, out, funcName)
		},
	}
)

type tfsubst struct{}

func (c *tfsubst) execute(
	ctx context.Context,
	stateLoc string,
	in io.Reader,
	out io.Writer,
	funcName string) error {

	funcMap, err := tfstate.FuncMapWithName(ctx, funcName, stateLoc)
	if err != nil {
		return err
	}

	b, err := io.ReadAll(in)
	if err != nil {
		return err
	}

	t, err := template.New("file").Funcs(funcMap).Parse(string(b))
	if err = t.Execute(out, nil); err != nil {
		return err
	}
	return nil
}

func init() {
	tfsubstCmd.Flags().StringVarP(&stateLoc, "state", "s", "", "local tfstate file path or remote URL")
	tfsubstCmd.Flags().StringVarP(&inputFile, "input", "i", "", "specify file input, otherwise use stdin")
	tfsubstCmd.Flags().StringVarP(&outputFile, "output", "o", "", "specify file output, otherwise use stdout")
	tfsubstCmd.Flags().StringVar(&funcName, "func-name", "tfstate", "func name to use in template")

	_ = tfsubstCmd.MarkFlagRequired("state")
}
