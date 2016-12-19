package cmd

import (
	"github.com/docker/ddc-opencontrol/nlp/cmd/analyze"
	"github.com/spf13/cobra"
)

var (
	// RootCmd is the nlp command root
	RootCmd = &cobra.Command{
		Use:   "nlp",
		Short: "nlp analyzes DDC component control text",
		Long: `
A tool to analyze DDC component control text and compare narrative content with
NIST 800-53 and FedRAMP control definitions`,
		SilenceUsage:  true,
		SilenceErrors: true,
	}
)

func init() {
	RootCmd.AddCommand(analyze.AnalyzeCmd)
}
