package main

import (
  "github.com/urfave/cli"
  "testing"
  "encoding/json"
	"strconv"
  "fmt"
)

var commandFlag = "-command"
var errorNotNullMessage = "expected no error from Run, got %s"

func createCliConfig() CliConfig{

	var cliConfig CliConfig
	cliConfig.Path = "../../config.xml"

	return cliConfig
}

func createCliWithJsonIn(endpoint string, path string) (error, map[string] interface{}) {

	flags := []cli.Flag{

		cli.StringFlag{
			Name: commandLong,
			Value: endpoint,
		},
		cli.StringFlag{
			Name: inputLong,
			Value: path,
		},
	}

  return doCreateCli(endpoint, flags)
}

func doCreateCli(endpoint string, flags []cli.Flag) (error, map[string] interface{}) {

  app := cli.NewApp()

  app.Flags = flags

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

	flags := []cli.Flag{

		cli.StringFlag{
			Name: commandLong,
			Value: endpoint,
		},
		cli.StringFlag{
			Name: inputLong,
			Value: path,
		},
		cli.StringFlag{
			Name: fromLong,
			Value: from,
		},
		cli.StringFlag{
			Name: toLong,
			Value: to,
		},
		cli.StringFlag{
			Name: textLong,
			Value: text,
		},
	}

  return doCreateCli(endpoint, flags)
}

func createCreateTrunkCliWithJsonIn(endpoint string, path string, trunkName string, trunkUri string, trunkConcurrentCalls int, trunkMaxMinutes int) (error, map[string] interface{}) {

	flags := []cli.Flag{

		cli.StringFlag{
			Name: commandLong,
			Value: endpoint,
		},
		cli.StringFlag{
			Name: inputLong,
			Value: path,
		},
		cli.StringFlag{
			Name: trunkNameLong,
			Value: trunkName,
		},
		cli.StringFlag{
			Name: trunkUriLong,
			Value: trunkUri,
		},
		cli.IntFlag{
			Name: maxConcurrentCallsLong,
			Value: trunkConcurrentCalls,
		},
		cli.IntFlag{
			Name: maxMinutesPerMonthLong,
			Value: trunkMaxMinutes,
		},
		cli.BoolTFlag{
			Name: verboseLong,
		},
	}

  return doCreateCli(endpoint, flags)
}

func createReplaceCliWithJsonIn(endpoint string, path string, id int) (error, map[string] interface{}) {

	flags := []cli.Flag{

		cli.StringFlag{
			Name: commandLong,
			Value: endpoint,
		},
		cli.StringFlag{
			Name: inputLong,
			Value: path,
		},
		cli.StringFlag{
			Name: idLong,
			Value: strconv.Itoa(id),
		},
		cli.BoolTFlag{
			Name: verboseLong,
		},
	}

  return doCreateCli(endpoint, flags)
}

func createCliWithIdIdSecondary(endpoint string, id int, idSecondary int) (error, map[string] interface{}) {

	flags := []cli.Flag{

		cli.StringFlag{
			Name: commandLong,
			Value: endpoint,
		},
		cli.StringFlag{
			Name: idLong,
			Value: strconv.Itoa(id),
		},
		cli.StringFlag{
			Name: idSecondaryLong,
			Value: strconv.Itoa(idSecondary),
		},
	}

  return doCreateCli(endpoint, flags)
}

func createReplaceContactCliWithJsonIn(endpoint string, path string, id int, idSecondary int) (error, map[string] interface{}) {

	flags := []cli.Flag{

		cli.StringFlag{
			Name: commandLong,
			Value: endpoint,
		},
		cli.StringFlag{
			Name: inputLong,
			Value: path,
		},
		cli.StringFlag{
			Name: idLong,
			Value: strconv.Itoa(id),
		},
		cli.StringFlag{
			Name: idSecondaryLong,
			Value: strconv.Itoa(idSecondary),
		},
	}

  return doCreateCli(endpoint, flags)
}

func createReplaceTrunkCliWithJsonIn(endpoint string, path string, trunkName string, trunkUri string, trunkConcurrentCalls int, trunkMaxMinutes int, id int) (error, map[string] interface{}) {

	flags := []cli.Flag{

		cli.StringFlag{
			Name: commandLong,
			Value: endpoint,
		},
		cli.StringFlag{
			Name: idLong,
			Value: strconv.Itoa(id),
		},
		cli.StringFlag{
			Name: inputLong,
			Value: path,
		},
		cli.StringFlag{
			Name: trunkNameLong,
			Value: trunkName,
		},
		cli.StringFlag{
			Name: trunkUriLong,
			Value: trunkUri,
		},
		cli.IntFlag{
			Name: maxConcurrentCallsLong,
			Value: trunkConcurrentCalls,
		},
		cli.IntFlag{
			Name: maxMinutesPerMonthLong,
			Value: trunkMaxMinutes,
		},
		cli.BoolTFlag{
			Name: verboseLong,
		},
	}

  return doCreateCli(endpoint, flags)
}

