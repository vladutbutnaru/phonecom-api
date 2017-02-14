/* 
 * Phone.com API
 *
 * This is a Phone.com api Swagger definition
 *
 * OpenAPI spec version: 1.0.0
 * Contact: apisupport@phone.com
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package swagger

import (
	"net/url"
	"encoding/json"
	"fmt"
	"strings"
)

type ApplicationsApi struct {
	Configuration Configuration
}

func NewApplicationsApi() *ApplicationsApi {
	configuration := NewConfiguration()
	return &ApplicationsApi{
		Configuration: *configuration,
	}
}

func NewApplicationsApiWithBasePath(basePath string) *ApplicationsApi {
	configuration := NewConfiguration()
	configuration.BasePath = basePath

	return &ApplicationsApi{
		Configuration: *configuration,
	}
}

/**
 * Show details of an individual application
 * 
 *
 * @param accountId Account ID
 * @param applicationId Application ID
 * @return *ApplicationFull
 */
func (a ApplicationsApi) GetAccountApplication(accountId int32, applicationId int32) (*ApplicationFull, *APIResponse, error) {

	var httpMethod = "Get"
	// create path and map variables
	path := a.Configuration.BasePath + "/accounts/{account_id}/applications/{application_id}"
	path = strings.Replace(path, "{"+"account_id"+"}", fmt.Sprintf("%v", accountId), -1)
	path = strings.Replace(path, "{"+"application_id"+"}", fmt.Sprintf("%v", applicationId), -1)


	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := make(map[string]string)
	var postBody interface{}
	var fileName string
	var fileBytes []byte
	// authentication '(apiKey)' required
	// set key with prefix in header
	headerParams["Authorization"] = a.Configuration.GetAPIKeyWithPrefix("Authorization")
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headerParams[key] = a.Configuration.DefaultHeader[key]
	}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload = new(ApplicationFull)
	httpResponse, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, fileName, fileBytes)
	if err != nil {
		return successPayload, NewAPIResponse(httpResponse.RawResponse), err
	}
	err = json.Unmarshal(httpResponse.Body(), &successPayload)
	return successPayload, NewAPIResponse(httpResponse.RawResponse), err
}

/**
 * Get a list of applications you have defined
 * Get a list of an account available applications
 *
 * @param accountId Account ID
 * @param filtersId ID filter
 * @param filtersName Name filter
 * @param sortId ID sorting
 * @param sortName Name sorting
 * @param limit Max results
 * @param offset Results to skip
 * @param fields Field set
 * @return *ListApplicationsFull
 */
func (a ApplicationsApi) ListAccountApplications(accountId int32, filtersId []string, filtersName []string, sortId string, sortName string, limit int32, offset int32, fields string) (*ListApplicationsFull, *APIResponse, error) {

	var httpMethod = "Get"
	// create path and map variables
	path := a.Configuration.BasePath + "/accounts/{account_id}/applications"
	path = strings.Replace(path, "{"+"account_id"+"}", fmt.Sprintf("%v", accountId), -1)


	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := make(map[string]string)
	var postBody interface{}
	var fileName string
	var fileBytes []byte
	// authentication '(apiKey)' required
	// set key with prefix in header
	headerParams["Authorization"] = a.Configuration.GetAPIKeyWithPrefix("Authorization")
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headerParams[key] = a.Configuration.DefaultHeader[key]
	}
	var collectionFormat = "multi"
	if collectionFormat == "multi" {
		for _, value := range filtersId {
			queryParams.Add("filters[id]", value)
		}
	} else {
		queryParams.Add("filters[id]", a.Configuration.APIClient.ParameterToString(filtersId, collectionFormat))
	}
		
	if collectionFormat == "multi" {
		for _, value := range filtersName {
			queryParams.Add("filters[name]", value)
		}
	} else {
		queryParams.Add("filters[name]", a.Configuration.APIClient.ParameterToString(filtersName, collectionFormat))
	}
			queryParams.Add("sort[id]", a.Configuration.APIClient.ParameterToString(sortId, ""))
			queryParams.Add("sort[name]", a.Configuration.APIClient.ParameterToString(sortName, ""))
			queryParams.Add("limit", a.Configuration.APIClient.ParameterToString(limit, ""))
			queryParams.Add("offset", a.Configuration.APIClient.ParameterToString(offset, ""))
			queryParams.Add("fields", a.Configuration.APIClient.ParameterToString(fields, ""))
	

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload = new(ListApplicationsFull)
	httpResponse, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, fileName, fileBytes)
	if err != nil {
		return successPayload, NewAPIResponse(httpResponse.RawResponse), err
	}
	err = json.Unmarshal(httpResponse.Body(), &successPayload)
	return successPayload, NewAPIResponse(httpResponse.RawResponse), err
}

