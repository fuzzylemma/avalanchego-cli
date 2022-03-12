package cmd

import (
   "fmt"
   "context"
   "encoding/json"
   "github.com/ava-labs/avalanchego/api/info"
   "github.com/spf13/cobra"
)

func AddInfoCommands(rootCmd *cobra.Command) {

   var infoCmd = &cobra.Command{
      Use: "info [method]",
      Short: "avalanchego/api/info",
      Long: `avalanchego/api/info`,
      Run: func(cmd *cobra.Command, args []string) {},
   }

   infoCmd.AddCommand(NodeIDCmd())
   infoCmd.AddCommand(NetworkIDCmd())
   infoCmd.AddCommand(NetworkNameCmd())
   infoCmd.AddCommand(BlockchainIDCmd())
   infoCmd.AddCommand(PeersCmd())
   infoCmd.AddCommand(IsBootstrappedCmd())
   infoCmd.AddCommand(TxFeeCmd())
   infoCmd.AddCommand(NodeInfomd())

   (*rootCmd).AddCommand(infoCmd)
}


func InfoClient() info.Client {
   uri := fmt.Sprintf("http://%s:%d", NodeAddress, NodePort)
   return info.NewClient(uri)
}

func NodeIDCmd() *cobra.Command {
   return &cobra.Command{
      Use: "nodeID",
      Short: "avalanchego/api/info GetNodeID method",
      Long: `avalanchego/api/info`,
      Run: nodeID,
   }
}

func NetworkIDCmd() *cobra.Command {
   return &cobra.Command{
      Use: "networkID",
      Short: "avalanchego/api/info GetNetworkID method",
      Long: `avalanchego/api/info`,
      Run: networkID,
   }
}

func NetworkNameCmd() *cobra.Command {
   return &cobra.Command{
      Use: "networkName",
      Short: "avalanchego/api/info GetNetworkName method",
      Long: `avalanchego/api/info`,
      Run: networkName,
   }
}

func BlockchainIDCmd() *cobra.Command {
   return &cobra.Command{
      Use: "nodeID",
      Short: "avalanchego/api/info GetBlockchainID method",
      Long: `avalanchego/api/info`,
      Run:  blockchainID,
      Args: cobra.ExactArgs(1),
   }
}

func PeersCmd() *cobra.Command {
   return &cobra.Command{
      Use: "peers",
      Short: "avalanchego/api/info Peers method",
      Long: `avalanchego/api/info`,
      Run: peers,
   }
}

func IsBootstrappedCmd() *cobra.Command {
   return &cobra.Command{
      Use: "isBootstrapped",
      Short: "avalanchego/api/info IsBootstrapped method",
      Long: `avalanchego/api/info`,
      Run: isBootstrapped,
      Args: cobra.ExactArgs(1),
   }
}

func TxFeeCmd() *cobra.Command {
   return &cobra.Command{
      Use: "txFee",
      Short: "avalanchego/api/info GetTxFee method",
      Long: `avalanchego/api/info`,
      Run: txFee,
   }
}

func NodeInfomd() *cobra.Command {
   return &cobra.Command{
      Use: "nodeIP",
      Short: "avalanchego/api/info GetNodeIP method",
      Long: `avalanchego/api/info`,
      Run: nodeIP,
   }
}

func nodeID(cmd *cobra.Command, args []string) {
   ctx := context.Background()
   out, err := InfoClient().GetNodeID(ctx)
   check(err)
   fout, ferr := json.MarshalIndent(out, "", "   ")
   check(ferr)
   fmt.Println(string(fout))
}
func networkID(cmd *cobra.Command, args []string) {
   ctx := context.Background()
   out, err := InfoClient().GetNetworkID(ctx)
   check(err)
   fout, ferr := json.MarshalIndent(out, "", "   ")
   check(ferr)
   fmt.Println(string(fout))
}
func networkName(cmd *cobra.Command, args []string) {
   ctx := context.Background()
   out, err := InfoClient().GetNetworkName(ctx)
   check(err)
   fout, ferr := json.MarshalIndent(out, "", "   ")
   check(ferr)
   fmt.Println(string(fout))
}
func blockchainID(cmd *cobra.Command, args []string) {
   ctx := context.Background()
   alias := args[0]
   out, err := InfoClient().GetBlockchainID(ctx, alias)
   check(err)
   fout, ferr := json.MarshalIndent(out, "", "   ")
   check(ferr)
   fmt.Println(string(fout))
}
func peers(cmd *cobra.Command, args []string) {
   ctx := context.Background()
   out, err := InfoClient().Peers(ctx)
   check(err)
   fout, ferr := json.MarshalIndent(out, "", "   ")
   check(ferr)
   fmt.Println(string(fout))
}
func isBootstrapped(cmd *cobra.Command, args []string) {
   ctx := context.Background()
   chain := args[0]
   out, err := InfoClient().IsBootstrapped(ctx, chain)
   check(err)
   fout, ferr := json.MarshalIndent(out, "", "   ")
   check(ferr)
   fmt.Println(string(fout))
}
func txFee(cmd *cobra.Command, args []string) {
   ctx := context.Background()
   out, err := InfoClient().GetTxFee(ctx)
   check(err)
   fout, ferr := json.MarshalIndent(out, "", "   ")
   check(ferr)
   fmt.Println(string(fout))
}
func nodeIP(cmd *cobra.Command, args []string) {
   ctx := context.Background()
   out, err := InfoClient().GetNodeIP(ctx)
   check(err)
   fout, ferr := json.MarshalIndent(out, "", "   ")
   check(ferr)
   fmt.Println(string(fout))
}
