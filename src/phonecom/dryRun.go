package main

import (
	"fmt"
)

func showDryRunVerbose(cli CliParams) string {

	if (cli.dryRun || cli.verbose) {

		var baseMessage string = "Calling '%+v' with parameters: limit: %+v, offset: %+v, "
		var endMessage string = "expecting JSON response\n"

		if (cli.id != 0 && cli.accountId != 0 && cli.idSecondary != 0 && cli.input != "") {

			fmt.Printf(
				baseMessage + "account ID: %+v, id: %+v, second id: %+v, input json: %+v, " + endMessage,
				cli.command,
				cli.limit,
				cli.offset,
				cli.accountId,
				cli.id,
				cli.idSecondary,
			  cli.input)

		} else if (cli.id != 0 && cli.accountId != 0 && cli.input != "") {

			fmt.Printf(
				baseMessage +  "account ID: %+v, id: %+v, input json: %+v, " + endMessage,
				cli.command,
				cli.limit,
				cli.offset,
				cli.accountId,
				cli.id,
				cli.input)

		} else if (cli.accountId != 0 && cli.input != "") {

			fmt.Printf(
				baseMessage + "account ID: %+v, input json: %+v, " + endMessage,
				cli.command,
				cli.limit,
				cli.offset,
				cli.accountId,
				cli.input)

		} else if (cli.id != 0 && cli.accountId != 0 && cli.idSecondary != 0) {

			fmt.Printf(
				baseMessage + "account ID: %+v, id: %+v, second id: %+v, " + endMessage,
				cli.command,
				cli.limit,
				cli.offset,
				cli.accountId,
				cli.id,
			  cli.idSecondary)

		} else if (cli.id != 0 && cli.accountId != 0) {

			fmt.Printf(
				baseMessage + "account ID: %+v, id: %+v, " + endMessage,
				cli.command,
				cli.limit,
				cli.offset,
				cli.accountId,
				cli.id)

		} else if (cli.id != 0) {

			fmt.Printf(
				baseMessage + "ID: %+v, " + endMessage,
				cli.command,
				cli.limit,
				cli.offset,
				cli.id)

		} else {

			fmt.Printf(
				baseMessage + endMessage,
				cli.command,
				cli.limit,
				cli.offset)
		}

		if (cli.dryRun) {
			return "dryrun"
		}
	}

	return "continue"
}