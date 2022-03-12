package cmd

import (
   "fmt"
   "time"
   "context"
   "strconv"
   "github.com/ava-labs/avalanchego/api/health"
   "encoding/json"
   "github.com/spf13/cobra"
)


func AddHealthCommands(rootCmd *cobra.Command) {

   var healthCmd = &cobra.Command{
      Use: "health [method]",
      Short: "avalanchego/api/health",
      Long: `avalanchego/api/health`,
      Run: func(cmd *cobra.Command, args []string) {},
   }

   healthCmd.AddCommand(GetLivenessCmd())
   healthCmd.AddCommand(HealthCmd())
   healthCmd.AddCommand(AwaitHealthyCmd())

   (*rootCmd).AddCommand(healthCmd)
}

func HealthClient() health.Client {
   uri := fmt.Sprintf("http://%s:%d", NodeAddress, NodePort)
   return health.NewClient(uri)
}


func GetLivenessCmd() *cobra.Command {
   return &cobra.Command{
      Use: "liveness",
      Short: "avalanchego/api/health GetLiveness method",
      Long: `avalanchego/api/health`,
      Run: liveness,
      Args: cobra.ExactArgs(0),
   }
}
func HealthCmd() *cobra.Command {
   return &cobra.Command{
      Use: "healthy",
      Short: "avalanchego/api/health Health method",
      Long: `avalanchego/api/health`,
      Run: healthy,
      Args: cobra.ExactArgs(0),
   }
}
func AwaitHealthyCmd() *cobra.Command {
   return &cobra.Command{
      Use: "waitHealthy",
      Short: "avalanchego/api/health GetNodeID method",
      Long: `avalanchego/api/health`,
      Run: awaitHealthy,
      Args: cobra.ExactArgs(1),
   }
}

func liveness(cmd *cobra.Command, args []string) {
   ctx := context.Background()
   out, err := HealthClient().Liveness(ctx)
   check(err)
   fout, ferr := json.MarshalIndent(out, "", "   ")
   check(ferr)
   fmt.Println(string(fout))
}

func healthy(cmd *cobra.Command, args []string) {
   ctx := context.Background()
   out, err := HealthClient().Health(ctx)
   check(err)
   fout, ferr := json.MarshalIndent(out, "", "   ")
   check(ferr)
   fmt.Println(string(fout))
}

func awaitHealthy(cmd *cobra.Command, args []string) {
   ctx := context.Background()
   temp, _ := strconv.ParseInt(args[0], 10, 64)
   var interval time.Duration = time.Duration(temp)

   out, err := HealthClient().AwaitHealthy(ctx, interval)
   check(err)

   fout, ferr := json.MarshalIndent(out, "", "   ")
   check(ferr)
   fmt.Println(string(fout))
}
