package main

import (
	"fmt"
)

func showDryrunVerbose(cli CliParams) string {

	if (cli.dryRun != "" || cli.verbose != "") {

		if (cli.id != 0 && cli.accountId != 0 && cli.idSecondary != 0 && cli.input != "") {

			fmt.Printf(
				"Calling '%+v' with parameters: limit: %+v, offset: %+v, account ID: %+v, id: %+v, second id: %+v, input json: %+v, expecting JSON response\n",
				cli.command,
				cli.limit,
				cli.offset,
				cli.accountId,
				cli.id,
				cli.idSecondary,
			  cli.input)

		} else if (cli.id != 0 && cli.accountId != 0 && cli.input != "") {

			fmt.Printf(
				"Calling '%+v' with parameters: limit: %+v, offset: %+v, account ID: %+v, id: %+v, input json: %+v, expecting JSON response\n",
				cli.command,
				cli.limit,
				cli.offset,
				cli.accountId,
				cli.id,
				cli.input)

		} else if (cli.accountId != 0 && cli.input != "") {

			fmt.Printf(
				"Calling '%+v' with parameters: limit: %+v, offset: %+v, account ID: %+v, input json: %+v, expecting JSON response\n",
				cli.command,
				cli.limit,
				cli.offset,
				cli.accountId,
				cli.input)

		} else if (cli.id != 0 && cli.accountId != 0 && cli.idSecondary != 0) {

			fmt.Printf(
				"Calling '%+v' with parameters: limit: %+v, offset: %+v, account ID: %+v, id: %+v, second id: %+v, expecting JSON response\n",
				cli.command,
				cli.limit,
				cli.offset,
				cli.accountId,
				cli.id,
			  cli.idSecondary)

		} else if (cli.id != 0 && cli.accountId != 0) {

			fmt.Printf(
				"Calling '%+v' with parameters: limit: %+v, offset: %+v, account ID: %+v, id: %+v, expecting JSON response\n",
				cli.command,
				cli.limit,
				cli.offset,
				cli.accountId,
				cli.id)

		} else if (cli.id != 0) {

			fmt.Printf(
				"Calling '%+v' with parameters: limit: %+v, offset: %+v, ID: %+v, expecting JSON response\n",
				cli.command,
				cli.limit,
				cli.offset,
				cli.id)

		} else {

			fmt.Printf(
				"Calling '%+v' with parameters: limit: %+v, offset: %+v, expecting JSON response\n",
				cli.command,
				cli.limit,
				cli.offset)
		}

		if (cli.dryRun != "") {
			return "dryrun"
		}
	}

	return "continue"
}