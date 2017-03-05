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

using System;
using System.Collections.Generic;
using System.Collections.ObjectModel;
using System.Linq;
using RestSharp;
using IO.Swagger.Client;
using IO.Swagger.Model;

namespace IO.Swagger.Api
{
    /// <summary>
    /// Represents a collection of functions to interact with the API endpoints
    /// </summary>
    public interface INumberregionsApi : IApiAccessor
    {
        #region Synchronous Operations
        /// <summary>
        /// 
        /// </summary>
        /// <remarks>
        /// This service lists the quantities of available phone numbers by region.
        /// </remarks>
        /// <exception cref="IO.Swagger.Client.ApiException">Thrown when fails to make API call</exception>
        /// <param name="filtersCountryCode">Country Code filter (optional)</param>
        /// <param name="filtersNpa">Area Code filter (North America only) (optional)</param>
        /// <param name="filtersNxx">2nd set of 3 digits filter (North America only) (optional)</param>
        /// <param name="filtersIsTollFree">Toll-free status filter (optional)</param>
        /// <param name="filtersCity">City filter (optional)</param>
        /// <param name="filtersProvincePostalCode">State or Province filter (optional)</param>
        /// <param name="filtersCountryPostalCode">Country filter (optional)</param>
        /// <param name="sortCountryCode">International calling code sorting (optional)</param>
        /// <param name="sortNpa">Area Code sorting (North America only) (optional)</param>
        /// <param name="sortNxx">2nd set of 3 digits sorting (North America) (optional)</param>
        /// <param name="sortIsTollFree">Toll Free status sorting (optional)</param>
        /// <param name="sortCity">City sorting (optional)</param>
        /// <param name="sortProvincePostalCode">State or Province sorting (optional)</param>
        /// <param name="sortCountryPostalCode">Country sorting (optional)</param>
        /// <param name="limit">Max results (optional)</param>
        /// <param name="offset">Results to skip (optional)</param>
        /// <param name="fields">Field set (optional)</param>
        /// <param name="groupBy">Fields to group by (supports the same set of fields as the filters and sorting) (optional)</param>
        /// <returns>ListPhoneNumbersRegionsFull</returns>
        ListPhoneNumbersRegionsFull ListAvailablePhoneNumberRegions (List<string> filtersCountryCode = null, List<string> filtersNpa = null, List<string> filtersNxx = null, List<string> filtersIsTollFree = null, List<string> filtersCity = null, List<string> filtersProvincePostalCode = null, List<string> filtersCountryPostalCode = null, string sortCountryCode = null, string sortNpa = null, string sortNxx = null, string sortIsTollFree = null, string sortCity = null, string sortProvincePostalCode = null, string sortCountryPostalCode = null, int? limit = null, int? offset = null, string fields = null, List<string> groupBy = null);

