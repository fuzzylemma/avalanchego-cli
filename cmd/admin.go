package cmd
import (
   "fmt"
   "time"
   "github.com/ava-labs/avalanchego/api/admin"
   "encoding/json"
   "github.com/spf13/cobra"
)

func AddAdminCommands(rootCmd *cobra.Command) {

   var adminCmd = &cobra.Command{
      Use: "admin [method]",
      Short: "avalanchego/api/admin",
      Long: `avalanchego/api/admin`,
      Run: func(cmd *cobra.Command, args []string) {},
   }

   adminCmd.AddCommand(StartCPUProfilerCmd())
   adminCmd.AddCommand(StopCPUProfilerCmd())
   adminCmd.AddCommand(MemoryProfileCmd())
   adminCmd.AddCommand(LockProfileCmd())
   adminCmd.AddCommand(AliasCmd())
   adminCmd.AddCommand(AliasChainCmd())
   adminCmd.AddCommand(ChainAliasesCmd())
   adminCmd.AddCommand(StackTraceCmd())

   (*rootCmd).AddCommand(adminCmd)
}

func AdminClient() *admin.Client {
   uri  := fmt.Sprintf("http://%s:%d", NodeAddress, NodePort)
   var timeout time.Duration = 1000000000
   return admin.NewClient(uri, timeout)
}

func StartCPUProfilerCmd() *cobra.Command {
   return &cobra.Command{
      Use: "nodeID",
      Short: "avalanchego/api/admin StartCPUProfiler method",
      Long: `avalanchego/api/admin`,
      Run: startCPUProfiler,
      Args: cobra.ExactArgs(0),
   }
}

func StopCPUProfilerCmd() *cobra.Command {
   return &cobra.Command{
      Use: "stop",
      Short: "avalanchego/api/admin StopCPUProfiler method",
      Long: `avalanchego/api/admin`,
      Run: stopCPUProfiler,
      Args: cobra.ExactArgs(0),
   }
}

func MemoryProfileCmd() *cobra.Command {
   return &cobra.Command{
      Use: "memory",
      Short: "avalanchego/api/admin MemoryProfile method",
      Long: `avalanchego/api/admin`,
      Run: memoryProfile,
      Args: cobra.ExactArgs(0),
   }
}

func LockProfileCmd() *cobra.Command {
   return &cobra.Command{
      Use: "lock",
      Short: "avalanchego/api/admin LockProfile method",
      Long: `avalanchego/api/admin`,
      Run: lockProfile,
      Args: cobra.ExactArgs(0),
   }
}


func AliasCmd() *cobra.Command {
   return &cobra.Command{
      Use: "alias",
      Short: "avalanchego/api/admin GetNodeID method",
      Long: `avalanchego/api/admin`,
      Run: alias,
      Args: cobra.ExactArgs(2),
   }
}

func AliasChainCmd() *cobra.Command {
   return &cobra.Command{
      Use: "aliasChain",
      Short: "avalanchego/api/admin AliasChain method",
      Long: `avalanchego/api/admin`,
      Run: aliasChain,
      Args: cobra.ExactArgs(2),
   }
}

func ChainAliasesCmd() *cobra.Command {
   return &cobra.Command{
      Use: "chainAliases",
      Short: "avalanchego/api/admin GetChainAliases method",
      Long: `avalanchego/api/admin`,
      Run: chainAliases,
      Args: cobra.ExactArgs(1),
   }
}

func StackTraceCmd() *cobra.Command {
   return &cobra.Command{
      Use: "stackTrace",
      Short: "avalanchego/api/admin StackTrace method",
      Long: `avalanchego/api/admin`,
      Run: stacktrace,
      Args: cobra.ExactArgs(0),
   }
}

func startCPUProfiler(cmd *cobra.Command, args []string) {
   out, err := AdminClient().StartCPUProfiler()
   check(err)
   fout, ferr := json.MarshalIndent(out, "", "   ")
   check(ferr)
   fmt.Println(string(fout))
}
func stopCPUProfiler(cmd *cobra.Command, args []string) {
   out, err := AdminClient().StopCPUProfiler()
   check(err)
   fout, ferr := json.MarshalIndent(out, "", "   ")
   check(ferr)
   fmt.Println(string(fout))
}
func memoryProfile(cmd *cobra.Command, args []string) {
   out, err := AdminClient().MemoryProfile()
   check(err)
   fout, ferr := json.MarshalIndent(out, "", "   ")
   check(ferr)
   fmt.Println(string(fout))
}

func lockProfile(cmd *cobra.Command, args []string) {
   out, err := AdminClient().LockProfile()
   check(err)
   fout, ferr := json.MarshalIndent(out, "", "   ")
   check(ferr)
   fmt.Println(string(fout))
}

func alias(cmd *cobra.Command, args []string) {
   endpoint, alas := args[0], args[1]
   out, err := AdminClient().Alias(endpoint, alas)
   check(err)
   fout, ferr := json.MarshalIndent(out, "", "   ")
   check(ferr)
   fmt.Println(string(fout))
}

func aliasChain(cmd *cobra.Command, args []string) {
   chain, alias := args[0], args[1]
   out, err := AdminClient().AliasChain(chain, alias)
   check(err)
   fout, ferr := json.MarshalIndent(out, "", "   ")
   check(ferr)
   fmt.Println(string(fout))
}


func chainAliases(cmd *cobra.Command, args []string) {
   chain := args[0]
   out, err := AdminClient().GetChainAliases(chain)
   check(err)
   fout, ferr := json.MarshalIndent(out, "", "   ")
   check(ferr)
   fmt.Println(string(fout))
}

func stacktrace(cmd *cobra.Command, args []string) {
   out, err := AdminClient().Stacktrace()
   check(err)
   fout, ferr := json.MarshalIndent(out, "", "   ")
   check(ferr)
   fmt.Println(string(fout))
}
