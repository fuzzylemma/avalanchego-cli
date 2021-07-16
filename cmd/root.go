package cmd


import (
   "log"
   "fmt"

   "github.com/spf13/cobra"
   //"github.com/spf13/viper"

)


var configFilename string

/*
var rootCmd = &cobra.Command{
   Use: "avaxgoclt",
   Short: "cli wrapper for github.com/ava-labs/avalanchgo/api",
   Long: `cli wrapper for github.com/ava-labs/avalanchego/api to help control avax nodes`,
}
*/


func check(err error) {
   if err != nil {
      log.Fatal(err)
   }
}


func init() {
   //cobra.OnInitialize(initConfig)

   // get home dir
   // home := homedirallocationa0
   // defaultConfigFilename := path.Join(home, ".config/avaxgoclt.yaml")

   //info.AddCommand(&rootCmd)
}

func initConfig() {
   //viper stuff
}


type AvaxNode struct {
   Name string
   Address string
   Port int
}

/*
type User struct {
   Name string
   User string
   Pass string
}
*/

func PrettyPrint(data []byte) {

}

//var avaxNode AvaxNode
func Execute () {

   var rootCmd = &cobra.Command{
      Use: "avaxgo [sub]",
      Short: "avaxgo api cli",
      Long: `avaxgo api cli`,
      Run: func(cmd *cobra.Command, args []string) {
         fmt.Printf("Inside rootCmd Run with args: %v\n", args)
      },
   }


   rootCmd.PersistentFlags().StringVarP(&NodeAddress, "address", "a", "localhost", "node address")
   rootCmd.PersistentFlags().IntVarP(&NodePort, "port", "p", 9650, "node port")
   rootCmd.PersistentFlags().StringVarP(&Username, "username", "u", "", "avax node username")
   rootCmd.PersistentFlags().StringVarP(&Password, "password", "w", "", "avax node password")

   AddInfoCommands(rootCmd)
   AddKeystoreCommands(rootCmd)
   AddIpcCommands(rootCmd)
   AddHealthCommands(rootCmd)
   AddAuthCommands(rootCmd)
   AddAdminCommands(rootCmd)

   err := rootCmd.Execute()
   check(err)
}