        /// <summary>
        /// 
        /// </summary>
        /// <remarks>
        /// This service lists the quantities of available phone numbers by region.
        /// </remarks>
        /// <exception cref="IO.Swagger.Client.ApiException">Thrown when fails to make API call</exception>
        /// <param name="filtersCountryCode">Country Code filter (optional)</param>
        /// <param name="filtersNpa">Area Code filter (North America only) (optional)</param>
        /// <param name="filtersNxx">2nd set of 3 digits filter (North America only) (optional)</param>
        /// <param name="filtersIsTollFree">Toll-free status filter (optional)</param>
        /// <param name="filtersCity">City filter (optional)</param>
        /// <param name="filtersProvincePostalCode">State or Province filter (optional)</param>
        /// <param name="filtersCountryPostalCode">Country filter (optional)</param>
        /// <param name="sortCountryCode">International calling code sorting (optional)</param>
        /// <param name="sortNpa">Area Code sorting (North America only) (optional)</param>
        /// <param name="sortNxx">2nd set of 3 digits sorting (North America) (optional)</param>
        /// <param name="sortIsTollFree">Toll Free status sorting (optional)</param>
        /// <param name="sortCity">City sorting (optional)</param>
        /// <param name="sortProvincePostalCode">State or Province sorting (optional)</param>
        /// <param name="sortCountryPostalCode">Country sorting (optional)</param>
        /// <param name="limit">Max results (optional)</param>
        /// <param name="offset">Results to skip (optional)</param>
        /// <param name="fields">Field set (optional)</param>
        /// <param name="groupBy">Fields to group by (supports the same set of fields as the filters and sorting) (optional)</param>
        /// <returns>ApiResponse of ListPhoneNumbersRegionsFull</returns>
        ApiResponse<ListPhoneNumbersRegionsFull> ListAvailablePhoneNumberRegionsWithHttpInfo (List<string> filtersCountryCode = null, List<string> filtersNpa = null, List<string> filtersNxx = null, List<string> filtersIsTollFree = null, List<string> filtersCity = null, List<string> filtersProvincePostalCode = null, List<string> filtersCountryPostalCode = null, string sortCountryCode = null, string sortNpa = null, string sortNxx = null, string sortIsTollFree = null, string sortCity = null, string sortProvincePostalCode = null, string sortCountryPostalCode = null, int? limit = null, int? offset = null, string fields = null, List<string> groupBy = null);
        #endregion Synchronous Operations
        #region Asynchronous Operations
        /// <summary>
        /// 
        /// </summary>
        /// <remarks>
        /// This service lists the quantities of available phone numbers by region.
        /// </remarks>
        /// <exception cref="IO.Swagger.Client.ApiException">Thrown when fails to make API call</exception>
        /// <param name="filtersCountryCode">Country Code filter (optional)</param>
        /// <param name="filtersNpa">Area Code filter (North America only) (optional)</param>
        /// <param name="filtersNxx">2nd set of 3 digits filter (North America only) (optional)</param>
        /// <param name="filtersIsTollFree">Toll-free status filter (optional)</param>
        /// <param name="filtersCity">City filter (optional)</param>
        /// <param name="filtersProvincePostalCode">State or Province filter (optional)</param>
        /// <param name="filtersCountryPostalCode">Country filter (optional)</param>
        /// <param name="sortCountryCode">International calling code sorting (optional)</param>
        /// <param name="sortNpa">Area Code sorting (North America only) (optional)</param>
        /// <param name="sortNxx">2nd set of 3 digits sorting (North America) (optional)</param>
        /// <param name="sortIsTollFree">Toll Free status sorting (optional)</param>
        /// <param name="sortCity">City sorting (optional)</param>
        /// <param name="sortProvincePostalCode">State or Province sorting (optional)</param>
        /// <param name="sortCountryPostalCode">Country sorting (optional)</param>
        /// <param name="limit">Max results (optional)</param>
        /// <param name="offset">Results to skip (optional)</param>
        /// <param name="fields">Field set (optional)</param>
        /// <param name="groupBy">Fields to group by (supports the same set of fields as the filters and sorting) (optional)</param>
        /// <returns>Task of ListPhoneNumbersRegionsFull</returns>
        System.Threading.Tasks.Task<ListPhoneNumbersRegionsFull> ListAvailablePhoneNumberRegionsAsync (List<string> filtersCountryCode = null, List<string> filtersNpa = null, List<string> filtersNxx = null, List<string> filtersIsTollFree = null, List<string> filtersCity = null, List<string> filtersProvincePostalCode = null, List<string> filtersCountryPostalCode = null, string sortCountryCode = null, string sortNpa = null, string sortNxx = null, string sortIsTollFree = null, string sortCity = null, string sortProvincePostalCode = null, string sortCountryPostalCode = null, int? limit = null, int? offset = null, string fields = null, List<string> groupBy = null);

