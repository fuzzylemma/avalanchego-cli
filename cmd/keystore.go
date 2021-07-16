package cmd
import (
   "fmt"
   "time"
   "encoding/json"
   "github.com/ava-labs/avalanchego/api/keystore"
   "github.com/ava-labs/avalanchego/api"
   "github.com/spf13/cobra"
)

func AddKeystoreCommands(rootCmd *cobra.Command) {

   var keystoreCmd = &cobra.Command{
      Use: "keystore [method]",
      Short: "avalanchego/api/keystore",
      Long: `avalanchego/api/keystore`,
      Run: func(cmd *cobra.Command, args []string) {},
   }

   keystoreCmd.AddCommand(ListUsers())
   keystoreCmd.AddCommand(CreateUser())
   keystoreCmd.AddCommand(ExportUser())
   keystoreCmd.AddCommand(ImportUser())
   keystoreCmd.AddCommand(DeleteUser())
   (*rootCmd).AddCommand(keystoreCmd)
}

func KeystoreClient() *keystore.Client {
   uri  := fmt.Sprintf("http://%s:%d", NodeAddress, NodePort)
   var timeout time.Duration = 1000000000
   return keystore.NewClient(uri, timeout)
}

func ListUsers() *cobra.Command {
   return &cobra.Command{
      Use: "list",
      Short: "avalanchego/api/keystore ListUsers method",
      Long: `avalanchego/api/keystore ListUsers method`,
      Args: cobra.ExactArgs(0),
      Run: listUsers,
   }
}

func CreateUser() *cobra.Command {
   return &cobra.Command{
      Use: "create user pass",
      Short: "avalanchego/api/keystore CreateUsers method",
      Long: `avalanchego/api/keystore CreateUsers method`,
      Args: cobra.ExactArgs(2),
      Run: createUser,
   }
}

func ExportUser() *cobra.Command {
   return &cobra.Command{
      Use: "export user pass",
      Short: "avalanchego/api/keystore ExportUser method",
      Long: `avalanchego/api/keystore ExportUser method`,
      Args: cobra.ExactArgs(2),
      Run: exportUser,
   }
}

func ImportUser() *cobra.Command {
   return &cobra.Command{
      Use: "import user pass encoding",
      Short: "avalanchego/api/keystore ImportUser method",
      Long: `avalanchego/api/keystore ImportUser method`,
      Args: cobra.ExactArgs(3),
      Run: importUser,
   }
}

func DeleteUser() *cobra.Command {
   return &cobra.Command{
      Use: "delete user pass",
      Short: "avalanchego/api/keystore DeleteUser method",
      Long: `avalanchego/api/keystore DeleteUser method`,
      Args: cobra.ExactArgs(2),
      Run: deleteUser,
   }
}

func listUsers(command *cobra.Command, args []string) {
   out, err := KeystoreClient().ListUsers()
   check(err)
   fout, ferr := json.MarshalIndent(out, "", "   ")
   check(ferr)
   fmt.Println(string(fout))
}

func createUser(command *cobra.Command, args []string) {
   userPass := api.UserPass{Username: args[0], Password: args[1] }
   out, err := KeystoreClient().CreateUser(userPass)
   check(err)
   fout, ferr := json.MarshalIndent(out, "", "   ")
   check(ferr)
   fmt.Println(string(fout))
}

func exportUser(command *cobra.Command, args []string) {
   userPass := api.UserPass{Username: args[0], Password: args[1] }
   out, err := KeystoreClient().ExportUser(userPass)
   check(err)
   fout, ferr := json.MarshalIndent(out, "", "   ")
   check(ferr)
   fmt.Println(string(fout))
}

func importUser(command *cobra.Command, args []string) {
   userPass := api.UserPass{Username: args[0], Password: args[1] }
   encoding := []byte(args[2])
   out, err := KeystoreClient().ImportUser(userPass, encoding)
   check(err)
   fout, ferr := json.MarshalIndent(out, "", "   ")
   check(ferr)
   fmt.Println(string(fout))
}

func deleteUser(command *cobra.Command, args []string) {
   userPass := api.UserPass{Username: args[0], Password: args[1] }
   out, err := KeystoreClient().DeleteUser(userPass)
   check(err)
   fout, ferr := json.MarshalIndent(out, "", "   ")
   check(ferr)
   fmt.Println(string(fout))
}