func createCliListAvailablePhoneNumbers(endpoint string, filtersPhoneNumber []string, filtersCountryCode []string, filtersNpa []string, filtersNxx []string, filtersXxxx []string, filtersCity []string, filtersProvince []string, filtersCountry []string, filtersPrice []string, filtersCategory []string) (error, map[string] interface{}) {

  values := filtersPhoneNumber[0] + ";" + filtersCountryCode[0] + ";" + filtersNpa[0] + ";" + filtersNxx[0] + ";" + filtersXxxx[0] + ";" + filtersCity[0] + ";" + filtersProvince[0] + ";" + filtersCountry[0] + ";" + filtersPrice[0]/* + ";" + filtersCategory[0]*/
  flags := []cli.Flag{
    cli.StringFlag{
      Name: commandLong,
      Value: endpoint,
    },
    cli.BoolTFlag{
      Name: verboseLong,
    },
    cli.IntFlag{
      Name: limitLong,
      Value: 25,
    },
    cli.StringFlag{
      Name: filtersTypeLong,
      Value: "phone_number;country_code;npa;nxx;xxxx;city;province;country;price",
    },
    cli.StringFlag{
      Name: filtersValueLong,
      Value: values,
    },
    cli.BoolTFlag{
      Name: fullListLong,
    },
  }

  return doCreateCli(endpoint, flags)
}

func createCliListAvailablePhoneNumberRegions(endpoint string, filtersCountryCode []string, filtersNpa []string, filtersNxx []string, filtersIsTollFree []string, filtersCity []string, filtersProvincePostalCode []string, filtersCountryPostalCode []string) (error, map[string] interface{}) {

  values := filtersCountryCode[0] + ";" + filtersNpa[0] + ";" + filtersNxx[0] + ";" + filtersIsTollFree[0] + ";" + filtersCity[0] + ";" + filtersProvincePostalCode[0] + ";" + filtersCountryPostalCode[0]
  fmt.Println(values)
  flags := []cli.Flag{
    cli.StringFlag{
      Name: commandLong,
      Value: endpoint,
    },
    cli.BoolTFlag{
      Name: verboseLong,
    },
    cli.IntFlag{
      Name: limitLong,
      Value: 25,
    },
    cli.StringFlag{
      Name: filtersTypeLong,
      Value: "country_code;npa;nxx;is_toll_free;city;province_postal_code;country_postal_code",
    },
    cli.StringFlag{
      Name: filtersValueLong,
      Value: values,
    },
    cli.BoolTFlag{
      Name: fullListLong,
    },
  }

  return doCreateCli(endpoint, flags)
}

func createCli(endpoint string) (error, map[string] interface{}) {

	flags := []cli.Flag{

		cli.StringFlag{
			Name: commandLong,
			Value: endpoint,
		},
		cli.BoolTFlag{
			Name: verboseLong,
		},
		cli.IntFlag{
			Name: limitLong,
			Value: 5,
		},
	}

  return doCreateCli(endpoint, flags)
}

func createGetCliStringId(endpoint string, id string) (error, map[string] interface{}) {

	if (id == "") {
		return nil, nil
	}

	flags := []cli.Flag{

		cli.StringFlag{
			Name: commandLong,
			Value: endpoint,
		},
		cli.StringFlag{
			Name: idLong,
			Value: id,
		},
	}

  return doCreateCli(endpoint, flags)
}

func createCliWithId(endpoint string, id int) (error, map[string] interface{}) {

  if (id <= 0) {
    return nil, nil
  }

	flags := []cli.Flag{

		cli.StringFlag{
			Name: commandLong,
			Value: endpoint,
		},
		cli.StringFlag{
			Name: idLong,
			Value: strconv.Itoa(id),
		},
		cli.IntFlag{
			Name: limitLong,
			Value: 5,
		},
	}

  return doCreateCli(endpoint, flags)
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

func getFilters(jsonObject map[string] interface{}) map[string] interface{} {

  filters := jsonObject["filters"].(map[string] interface{})

  return filters
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