        /// <summary>
        /// 
        /// </summary>
        /// <remarks>
        /// This service lists the quantities of available phone numbers by region.
        /// </remarks>
        /// <exception cref="IO.Swagger.Client.ApiException">Thrown when fails to make API call</exception>
        /// <param name="filtersCountryCode">Country Code filter (optional)</param>
        /// <param name="filtersNpa">Area Code filter (North America only) (optional)</param>
        /// <param name="filtersNxx">2nd set of 3 digits filter (North America only) (optional)</param>
        /// <param name="filtersIsTollFree">Toll-free status filter (optional)</param>
        /// <param name="filtersCity">City filter (optional)</param>
        /// <param name="filtersProvincePostalCode">State or Province filter (optional)</param>
        /// <param name="filtersCountryPostalCode">Country filter (optional)</param>
        /// <param name="sortCountryCode">International calling code sorting (optional)</param>
        /// <param name="sortNpa">Area Code sorting (North America only) (optional)</param>
        /// <param name="sortNxx">2nd set of 3 digits sorting (North America) (optional)</param>
        /// <param name="sortIsTollFree">Toll Free status sorting (optional)</param>
        /// <param name="sortCity">City sorting (optional)</param>
        /// <param name="sortProvincePostalCode">State or Province sorting (optional)</param>
        /// <param name="sortCountryPostalCode">Country sorting (optional)</param>
        /// <param name="limit">Max results (optional)</param>
        /// <param name="offset">Results to skip (optional)</param>
        /// <param name="fields">Field set (optional)</param>
        /// <param name="groupBy">Fields to group by (supports the same set of fields as the filters and sorting) (optional)</param>
        /// <returns>Task of ApiResponse (ListPhoneNumbersRegionsFull)</returns>
        System.Threading.Tasks.Task<ApiResponse<ListPhoneNumbersRegionsFull>> ListAvailablePhoneNumberRegionsAsyncWithHttpInfo (List<string> filtersCountryCode = null, List<string> filtersNpa = null, List<string> filtersNxx = null, List<string> filtersIsTollFree = null, List<string> filtersCity = null, List<string> filtersProvincePostalCode = null, List<string> filtersCountryPostalCode = null, string sortCountryCode = null, string sortNpa = null, string sortNxx = null, string sortIsTollFree = null, string sortCity = null, string sortProvincePostalCode = null, string sortCountryPostalCode = null, int? limit = null, int? offset = null, string fields = null, List<string> groupBy = null);
        #endregion Asynchronous Operations
    }

    /// <summary>
    /// Represents a collection of functions to interact with the API endpoints
    /// </summary>
    public partial class NumberregionsApi : INumberregionsApi
    {
        private IO.Swagger.Client.ExceptionFactory _exceptionFactory = (name, response) => null;

        /// <summary>
        /// Initializes a new instance of the <see cref="NumberregionsApi"/> class.
        /// </summary>
        /// <returns></returns>
        public NumberregionsApi(String basePath)
        {
            this.Configuration = new Configuration(new ApiClient(basePath));

            ExceptionFactory = IO.Swagger.Client.Configuration.DefaultExceptionFactory;

            // ensure API client has configuration ready
            if (Configuration.ApiClient.Configuration == null)
            {
                this.Configuration.ApiClient.Configuration = this.Configuration;
            }
        }

        /// <summary>
        /// Initializes a new instance of the <see cref="NumberregionsApi"/> class
        /// using Configuration object
        /// </summary>
        /// <param name="configuration">An instance of Configuration</param>
        /// <returns></returns>
        public NumberregionsApi(Configuration configuration = null)
        {
            if (configuration == null) // use the default one in Configuration
                this.Configuration = Configuration.Default;
            else
                this.Configuration = configuration;

            ExceptionFactory = IO.Swagger.Client.Configuration.DefaultExceptionFactory;

            // ensure API client has configuration ready
            if (Configuration.ApiClient.Configuration == null)
            {
                this.Configuration.ApiClient.Configuration = this.Configuration;
            }
        }

        /// <summary>
        /// Gets the base path of the API client.
        /// </summary>
        /// <value>The base path</value>
        public String GetBasePath()
        {
            return this.Configuration.ApiClient.RestClient.BaseUrl.ToString();
        }

        /// <summary>
        /// Sets the base path of the API client.
        /// </summary>
        /// <value>The base path</value>
        [Obsolete("SetBasePath is deprecated, please do 'Configuration.ApiClient = new ApiClient(\"http://new-path\")' instead.")]
        public void SetBasePath(String basePath)
        {
            // do nothing
        }

        /// <summary>
        /// Gets or sets the configuration object
        /// </summary>
        /// <value>An instance of the Configuration</value>
        public Configuration Configuration {get; set;}

        /// <summary>
        /// Provides a factory method hook for the creation of exceptions.
        /// </summary>
        public IO.Swagger.Client.ExceptionFactory ExceptionFactory
        {
            get
            {
                if (_exceptionFactory != null && _exceptionFactory.GetInvocationList().Length > 1)
                {
                    throw new InvalidOperationException("Multicast delegate for ExceptionFactory is unsupported.");
                }
                return _exceptionFactory;
            }
            set { _exceptionFactory = value; }
        }

