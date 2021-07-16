package cmd

import (
   //"github.com/spf13/viper"
)

var NodeAddress string
var NodePort int
var Username string
var Password string

var configBaseDir = "~/.avaxgoclt/"
var nodeConfig    = configBaseDir + "node"
var userConfig    = configBaseDir + "user"
var tokenConfig   = configBaseDir + "token"

//func initConfig() {}
//func Add(config string, entity func) {}
//func Delete(config string, entity func) {}
