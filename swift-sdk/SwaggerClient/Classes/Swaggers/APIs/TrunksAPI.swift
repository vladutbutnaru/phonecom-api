//
// TrunksAPI.swift
//
// Generated by swagger-codegen
// https://github.com/swagger-api/swagger-codegen
//

import Alamofire



public class TrunksAPI: APIBase {
    /**
     Add a trunk record with SIP information
     
     - parameter accountId: (path) Account ID 
     - parameter data: (body) Trunk data 
     - parameter completion: completion handler to receive the data and the error objects
     */
    public class func createAccountTrunk(accountId accountId: Int32, data: CreateTrunkParams, completion: ((data: TrunkFull?, error: ErrorType?) -> Void)) {
        createAccountTrunkWithRequestBuilder(accountId: accountId, data: data).execute { (response, error) -> Void in
            completion(data: response?.body, error: error);
        }
    }


    /**
     Add a trunk record with SIP information
     - POST /accounts/{accountId}/trunks
     - For more on the input fields, see Account Trunks.
     - API Key:
       - type: apiKey Authorization 
       - name: apiKey
     - examples: [{contentType=application/json, example={
  "error_message" : "",
  "max_minutes_per_month" : "",
  "greeting" : {
    "name" : "aeiou",
    "id" : ""
  },
  "codecs" : [ "aeiou" ],
  "name" : "aeiou",
  "id" : "",
  "max_concurrent_calls" : "",
  "uri" : "aeiou"
}}]
     
     - parameter accountId: (path) Account ID 
     - parameter data: (body) Trunk data 

     - returns: RequestBuilder<TrunkFull> 
     */
    public class func createAccountTrunkWithRequestBuilder(accountId accountId: Int32, data: CreateTrunkParams) -> RequestBuilder<TrunkFull> {
        var path = "/accounts/{accountId}/trunks"
        path = path.stringByReplacingOccurrencesOfString("{accountId}", withString: "\(accountId)", options: .LiteralSearch, range: nil)
        let URLString = SwaggerClientAPI.basePath + path
        let parameters = data.encodeToJSON() as? [String:AnyObject]
 
        let convertedParameters = APIHelper.convertBoolToString(parameters)
 
        let requestBuilder: RequestBuilder<TrunkFull>.Type = SwaggerClientAPI.requestBuilderFactory.getBuilder()

        return requestBuilder.init(method: "POST", URLString: URLString, parameters: convertedParameters, isBody: true)
    }

    /**
     Delete a trunk from account
     
     - parameter accountId: (path) Account ID 
     - parameter trunkId: (path) Trunk ID 
     - parameter completion: completion handler to receive the data and the error objects
     */
    public class func deleteAccountTrunk(accountId accountId: Int32, trunkId: Int32, completion: ((data: DeleteTrunk?, error: ErrorType?) -> Void)) {
        deleteAccountTrunkWithRequestBuilder(accountId: accountId, trunkId: trunkId).execute { (response, error) -> Void in
            completion(data: response?.body, error: error);
        }
    }


    /**
     Delete a trunk from account
     - DELETE /accounts/{accountId}/trunks/{trunkId}
     - This service deletes a trunk from the account. For more on the properties of trunks, see Account Trunks.
     - API Key:
       - type: apiKey Authorization 
       - name: apiKey
     - examples: [{contentType=application/json, example={
  "success" : true
}}]
     
     - parameter accountId: (path) Account ID 
     - parameter trunkId: (path) Trunk ID 

     - returns: RequestBuilder<DeleteTrunk> 
     */
    public class func deleteAccountTrunkWithRequestBuilder(accountId accountId: Int32, trunkId: Int32) -> RequestBuilder<DeleteTrunk> {
        var path = "/accounts/{accountId}/trunks/{trunkId}"
        path = path.stringByReplacingOccurrencesOfString("{accountId}", withString: "\(accountId)", options: .LiteralSearch, range: nil)
        path = path.stringByReplacingOccurrencesOfString("{trunkId}", withString: "\(trunkId)", options: .LiteralSearch, range: nil)
        let URLString = SwaggerClientAPI.basePath + path

        let nillableParameters: [String:AnyObject?] = [:]
 
        let parameters = APIHelper.rejectNil(nillableParameters)
 
        let convertedParameters = APIHelper.convertBoolToString(parameters)
 
        let requestBuilder: RequestBuilder<DeleteTrunk>.Type = SwaggerClientAPI.requestBuilderFactory.getBuilder()

