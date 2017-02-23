package main

import (
  "github.com/urfave/cli"
  "testing"
)

var commandFlag = "-command"
var errorNotNullMessage = "expected no error from Run, got %s"

func createCli(endpoint string) error {

  app := cli.NewApp()

  defaultCommand = endpoint

  app.Flags = getCliFlags()
  configPath = "../../config.xml"

  app.Action = func(c *cli.Context) error {
    return execute(c)
  }

  app.Run([]string{commandFlag, endpoint})

  return nil
}

func assertErrorNotNull(t *testing.T, err error) {

  if err != nil {
    t.Fatalf(errorNotNullMessage, err)
  }
}