        /// <summary>
        /// Gets the default header.
        /// </summary>
        /// <returns>Dictionary of HTTP header</returns>
        [Obsolete("DefaultHeader is deprecated, please use Configuration.DefaultHeader instead.")]
        public Dictionary<String, String> DefaultHeader()
        {
            return this.Configuration.DefaultHeader;
        }

        /// <summary>
        /// Add default header.
        /// </summary>
        /// <param name="key">Header field name.</param>
        /// <param name="value">Header field value.</param>
        /// <returns></returns>
        [Obsolete("AddDefaultHeader is deprecated, please use Configuration.AddDefaultHeader instead.")]
        public void AddDefaultHeader(string key, string value)
        {
            this.Configuration.AddDefaultHeader(key, value);
        }

        /// <summary>
        ///  This service lists the quantities of available phone numbers by region.
        /// </summary>
        /// <exception cref="IO.Swagger.Client.ApiException">Thrown when fails to make API call</exception>
        /// <param name="filtersCountryCode">Country Code filter (optional)</param>
        /// <param name="filtersNpa">Area Code filter (North America only) (optional)</param>
        /// <param name="filtersNxx">2nd set of 3 digits filter (North America only) (optional)</param>
        /// <param name="filtersIsTollFree">Toll-free status filter (optional)</param>
        /// <param name="filtersCity">City filter (optional)</param>
        /// <param name="filtersProvincePostalCode">State or Province filter (optional)</param>
        /// <param name="filtersCountryPostalCode">Country filter (optional)</param>
        /// <param name="sortCountryCode">International calling code sorting (optional)</param>
        /// <param name="sortNpa">Area Code sorting (North America only) (optional)</param>
        /// <param name="sortNxx">2nd set of 3 digits sorting (North America) (optional)</param>
        /// <param name="sortIsTollFree">Toll Free status sorting (optional)</param>
        /// <param name="sortCity">City sorting (optional)</param>
        /// <param name="sortProvincePostalCode">State or Province sorting (optional)</param>
        /// <param name="sortCountryPostalCode">Country sorting (optional)</param>
        /// <param name="limit">Max results (optional)</param>
        /// <param name="offset">Results to skip (optional)</param>
        /// <param name="fields">Field set (optional)</param>
        /// <param name="groupBy">Fields to group by (supports the same set of fields as the filters and sorting) (optional)</param>
        /// <returns>ListPhoneNumbersRegionsFull</returns>
        public ListPhoneNumbersRegionsFull ListAvailablePhoneNumberRegions (List<string> filtersCountryCode = null, List<string> filtersNpa = null, List<string> filtersNxx = null, List<string> filtersIsTollFree = null, List<string> filtersCity = null, List<string> filtersProvincePostalCode = null, List<string> filtersCountryPostalCode = null, string sortCountryCode = null, string sortNpa = null, string sortNxx = null, string sortIsTollFree = null, string sortCity = null, string sortProvincePostalCode = null, string sortCountryPostalCode = null, int? limit = null, int? offset = null, string fields = null, List<string> groupBy = null)
        {
             ApiResponse<ListPhoneNumbersRegionsFull> localVarResponse = ListAvailablePhoneNumberRegionsWithHttpInfo(filtersCountryCode, filtersNpa, filtersNxx, filtersIsTollFree, filtersCity, filtersProvincePostalCode, filtersCountryPostalCode, sortCountryCode, sortNpa, sortNxx, sortIsTollFree, sortCity, sortProvincePostalCode, sortCountryPostalCode, limit, offset, fields, groupBy);
             return localVarResponse.Data;
        }

