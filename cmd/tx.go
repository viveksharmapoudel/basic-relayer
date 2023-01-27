package cmd

import (
	"strconv"

	"github.com/basic-relayer/src"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(newTxCmd())
}

func newTxCmd() *cobra.Command {
	txCmd := cobra.Command{
		Use:     "transaction",
		Aliases: []string{"tx"},
		Short:   "this function gives transaction data",
		Long:    "this is long form of transaction data massge",
	}

	txCmd.AddCommand(
		txRaw(),
		txBalanceTransfer(),
	)

	return &txCmd
}

func txRaw() *cobra.Command {
	return &cobra.Command{
		Use:  "raw",
		Args: cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			raw := args[0]
			src.RawTransaction(raw)
		},
	}

}
func txBalanceTransfer() *cobra.Command {

	return &cobra.Command{
		Use:   "balance-transfer",
		Short: "transfer balance ",
		Long: `tranfer balance to another account:
			// pass 2 argument 
			// First -> accountAddress
			Second -> balance`,
		Args: cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			bal, _ := strconv.Atoi(args[1])
			src.BalanceTransfer(args[0], bal)
		},
	}

}
