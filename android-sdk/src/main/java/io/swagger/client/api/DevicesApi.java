/**
 * Phone.com API
 * This is a Phone.com api Swagger definition
 *
 * OpenAPI spec version: 1.0.0
 * Contact: apisupport@phone.com
 *
 * NOTE: This class is auto generated by the swagger code generator program.
 * https://github.com/swagger-api/swagger-codegen.git
 * Do not edit the class manually.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package io.swagger.client.api;

import io.swagger.client.ApiInvoker;
import io.swagger.client.ApiException;
import io.swagger.client.Pair;

import io.swagger.client.model.*;

import java.util.*;

import com.android.volley.Response;
import com.android.volley.VolleyError;

import io.swagger.client.model.CreateDeviceParams;
import io.swagger.client.model.DeviceFull;
import io.swagger.client.model.ListDevicesFull;

import org.apache.http.HttpEntity;
import org.apache.http.entity.mime.MultipartEntityBuilder;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;
import java.util.concurrent.ExecutionException;
import java.util.concurrent.TimeoutException;

public class DevicesApi {
  String basePath = "https://api.phone.com/v4";
  ApiInvoker apiInvoker = ApiInvoker.getInstance();

  public void addHeader(String key, String value) {
    getInvoker().addDefaultHeader(key, value);
  }

  public ApiInvoker getInvoker() {
    return apiInvoker;
  }

  public void setBasePath(String basePath) {
    this.basePath = basePath;
  }

  public String getBasePath() {
    return basePath;
  }

  /**
  * Register a generic VoIP device
  * 
   * @param accountId Account ID
   * @param data Device data
   * @return DeviceFull
  */
  public DeviceFull createAccountDevice (Integer accountId, CreateDeviceParams data) throws TimeoutException, ExecutionException, InterruptedException, ApiException {
     Object postBody = data;
  
      // verify the required parameter 'accountId' is set
      if (accountId == null) {
      VolleyError error = new VolleyError("Missing the required parameter 'accountId' when calling createAccountDevice",
      new ApiException(400, "Missing the required parameter 'accountId' when calling createAccountDevice"));
      }
  

  // create path and map variables
  String path = "/accounts/{account_id}/devices".replaceAll("\\{format\\}","json").replaceAll("\\{" + "account_id" + "\\}", apiInvoker.escapeString(accountId.toString()));

  // query params
  List<Pair> queryParams = new ArrayList<Pair>();
      // header params
      Map<String, String> headerParams = new HashMap<String, String>();
      // form params
      Map<String, String> formParams = new HashMap<String, String>();



      String[] contentTypes = {
  "application/json"
      };
      String contentType = contentTypes.length > 0 ? contentTypes[0] : "application/json";

      if (contentType.startsWith("multipart/form-data")) {
      // file uploading
      MultipartEntityBuilder localVarBuilder = MultipartEntityBuilder.create();
  

      HttpEntity httpEntity = localVarBuilder.build();
      postBody = httpEntity;
      } else {
      // normal form params
        }

      String[] authNames = new String[] { "apiKey" };

      try {
        String localVarResponse = apiInvoker.invokeAPI (basePath, path, "POST", queryParams, postBody, headerParams, formParams, contentType, authNames);
        if(localVarResponse != null){
           return (DeviceFull) ApiInvoker.deserialize(localVarResponse, "", DeviceFull.class);
        } else {
           return null;
        }
      } catch (ApiException ex) {
         throw ex;
      } catch (InterruptedException ex) {
         throw ex;
      } catch (ExecutionException ex) {
         if(ex.getCause() instanceof VolleyError) {
	    VolleyError volleyError = (VolleyError)ex.getCause();
	    if (volleyError.networkResponse != null) {
	       throw new ApiException(volleyError.networkResponse.statusCode, volleyError.getMessage());
	    }
         }
         throw ex;
      } catch (TimeoutException ex) {
         throw ex;
      }
  }

      /**
   * Register a generic VoIP device
   * 
   * @param accountId Account ID   * @param data Device data
  */
  public void createAccountDevice (Integer accountId, CreateDeviceParams data, final Response.Listener<DeviceFull> responseListener, final Response.ErrorListener errorListener) {
    Object postBody = data;

  
    // verify the required parameter 'accountId' is set
    if (accountId == null) {
       VolleyError error = new VolleyError("Missing the required parameter 'accountId' when calling createAccountDevice",
         new ApiException(400, "Missing the required parameter 'accountId' when calling createAccountDevice"));
    }
    

    // create path and map variables
    String path = "/accounts/{account_id}/devices".replaceAll("\\{format\\}","json").replaceAll("\\{" + "account_id" + "\\}", apiInvoker.escapeString(accountId.toString()));

    // query params
    List<Pair> queryParams = new ArrayList<Pair>();
    // header params
    Map<String, String> headerParams = new HashMap<String, String>();
    // form params
    Map<String, String> formParams = new HashMap<String, String>();



    String[] contentTypes = {
      "application/json"
    };
    String contentType = contentTypes.length > 0 ? contentTypes[0] : "application/json";

    if (contentType.startsWith("multipart/form-data")) {
      // file uploading
      MultipartEntityBuilder localVarBuilder = MultipartEntityBuilder.create();
      

      HttpEntity httpEntity = localVarBuilder.build();
      postBody = httpEntity;
    } else {
      // normal form params
          }

      String[] authNames = new String[] { "apiKey" };

    try {
      apiInvoker.invokeAPI(basePath, path, "POST", queryParams, postBody, headerParams, formParams, contentType, authNames,
        new Response.Listener<String>() {
          @Override
          public void onResponse(String localVarResponse) {
            try {
              responseListener.onResponse((DeviceFull) ApiInvoker.deserialize(localVarResponse,  "", DeviceFull.class));
            } catch (ApiException exception) {
               errorListener.onErrorResponse(new VolleyError(exception));
            }
          }
      }, new Response.ErrorListener() {
          @Override
          public void onErrorResponse(VolleyError error) {
            errorListener.onErrorResponse(error);
          }
      });
    } catch (ApiException ex) {
      errorListener.onErrorResponse(new VolleyError(ex));
    }
  }
  /**
  * Show details of an individual VoIP device
  * 
   * @param accountId Account ID
   * @param deviceId Device ID
   * @return DeviceFull
  */
  public DeviceFull getAccountDevice (Integer accountId, Integer deviceId) throws TimeoutException, ExecutionException, InterruptedException, ApiException {
     Object postBody = null;
  
      // verify the required parameter 'accountId' is set
      if (accountId == null) {
      VolleyError error = new VolleyError("Missing the required parameter 'accountId' when calling getAccountDevice",
      new ApiException(400, "Missing the required parameter 'accountId' when calling getAccountDevice"));
      }
  
      // verify the required parameter 'deviceId' is set
      if (deviceId == null) {
      VolleyError error = new VolleyError("Missing the required parameter 'deviceId' when calling getAccountDevice",
      new ApiException(400, "Missing the required parameter 'deviceId' when calling getAccountDevice"));
      }
  

  // create path and map variables
  String path = "/accounts/{account_id}/devices/{device_id}".replaceAll("\\{format\\}","json").replaceAll("\\{" + "account_id" + "\\}", apiInvoker.escapeString(accountId.toString())).replaceAll("\\{" + "device_id" + "\\}", apiInvoker.escapeString(deviceId.toString()));

  // query params
  List<Pair> queryParams = new ArrayList<Pair>();
      // header params
      Map<String, String> headerParams = new HashMap<String, String>();
      // form params
      Map<String, String> formParams = new HashMap<String, String>();



      String[] contentTypes = {
  "application/json"
      };
      String contentType = contentTypes.length > 0 ? contentTypes[0] : "application/json";

      if (contentType.startsWith("multipart/form-data")) {
      // file uploading
      MultipartEntityBuilder localVarBuilder = MultipartEntityBuilder.create();
  

      HttpEntity httpEntity = localVarBuilder.build();
      postBody = httpEntity;
      } else {
      // normal form params
        }

      String[] authNames = new String[] { "apiKey" };

      try {
        String localVarResponse = apiInvoker.invokeAPI (basePath, path, "GET", queryParams, postBody, headerParams, formParams, contentType, authNames);
        if(localVarResponse != null){
           return (DeviceFull) ApiInvoker.deserialize(localVarResponse, "", DeviceFull.class);
        } else {
           return null;
        }
      } catch (ApiException ex) {
         throw ex;
      } catch (InterruptedException ex) {
         throw ex;
      } catch (ExecutionException ex) {
         if(ex.getCause() instanceof VolleyError) {
	    VolleyError volleyError = (VolleyError)ex.getCause();
	    if (volleyError.networkResponse != null) {
	       throw new ApiException(volleyError.networkResponse.statusCode, volleyError.getMessage());
	    }
         }
         throw ex;
      } catch (TimeoutException ex) {
         throw ex;
      }
  }

      /**
   * Show details of an individual VoIP device
   * 
   * @param accountId Account ID   * @param deviceId Device ID
  */
  public void getAccountDevice (Integer accountId, Integer deviceId, final Response.Listener<DeviceFull> responseListener, final Response.ErrorListener errorListener) {
    Object postBody = null;

  
    // verify the required parameter 'accountId' is set
    if (accountId == null) {
       VolleyError error = new VolleyError("Missing the required parameter 'accountId' when calling getAccountDevice",
         new ApiException(400, "Missing the required parameter 'accountId' when calling getAccountDevice"));
    }
    
    // verify the required parameter 'deviceId' is set
    if (deviceId == null) {
       VolleyError error = new VolleyError("Missing the required parameter 'deviceId' when calling getAccountDevice",
         new ApiException(400, "Missing the required parameter 'deviceId' when calling getAccountDevice"));
    }
    

    // create path and map variables
    String path = "/accounts/{account_id}/devices/{device_id}".replaceAll("\\{format\\}","json").replaceAll("\\{" + "account_id" + "\\}", apiInvoker.escapeString(accountId.toString())).replaceAll("\\{" + "device_id" + "\\}", apiInvoker.escapeString(deviceId.toString()));

    // query params
    List<Pair> queryParams = new ArrayList<Pair>();
    // header params
    Map<String, String> headerParams = new HashMap<String, String>();
    // form params
    Map<String, String> formParams = new HashMap<String, String>();



    String[] contentTypes = {
      "application/json"
    };
    String contentType = contentTypes.length > 0 ? contentTypes[0] : "application/json";

    if (contentType.startsWith("multipart/form-data")) {
      // file uploading
      MultipartEntityBuilder localVarBuilder = MultipartEntityBuilder.create();
      

      HttpEntity httpEntity = localVarBuilder.build();
      postBody = httpEntity;
    } else {
      // normal form params
          }

      String[] authNames = new String[] { "apiKey" };

    try {
      apiInvoker.invokeAPI(basePath, path, "GET", queryParams, postBody, headerParams, formParams, contentType, authNames,
        new Response.Listener<String>() {
          @Override
          public void onResponse(String localVarResponse) {
            try {
              responseListener.onResponse((DeviceFull) ApiInvoker.deserialize(localVarResponse,  "", DeviceFull.class));
            } catch (ApiException exception) {
               errorListener.onErrorResponse(new VolleyError(exception));
            }
          }
      }, new Response.ErrorListener() {
          @Override
          public void onErrorResponse(VolleyError error) {
            errorListener.onErrorResponse(error);
          }
      });
    } catch (ApiException ex) {
      errorListener.onErrorResponse(new VolleyError(ex));
    }
  }
  /**
  * Get a list of VoIP devices associated with your account
  * 
   * @param accountId Account ID
   * @param filtersId ID filter
   * @param filtersName Name filter
   * @param sortId ID sorting
   * @param sortName Name sorting
   * @param limit Max results
   * @param offset Results to skip
   * @param fields Field set
   * @return ListDevicesFull
  */
  public ListDevicesFull listAccountDevices (Integer accountId, List<String> filtersId, List<String> filtersName, String sortId, String sortName, Integer limit, Integer offset, String fields) throws TimeoutException, ExecutionException, InterruptedException, ApiException {
     Object postBody = null;
  
      // verify the required parameter 'accountId' is set
      if (accountId == null) {
      VolleyError error = new VolleyError("Missing the required parameter 'accountId' when calling listAccountDevices",
      new ApiException(400, "Missing the required parameter 'accountId' when calling listAccountDevices"));
      }
  

  // create path and map variables
  String path = "/accounts/{account_id}/devices".replaceAll("\\{format\\}","json").replaceAll("\\{" + "account_id" + "\\}", apiInvoker.escapeString(accountId.toString()));

  // query params
  List<Pair> queryParams = new ArrayList<Pair>();
      // header params
      Map<String, String> headerParams = new HashMap<String, String>();
      // form params
      Map<String, String> formParams = new HashMap<String, String>();

          queryParams.addAll(ApiInvoker.parameterToPairs("multi", "filters[id]", filtersId));
          queryParams.addAll(ApiInvoker.parameterToPairs("multi", "filters[name]", filtersName));
          queryParams.addAll(ApiInvoker.parameterToPairs("", "sort[id]", sortId));
          queryParams.addAll(ApiInvoker.parameterToPairs("", "sort[name]", sortName));
          queryParams.addAll(ApiInvoker.parameterToPairs("", "limit", limit));
          queryParams.addAll(ApiInvoker.parameterToPairs("", "offset", offset));
          queryParams.addAll(ApiInvoker.parameterToPairs("", "fields", fields));


      String[] contentTypes = {
  "application/json"
      };
      String contentType = contentTypes.length > 0 ? contentTypes[0] : "application/json";

      if (contentType.startsWith("multipart/form-data")) {
      // file uploading
      MultipartEntityBuilder localVarBuilder = MultipartEntityBuilder.create();
  

      HttpEntity httpEntity = localVarBuilder.build();
      postBody = httpEntity;
      } else {
      // normal form params
        }

      String[] authNames = new String[] { "apiKey" };

      try {
        String localVarResponse = apiInvoker.invokeAPI (basePath, path, "GET", queryParams, postBody, headerParams, formParams, contentType, authNames);
        if(localVarResponse != null){
           return (ListDevicesFull) ApiInvoker.deserialize(localVarResponse, "", ListDevicesFull.class);
        } else {
           return null;
        }
      } catch (ApiException ex) {
         throw ex;
      } catch (InterruptedException ex) {
         throw ex;
      } catch (ExecutionException ex) {
         if(ex.getCause() instanceof VolleyError) {
	    VolleyError volleyError = (VolleyError)ex.getCause();
	    if (volleyError.networkResponse != null) {
	       throw new ApiException(volleyError.networkResponse.statusCode, volleyError.getMessage());
	    }
         }
         throw ex;
      } catch (TimeoutException ex) {
         throw ex;
      }
  }

      /**
   * Get a list of VoIP devices associated with your account
   * 
   * @param accountId Account ID   * @param filtersId ID filter   * @param filtersName Name filter   * @param sortId ID sorting   * @param sortName Name sorting   * @param limit Max results   * @param offset Results to skip   * @param fields Field set
  */
  public void listAccountDevices (Integer accountId, List<String> filtersId, List<String> filtersName, String sortId, String sortName, Integer limit, Integer offset, String fields, final Response.Listener<ListDevicesFull> responseListener, final Response.ErrorListener errorListener) {
    Object postBody = null;

  
    // verify the required parameter 'accountId' is set
    if (accountId == null) {
       VolleyError error = new VolleyError("Missing the required parameter 'accountId' when calling listAccountDevices",
         new ApiException(400, "Missing the required parameter 'accountId' when calling listAccountDevices"));
    }
    

    // create path and map variables
    String path = "/accounts/{account_id}/devices".replaceAll("\\{format\\}","json").replaceAll("\\{" + "account_id" + "\\}", apiInvoker.escapeString(accountId.toString()));

    // query params
    List<Pair> queryParams = new ArrayList<Pair>();
    // header params
    Map<String, String> headerParams = new HashMap<String, String>();
    // form params
    Map<String, String> formParams = new HashMap<String, String>();

    queryParams.addAll(ApiInvoker.parameterToPairs("multi", "filters[id]", filtersId));
    queryParams.addAll(ApiInvoker.parameterToPairs("multi", "filters[name]", filtersName));
    queryParams.addAll(ApiInvoker.parameterToPairs("", "sort[id]", sortId));
    queryParams.addAll(ApiInvoker.parameterToPairs("", "sort[name]", sortName));
    queryParams.addAll(ApiInvoker.parameterToPairs("", "limit", limit));
    queryParams.addAll(ApiInvoker.parameterToPairs("", "offset", offset));
    queryParams.addAll(ApiInvoker.parameterToPairs("", "fields", fields));


    String[] contentTypes = {
      "application/json"
    };
    String contentType = contentTypes.length > 0 ? contentTypes[0] : "application/json";

    if (contentType.startsWith("multipart/form-data")) {
      // file uploading
      MultipartEntityBuilder localVarBuilder = MultipartEntityBuilder.create();
      

      HttpEntity httpEntity = localVarBuilder.build();
      postBody = httpEntity;
    } else {
      // normal form params
          }

      String[] authNames = new String[] { "apiKey" };

    try {
      apiInvoker.invokeAPI(basePath, path, "GET", queryParams, postBody, headerParams, formParams, contentType, authNames,
        new Response.Listener<String>() {
          @Override
          public void onResponse(String localVarResponse) {
            try {
              responseListener.onResponse((ListDevicesFull) ApiInvoker.deserialize(localVarResponse,  "", ListDevicesFull.class));
            } catch (ApiException exception) {
               errorListener.onErrorResponse(new VolleyError(exception));
            }
          }
      }, new Response.ErrorListener() {
          @Override
          public void onErrorResponse(VolleyError error) {
            errorListener.onErrorResponse(error);
          }
      });
    } catch (ApiException ex) {
      errorListener.onErrorResponse(new VolleyError(ex));
    }
  }
  /**
  * Update the settings for an individual VoIP device
  * 
   * @param accountId Account ID
   * @param deviceId Device ID
   * @param data Device data
   * @return DeviceFull
  */
  public DeviceFull replaceAccountDevice (Integer accountId, Integer deviceId, CreateDeviceParams data) throws TimeoutException, ExecutionException, InterruptedException, ApiException {
     Object postBody = data;
  
      // verify the required parameter 'accountId' is set
      if (accountId == null) {
      VolleyError error = new VolleyError("Missing the required parameter 'accountId' when calling replaceAccountDevice",
      new ApiException(400, "Missing the required parameter 'accountId' when calling replaceAccountDevice"));
      }
  
      // verify the required parameter 'deviceId' is set
      if (deviceId == null) {
      VolleyError error = new VolleyError("Missing the required parameter 'deviceId' when calling replaceAccountDevice",
      new ApiException(400, "Missing the required parameter 'deviceId' when calling replaceAccountDevice"));
      }
  

  // create path and map variables
  String path = "/accounts/{account_id}/devices/{device_id}".replaceAll("\\{format\\}","json").replaceAll("\\{" + "account_id" + "\\}", apiInvoker.escapeString(accountId.toString())).replaceAll("\\{" + "device_id" + "\\}", apiInvoker.escapeString(deviceId.toString()));

  // query params
  List<Pair> queryParams = new ArrayList<Pair>();
      // header params
      Map<String, String> headerParams = new HashMap<String, String>();
      // form params
      Map<String, String> formParams = new HashMap<String, String>();



      String[] contentTypes = {
  "application/json"
      };
      String contentType = contentTypes.length > 0 ? contentTypes[0] : "application/json";

      if (contentType.startsWith("multipart/form-data")) {
      // file uploading
      MultipartEntityBuilder localVarBuilder = MultipartEntityBuilder.create();
  

      HttpEntity httpEntity = localVarBuilder.build();
      postBody = httpEntity;
      } else {
      // normal form params
        }

      String[] authNames = new String[] { "apiKey" };

      try {
        String localVarResponse = apiInvoker.invokeAPI (basePath, path, "PUT", queryParams, postBody, headerParams, formParams, contentType, authNames);
        if(localVarResponse != null){
           return (DeviceFull) ApiInvoker.deserialize(localVarResponse, "", DeviceFull.class);
        } else {
           return null;
        }
      } catch (ApiException ex) {
         throw ex;
      } catch (InterruptedException ex) {
         throw ex;
      } catch (ExecutionException ex) {
         if(ex.getCause() instanceof VolleyError) {
	    VolleyError volleyError = (VolleyError)ex.getCause();
	    if (volleyError.networkResponse != null) {
	       throw new ApiException(volleyError.networkResponse.statusCode, volleyError.getMessage());
	    }
         }
         throw ex;
      } catch (TimeoutException ex) {
         throw ex;
      }
  }

      /**
   * Update the settings for an individual VoIP device
   * 
   * @param accountId Account ID   * @param deviceId Device ID   * @param data Device data
  */
  public void replaceAccountDevice (Integer accountId, Integer deviceId, CreateDeviceParams data, final Response.Listener<DeviceFull> responseListener, final Response.ErrorListener errorListener) {
    Object postBody = data;

  
    // verify the required parameter 'accountId' is set
    if (accountId == null) {
       VolleyError error = new VolleyError("Missing the required parameter 'accountId' when calling replaceAccountDevice",
         new ApiException(400, "Missing the required parameter 'accountId' when calling replaceAccountDevice"));
    }
    
    // verify the required parameter 'deviceId' is set
    if (deviceId == null) {
       VolleyError error = new VolleyError("Missing the required parameter 'deviceId' when calling replaceAccountDevice",
         new ApiException(400, "Missing the required parameter 'deviceId' when calling replaceAccountDevice"));
    }
    

    // create path and map variables
    String path = "/accounts/{account_id}/devices/{device_id}".replaceAll("\\{format\\}","json").replaceAll("\\{" + "account_id" + "\\}", apiInvoker.escapeString(accountId.toString())).replaceAll("\\{" + "device_id" + "\\}", apiInvoker.escapeString(deviceId.toString()));

    // query params
    List<Pair> queryParams = new ArrayList<Pair>();
    // header params
    Map<String, String> headerParams = new HashMap<String, String>();
    // form params
    Map<String, String> formParams = new HashMap<String, String>();



    String[] contentTypes = {
      "application/json"
    };
    String contentType = contentTypes.length > 0 ? contentTypes[0] : "application/json";

    if (contentType.startsWith("multipart/form-data")) {
      // file uploading
      MultipartEntityBuilder localVarBuilder = MultipartEntityBuilder.create();
      

      HttpEntity httpEntity = localVarBuilder.build();
      postBody = httpEntity;
    } else {
      // normal form params
          }

      String[] authNames = new String[] { "apiKey" };

    try {
      apiInvoker.invokeAPI(basePath, path, "PUT", queryParams, postBody, headerParams, formParams, contentType, authNames,
        new Response.Listener<String>() {
          @Override
          public void onResponse(String localVarResponse) {
            try {
              responseListener.onResponse((DeviceFull) ApiInvoker.deserialize(localVarResponse,  "", DeviceFull.class));
            } catch (ApiException exception) {
               errorListener.onErrorResponse(new VolleyError(exception));
            }
          }
      }, new Response.ErrorListener() {
          @Override
          public void onErrorResponse(VolleyError error) {
            errorListener.onErrorResponse(error);
          }
      });
    } catch (ApiException ex) {
      errorListener.onErrorResponse(new VolleyError(ex));
    }
  }
}