        /// <summary>
        ///  This service lists the quantities of available phone numbers by region.
        /// </summary>
        /// <exception cref="IO.Swagger.Client.ApiException">Thrown when fails to make API call</exception>
        /// <param name="filtersCountryCode">Country Code filter (optional)</param>
        /// <param name="filtersNpa">Area Code filter (North America only) (optional)</param>
        /// <param name="filtersNxx">2nd set of 3 digits filter (North America only) (optional)</param>
        /// <param name="filtersIsTollFree">Toll-free status filter (optional)</param>
        /// <param name="filtersCity">City filter (optional)</param>
        /// <param name="filtersProvincePostalCode">State or Province filter (optional)</param>
        /// <param name="filtersCountryPostalCode">Country filter (optional)</param>
        /// <param name="sortCountryCode">International calling code sorting (optional)</param>
        /// <param name="sortNpa">Area Code sorting (North America only) (optional)</param>
        /// <param name="sortNxx">2nd set of 3 digits sorting (North America) (optional)</param>
        /// <param name="sortIsTollFree">Toll Free status sorting (optional)</param>
        /// <param name="sortCity">City sorting (optional)</param>
        /// <param name="sortProvincePostalCode">State or Province sorting (optional)</param>
        /// <param name="sortCountryPostalCode">Country sorting (optional)</param>
        /// <param name="limit">Max results (optional)</param>
        /// <param name="offset">Results to skip (optional)</param>
        /// <param name="fields">Field set (optional)</param>
        /// <param name="groupBy">Fields to group by (supports the same set of fields as the filters and sorting) (optional)</param>
        /// <returns>ApiResponse of ListPhoneNumbersRegionsFull</returns>
        public ApiResponse< ListPhoneNumbersRegionsFull > ListAvailablePhoneNumberRegionsWithHttpInfo (List<string> filtersCountryCode = null, List<string> filtersNpa = null, List<string> filtersNxx = null, List<string> filtersIsTollFree = null, List<string> filtersCity = null, List<string> filtersProvincePostalCode = null, List<string> filtersCountryPostalCode = null, string sortCountryCode = null, string sortNpa = null, string sortNxx = null, string sortIsTollFree = null, string sortCity = null, string sortProvincePostalCode = null, string sortCountryPostalCode = null, int? limit = null, int? offset = null, string fields = null, List<string> groupBy = null)
        {

            var localVarPath = "/phone-numbers/available/regions";
            var localVarPathParams = new Dictionary<String, String>();
            var localVarQueryParams = new Dictionary<String, String>();
            var localVarHeaderParams = new Dictionary<String, String>(Configuration.DefaultHeader);
            var localVarFormParams = new Dictionary<String, String>();
            var localVarFileParams = new Dictionary<String, FileParameter>();
            Object localVarPostBody = null;

            // to determine the Content-Type header
            String[] localVarHttpContentTypes = new String[] {
                "application/json"
            };
            String localVarHttpContentType = Configuration.ApiClient.SelectHeaderContentType(localVarHttpContentTypes);

            // to determine the Accept header
            String[] localVarHttpHeaderAccepts = new String[] {
                "application/json"
            };
            String localVarHttpHeaderAccept = Configuration.ApiClient.SelectHeaderAccept(localVarHttpHeaderAccepts);
            if (localVarHttpHeaderAccept != null)
                localVarHeaderParams.Add("Accept", localVarHttpHeaderAccept);

            // set "format" to json by default
            // e.g. /pet/{petId}.{format} becomes /pet/{petId}.json
            localVarPathParams.Add("format", "json");
            if (filtersCountryCode != null) localVarQueryParams.Add("filters[country_code]", Configuration.ApiClient.ParameterToString(filtersCountryCode)); // query parameter
            if (filtersNpa != null) localVarQueryParams.Add("filters[npa]", Configuration.ApiClient.ParameterToString(filtersNpa)); // query parameter
            if (filtersNxx != null) localVarQueryParams.Add("filters[nxx]", Configuration.ApiClient.ParameterToString(filtersNxx)); // query parameter
            if (filtersIsTollFree != null) localVarQueryParams.Add("filters[is_toll_free]", Configuration.ApiClient.ParameterToString(filtersIsTollFree)); // query parameter
            if (filtersCity != null) localVarQueryParams.Add("filters[city]", Configuration.ApiClient.ParameterToString(filtersCity)); // query parameter
            if (filtersProvincePostalCode != null) localVarQueryParams.Add("filters[province_postal_code]", Configuration.ApiClient.ParameterToString(filtersProvincePostalCode)); // query parameter
            if (filtersCountryPostalCode != null) localVarQueryParams.Add("filters[country_postal_code]", Configuration.ApiClient.ParameterToString(filtersCountryPostalCode)); // query parameter
            if (sortCountryCode != null) localVarQueryParams.Add("sort[country_code]", Configuration.ApiClient.ParameterToString(sortCountryCode)); // query parameter
            if (sortNpa != null) localVarQueryParams.Add("sort[npa]", Configuration.ApiClient.ParameterToString(sortNpa)); // query parameter
            if (sortNxx != null) localVarQueryParams.Add("sort[nxx]", Configuration.ApiClient.ParameterToString(sortNxx)); // query parameter
            if (sortIsTollFree != null) localVarQueryParams.Add("sort[is_toll_free]", Configuration.ApiClient.ParameterToString(sortIsTollFree)); // query parameter
            if (sortCity != null) localVarQueryParams.Add("sort[city]", Configuration.ApiClient.ParameterToString(sortCity)); // query parameter
            if (sortProvincePostalCode != null) localVarQueryParams.Add("sort[province_postal_code]", Configuration.ApiClient.ParameterToString(sortProvincePostalCode)); // query parameter
            if (sortCountryPostalCode != null) localVarQueryParams.Add("sort[country_postal_code]", Configuration.ApiClient.ParameterToString(sortCountryPostalCode)); // query parameter
            if (limit != null) localVarQueryParams.Add("limit", Configuration.ApiClient.ParameterToString(limit)); // query parameter
            if (offset != null) localVarQueryParams.Add("offset", Configuration.ApiClient.ParameterToString(offset)); // query parameter
            if (fields != null) localVarQueryParams.Add("fields", Configuration.ApiClient.ParameterToString(fields)); // query parameter
            if (groupBy != null) localVarQueryParams.Add("group_by", Configuration.ApiClient.ParameterToString(groupBy)); // query parameter

            // authentication (apiKey) required
            if (!String.IsNullOrEmpty(Configuration.GetApiKeyWithPrefix("Authorization")))
            {
                localVarHeaderParams["Authorization"] = Configuration.GetApiKeyWithPrefix("Authorization");
            }


            // make the HTTP request
            IRestResponse localVarResponse = (IRestResponse) Configuration.ApiClient.CallApi(localVarPath,
                Method.GET, localVarQueryParams, localVarPostBody, localVarHeaderParams, localVarFormParams, localVarFileParams,
                localVarPathParams, localVarHttpContentType);

            int localVarStatusCode = (int) localVarResponse.StatusCode;

            if (ExceptionFactory != null)
            {
                Exception exception = ExceptionFactory("ListAvailablePhoneNumberRegions", localVarResponse);
                if (exception != null) throw exception;
            }

            return new ApiResponse<ListPhoneNumbersRegionsFull>(localVarStatusCode,
                localVarResponse.Headers.ToDictionary(x => x.Name, x => x.Value.ToString()),
                (ListPhoneNumbersRegionsFull) Configuration.ApiClient.Deserialize(localVarResponse, typeof(ListPhoneNumbersRegionsFull)));
            
        }

