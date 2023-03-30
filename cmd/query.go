package cmd

// import (
// 	"log"
// 	"strconv"

// 	"github.com/spf13/cobra"
// )

// func init() {
// 	rootCmd.AddCommand(newQueryCmd())
// }

// func newQueryCmd() *cobra.Command {
// 	rootCmd := cobra.Command{
// 		Use:     "query",
// 		Aliases: []string{"q"},
// 	}
// 	rootCmd.AddCommand(
// 		queryBalance(),
// 		queryBlockHash(),
// 	)

// 	return &rootCmd
// }

// func queryBalance() *cobra.Command {
// 	return &cobra.Command{
// 		Use:   "balance",
// 		Short: "Query balance of an account ",
// 		Long: `
// 			check the balance of a query:
// 			argument address
// 		`,
// 		Aliases: []string{"bal"},
// 		Args:    cobra.ExactArgs(1),
// 		Run: func(cmd *cobra.Command, args []string) {
// 			src.GetBalance(args[0])
// 		},
// 	}
// }

// func queryBlockHash() *cobra.Command {
// 	return &cobra.Command{
// 		Use:   "blockhash",
// 		Short: "Get blockhash ",
// 		Long:  "Get blockhash of given block number ",
// 		Run: func(cmd *cobra.Command, args []string) {

// 			blockNumber, err := strconv.Atoi(args[0])
// 			if err != nil {
// 				log.Fatal("Invalid BlockNumber")
// 				return
// 			}
// 			src.GetBlockHash(blockNumber)
// 		},
// 	}

// }
