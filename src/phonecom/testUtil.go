package main

import (
  "github.com/urfave/cli"
  "testing"
	"strconv"
  "encoding/json"
)

var commandFlag = "-command"
var errorNotNullMessage = "expected no error from Run, got %s"

func createCliConfig() CliConfig{

	var cliConfig CliConfig
	cliConfig.Path = "../../config.xml"

	return cliConfig
}

func createCliWithJsonIn(endpoint string, path string) (error, map[string] interface{}) {

  defaultCommand = endpoint
  defaultInput = path

  return doCreateCli(endpoint)
}

func doCreateCli(endpoint string) (error, map[string] interface{}) {

  app := cli.NewApp()

  app.Flags = getCliFlags()

  var response (map[string] interface{})
  var err error

  app.Action = func(c *cli.Context) error {
    err, response = execute(c, createCliConfig())
    return err
  }

  app.Run([]string{commandFlag, endpoint})

  return err, response
}

func createCreateSmsCliWithJsonIn(endpoint string, path string, from string, to string, text string) (error, map[string] interface{}) {

  defaultCommand = endpoint
  defaultInput = path
  defaultFrom = from
  defaultTo = to
  defaultText = text

  return doCreateCli(endpoint)
}

func createCreateTrunkCliWithJsonIn(endpoint string, path string, trunkName string, trunkUri string, trunkConcurrentCalls int32, trunkMaxMinutes int32) (error, map[string] interface{}) {

  defaultCommand = endpoint
  defaultInput = path
  defaultTrunkName = trunkName
  defaultTrunkUri = trunkUri
  defaultTrunkConcurrentCalls = int(trunkConcurrentCalls)
  defaultTrunkMaxMinutes = int(trunkMaxMinutes)

  return doCreateCli(endpoint)
}

func createReplaceCliWithJsonIn(endpoint string, path string, id int) (error, map[string] interface{}) {

  defaultCommand = endpoint
  defaultInput = path
  defaultId = strconv.Itoa(id)

  return doCreateCli(endpoint)
}

func createCliWithIdIdSecondary(endpoint string, id int, idSecondary int) (error, map[string] interface{}) {

  defaultCommand = endpoint
  defaultId = strconv.Itoa(id)
  defaultIdSecondary = strconv.Itoa(idSecondary)

  return doCreateCli(endpoint)
}

func createReplaceContactCliWithJsonIn(endpoint string, path string, id int, idSecondary int) (error, map[string] interface{}) {

  defaultCommand = endpoint
  defaultInput = path
  defaultId = strconv.Itoa(id)
  defaultIdSecondary = strconv.Itoa(idSecondary)

  return doCreateCli(endpoint)
}

func createReplaceTrunkCliWithJsonIn(endpoint string, path string, trunkName string, trunkUri string, trunkConcurrentCalls int32, trunkMaxMinutes int32, id int) (error, map[string] interface{}) {

  defaultCommand = endpoint
  defaultInput = path
  defaultId = strconv.Itoa(id)
  defaultTrunkName = trunkName
  defaultTrunkUri = trunkUri
  defaultTrunkConcurrentCalls = int(trunkConcurrentCalls)
  defaultTrunkMaxMinutes = int(trunkMaxMinutes)

  return doCreateCli(endpoint)
}

func createCli(endpoint string) (error, map[string] interface{}) {

  defaultCommand = endpoint

  return doCreateCli(endpoint)
}

func createGetCliStringId(endpoint string, id string) (error, map[string] interface{}) {

	if (id == "") {
		return nil, nil
	}

	defaultCommand = endpoint
	defaultId = id

  return doCreateCli(endpoint)
}

func createCliWithId(endpoint string, id int) (error, map[string] interface{}) {

  if (id <= 0) {
    return nil, nil
  }

  defaultCommand = endpoint
  defaultId = strconv.Itoa(id)

  return doCreateCli(endpoint)
}

func assertErrorNotNull(t *testing.T, err error) {

  if err != nil {
    t.Fatalf(errorNotNullMessage, err)
  }
}

func getFirstIdString(json map[string] interface{}) string {

	items := json["items"].([]interface{})
	if (len(items) > 0) {
		firstItem := items[0].(map[string] interface{})
		firstId := firstItem["id"].(string)
		return firstId
	}

	return ""
}

func getFirstId(jsonObject map[string] interface{}) int {

  items := jsonObject["items"].([]interface{})
  if (len(items) > 0) {
    firstItem := items[0].(map[string] interface{})
    firstId := firstItem["id"].(json.Number)
    idToReturn, _ := json.Number.Int64(firstId)
    return int(idToReturn)
  }

  return 0
}

func getFirstAvailablePhoneNumber(json map[string] interface{}) string {

  items := json["items"].([]interface{})
  if (len(items) > 0) {
    firstItem := items[0].(map[string] interface{})
    firstNumber := firstItem["phone_number"].(string)
    return string(firstNumber)
  }

  return ""
}

func getId(jsonObject map[string] interface{}) int {

  id := jsonObject["id"].(json.Number)
  idToReturn, _ := json.Number.Int64(id)
  return int(idToReturn)
}

func getName(json map[string] interface{}) string {

  name := json["name"].(string)
  return name
}
