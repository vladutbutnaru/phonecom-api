package main

import (
  "github.com/urfave/cli"
  "testing"
)

var commandFlag = "-command"
var errorNotNullMessage = "expected no error from Run, got %s"

func createCli(endpoint string) (error, map[string] interface{}) {

  app := cli.NewApp()

  defaultCommand = endpoint

  app.Flags = getCliFlags()
  configPath = "../../config.xml"
  var json (map[string] interface{})
  var err error

  app.Action = func(c *cli.Context) error {
    err, json = execute(c)
    return err
  }

  app.Run([]string{commandFlag, endpoint})

  return err, json
}

func assertErrorNotNull(t *testing.T, err error) {

  if err != nil {
    t.Fatalf(errorNotNullMessage, err)
  }
}
