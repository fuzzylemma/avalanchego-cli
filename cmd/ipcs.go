package cmd
import (
   "fmt"
   "context"
   "github.com/ava-labs/avalanchego/api/ipcs"
   "encoding/json"
   "github.com/spf13/cobra"
)

func AddIpcCommands(rootCmd *cobra.Command) {

   var ipcCmd = &cobra.Command{
      Use: "ipc [method]",
      Short: "avalanchego/api/ipcs",
      Long: `avalanchego/api/ipcs`,
      Run: func(cmd *cobra.Command, args []string) {},
   }

   ipcCmd.AddCommand(PublishBlockchainCmd())
   ipcCmd.AddCommand(UnpublishBlockchainCmd())
   ipcCmd.AddCommand(PublishedBlockchainsCmd())

   (*rootCmd).AddCommand(ipcCmd)
}

func IPCClient() ipcs.Client {
   uri := fmt.Sprintf("http://%s:%d", NodeAddress, NodePort)
   return ipcs.NewClient(uri)
}


func PublishBlockchainCmd() *cobra.Command {
   return &cobra.Command{
      Use: "publish blockchainID",
      Short: "avalanchego/api/ipcs PublishBlokchain method",
      Long: `avalanchego/api/ipcs`,
      Run: publishBlockchain,
      Args: cobra.ExactArgs(1),
   }
}
func UnpublishBlockchainCmd() *cobra.Command {
   return &cobra.Command{
      Use: "unpublish blockchainID",
      Short: "avalanchego/api/ipcs UnpublishedBlockchain method",
      Long: `avalanchego/api/ipcs`,
      Run: unpublishBlockchain,
      Args: cobra.ExactArgs(1),
   }
}
func PublishedBlockchainsCmd() *cobra.Command {
   return &cobra.Command{
      Use: "blockchains",
      Short: "avalanchego/api/ipcs GetPublishedBlockchains method",
      Long: `avalanchego/api/ipcs`,
      Run: publishedBlockchains,
      Args: cobra.ExactArgs(0),
   }
}


func publishBlockchain(cmd *cobra.Command, args []string) {
   ctx := context.Background()
   blockchainID := args[0]
   out, err := IPCClient().PublishBlockchain(ctx, blockchainID)
   check(err)
   fout, ferr := json.MarshalIndent(out, "", "   ")
   check(ferr)
   fmt.Println(string(fout))
}
func unpublishBlockchain(cmd *cobra.Command, args []string) {
   ctx := context.Background()
   blockchainID := args[0]
   out, err := IPCClient().UnpublishBlockchain(ctx, blockchainID)
   check(err)
   fout, ferr := json.MarshalIndent(out, "", "   ")
   check(ferr)
   fmt.Println(string(fout))
}
func publishedBlockchains(cmd *cobra.Command, args []string) {
   ctx := context.Background()
   out, err := IPCClient().GetPublishedBlockchains(ctx)
   check(err)
   fout, ferr := json.MarshalIndent(out, "", "   ")
   check(ferr)
   fmt.Println(string(fout))
}
