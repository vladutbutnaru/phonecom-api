//
// NumberregionsAPI.swift
//
// Generated by swagger-codegen
// https://github.com/swagger-api/swagger-codegen
//

import Alamofire



public class NumberregionsAPI: APIBase {
    /**
     
     
     - parameter filtersCountryCode: (query) Country Code filter (optional)
     - parameter filtersNpa: (query) Area Code filter (North America only) (optional)
     - parameter filtersNxx: (query) 2nd set of 3 digits filter (North America only) (optional)
     - parameter filtersIsTollFree: (query) Toll-free status filter (optional)
     - parameter filtersCity: (query) City filter (optional)
     - parameter filtersProvincePostalCode: (query) State or Province filter (optional)
     - parameter filtersCountryPostalCode: (query) Country filter (optional)
     - parameter sortCountryCode: (query) International calling code sorting (optional)
     - parameter sortNpa: (query) Area Code sorting (North America only) (optional)
     - parameter sortNxx: (query) 2nd set of 3 digits sorting (North America) (optional)
     - parameter sortIsTollFree: (query) Toll Free status sorting (optional)
     - parameter sortCity: (query) City sorting (optional)
     - parameter sortProvincePostalCode: (query) State or Province sorting (optional)
     - parameter sortCountryPostalCode: (query) Country sorting (optional)
     - parameter limit: (query) Max results (optional)
     - parameter offset: (query) Results to skip (optional)
     - parameter fields: (query) Field set (optional)
     - parameter groupBy: (query) Fields to group by (supports the same set of fields as the filters and sorting) (optional)
     - parameter completion: completion handler to receive the data and the error objects
     */
    public class func listAvailablePhoneNumberRegions(filtersCountryCode filtersCountryCode: [String]? = nil, filtersNpa: [String]? = nil, filtersNxx: [String]? = nil, filtersIsTollFree: [String]? = nil, filtersCity: [String]? = nil, filtersProvincePostalCode: [String]? = nil, filtersCountryPostalCode: [String]? = nil, sortCountryCode: String? = nil, sortNpa: String? = nil, sortNxx: String? = nil, sortIsTollFree: String? = nil, sortCity: String? = nil, sortProvincePostalCode: String? = nil, sortCountryPostalCode: String? = nil, limit: Int32? = nil, offset: Int32? = nil, fields: String? = nil, groupBy: [String]? = nil, completion: ((data: ListPhoneNumbersRegionsFull?, error: ErrorType?) -> Void)) {
        listAvailablePhoneNumberRegionsWithRequestBuilder(filtersCountryCode: filtersCountryCode, filtersNpa: filtersNpa, filtersNxx: filtersNxx, filtersIsTollFree: filtersIsTollFree, filtersCity: filtersCity, filtersProvincePostalCode: filtersProvincePostalCode, filtersCountryPostalCode: filtersCountryPostalCode, sortCountryCode: sortCountryCode, sortNpa: sortNpa, sortNxx: sortNxx, sortIsTollFree: sortIsTollFree, sortCity: sortCity, sortProvincePostalCode: sortProvincePostalCode, sortCountryPostalCode: sortCountryPostalCode, limit: limit, offset: offset, fields: fields, groupBy: groupBy).execute { (response, error) -> Void in
            completion(data: response?.body, error: error);
        }
    }