        /// <summary>
        ///  This service lists the quantities of available phone numbers by region.
        /// </summary>
        /// <exception cref="IO.Swagger.Client.ApiException">Thrown when fails to make API call</exception>
        /// <param name="filtersCountryCode">Country Code filter (optional)</param>
        /// <param name="filtersNpa">Area Code filter (North America only) (optional)</param>
        /// <param name="filtersNxx">2nd set of 3 digits filter (North America only) (optional)</param>
        /// <param name="filtersIsTollFree">Toll-free status filter (optional)</param>
        /// <param name="filtersCity">City filter (optional)</param>
        /// <param name="filtersProvincePostalCode">State or Province filter (optional)</param>
        /// <param name="filtersCountryPostalCode">Country filter (optional)</param>
        /// <param name="sortCountryCode">International calling code sorting (optional)</param>
        /// <param name="sortNpa">Area Code sorting (North America only) (optional)</param>
        /// <param name="sortNxx">2nd set of 3 digits sorting (North America) (optional)</param>
        /// <param name="sortIsTollFree">Toll Free status sorting (optional)</param>
        /// <param name="sortCity">City sorting (optional)</param>
        /// <param name="sortProvincePostalCode">State or Province sorting (optional)</param>
        /// <param name="sortCountryPostalCode">Country sorting (optional)</param>
        /// <param name="limit">Max results (optional)</param>
        /// <param name="offset">Results to skip (optional)</param>
        /// <param name="fields">Field set (optional)</param>
        /// <param name="groupBy">Fields to group by (supports the same set of fields as the filters and sorting) (optional)</param>
        /// <returns>Task of ListPhoneNumbersRegionsFull</returns>
        public async System.Threading.Tasks.Task<ListPhoneNumbersRegionsFull> ListAvailablePhoneNumberRegionsAsync (List<string> filtersCountryCode = null, List<string> filtersNpa = null, List<string> filtersNxx = null, List<string> filtersIsTollFree = null, List<string> filtersCity = null, List<string> filtersProvincePostalCode = null, List<string> filtersCountryPostalCode = null, string sortCountryCode = null, string sortNpa = null, string sortNxx = null, string sortIsTollFree = null, string sortCity = null, string sortProvincePostalCode = null, string sortCountryPostalCode = null, int? limit = null, int? offset = null, string fields = null, List<string> groupBy = null)
        {
             ApiResponse<ListPhoneNumbersRegionsFull> localVarResponse = await ListAvailablePhoneNumberRegionsAsyncWithHttpInfo(filtersCountryCode, filtersNpa, filtersNxx, filtersIsTollFree, filtersCity, filtersProvincePostalCode, filtersCountryPostalCode, sortCountryCode, sortNpa, sortNxx, sortIsTollFree, sortCity, sortProvincePostalCode, sortCountryPostalCode, limit, offset, fields, groupBy);
             return localVarResponse.Data;

        }

