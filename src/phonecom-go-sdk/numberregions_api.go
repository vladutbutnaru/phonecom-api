/* 
 * Phone.com API
 *
 * This is a Phone.com api Swagger definition
 *
 * OpenAPI spec version: 1.0.0
 * Contact: apisupport@phone.com
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 */

package swagger

import (
	"net/url"
	"strings"
	"encoding/json"
)

type NumberregionsApi struct {
	Configuration *Configuration
}

func NewNumberregionsApi() *NumberregionsApi {
	configuration := NewConfiguration()
	return &NumberregionsApi{
		Configuration: configuration,
	}
}

func NewNumberregionsApiWithBasePath(basePath string) *NumberregionsApi {
	configuration := NewConfiguration()
	configuration.BasePath = basePath

	return &NumberregionsApi{
		Configuration: configuration,
	}
}

/**
 * 
 * This service lists the quantities of available phone numbers by region.
 *
 * @param filtersCountryCode Country Code filter
 * @param filtersNpa Area Code filter (North America only)
 * @param filtersNxx 2nd set of 3 digits filter (North America only)
 * @param filtersIsTollFree Toll-free status filter
 * @param filtersCity City filter
 * @param filtersProvincePostalCode State or Province filter
 * @param filtersCountryPostalCode Country filter
 * @param sortCountryCode International calling code sorting
 * @param sortNpa Area Code sorting (North America only)
 * @param sortNxx 2nd set of 3 digits sorting (North America)
 * @param sortIsTollFree Toll Free status sorting
 * @param sortCity City sorting
 * @param sortProvincePostalCode State or Province sorting
 * @param sortCountryPostalCode Country sorting
 * @param limit Max results
 * @param offset Results to skip
 * @param fields Field set
 * @param groupBy Fields to group by (supports the same set of fields as the filters and sorting)
 * @return *ListPhoneNumbersRegions
 */
func (a NumberregionsApi) ListAvailablePhoneNumberRegions(filtersCountryCode []string, filtersNpa []string, filtersNxx []string, filtersIsTollFree []string, filtersCity []string, filtersProvincePostalCode []string, filtersCountryPostalCode []string, sortCountryCode string, sortNpa string, sortNxx string, sortIsTollFree string, sortCity string, sortProvincePostalCode string, sortCountryPostalCode string, limit int32, offset int32, fields string, groupBy []string) (*ListPhoneNumbersRegions, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Get")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/phone-numbers/available/regions"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// authentication '(apiKey)' required
	// set key with prefix in header
	localVarHeaderParams["Authorization"] = a.Configuration.GetAPIKeyWithPrefix("Authorization")
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}
	var filtersCountryCodeCollectionFormat = "multi"
	localVarQueryParams.Add("filters[country_code]", a.Configuration.APIClient.ParameterToString(filtersCountryCode, filtersCountryCodeCollectionFormat))

	var filtersNpaCollectionFormat = "multi"
	localVarQueryParams.Add("filters[npa]", a.Configuration.APIClient.ParameterToString(filtersNpa, filtersNpaCollectionFormat))

	var filtersNxxCollectionFormat = "multi"
	localVarQueryParams.Add("filters[nxx]", a.Configuration.APIClient.ParameterToString(filtersNxx, filtersNxxCollectionFormat))

	var filtersIsTollFreeCollectionFormat = "multi"
	localVarQueryParams.Add("filters[is_toll_free]", a.Configuration.APIClient.ParameterToString(filtersIsTollFree, filtersIsTollFreeCollectionFormat))

	var filtersCityCollectionFormat = "multi"
	localVarQueryParams.Add("filters[city]", a.Configuration.APIClient.ParameterToString(filtersCity, filtersCityCollectionFormat))

	var filtersProvincePostalCodeCollectionFormat = "multi"
	localVarQueryParams.Add("filters[province_postal_code]", a.Configuration.APIClient.ParameterToString(filtersProvincePostalCode, filtersProvincePostalCodeCollectionFormat))

	var filtersCountryPostalCodeCollectionFormat = "multi"
	localVarQueryParams.Add("filters[country_postal_code]", a.Configuration.APIClient.ParameterToString(filtersCountryPostalCode, filtersCountryPostalCodeCollectionFormat))

	localVarQueryParams.Add("sort[country_code]", a.Configuration.APIClient.ParameterToString(sortCountryCode, ""))
	localVarQueryParams.Add("sort[npa]", a.Configuration.APIClient.ParameterToString(sortNpa, ""))
	localVarQueryParams.Add("sort[nxx]", a.Configuration.APIClient.ParameterToString(sortNxx, ""))
	localVarQueryParams.Add("sort[is_toll_free]", a.Configuration.APIClient.ParameterToString(sortIsTollFree, ""))
	localVarQueryParams.Add("sort[city]", a.Configuration.APIClient.ParameterToString(sortCity, ""))
	localVarQueryParams.Add("sort[province_postal_code]", a.Configuration.APIClient.ParameterToString(sortProvincePostalCode, ""))
	localVarQueryParams.Add("sort[country_postal_code]", a.Configuration.APIClient.ParameterToString(sortCountryPostalCode, ""))
	localVarQueryParams.Add("limit", a.Configuration.APIClient.ParameterToString(limit, ""))
	localVarQueryParams.Add("offset", a.Configuration.APIClient.ParameterToString(offset, ""))
	localVarQueryParams.Add("fields", a.Configuration.APIClient.ParameterToString(fields, ""))
	var groupByCollectionFormat = "multi"
	localVarQueryParams.Add("group_by", a.Configuration.APIClient.ParameterToString(groupBy, groupByCollectionFormat))


	clearEmptyParams(localVarQueryParams)

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload = new(ListPhoneNumbersRegions)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "ListAvailablePhoneNumberRegions", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

