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

import io.swagger.client.model.SmsFull;
import io.swagger.client.model.CreateSmsParams;
import io.swagger.client.model.ListSmsFull;

import org.apache.http.HttpEntity;
import org.apache.http.entity.mime.MultipartEntityBuilder;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;
import java.util.concurrent.ExecutionException;
import java.util.concurrent.TimeoutException;

public class SmsApi {
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
  * Send a SMS to one or a group of recipients
  * For more on the input fields, see Intro to SMS.
   * @param accountId Account ID
   * @param data SMS data
   * @return SmsFull
  */
  public SmsFull createAccountSms (Integer accountId, CreateSmsParams data) throws TimeoutException, ExecutionException, InterruptedException, ApiException {
     Object postBody = data;
  
      // verify the required parameter 'accountId' is set
      if (accountId == null) {
      VolleyError error = new VolleyError("Missing the required parameter 'accountId' when calling createAccountSms",
      new ApiException(400, "Missing the required parameter 'accountId' when calling createAccountSms"));
      }
  
      // verify the required parameter 'data' is set
      if (data == null) {
      VolleyError error = new VolleyError("Missing the required parameter 'data' when calling createAccountSms",
      new ApiException(400, "Missing the required parameter 'data' when calling createAccountSms"));
      }
  

  // create path and map variables
  String path = "/accounts/{account_id}/sms".replaceAll("\\{format\\}","json").replaceAll("\\{" + "account_id" + "\\}", apiInvoker.escapeString(accountId.toString()));

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
           return (SmsFull) ApiInvoker.deserialize(localVarResponse, "", SmsFull.class);
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
   * Send a SMS to one or a group of recipients
   * For more on the input fields, see Intro to SMS.
   * @param accountId Account ID   * @param data SMS data
  */
  public void createAccountSms (Integer accountId, CreateSmsParams data, final Response.Listener<SmsFull> responseListener, final Response.ErrorListener errorListener) {
    Object postBody = data;

  
    // verify the required parameter 'accountId' is set
    if (accountId == null) {
       VolleyError error = new VolleyError("Missing the required parameter 'accountId' when calling createAccountSms",
         new ApiException(400, "Missing the required parameter 'accountId' when calling createAccountSms"));
    }
    
    // verify the required parameter 'data' is set
    if (data == null) {
       VolleyError error = new VolleyError("Missing the required parameter 'data' when calling createAccountSms",
         new ApiException(400, "Missing the required parameter 'data' when calling createAccountSms"));
    }
    

    // create path and map variables
    String path = "/accounts/{account_id}/sms".replaceAll("\\{format\\}","json").replaceAll("\\{" + "account_id" + "\\}", apiInvoker.escapeString(accountId.toString()));

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
              responseListener.onResponse((SmsFull) ApiInvoker.deserialize(localVarResponse,  "", SmsFull.class));
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
  * Show details of an individual SMS
  * This service shows the details of an individual sms. See Intro to SMS for more info on the properties.
   * @param accountId Account ID
   * @param smsId SMS ID
   * @return SmsFull
  */
  public SmsFull getAccountSms (Integer accountId, Integer smsId) throws TimeoutException, ExecutionException, InterruptedException, ApiException {
     Object postBody = null;
  
      // verify the required parameter 'accountId' is set
      if (accountId == null) {
      VolleyError error = new VolleyError("Missing the required parameter 'accountId' when calling getAccountSms",
      new ApiException(400, "Missing the required parameter 'accountId' when calling getAccountSms"));
      }
  
      // verify the required parameter 'smsId' is set
      if (smsId == null) {
      VolleyError error = new VolleyError("Missing the required parameter 'smsId' when calling getAccountSms",
      new ApiException(400, "Missing the required parameter 'smsId' when calling getAccountSms"));
      }
  

  // create path and map variables
  String path = "/accounts/{account_id}/sms/{sms_id}".replaceAll("\\{format\\}","json").replaceAll("\\{" + "account_id" + "\\}", apiInvoker.escapeString(accountId.toString())).replaceAll("\\{" + "sms_id" + "\\}", apiInvoker.escapeString(smsId.toString()));

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
           return (SmsFull) ApiInvoker.deserialize(localVarResponse, "", SmsFull.class);
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
   * Show details of an individual SMS
   * This service shows the details of an individual sms. See Intro to SMS for more info on the properties.
   * @param accountId Account ID   * @param smsId SMS ID
  */
  public void getAccountSms (Integer accountId, Integer smsId, final Response.Listener<SmsFull> responseListener, final Response.ErrorListener errorListener) {
    Object postBody = null;

  
    // verify the required parameter 'accountId' is set
    if (accountId == null) {
       VolleyError error = new VolleyError("Missing the required parameter 'accountId' when calling getAccountSms",
         new ApiException(400, "Missing the required parameter 'accountId' when calling getAccountSms"));
    }
    
    // verify the required parameter 'smsId' is set
    if (smsId == null) {
       VolleyError error = new VolleyError("Missing the required parameter 'smsId' when calling getAccountSms",
         new ApiException(400, "Missing the required parameter 'smsId' when calling getAccountSms"));
    }
    

    // create path and map variables
    String path = "/accounts/{account_id}/sms/{sms_id}".replaceAll("\\{format\\}","json").replaceAll("\\{" + "account_id" + "\\}", apiInvoker.escapeString(accountId.toString())).replaceAll("\\{" + "sms_id" + "\\}", apiInvoker.escapeString(smsId.toString()));

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
              responseListener.onResponse((SmsFull) ApiInvoker.deserialize(localVarResponse,  "", SmsFull.class));
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
  * Get a list of SMS messages for an account
  * See Intro to SMS for more info on the properties.
   * @param accountId Account ID
   * @param filtersId ID filter
   * @param filtersDirection Direction filter
   * @param filtersFrom Caller ID filter
   * @param sortId ID sorting
   * @param sortCreatedAt Sort by created time of message
   * @param limit Max results
   * @param offset Results to skip
   * @param fields Field set
   * @return ListSmsFull
  */
  public ListSmsFull listAccountSms (Integer accountId, List<String> filtersId, String filtersDirection, String filtersFrom, String sortId, String sortCreatedAt, Integer limit, Integer offset, String fields) throws TimeoutException, ExecutionException, InterruptedException, ApiException {
     Object postBody = null;
  
      // verify the required parameter 'accountId' is set
      if (accountId == null) {
      VolleyError error = new VolleyError("Missing the required parameter 'accountId' when calling listAccountSms",
      new ApiException(400, "Missing the required parameter 'accountId' when calling listAccountSms"));
      }
  

  // create path and map variables
  String path = "/accounts/{account_id}/sms".replaceAll("\\{format\\}","json").replaceAll("\\{" + "account_id" + "\\}", apiInvoker.escapeString(accountId.toString()));

  // query params
  List<Pair> queryParams = new ArrayList<Pair>();
      // header params
      Map<String, String> headerParams = new HashMap<String, String>();
      // form params
      Map<String, String> formParams = new HashMap<String, String>();

          queryParams.addAll(ApiInvoker.parameterToPairs("multi", "filters[id]", filtersId));
          queryParams.addAll(ApiInvoker.parameterToPairs("", "filters[direction]", filtersDirection));
          queryParams.addAll(ApiInvoker.parameterToPairs("", "filters[from]", filtersFrom));
          queryParams.addAll(ApiInvoker.parameterToPairs("", "sort[id]", sortId));
          queryParams.addAll(ApiInvoker.parameterToPairs("", "sort[created_at]", sortCreatedAt));
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
           return (ListSmsFull) ApiInvoker.deserialize(localVarResponse, "", ListSmsFull.class);
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
   * Get a list of SMS messages for an account
   * See Intro to SMS for more info on the properties.
   * @param accountId Account ID   * @param filtersId ID filter   * @param filtersDirection Direction filter   * @param filtersFrom Caller ID filter   * @param sortId ID sorting   * @param sortCreatedAt Sort by created time of message   * @param limit Max results   * @param offset Results to skip   * @param fields Field set
  */
  public void listAccountSms (Integer accountId, List<String> filtersId, String filtersDirection, String filtersFrom, String sortId, String sortCreatedAt, Integer limit, Integer offset, String fields, final Response.Listener<ListSmsFull> responseListener, final Response.ErrorListener errorListener) {
    Object postBody = null;

  
    // verify the required parameter 'accountId' is set
    if (accountId == null) {
       VolleyError error = new VolleyError("Missing the required parameter 'accountId' when calling listAccountSms",
         new ApiException(400, "Missing the required parameter 'accountId' when calling listAccountSms"));
    }
    

    // create path and map variables
    String path = "/accounts/{account_id}/sms".replaceAll("\\{format\\}","json").replaceAll("\\{" + "account_id" + "\\}", apiInvoker.escapeString(accountId.toString()));

    // query params
    List<Pair> queryParams = new ArrayList<Pair>();
    // header params
    Map<String, String> headerParams = new HashMap<String, String>();
    // form params
    Map<String, String> formParams = new HashMap<String, String>();

    queryParams.addAll(ApiInvoker.parameterToPairs("multi", "filters[id]", filtersId));
    queryParams.addAll(ApiInvoker.parameterToPairs("", "filters[direction]", filtersDirection));
    queryParams.addAll(ApiInvoker.parameterToPairs("", "filters[from]", filtersFrom));
    queryParams.addAll(ApiInvoker.parameterToPairs("", "sort[id]", sortId));
    queryParams.addAll(ApiInvoker.parameterToPairs("", "sort[created_at]", sortCreatedAt));
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
              responseListener.onResponse((ListSmsFull) ApiInvoker.deserialize(localVarResponse,  "", ListSmsFull.class));
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