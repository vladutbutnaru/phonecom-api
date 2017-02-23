package main

import "github.com/urfave/cli"

func createCli() *cli.App {

  app := cli.NewApp()

  app.Flags = cliFlags
  configPath = "../../config.xml"

  app.Action = func(c *cli.Context) error {
    return execute(c)
  }

  return app
}
