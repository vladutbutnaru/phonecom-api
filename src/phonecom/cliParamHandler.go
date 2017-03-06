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

func createCliParams(context *cli.Context) (CliParams, error) {

	slice := make([]string, 0)
	limit := int32(context.Int("limit"))
	offset := int32(context.Int("offset"))
	idString := context.String("id")
	var id int32 = 0

	if _, err := strconv.Atoi(idString); err == nil {
		idInt := 0
		idInt, err = strconv.Atoi(idString);
		id = int32(idInt);
	}

	verbose := context.Bool("verbose")
	dryRun := context.Bool("dryrun")
	input := context.String("input")
	command := context.String("command")
	idSecondary := int32(context.Int("id-secondary"))
	accountId := int32(context.Int("account"))
	contact := context.String("contact")
	billingContact := context.String("billing-contact")
	fields := ""
	filtersType := context.String("filtersType")
	filtersValue := context.String("filtersValue")
	sortType := context.String("sortType")
	sortValue := context.String("sortValue")
	samplein := context.String("samplein")
	sampleout := context.String("sampleout")
	fullList := context.Bool("fullList")
	inputFormat := context.String("inputFormat")
	outputFormat := context.String("outputFormat")
	apiKey := context.String("api-key")
	apiKeyPrefix := context.String("api-key-prefix")

	var filtersId []string

	var from string
	var to string
	var text string
	from = context.String("from")
	to = context.String("to")
	text = context.String("text")

	var trunkName = context.String("name")
	var trunkUri = context.String("uri")
	var trunkConcurrentCalls = int32(context.Int("max-concurrent-calls"))
	var trunkMaxMinutes = int32(context.Int("max-minutes-per-month"))

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

	if par.sortType != "" && par.sortValue != "" {
		switch par.sortType {
		case "id":
			par.sortParams.sortId = par.sortValue
		case "name":
			par.sortParams.sortName = par.sortValue
		case "start_time":
			par.sortParams.sortStartTime = par.sortValue
		case "created_at":
			par.sortParams.sortCreatedAt = par.sortValue
		case "extension":
			par.sortParams.sortExtension = par.sortValue
		case "number":
			par.sortParams.sortNumber = par.sortValue
		case "updated_at":
			par.sortParams.sortUpdatedAt = par.sortValue
		case "phone_number":
			par.sortParams.sortPhoneNumber = par.sortValue
		case "internal":
			par.sortParams.sortInternal = par.sortValue
		case "price":
			par.sortParams.sortPrice = par.sortValue
		case "npa":
			par.sortParams.sortNpa = par.sortValue
		case "nxx":
			par.sortParams.sortNxx = par.sortValue
		case "is_toll_free":
			par.sortParams.sortIsTollFree = par.sortValue
		case "city":
			par.sortParams.sortCity = par.sortValue
		case "province_postal_code":
			par.sortParams.sortProvincePostalCode = par.sortValue
		case "country_postal_code":
			par.sortParams.sortCountryPostalCode = par.sortValue
		}
	}

	if par.filterType != "" && par.filterValue != "" {
		slice1 := make([]string, 0)
		slice1 = append(slice1, par.filterValue)
		switch par.filterType {
		case "id":
			par.filtersId = slice1
		case "name":
			par.filterParams.filtersName = slice1
		case "start_time":
			par.filterParams.filtersStartTime = slice1
		case "created_at":
			par.filterParams.filtersCreatedAt = par.filterValue
		case "direction":
			par.filterParams.filtersDirection = par.filterValue
		case "called_number":
			par.filterParams.filtersCalledNumber = par.filterValue
		case "type":
			par.filterParams.filtersType = par.filterValue
		case "extension":
			par.filterParams.filtersExtension = slice1
		case "number":
			par.filterParams.filtersNumber = slice1
		case "group_id":
			par.filterParams.filtersGroupId = slice1
		case "updated_at":
			par.filterParams.filtersUpdatedAt = slice1
		case "phone_number":
			par.filterParams.filtersPhoneNumber = slice1
		case "from":
			par.filterParams.filtersFrom = par.filterValue
		case "to":
			par.filterParams.filtersTo = slice1
		case "country_code":
			par.filterParams.filtersCountryCode = slice1
		case "npa":
			par.filterParams.filtersNpa = slice1
		case "nxx":
			par.filterParams.filtersNxx = slice1
		case "xxxx":
			par.filterParams.filtersXxxx = slice1
		case "city":
			par.filterParams.filtersCity = slice1
		case "province":
			par.filterParams.filtersProvince = slice1
		case "country":
			par.filterParams.filtersCountry = slice1
		case "price":
			par.filterParams.filtersPrice = slice1
		case "category":
			par.filterParams.filtersCategory = slice1
		case "is_toll_free":
			par.filterParams.filtersIsTollFree = slice1
		case "province_postal_code":
			par.filterParams.filtersProvincePostalCode = slice1
		case "country_postal_code":
			par.filterParams.filtersCountryPostalCode = slice1
		}
	}

	return par, nil
}