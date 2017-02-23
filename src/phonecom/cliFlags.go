package main

import "github.com/urfave/cli"

var defaultCommand = listAccounts
var defaultId = 0

func getCliFlags() []cli.Flag {

  return []cli.Flag{

    cli.StringFlag{
      Name: "command, c",
      Value: defaultCommand,
      Usage: "Phone.com API command that you want to execute",
    },
    cli.IntFlag{
      Name: "id",
      Value: defaultId,
      Usage: "ID of entity you want to operate",
    },
    cli.IntFlag{
      Name: "id-secondary, is",
      Value: 0,
      Usage: "Secondary ID of entity you want to operate",
    },
    cli.IntFlag{
      Name: "limit, l",
      Value: 5,
      Usage: "Upper limit of results you want to fetch",
    },
    cli.IntFlag{
      Name: "offset, o",
      Value: 0,
      Usage: "Offset of results you want to fetch",
    },
    cli.IntFlag{
      Name: "account, a",
      Value: 1315091,
      Usage: "Phone.com API account to use in API calls",
    },
    cli.StringFlag{
      Name: "dryrun, d",
      Value: "",
      Usage: "Print the expected action without executing the API command",
    },
    cli.StringFlag{
      Name: "input, i",
      Value: "",
      Usage: "Specify the path to the JSON file for making the API call",
    },
    cli.StringFlag{
      Name: "verbose, vr",
      Value: "",
      Usage: "Activate verbose mode",
    },

  }
}
