# avaxgo
cli wrapper for https://github.com/ava-labs/avalanchego/api
avaxgo api cli

## Basic usage
```
Usage:
  avaxgo [sub] [flags]
  avaxgo [command]

Available Commands:
  admin       avalanchego/api/admin
  auth        avalanchego/api/auth
  completion  Generate the autocompletion script for the specified shell
  health      avalanchego/api/health
  help        Help about any command
  info        avalanchego/api/info
  ipc         avalanchego/api/ipcs
  keystore    avalanchego/api/keystore

Flags:
  -a, --address string    node address (default "localhost")
  -h, --help              help for avaxgo
  -w, --password string   avax node password
  -p, --port int          node port (default 9650)
  -u, --username string   avax node username

Use "avaxgo [command] --help" for more information about a command.
```

If needed, parameters follow same ordering as in avalanchego. See their [api](https://github.com/ava-labs/avalanche-docs/tree/master/docs/build/avalanchego-apis) for more info.
