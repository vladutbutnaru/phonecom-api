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

type AvailablenumbersApi struct {
	Configuration *Configuration
}

func NewAvailablenumbersApi() *AvailablenumbersApi {
	configuration := NewConfiguration()
	return &AvailablenumbersApi{
		Configuration: configuration,
	}
}

func NewAvailablenumbersApiWithBasePath(basePath string) *AvailablenumbersApi {
	configuration := NewConfiguration()
	configuration.BasePath = basePath

	return &AvailablenumbersApi{
		Configuration: configuration,
	}
}

/**
 * 
 * 
 *
 * @param filtersPhoneNumber Phone number filter
 * @param filtersCountryCode Country Code filter
 * @param filtersNpa Area Code filter (North America only)
 * @param filtersNxx 2nd set of 3 digits filter (North America only)
 * @param filtersXxxx NANP XXXX filter
 * @param filtersCity City filter
 * @param filtersProvince State or Province (postal code) filter
 * @param filtersCountry Country (postal code) filter
 * @param filtersPrice Price filter
 * @param filtersCategory Category filter
 * @param filtersIsTollFree Toll-free status filter
 * @param sortInternal Internal (quasi-random) sorting
 * @param sortPrice Price sorting
 * @param sortPhoneNumber Phone number sorting
 * @param limit Max results
 * @param offset Results to skip
 * @param fields Field set
 * @return *ListAvailableNumbersFull
 */
func (a AvailablenumbersApi) ListAvailablePhoneNumbers(filtersPhoneNumber []string, filtersCountryCode []string, filtersNpa []string, filtersNxx []string, filtersXxxx []string, filtersCity []string, filtersProvince []string, filtersCountry []string, filtersPrice []string, filtersCategory []string, filtersIsTollFree []string, sortInternal string, sortPrice string, sortPhoneNumber string, limit int32, offset int32, fields string) (*ListAvailableNumbersFull, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Get")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/phone-numbers/available"

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
	var collectionFormat = "multi"
	localVarQueryParams.Add("filters[phone_number]", a.Configuration.APIClient.ParameterToString(filtersPhoneNumber, collectionFormat))

	/*var */ collectionFormat = "multi"
	localVarQueryParams.Add("filters[country_code]", a.Configuration.APIClient.ParameterToString(filtersCountryCode, collectionFormat))

	/*var */ collectionFormat = "multi"
	localVarQueryParams.Add("filters[npa]", a.Configuration.APIClient.ParameterToString(filtersNpa, collectionFormat))

	/*var */ collectionFormat = "multi"
	localVarQueryParams.Add("filters[nxx]", a.Configuration.APIClient.ParameterToString(filtersNxx, collectionFormat))

	/*var */ collectionFormat = "multi"
	localVarQueryParams.Add("filters[xxxx]", a.Configuration.APIClient.ParameterToString(filtersXxxx, collectionFormat))

	/*var */ collectionFormat = "multi"
	localVarQueryParams.Add("filters[city]", a.Configuration.APIClient.ParameterToString(filtersCity, collectionFormat))

	/*var */ collectionFormat = "multi"
	localVarQueryParams.Add("filters[province]", a.Configuration.APIClient.ParameterToString(filtersProvince, collectionFormat))

	/*var */ collectionFormat = "multi"
	localVarQueryParams.Add("filters[country]", a.Configuration.APIClient.ParameterToString(filtersCountry, collectionFormat))

	/*var */ collectionFormat = "multi"
	localVarQueryParams.Add("filters[price]", a.Configuration.APIClient.ParameterToString(filtersPrice, collectionFormat))

	/*var */ collectionFormat = "multi"
	localVarQueryParams.Add("filters[category]", a.Configuration.APIClient.ParameterToString(filtersCategory, collectionFormat))

	/*var */ collectionFormat = "multi"
	localVarQueryParams.Add("filters[is_toll_free]", a.Configuration.APIClient.ParameterToString(filtersIsTollFree, collectionFormat))

	localVarQueryParams.Add("sort[internal]", a.Configuration.APIClient.ParameterToString(sortInternal, ""))
	localVarQueryParams.Add("sort[price]", a.Configuration.APIClient.ParameterToString(sortPrice, ""))
	localVarQueryParams.Add("sort[phone_number]", a.Configuration.APIClient.ParameterToString(sortPhoneNumber, ""))
	localVarQueryParams.Add("limit", a.Configuration.APIClient.ParameterToString(limit, ""))
	localVarQueryParams.Add("offset", a.Configuration.APIClient.ParameterToString(offset, ""))
	localVarQueryParams.Add("fields", a.Configuration.APIClient.ParameterToString(fields, ""))

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
	var successPayload = new(ListAvailableNumbersFull)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "ListAvailablePhoneNumbers", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
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
