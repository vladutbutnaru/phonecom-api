<?php
/**
 * CalleridsApi
 * PHP version 5
 *
 * @category Class
 * @package  Swagger\Client
 * @author   http://github.com/swagger-api/swagger-codegen
 * @license  http://www.apache.org/licenses/LICENSE-2.0 Apache Licene v2
 * @link     https://github.com/swagger-api/swagger-codegen
 */

/**
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

/**
 * NOTE: This class is auto generated by the swagger code generator program.
 * https://github.com/swagger-api/swagger-codegen
 * Do not edit the class manually.
 */

namespace Swagger\Client\Api;

use \Swagger\Client\Configuration;
use \Swagger\Client\ApiClient;
use \Swagger\Client\ApiException;
use \Swagger\Client\ObjectSerializer;

/**
 * CalleridsApi Class Doc Comment
 *
 * @category Class
 * @package  Swagger\Client
 * @author   http://github.com/swagger-api/swagger-codegen
 * @license  http://www.apache.org/licenses/LICENSE-2.0 Apache Licene v2
 * @link     https://github.com/swagger-api/swagger-codegen
 */
class CalleridsApi
{

    /**
     * API Client
     *
     * @var \Swagger\Client\ApiClient instance of the ApiClient
     */
    protected $apiClient;

    /**
     * Constructor
     *
     * @param \Swagger\Client\ApiClient|null $apiClient The api client to use
     */
    public function __construct(\Swagger\Client\ApiClient $apiClient = null)
    {
        if ($apiClient == null) {
            $apiClient = new ApiClient();
            $apiClient->getConfig()->setHost('https://api.phone.com/v4');
        }

        $this->apiClient = $apiClient;
    }

    /**
     * Get API client
     *
     * @return \Swagger\Client\ApiClient get the API client
     */
    public function getApiClient()
    {
        return $this->apiClient;
    }

    /**
     * Set the API client
     *
     * @param \Swagger\Client\ApiClient $apiClient set the API client
     *
     * @return CalleridsApi
     */
    public function setApiClient(\Swagger\Client\ApiClient $apiClient)
    {
        $this->apiClient = $apiClient;
        return $this;
    }

    /**
     * Operation getCallerIds
     *
     * Show the Caller ID options a given extension can use
     *
     * @param int $account_id Account ID (required)
     * @param int $extension_id Extension ID (required)
     * @param string[] $filters_number Number filter (optional)
     * @param string[] $filters_name Name filter (optional)
     * @param string $sort_number Number sorting (optional)
     * @param string $sort_name Name sorting (optional)
     * @param int $limit Max results (optional)
     * @param int $offset Results to skip (optional)
     * @param string $fields Field set (optional)
     * @return \Swagger\Client\Model\ListCallerIdsFull
     * @throws \Swagger\Client\ApiException on non-2xx response
     */
    public function getCallerIds($account_id, $extension_id, $filters_number = null, $filters_name = null, $sort_number = null, $sort_name = null, $limit = null, $offset = null, $fields = null)
    {
        list($response) = $this->getCallerIdsWithHttpInfo($account_id, $extension_id, $filters_number, $filters_name, $sort_number, $sort_name, $limit, $offset, $fields);
        return $response;
    }

