package main

import "github.com/urfave/cli"

var defaultCommand = listAccounts
var defaultId = ""
var defaultInput = ""
var defaultFrom = ""
var defaultTo = ""
var defaultText = ""
var defaultTrunkName = ""
var defaultTrunkUri = ""
var defaultTrunkConcurrentCalls = -1
var defaultTrunkMaxMinutes = -1
const defaultAccountId = 1315091


func getCliFlags() []cli.Flag {

  return []cli.Flag{

    cli.StringFlag{
      Name: "command, c",
      Value: defaultCommand,
      Usage: "Phone.com API command that you want to execute",
    },
    cli.StringFlag{
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
      Value: defaultAccountId,
      Usage: "Phone.com API account to use in API calls",
    },
    cli.BoolFlag{
      Name: "dryrun, d",
      Usage: "Print the expected action without executing the API command",
    },
    cli.StringFlag{
      Name: "input, i",
      Value: defaultInput,
      Usage: "Specify the path to the JSON file for making the API call",
    },
    cli.BoolFlag{
      Name: "verbose, vr",
      Usage: "Activate verbose mode",
    },
    cli.StringFlag{
      Name: "contact",
      Usage: "Path to the local JSON descriptor to a contact object",
    },
    cli.StringFlag{
      Name: "billing-contact",
      Usage: "Path to the local JSON descriptor to a billing contact object",
    },
    cli.StringFlag{
      Name: "from",
      Usage: "The phone number of the SMS sender",
    },
    cli.StringFlag{
      Name: "to",
      Usage: "The phone number of the SMS receiver",
    },
    cli.StringFlag{
      Name: "text",
      Usage: "SMS body",
    },
		cli.StringFlag{
      Name: "name",
      Usage: "Trunk name",
    },
		cli.StringFlag{
      Name: "uri",
      Usage: "Trunk uri",
    },
		cli.IntFlag{
      Name: "max-concurrent-calls",
      Value: 60,
      Usage: "Maximum concurrent calls for the trunk",
    },
    cli.IntFlag{
      Name: "max-minutes-per-month",
      Value: 800,
      Usage: "Maximum of minutes per month for the trunk",
    },
    cli.StringFlag{
      Name: "filtersType, ft",
      Usage: "Type of filter",
    },
    cli.StringFlag{
      Name: "filtersValue, fv",
      Usage: "Type of filter",
    },
    cli.StringFlag{
      Name: "sortType, st",
      Usage: "Type of filter",
    },
    cli.StringFlag{
      Name: "sortValue, sv",
      Usage: "Type of filter",
    },
    cli.StringFlag{
      Name: "samplein, si",
      Usage: "Generate sample input json",
    },
    cli.StringFlag{
      Name: "sampleout, so",
      Usage: "Generate sample output json",
    },
  }
}
