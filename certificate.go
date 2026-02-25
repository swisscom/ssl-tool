package main

import "github.com/spf13/cobra"

var certCmd = &cobra.Command{
	Use:   "cert",
	Short: "Certificate operations",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

func init() {
	certCmd.AddCommand(parseCmd)
}