        return requestBuilder.init(method: "DELETE", URLString: URLString, parameters: convertedParameters, isBody: true)
    }

    /**
     Show details of an individual trunk
     
     - parameter accountId: (path) Account ID 
     - parameter trunkId: (path) Trunk ID 
     - parameter completion: completion handler to receive the data and the error objects
     */
    public class func getAccountTrunk(accountId accountId: Int32, trunkId: Int32, completion: ((data: TrunkFull?, error: ErrorType?) -> Void)) {
        getAccountTrunkWithRequestBuilder(accountId: accountId, trunkId: trunkId).execute { (response, error) -> Void in
            completion(data: response?.body, error: error);
        }
    }


    /**
     Show details of an individual trunk
     - GET /accounts/{accountId}/trunks/{trunkId}
     - This service shows the details of an individual Trunk.
     - API Key:
       - type: apiKey Authorization 
       - name: apiKey
     - examples: [{contentType=application/json, example={
  "error_message" : "",
  "max_minutes_per_month" : "",
  "greeting" : {
    "name" : "aeiou",
    "id" : ""
  },
  "codecs" : [ "aeiou" ],
  "name" : "aeiou",
  "id" : "",
  "max_concurrent_calls" : "",
  "uri" : "aeiou"
}}]
     
     - parameter accountId: (path) Account ID 
     - parameter trunkId: (path) Trunk ID 

     - returns: RequestBuilder<TrunkFull> 
     */
    public class func getAccountTrunkWithRequestBuilder(accountId accountId: Int32, trunkId: Int32) -> RequestBuilder<TrunkFull> {
        var path = "/accounts/{accountId}/trunks/{trunkId}"
        path = path.stringByReplacingOccurrencesOfString("{accountId}", withString: "\(accountId)", options: .LiteralSearch, range: nil)
        path = path.stringByReplacingOccurrencesOfString("{trunkId}", withString: "\(trunkId)", options: .LiteralSearch, range: nil)
        let URLString = SwaggerClientAPI.basePath + path

        let nillableParameters: [String:AnyObject?] = [:]
 
        let parameters = APIHelper.rejectNil(nillableParameters)
 
        let convertedParameters = APIHelper.convertBoolToString(parameters)
 
        let requestBuilder: RequestBuilder<TrunkFull>.Type = SwaggerClientAPI.requestBuilderFactory.getBuilder()

        return requestBuilder.init(method: "GET", URLString: URLString, parameters: convertedParameters, isBody: true)
    }

    /**
     Get a list of trunks for an account
     
     - parameter accountId: (path) Account ID 
     - parameter filtersId: (query) ID filter (optional)
     - parameter filtersName: (query) Name filter (optional)
     - parameter sortId: (query) ID sorting (optional)
     - parameter sortName: (query) Name sorting (optional)
     - parameter limit: (query) Max results (optional)
     - parameter offset: (query) Results to skip (optional)
     - parameter fields: (query) Field set (optional)
     - parameter completion: completion handler to receive the data and the error objects
     */
    public class func listAccountTrunks(accountId accountId: Int32, filtersId: [String]? = nil, filtersName: [String]? = nil, sortId: String? = nil, sortName: String? = nil, limit: Int32? = nil, offset: Int32? = nil, fields: String? = nil, completion: ((data: ListTrunksFull?, error: ErrorType?) -> Void)) {
        listAccountTrunksWithRequestBuilder(accountId: accountId, filtersId: filtersId, filtersName: filtersName, sortId: sortId, sortName: sortName, limit: limit, offset: offset, fields: fields).execute { (response, error) -> Void in
            completion(data: response?.body, error: error);
        }
    }


