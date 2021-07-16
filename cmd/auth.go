package cmd

import (
   "fmt"
   "time"
   "github.com/ava-labs/avalanchego/api/auth"
   "github.com/ava-labs/avalanchego/utils/rpc"
   "github.com/ava-labs/avalanchego/api"
   "encoding/json"
   "github.com/spf13/cobra"
)



func AddAuthCommands(rootCmd *cobra.Command) {

   var authCmd = &cobra.Command{
      Use: "auth [method]",
      Short: "avalanchego/api/auth",
      Long: `avalanchego/api/auth`,
      Run: func(cmd *cobra.Command, args []string) {},
   }

   authCmd.AddCommand(NewTokenCmd())
   authCmd.AddCommand(RevokeTokenCmd())
   authCmd.AddCommand(ChangePasswordCmd())

   (*rootCmd).AddCommand(authCmd)
}

func AuthClient() *Client {
   uri  := fmt.Sprintf("http://%s:%d", NodeAddress, NodePort)
   var timeout time.Duration = 1000000000
   return ANewClient(uri, timeout)
}


func NewTokenCmd() *cobra.Command {
   return &cobra.Command{
      Use: "newToken",
      Short: "avalanchego/api/auth NewToken method",
      Long: `avalanchego/api/auth`,
      Run: newToken,
   }
}

func RevokeTokenCmd() *cobra.Command {
   return &cobra.Command{
      Use: "revokeToken token password",
      Short: "avalanchego/api/auth NewToken method",
      Long: `avalanchego/api/auth`,
      Run: revokeToken,
      Args: cobra.ExactArgs(2),
   }
}

func ChangePasswordCmd() *cobra.Command {
   return &cobra.Command{
      Use: "password old new",
      Short: "avalanchego/api/auth NewToken method",
      Long: `avalanchego/api/auth`,
      Run: changePassword,
      Args: cobra.ExactArgs(2),
   }
}



func newToken(cmd *cobra.Command, args []string) {
   password := args[0]
   endpoints := args[1:]
   out, err := AuthClient().GetNewToken(password, endpoints)
   check(err)
   fout, ferr := json.MarshalIndent(out, "", "   ")
   check(ferr)
   fmt.Println(string(fout))
}
func revokeToken(cmd *cobra.Command, args []string) {
   token := args[0]
   password := args[1]
   out, err := AuthClient().RevokeToken(token, password)
   check(err)
   fout, ferr := json.MarshalIndent(out, "", "   ")
   check(ferr)
   fmt.Println(string(fout))
}
func changePassword(cmd *cobra.Command, args []string) {
   oldPassword := args[0]
   newPassword := args[1]
   out, err := AuthClient().ChangePassword(oldPassword, newPassword)
   check(err)
   fout, ferr := json.MarshalIndent(out, "", "   ")
   check(ferr)
   fmt.Println(string(fout))
}
// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

type Client struct {
   requester rpc.EndpointRequester
}

func ANewClient(uri string, requestTimeout time.Duration) *Client {
   return &Client{
      requester: rpc.NewEndpointRequester(uri, "/ext/auth", "auth", requestTimeout),
   }
}

func (c *Client) GetNewToken(password string, endpoints []string) (string, error) {
   res := &auth.Token{}
   err := c.requester.SendRequest("newToken", &auth.NewTokenArgs{
            Password: auth.Password{password},
            Endpoints: endpoints,
   }, res)
   return res.Token, err
}


func (c *Client) RevokeToken(token, password string) (bool, error) {
   res := &api.SuccessResponse{}
   err := c.requester.SendRequest("revokeToken", &auth.RevokeTokenArgs{
            Password: auth.Password{password},
            Token: auth.Token{token},
   }, res)
   return res.Success, err
}

func (c *Client) ChangePassword(oldPassword, newPassword string) (bool, error) {
   res := &api.SuccessResponse{}
   err := c.requester.SendRequest("changePassword", &auth.ChangePasswordArgs{
            OldPassword: oldPassword,
            NewPassword: newPassword,
   }, res)
   return res.Success, err
}


