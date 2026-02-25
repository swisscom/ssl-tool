package main

import "github.com/spf13/cobra"

var clientCmd = &cobra.Command{
	Use:   "client",
	Short: "Client operations",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

func init() {
	clientCmd.AddCommand(getCertificateCmd)
}
