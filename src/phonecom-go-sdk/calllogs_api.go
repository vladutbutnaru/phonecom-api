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
	"fmt"
)

type CalllogsApi struct {
	Configuration *Configuration
}

func NewCalllogsApi() *CalllogsApi {
	configuration := NewConfiguration()
	return &CalllogsApi{
		Configuration: configuration,
	}
}

func NewCalllogsApiWithBasePath(basePath string) *CalllogsApi {
	configuration := NewConfiguration()
	configuration.BasePath = basePath

	return &CalllogsApi{
		Configuration: configuration,
	}
}

/**
 * Show details of an individual Call Log entry
 * This service shows the details of an individual Call Logs entry.
 *
 * @param accountId Account ID
 * @param callLogId Call Log ID
 * @return *CallLogFull
 */
func (a CalllogsApi) GetAccountCallLog(accountId int32, callLogId int32) (*CallLogFull, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Get")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/accounts/{account_id}/call-logs/{callLog_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"account_id"+"}", fmt.Sprintf("%v", accountId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"callLog_id"+"}", fmt.Sprintf("%v", callLogId), -1)

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
	var successPayload = new(CallLogFull)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "GetAccountCallLog", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
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

/**
 * Get a list of call details associated with your account
 * Fetch a list of call logs associated with your account.
 *
 * @param accountId Account ID
 * @param filtersId ID filter
 * @param filtersStartTime Call start time filter
 * @param filtersCreatedAt Call log creation time filter
 * @param filtersDirection Call direction filter: in or out
 * @param filtersCalledNumber Called number
 * @param filtersType Call type, such as &#39;call&#39;, &#39;fax&#39;, &#39;audiogram&#39;
 * @param filtersExtension Extension filter
 * @param sortId ID sorting
 * @param sortStartTime Sorting by call start time: asc or desc
 * @param sortCreatedAt Sorting by call log creation time: asc or desc
 * @param limit Max results
 * @param offset Results to skip
 * @param fields Field set
 * @return *ListCallLogsFull
 */
func (a CalllogsApi) ListAccountCallLogs(accountId int32, filtersId []string, filtersStartTime []string, filtersCreatedAt string, filtersDirection string, filtersCalledNumber string, filtersType string, filtersExtension []string, sortId string, sortStartTime string, sortCreatedAt string, limit int32, offset int32, fields string) (*ListCallLogsFull, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Get")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/accounts/{account_id}/call-logs"
	localVarPath = strings.Replace(localVarPath, "{"+"account_id"+"}", fmt.Sprintf("%v", accountId), -1)

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
	if len(filtersId) > 0 {
		localVarQueryParams.Add("filters[id]", a.Configuration.APIClient.ParameterToString(filtersId, collectionFormat))
	}

	if len(filtersStartTime) > 0 {
		localVarQueryParams.Add("filters[start_time]", a.Configuration.APIClient.ParameterToString(filtersStartTime, collectionFormat))
	}

	if filtersCreatedAt != "" {
		localVarQueryParams.Add("filters[created_at]", a.Configuration.APIClient.ParameterToString(filtersCreatedAt, ""))
	}
	if filtersDirection != "" {
		localVarQueryParams.Add("filters[direction]", a.Configuration.APIClient.ParameterToString(filtersDirection, ""))
	}
	if filtersCalledNumber != "" {
		localVarQueryParams.Add("filters[called_number]", a.Configuration.APIClient.ParameterToString(filtersCalledNumber, ""))
	}
	if filtersType != "" {
		localVarQueryParams.Add("filters[type]", a.Configuration.APIClient.ParameterToString(filtersType, ""))
	}
	if len(filtersExtension) > 0 {
		localVarQueryParams.Add("filters[extension]", a.Configuration.APIClient.ParameterToString(filtersExtension, collectionFormat))
	}

	if sortId != "" {
		localVarQueryParams.Add("sort[id]", a.Configuration.APIClient.ParameterToString(sortId, ""))
	}
	if sortStartTime != "" {
		localVarQueryParams.Add("sort[start_time]", a.Configuration.APIClient.ParameterToString(sortStartTime, ""))
	}
	if sortCreatedAt != "" {
		localVarQueryParams.Add("sort[created_at]", a.Configuration.APIClient.ParameterToString(sortCreatedAt, ""))
	}
	localVarQueryParams.Add("limit", a.Configuration.APIClient.ParameterToString(limit, ""))
	localVarQueryParams.Add("offset", a.Configuration.APIClient.ParameterToString(offset, ""))
	if fields != "" {
		localVarQueryParams.Add("fields", a.Configuration.APIClient.ParameterToString(fields, ""))
	}

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
	var successPayload = new(ListCallLogsFull)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "ListAccountCallLogs", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
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