        /// <summary>
        ///  This service lists the quantities of available phone numbers by region.
        /// </summary>
        /// <exception cref="IO.Swagger.Client.ApiException">Thrown when fails to make API call</exception>
        /// <param name="filtersCountryCode">Country Code filter (optional)</param>
        /// <param name="filtersNpa">Area Code filter (North America only) (optional)</param>
        /// <param name="filtersNxx">2nd set of 3 digits filter (North America only) (optional)</param>
        /// <param name="filtersIsTollFree">Toll-free status filter (optional)</param>
        /// <param name="filtersCity">City filter (optional)</param>
        /// <param name="filtersProvincePostalCode">State or Province filter (optional)</param>
        /// <param name="filtersCountryPostalCode">Country filter (optional)</param>
        /// <param name="sortCountryCode">International calling code sorting (optional)</param>
        /// <param name="sortNpa">Area Code sorting (North America only) (optional)</param>
        /// <param name="sortNxx">2nd set of 3 digits sorting (North America) (optional)</param>
        /// <param name="sortIsTollFree">Toll Free status sorting (optional)</param>
        /// <param name="sortCity">City sorting (optional)</param>
        /// <param name="sortProvincePostalCode">State or Province sorting (optional)</param>
        /// <param name="sortCountryPostalCode">Country sorting (optional)</param>
        /// <param name="limit">Max results (optional)</param>
        /// <param name="offset">Results to skip (optional)</param>
        /// <param name="fields">Field set (optional)</param>
        /// <param name="groupBy">Fields to group by (supports the same set of fields as the filters and sorting) (optional)</param>
        /// <returns>Task of ApiResponse (ListPhoneNumbersRegionsFull)</returns>
        public async System.Threading.Tasks.Task<ApiResponse<ListPhoneNumbersRegionsFull>> ListAvailablePhoneNumberRegionsAsyncWithHttpInfo (List<string> filtersCountryCode = null, List<string> filtersNpa = null, List<string> filtersNxx = null, List<string> filtersIsTollFree = null, List<string> filtersCity = null, List<string> filtersProvincePostalCode = null, List<string> filtersCountryPostalCode = null, string sortCountryCode = null, string sortNpa = null, string sortNxx = null, string sortIsTollFree = null, string sortCity = null, string sortProvincePostalCode = null, string sortCountryPostalCode = null, int? limit = null, int? offset = null, string fields = null, List<string> groupBy = null)
        {

            var localVarPath = "/phone-numbers/available/regions";
            var localVarPathParams = new Dictionary<String, String>();
            var localVarQueryParams = new Dictionary<String, String>();
            var localVarHeaderParams = new Dictionary<String, String>(Configuration.DefaultHeader);
            var localVarFormParams = new Dictionary<String, String>();
            var localVarFileParams = new Dictionary<String, FileParameter>();
            Object localVarPostBody = null;

            // to determine the Content-Type header
            String[] localVarHttpContentTypes = new String[] {
                "application/json"
            };
            String localVarHttpContentType = Configuration.ApiClient.SelectHeaderContentType(localVarHttpContentTypes);

            // to determine the Accept header
            String[] localVarHttpHeaderAccepts = new String[] {
                "application/json"
            };
            String localVarHttpHeaderAccept = Configuration.ApiClient.SelectHeaderAccept(localVarHttpHeaderAccepts);
            if (localVarHttpHeaderAccept != null)
                localVarHeaderParams.Add("Accept", localVarHttpHeaderAccept);

            // set "format" to json by default
            // e.g. /pet/{petId}.{format} becomes /pet/{petId}.json
            localVarPathParams.Add("format", "json");
            if (filtersCountryCode != null) localVarQueryParams.Add("filters[country_code]", Configuration.ApiClient.ParameterToString(filtersCountryCode)); // query parameter
            if (filtersNpa != null) localVarQueryParams.Add("filters[npa]", Configuration.ApiClient.ParameterToString(filtersNpa)); // query parameter
            if (filtersNxx != null) localVarQueryParams.Add("filters[nxx]", Configuration.ApiClient.ParameterToString(filtersNxx)); // query parameter
            if (filtersIsTollFree != null) localVarQueryParams.Add("filters[is_toll_free]", Configuration.ApiClient.ParameterToString(filtersIsTollFree)); // query parameter
            if (filtersCity != null) localVarQueryParams.Add("filters[city]", Configuration.ApiClient.ParameterToString(filtersCity)); // query parameter
            if (filtersProvincePostalCode != null) localVarQueryParams.Add("filters[province_postal_code]", Configuration.ApiClient.ParameterToString(filtersProvincePostalCode)); // query parameter
            if (filtersCountryPostalCode != null) localVarQueryParams.Add("filters[country_postal_code]", Configuration.ApiClient.ParameterToString(filtersCountryPostalCode)); // query parameter
            if (sortCountryCode != null) localVarQueryParams.Add("sort[country_code]", Configuration.ApiClient.ParameterToString(sortCountryCode)); // query parameter
            if (sortNpa != null) localVarQueryParams.Add("sort[npa]", Configuration.ApiClient.ParameterToString(sortNpa)); // query parameter
            if (sortNxx != null) localVarQueryParams.Add("sort[nxx]", Configuration.ApiClient.ParameterToString(sortNxx)); // query parameter
            if (sortIsTollFree != null) localVarQueryParams.Add("sort[is_toll_free]", Configuration.ApiClient.ParameterToString(sortIsTollFree)); // query parameter
            if (sortCity != null) localVarQueryParams.Add("sort[city]", Configuration.ApiClient.ParameterToString(sortCity)); // query parameter
            if (sortProvincePostalCode != null) localVarQueryParams.Add("sort[province_postal_code]", Configuration.ApiClient.ParameterToString(sortProvincePostalCode)); // query parameter
            if (sortCountryPostalCode != null) localVarQueryParams.Add("sort[country_postal_code]", Configuration.ApiClient.ParameterToString(sortCountryPostalCode)); // query parameter
            if (limit != null) localVarQueryParams.Add("limit", Configuration.ApiClient.ParameterToString(limit)); // query parameter
            if (offset != null) localVarQueryParams.Add("offset", Configuration.ApiClient.ParameterToString(offset)); // query parameter
            if (fields != null) localVarQueryParams.Add("fields", Configuration.ApiClient.ParameterToString(fields)); // query parameter
            if (groupBy != null) localVarQueryParams.Add("group_by", Configuration.ApiClient.ParameterToString(groupBy)); // query parameter

            // authentication (apiKey) required
            if (!String.IsNullOrEmpty(Configuration.GetApiKeyWithPrefix("Authorization")))
            {
                localVarHeaderParams["Authorization"] = Configuration.GetApiKeyWithPrefix("Authorization");
            }

            // make the HTTP request
            IRestResponse localVarResponse = (IRestResponse) await Configuration.ApiClient.CallApiAsync(localVarPath,
                Method.GET, localVarQueryParams, localVarPostBody, localVarHeaderParams, localVarFormParams, localVarFileParams,
                localVarPathParams, localVarHttpContentType);

            int localVarStatusCode = (int) localVarResponse.StatusCode;

            if (ExceptionFactory != null)
            {
                Exception exception = ExceptionFactory("ListAvailablePhoneNumberRegions", localVarResponse);
                if (exception != null) throw exception;
            }

            return new ApiResponse<ListPhoneNumbersRegionsFull>(localVarStatusCode,
                localVarResponse.Headers.ToDictionary(x => x.Name, x => x.Value.ToString()),
                (ListPhoneNumbersRegionsFull) Configuration.ApiClient.Deserialize(localVarResponse, typeof(ListPhoneNumbersRegionsFull)));
            
        }

    }
}