    /**
     Get a list of trunks for an account
     - GET /accounts/{accountId}/trunks
     - See Account Trunks for more info on the properties.
     - API Key:
       - type: apiKey Authorization 
       - name: apiKey
     - examples: [{contentType=application/json, example={
  "total" : "",
  "offset" : "",
  "limit" : "",
  "filters" : {
    "name" : "aeiou",
    "id" : "aeiou"
  },
  "sort" : {
    "name" : "aeiou",
    "id" : "aeiou"
  },
  "items" : [ {
    "error_message" : "",
    "max_minutes_per_month" : "",
    "greeting" : {
      "name" : "aeiou",
      "id" : ""
    },
    "codecs" : [ "aeiou" ],
    "name" : "aeiou",
    "id" : "",
    "max_concurrent_calls" : "",
    "uri" : "aeiou"
  } ]
}}]
     
     - parameter accountId: (path) Account ID 
     - parameter filtersId: (query) ID filter (optional)
     - parameter filtersName: (query) Name filter (optional)
     - parameter sortId: (query) ID sorting (optional)
     - parameter sortName: (query) Name sorting (optional)
     - parameter limit: (query) Max results (optional)
     - parameter offset: (query) Results to skip (optional)
     - parameter fields: (query) Field set (optional)

     - returns: RequestBuilder<ListTrunksFull> 
     */
    public class func listAccountTrunksWithRequestBuilder(accountId accountId: Int32, filtersId: [String]? = nil, filtersName: [String]? = nil, sortId: String? = nil, sortName: String? = nil, limit: Int32? = nil, offset: Int32? = nil, fields: String? = nil) -> RequestBuilder<ListTrunksFull> {
        var path = "/accounts/{accountId}/trunks"
        path = path.stringByReplacingOccurrencesOfString("{accountId}", withString: "\(accountId)", options: .LiteralSearch, range: nil)
        let URLString = SwaggerClientAPI.basePath + path

        let nillableParameters: [String:AnyObject?] = [
            "filters[id]": filtersId,
            "filters[name]": filtersName,
            "sort[id]": sortId,
            "sort[name]": sortName,
            "limit": limit?.encodeToJSON(),
            "offset": offset?.encodeToJSON(),
            "fields": fields
        ]
 
        let parameters = APIHelper.rejectNil(nillableParameters)
 
        let convertedParameters = APIHelper.convertBoolToString(parameters)
 
        let requestBuilder: RequestBuilder<ListTrunksFull>.Type = SwaggerClientAPI.requestBuilderFactory.getBuilder()

        return requestBuilder.init(method: "GET", URLString: URLString, parameters: convertedParameters, isBody: false)
    }

    /**
     Replace parameters in a trunk
     
     - parameter accountId: (path) Account ID 
     - parameter trunkId: (path) Trunk ID 
     - parameter data: (body) Trunk data 
     - parameter completion: completion handler to receive the data and the error objects
     */
    public class func replaceAccountTrunk(accountId accountId: Int32, trunkId: Int32, data: CreateTrunkParams, completion: ((data: TrunkFull?, error: ErrorType?) -> Void)) {
        replaceAccountTrunkWithRequestBuilder(accountId: accountId, trunkId: trunkId, data: data).execute { (response, error) -> Void in
            completion(data: response?.body, error: error);
        }
    }


    /**
     Replace parameters in a trunk
     - PUT /accounts/{accountId}/trunks/{trunkId}
     - For more on the input fields, see Account Trunks.
     - API Key:
       - type: apiKey Authorization 
       - name: apiKey
     - examples: [{contentType=application/json, example={
  "error_message" : "",
  "max_minutes_per_month" : "",
  "greeting" : {
    "name" : "aeiou",
    "id" : ""
  },
  "codecs" : [ "aeiou" ],
  "name" : "aeiou",
  "id" : "",
  "max_concurrent_calls" : "",
  "uri" : "aeiou"
}}]
     
     - parameter accountId: (path) Account ID 
     - parameter trunkId: (path) Trunk ID 
     - parameter data: (body) Trunk data 

     - returns: RequestBuilder<TrunkFull> 
     */
    public class func replaceAccountTrunkWithRequestBuilder(accountId accountId: Int32, trunkId: Int32, data: CreateTrunkParams) -> RequestBuilder<TrunkFull> {
        var path = "/accounts/{accountId}/trunks/{trunkId}"
        path = path.stringByReplacingOccurrencesOfString("{accountId}", withString: "\(accountId)", options: .LiteralSearch, range: nil)
        path = path.stringByReplacingOccurrencesOfString("{trunkId}", withString: "\(trunkId)", options: .LiteralSearch, range: nil)
        let URLString = SwaggerClientAPI.basePath + path
        let parameters = data.encodeToJSON() as? [String:AnyObject]
 
        let convertedParameters = APIHelper.convertBoolToString(parameters)
 
        let requestBuilder: RequestBuilder<TrunkFull>.Type = SwaggerClientAPI.requestBuilderFactory.getBuilder()

        return requestBuilder.init(method: "PUT", URLString: URLString, parameters: convertedParameters, isBody: true)
    }

}