    /**
     
     - GET /phone-numbers/available/regions
     - This service lists the quantities of available phone numbers by region.
     - API Key:
       - type: apiKey Authorization 
       - name: apiKey
     - examples: [{contentType=application/json, example={
  "total" : "",
  "offset" : "",
  "limit" : "",
  "group_by" : "",
  "filters" : {
    "country_code" : "aeiou",
    "province_postal_code" : "aeiou",
    "country_postal_code" : "aeiou",
    "city" : "aeiou",
    "is_toll_free" : "aeiou",
    "npa" : [ "" ],
    "nxx" : "aeiou"
  },
  "sort" : {
    "country_code" : "aeiou",
    "province_postal_code" : "aeiou",
    "country_postal_code" : "aeiou",
    "city" : "aeiou",
    "is_toll_free" : "aeiou",
    "npa" : "aeiou",
    "nxx" : "aeiou"
  },
  "items" : ""
}}]
     
     - parameter filtersCountryCode: (query) Country Code filter (optional)
     - parameter filtersNpa: (query) Area Code filter (North America only) (optional)
     - parameter filtersNxx: (query) 2nd set of 3 digits filter (North America only) (optional)
     - parameter filtersIsTollFree: (query) Toll-free status filter (optional)
     - parameter filtersCity: (query) City filter (optional)
     - parameter filtersProvincePostalCode: (query) State or Province filter (optional)
     - parameter filtersCountryPostalCode: (query) Country filter (optional)
     - parameter sortCountryCode: (query) International calling code sorting (optional)
     - parameter sortNpa: (query) Area Code sorting (North America only) (optional)
     - parameter sortNxx: (query) 2nd set of 3 digits sorting (North America) (optional)
     - parameter sortIsTollFree: (query) Toll Free status sorting (optional)
     - parameter sortCity: (query) City sorting (optional)
     - parameter sortProvincePostalCode: (query) State or Province sorting (optional)
     - parameter sortCountryPostalCode: (query) Country sorting (optional)
     - parameter limit: (query) Max results (optional)
     - parameter offset: (query) Results to skip (optional)
     - parameter fields: (query) Field set (optional)
     - parameter groupBy: (query) Fields to group by (supports the same set of fields as the filters and sorting) (optional)

     - returns: RequestBuilder<ListPhoneNumbersRegionsFull> 
     */
    public class func listAvailablePhoneNumberRegionsWithRequestBuilder(filtersCountryCode filtersCountryCode: [String]? = nil, filtersNpa: [String]? = nil, filtersNxx: [String]? = nil, filtersIsTollFree: [String]? = nil, filtersCity: [String]? = nil, filtersProvincePostalCode: [String]? = nil, filtersCountryPostalCode: [String]? = nil, sortCountryCode: String? = nil, sortNpa: String? = nil, sortNxx: String? = nil, sortIsTollFree: String? = nil, sortCity: String? = nil, sortProvincePostalCode: String? = nil, sortCountryPostalCode: String? = nil, limit: Int32? = nil, offset: Int32? = nil, fields: String? = nil, groupBy: [String]? = nil) -> RequestBuilder<ListPhoneNumbersRegionsFull> {
        let path = "/phone-numbers/available/regions"
        let URLString = SwaggerClientAPI.basePath + path

        let nillableParameters: [String:AnyObject?] = [
            "filters[country_code]": filtersCountryCode,
            "filters[npa]": filtersNpa,
            "filters[nxx]": filtersNxx,
            "filters[is_toll_free]": filtersIsTollFree,
            "filters[city]": filtersCity,
            "filters[province_postal_code]": filtersProvincePostalCode,
            "filters[country_postal_code]": filtersCountryPostalCode,
            "sort[country_code]": sortCountryCode,
            "sort[npa]": sortNpa,
            "sort[nxx]": sortNxx,
            "sort[is_toll_free]": sortIsTollFree,
            "sort[city]": sortCity,
            "sort[province_postal_code]": sortProvincePostalCode,
            "sort[country_postal_code]": sortCountryPostalCode,
            "limit": limit?.encodeToJSON(),
            "offset": offset?.encodeToJSON(),
            "fields": fields,
            "group_by": groupBy
        ]
 
        let parameters = APIHelper.rejectNil(nillableParameters)
 
        let convertedParameters = APIHelper.convertBoolToString(parameters)
 
        let requestBuilder: RequestBuilder<ListPhoneNumbersRegionsFull>.Type = SwaggerClientAPI.requestBuilderFactory.getBuilder()

        return requestBuilder.init(method: "GET", URLString: URLString, parameters: convertedParameters, isBody: false)
    }

}