    /**
     * Operation getCallerIdsWithHttpInfo
     *
     * Show the Caller ID options a given extension can use
     *
     * @param int $account_id Account ID (required)
     * @param int $extension_id Extension ID (required)
     * @param string[] $filters_number Number filter (optional)
     * @param string[] $filters_name Name filter (optional)
     * @param string $sort_number Number sorting (optional)
     * @param string $sort_name Name sorting (optional)
     * @param int $limit Max results (optional)
     * @param int $offset Results to skip (optional)
     * @param string $fields Field set (optional)
     * @return Array of \Swagger\Client\Model\ListCallerIdsFull, HTTP status code, HTTP response headers (array of strings)
     * @throws \Swagger\Client\ApiException on non-2xx response
     */
    public function getCallerIdsWithHttpInfo($account_id, $extension_id, $filters_number = null, $filters_name = null, $sort_number = null, $sort_name = null, $limit = null, $offset = null, $fields = null)
    {
        // verify the required parameter 'account_id' is set
        if ($account_id === null) {
            throw new \InvalidArgumentException('Missing the required parameter $account_id when calling getCallerIds');
        }
        // verify the required parameter 'extension_id' is set
        if ($extension_id === null) {
            throw new \InvalidArgumentException('Missing the required parameter $extension_id when calling getCallerIds');
        }
        if (!is_null($sort_number) && !preg_match("asc|desc", $sort_number)) {
            throw new \InvalidArgumentException('invalid value for "sort_number" when calling CalleridsApi.getCallerIds, must conform to the pattern asc|desc.');
        }

        if (!is_null($sort_name) && !preg_match("asc|desc", $sort_name)) {
            throw new \InvalidArgumentException('invalid value for "sort_name" when calling CalleridsApi.getCallerIds, must conform to the pattern asc|desc.');
        }

        // parse inputs
        $resourcePath = "/accounts/{account_id}/extensions/{extension_id}/caller-ids";
        $httpBody = '';
        $queryParams = array();
        $headerParams = array();
        $formParams = array();
        $_header_accept = $this->apiClient->selectHeaderAccept(array('application/json'));
        if (!is_null($_header_accept)) {
            $headerParams['Accept'] = $_header_accept;
        }
        $headerParams['Content-Type'] = $this->apiClient->selectHeaderContentType(array('application/json'));

        // query params
        if (is_array($filters_number)) {
            $filters_number = $this->apiClient->getSerializer()->serializeCollection($filters_number, 'multi', true);
        }
        if ($filters_number !== null) {
            $queryParams['filters[number]'] = $this->apiClient->getSerializer()->toQueryValue($filters_number);
        }
        // query params
        if (is_array($filters_name)) {
            $filters_name = $this->apiClient->getSerializer()->serializeCollection($filters_name, 'multi', true);
        }
        if ($filters_name !== null) {
            $queryParams['filters[name]'] = $this->apiClient->getSerializer()->toQueryValue($filters_name);
        }
        // query params
        if ($sort_number !== null) {
            $queryParams['sort[number]'] = $this->apiClient->getSerializer()->toQueryValue($sort_number);
        }
        // query params
        if ($sort_name !== null) {
            $queryParams['sort[name]'] = $this->apiClient->getSerializer()->toQueryValue($sort_name);
        }
        // query params
        if ($limit !== null) {
            $queryParams['limit'] = $this->apiClient->getSerializer()->toQueryValue($limit);
        }
        // query params
        if ($offset !== null) {
            $queryParams['offset'] = $this->apiClient->getSerializer()->toQueryValue($offset);
        }
        // query params
        if ($fields !== null) {
            $queryParams['fields'] = $this->apiClient->getSerializer()->toQueryValue($fields);
        }
        // path params
        if ($account_id !== null) {
            $resourcePath = str_replace(
                "{" . "account_id" . "}",
                $this->apiClient->getSerializer()->toPathValue($account_id),
                $resourcePath
            );
        }
        // path params
        if ($extension_id !== null) {
            $resourcePath = str_replace(
                "{" . "extension_id" . "}",
                $this->apiClient->getSerializer()->toPathValue($extension_id),
                $resourcePath
            );
        }
        // default format to json
        $resourcePath = str_replace("{format}", "json", $resourcePath);

        
        // for model (json/xml)
        if (isset($_tempBody)) {
            $httpBody = $_tempBody; // $_tempBody is the method argument, if present
        } elseif (count($formParams) > 0) {
            $httpBody = $formParams; // for HTTP post (form)
        }
        // this endpoint requires API key authentication
        $apiKey = $this->apiClient->getApiKeyWithPrefix('Authorization');
        if (strlen($apiKey) !== 0) {
            $headerParams['Authorization'] = $apiKey;
        }
        // make the API Call
        try {
            list($response, $statusCode, $httpHeader) = $this->apiClient->callApi(
                $resourcePath,
                'GET',
                $queryParams,
                $httpBody,
                $headerParams,
                '\Swagger\Client\Model\ListCallerIdsFull',
                '/accounts/{account_id}/extensions/{extension_id}/caller-ids'
            );

            return array($this->apiClient->getSerializer()->deserialize($response, '\Swagger\Client\Model\ListCallerIdsFull', $httpHeader), $statusCode, $httpHeader);
        } catch (ApiException $e) {
            switch ($e->getCode()) {
                case 200:
                    $data = $this->apiClient->getSerializer()->deserialize($e->getResponseBody(), '\Swagger\Client\Model\ListCallerIdsFull', $e->getResponseHeaders());
                    $e->setResponseObject($data);
                    break;
            }

            throw $e;
        }
    }

}