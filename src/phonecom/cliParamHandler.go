package main

import (
	"github.com/urfave/cli"
	"strconv"
)

type CliParams struct {
	slice       []string
	filtersId   []string
	limit       int32
	offset      int32
	id          int32
	idString    string
	dryRun      bool
	verbose     bool
	input       string
	command     string
	idSecondary int32
	accountId   int32
	fields      string
	contact     string
	billingContact string
	filterType  string
	filterValue string
	fullList    bool
	inputFormat string
	outputFormat string
	apiKey string
	apiKeyPrefix string
	sortType string
	sortValue string
	samplein string
	sampleout string

  filterParams FilterParams
  sortParams SortParams
  otherParams OtherParams

	from string
  to string
  text string

  trunkName string
  trunkUri string
  trunkConcurrentCalls int32
  trunkMaxMinutes int32
}

func createCliParams(c *cli.Context) (CliParams, error) {

	slice := make([]string, 0)
	limit := int32(c.Int("limit"))
	offset := int32(c.Int("offset"))
	idString := c.String("id")
	var id int32 = 0

	if _, err := strconv.Atoi(idString); err == nil {
		idInt := 0
		idInt, err = strconv.Atoi(idString);
		id = int32(idInt);
	}

	verbose := c.Bool("verbose")
	dryRun := c.Bool("dryrun")
	input := c.String("input")
	command := c.String("command")
	idSecondary := int32(c.Int("id-secondary"))
	accountId := int32(c.Int("account"))
	contact := c.String("contact")
	billingContact := c.String("billing-contact")
	fields := ""
	filtersType := c.String("filtersType")
	filtersValue := c.String("filtersValue")
	sortType := c.String("sortType")
	sortValue := c.String("sortValue")
	samplein := c.String("samplein")
	sampleout := c.String("sampleout")
	fullList := c.Bool("fullList")
	inputFormat := c.String("inputFormat")
	outputFormat := c.String("outputFormat")
	apiKey := c.String("api-key")
	apiKeyPrefix := c.String("api-key-prefix")

	var filtersId []string

	var from string
	var to string
	var text string
	from = c.String("from")
	to = c.String("to")
	text = c.String("text")

	var trunkName = c.String("name")
	var trunkUri = c.String("uri")
	var trunkConcurrentCalls = int32(c.Int("max-concurrent-calls"))
	var trunkMaxMinutes = int32(c.Int("max-minutes-per-month"))

	trunkName = defaultTrunkName
	trunkUri = defaultTrunkUri
	trunkConcurrentCalls = int32(defaultTrunkConcurrentCalls)
	trunkMaxMinutes = int32(defaultTrunkMaxMinutes)

	var filterParams FilterParams
	var sortParams SortParams
	var otherParams OtherParams

	var par CliParams

	if (input != "") {
		var err error
		var listParams ListParams
		err, listParams = getListParams(input)
		if (err != nil) {
			return par, err
		}

		accountId = listParams.accountId
		limit = listParams.limit
		offset = listParams.offset
		fields = listParams.fields

		err, filterParams = getFiltersParams(input)

		if (err != nil) {
			return par, err
		}

		filtersId = filterParams.filtersId

		err, sortParams = getSortParams(input)

		if (err != nil) {
			return par, err
		}

		err, otherParams = getOtherParams(input)

		if (err != nil) {
			return par, err
		}
	}

	par.verbose = verbose
	par.dryRun = dryRun
	par.input = input
	par.command = command
	par.idSecondary = idSecondary
	par.accountId = accountId
	par.contact = contact
	par.billingContact = billingContact
	par.fields = fields
	par.filterType = filtersType
	par.filterValue = filtersValue
	par.sortType = sortType
	par.sortValue = sortValue
	par.samplein = samplein
	par.sampleout = sampleout
	par.fullList = fullList
	par.slice = slice
	par.limit = limit
	par.offset = offset
	par.id = id
	par.idString = idString
	par.apiKey = apiKey
	par.apiKeyPrefix = apiKeyPrefix
	par.inputFormat = inputFormat
	par.outputFormat = outputFormat

	par.filterParams = filterParams
	par.sortParams = sortParams
	par.otherParams = otherParams

	par.filtersId = filtersId

	par.from = from
	par.to = to
	par.text = text

	par.trunkName = trunkName
	par.trunkUri = trunkUri
	par.trunkConcurrentCalls = trunkConcurrentCalls
	par.trunkMaxMinutes = trunkMaxMinutes

	return par, nil
}