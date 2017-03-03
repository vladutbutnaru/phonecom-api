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

import io.swagger.client.model.CreateGroupParams;
import io.swagger.client.model.GroupFull;
import io.swagger.client.model.DeleteGroup;
import io.swagger.client.model.ListGroupsFull;

import org.apache.http.HttpEntity;
import org.apache.http.entity.mime.MultipartEntityBuilder;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;
import java.util.concurrent.ExecutionException;
import java.util.concurrent.TimeoutException;

public class GroupsApi {
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
  * 
  * See Account Contact Groups for more info on the properties.
   * @param accountId Account ID
   * @param extensionId Extension ID
   * @param data Group name
   * @return GroupFull
  */
  public GroupFull createAccountExtensionContactGroup (Integer accountId, Integer extensionId, CreateGroupParams data) throws TimeoutException, ExecutionException, InterruptedException, ApiException {
     Object postBody = data;
  
      // verify the required parameter 'accountId' is set
      if (accountId == null) {
      VolleyError error = new VolleyError("Missing the required parameter 'accountId' when calling createAccountExtensionContactGroup",
      new ApiException(400, "Missing the required parameter 'accountId' when calling createAccountExtensionContactGroup"));
      }
  
      // verify the required parameter 'extensionId' is set
      if (extensionId == null) {
      VolleyError error = new VolleyError("Missing the required parameter 'extensionId' when calling createAccountExtensionContactGroup",
      new ApiException(400, "Missing the required parameter 'extensionId' when calling createAccountExtensionContactGroup"));
      }
  
      // verify the required parameter 'data' is set
      if (data == null) {
      VolleyError error = new VolleyError("Missing the required parameter 'data' when calling createAccountExtensionContactGroup",
      new ApiException(400, "Missing the required parameter 'data' when calling createAccountExtensionContactGroup"));
      }
  

  // create path and map variables
  String path = "/accounts/{account_id}/extensions/{extension_id}/contact-groups".replaceAll("\\{format\\}","json").replaceAll("\\{" + "account_id" + "\\}", apiInvoker.escapeString(accountId.toString())).replaceAll("\\{" + "extension_id" + "\\}", apiInvoker.escapeString(extensionId.toString()));

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
           return (GroupFull) ApiInvoker.deserialize(localVarResponse, "", GroupFull.class);
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
   * 
   * See Account Contact Groups for more info on the properties.
   * @param accountId Account ID   * @param extensionId Extension ID   * @param data Group name
  */
  public void createAccountExtensionContactGroup (Integer accountId, Integer extensionId, CreateGroupParams data, final Response.Listener<GroupFull> responseListener, final Response.ErrorListener errorListener) {
    Object postBody = data;

  
    // verify the required parameter 'accountId' is set
    if (accountId == null) {
       VolleyError error = new VolleyError("Missing the required parameter 'accountId' when calling createAccountExtensionContactGroup",
         new ApiException(400, "Missing the required parameter 'accountId' when calling createAccountExtensionContactGroup"));
    }
    
    // verify the required parameter 'extensionId' is set
    if (extensionId == null) {
       VolleyError error = new VolleyError("Missing the required parameter 'extensionId' when calling createAccountExtensionContactGroup",
         new ApiException(400, "Missing the required parameter 'extensionId' when calling createAccountExtensionContactGroup"));
    }
    
    // verify the required parameter 'data' is set
    if (data == null) {
       VolleyError error = new VolleyError("Missing the required parameter 'data' when calling createAccountExtensionContactGroup",
         new ApiException(400, "Missing the required parameter 'data' when calling createAccountExtensionContactGroup"));
    }
    

    // create path and map variables
    String path = "/accounts/{account_id}/extensions/{extension_id}/contact-groups".replaceAll("\\{format\\}","json").replaceAll("\\{" + "account_id" + "\\}", apiInvoker.escapeString(accountId.toString())).replaceAll("\\{" + "extension_id" + "\\}", apiInvoker.escapeString(extensionId.toString()));

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
              responseListener.onResponse((GroupFull) ApiInvoker.deserialize(localVarResponse,  "", GroupFull.class));
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
  * Delete an addressbook group
  * 
   * @param accountId Account ID
   * @param extensionId Extension ID
   * @param groupId Group ID
   * @return DeleteGroup
  */
  public DeleteGroup deleteAccountExtensionContactGroup (Integer accountId, Integer extensionId, Integer groupId) throws TimeoutException, ExecutionException, InterruptedException, ApiException {
     Object postBody = null;
  
      // verify the required parameter 'accountId' is set
      if (accountId == null) {
      VolleyError error = new VolleyError("Missing the required parameter 'accountId' when calling deleteAccountExtensionContactGroup",
      new ApiException(400, "Missing the required parameter 'accountId' when calling deleteAccountExtensionContactGroup"));
      }
  
      // verify the required parameter 'extensionId' is set
      if (extensionId == null) {
      VolleyError error = new VolleyError("Missing the required parameter 'extensionId' when calling deleteAccountExtensionContactGroup",
      new ApiException(400, "Missing the required parameter 'extensionId' when calling deleteAccountExtensionContactGroup"));
      }
  
      // verify the required parameter 'groupId' is set
      if (groupId == null) {
      VolleyError error = new VolleyError("Missing the required parameter 'groupId' when calling deleteAccountExtensionContactGroup",
      new ApiException(400, "Missing the required parameter 'groupId' when calling deleteAccountExtensionContactGroup"));
      }
  

  // create path and map variables
  String path = "/accounts/{account_id}/extensions/{extension_id}/contact-groups/{group_id}".replaceAll("\\{format\\}","json").replaceAll("\\{" + "account_id" + "\\}", apiInvoker.escapeString(accountId.toString())).replaceAll("\\{" + "extension_id" + "\\}", apiInvoker.escapeString(extensionId.toString())).replaceAll("\\{" + "group_id" + "\\}", apiInvoker.escapeString(groupId.toString()));

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
        String localVarResponse = apiInvoker.invokeAPI (basePath, path, "DELETE", queryParams, postBody, headerParams, formParams, contentType, authNames);
        if(localVarResponse != null){
           return (DeleteGroup) ApiInvoker.deserialize(localVarResponse, "", DeleteGroup.class);
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
   * Delete an addressbook group
   * 
   * @param accountId Account ID   * @param extensionId Extension ID   * @param groupId Group ID
  */
  public void deleteAccountExtensionContactGroup (Integer accountId, Integer extensionId, Integer groupId, final Response.Listener<DeleteGroup> responseListener, final Response.ErrorListener errorListener) {
    Object postBody = null;

  
    // verify the required parameter 'accountId' is set
    if (accountId == null) {
       VolleyError error = new VolleyError("Missing the required parameter 'accountId' when calling deleteAccountExtensionContactGroup",
         new ApiException(400, "Missing the required parameter 'accountId' when calling deleteAccountExtensionContactGroup"));
    }
    
    // verify the required parameter 'extensionId' is set
    if (extensionId == null) {
       VolleyError error = new VolleyError("Missing the required parameter 'extensionId' when calling deleteAccountExtensionContactGroup",
         new ApiException(400, "Missing the required parameter 'extensionId' when calling deleteAccountExtensionContactGroup"));
    }
    
    // verify the required parameter 'groupId' is set
    if (groupId == null) {
       VolleyError error = new VolleyError("Missing the required parameter 'groupId' when calling deleteAccountExtensionContactGroup",
         new ApiException(400, "Missing the required parameter 'groupId' when calling deleteAccountExtensionContactGroup"));
    }
    

    // create path and map variables
    String path = "/accounts/{account_id}/extensions/{extension_id}/contact-groups/{group_id}".replaceAll("\\{format\\}","json").replaceAll("\\{" + "account_id" + "\\}", apiInvoker.escapeString(accountId.toString())).replaceAll("\\{" + "extension_id" + "\\}", apiInvoker.escapeString(extensionId.toString())).replaceAll("\\{" + "group_id" + "\\}", apiInvoker.escapeString(groupId.toString()));

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
      apiInvoker.invokeAPI(basePath, path, "DELETE", queryParams, postBody, headerParams, formParams, contentType, authNames,
        new Response.Listener<String>() {
          @Override
          public void onResponse(String localVarResponse) {
            try {
              responseListener.onResponse((DeleteGroup) ApiInvoker.deserialize(localVarResponse,  "", DeleteGroup.class));
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
  * 
  * See Account Contact Groups for more info on the properties.
   * @param accountId Account ID
   * @param extensionId Extension ID
   * @param groupId Group ID
   * @return GroupFull
  */
  public GroupFull getAccountExtensionContactGroup (Integer accountId, Integer extensionId, Integer groupId) throws TimeoutException, ExecutionException, InterruptedException, ApiException {
     Object postBody = null;
  
      // verify the required parameter 'accountId' is set
      if (accountId == null) {
      VolleyError error = new VolleyError("Missing the required parameter 'accountId' when calling getAccountExtensionContactGroup",
      new ApiException(400, "Missing the required parameter 'accountId' when calling getAccountExtensionContactGroup"));
      }
  
      // verify the required parameter 'extensionId' is set
      if (extensionId == null) {
      VolleyError error = new VolleyError("Missing the required parameter 'extensionId' when calling getAccountExtensionContactGroup",
      new ApiException(400, "Missing the required parameter 'extensionId' when calling getAccountExtensionContactGroup"));
      }
  
      // verify the required parameter 'groupId' is set
      if (groupId == null) {
      VolleyError error = new VolleyError("Missing the required parameter 'groupId' when calling getAccountExtensionContactGroup",
      new ApiException(400, "Missing the required parameter 'groupId' when calling getAccountExtensionContactGroup"));
      }
  

  // create path and map variables
  String path = "/accounts/{account_id}/extensions/{extension_id}/contact-groups/{group_id}".replaceAll("\\{format\\}","json").replaceAll("\\{" + "account_id" + "\\}", apiInvoker.escapeString(accountId.toString())).replaceAll("\\{" + "extension_id" + "\\}", apiInvoker.escapeString(extensionId.toString())).replaceAll("\\{" + "group_id" + "\\}", apiInvoker.escapeString(groupId.toString()));

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
           return (GroupFull) ApiInvoker.deserialize(localVarResponse, "", GroupFull.class);
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
   * 
   * See Account Contact Groups for more info on the properties.
   * @param accountId Account ID   * @param extensionId Extension ID   * @param groupId Group ID
  */
  public void getAccountExtensionContactGroup (Integer accountId, Integer extensionId, Integer groupId, final Response.Listener<GroupFull> responseListener, final Response.ErrorListener errorListener) {
    Object postBody = null;

  
    // verify the required parameter 'accountId' is set
    if (accountId == null) {
       VolleyError error = new VolleyError("Missing the required parameter 'accountId' when calling getAccountExtensionContactGroup",
         new ApiException(400, "Missing the required parameter 'accountId' when calling getAccountExtensionContactGroup"));
    }
    
    // verify the required parameter 'extensionId' is set
    if (extensionId == null) {
       VolleyError error = new VolleyError("Missing the required parameter 'extensionId' when calling getAccountExtensionContactGroup",
         new ApiException(400, "Missing the required parameter 'extensionId' when calling getAccountExtensionContactGroup"));
    }
    
    // verify the required parameter 'groupId' is set
    if (groupId == null) {
       VolleyError error = new VolleyError("Missing the required parameter 'groupId' when calling getAccountExtensionContactGroup",
         new ApiException(400, "Missing the required parameter 'groupId' when calling getAccountExtensionContactGroup"));
    }
    

    // create path and map variables
    String path = "/accounts/{account_id}/extensions/{extension_id}/contact-groups/{group_id}".replaceAll("\\{format\\}","json").replaceAll("\\{" + "account_id" + "\\}", apiInvoker.escapeString(accountId.toString())).replaceAll("\\{" + "extension_id" + "\\}", apiInvoker.escapeString(extensionId.toString())).replaceAll("\\{" + "group_id" + "\\}", apiInvoker.escapeString(groupId.toString()));

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
              responseListener.onResponse((GroupFull) ApiInvoker.deserialize(localVarResponse,  "", GroupFull.class));
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
  * Show a list of contact groups belonging to an extension
  * See Account Contact Groups for details on the properties.
   * @param accountId Account ID
   * @param extensionId Extension ID
   * @param filtersId ID filter
   * @param filtersName Name filter
   * @param sortId ID sorting
   * @param sortName Name sorting
   * @param limit Max results
   * @param offset Results to skip
   * @param fields Field set
   * @return ListGroupsFull
  */
  public ListGroupsFull listAccountExtensionContactGroups (Integer accountId, Integer extensionId, List<String> filtersId, List<String> filtersName, String sortId, String sortName, Integer limit, Integer offset, String fields) throws TimeoutException, ExecutionException, InterruptedException, ApiException {
     Object postBody = null;
  
      // verify the required parameter 'accountId' is set
      if (accountId == null) {
      VolleyError error = new VolleyError("Missing the required parameter 'accountId' when calling listAccountExtensionContactGroups",
      new ApiException(400, "Missing the required parameter 'accountId' when calling listAccountExtensionContactGroups"));
      }
  
      // verify the required parameter 'extensionId' is set
      if (extensionId == null) {
      VolleyError error = new VolleyError("Missing the required parameter 'extensionId' when calling listAccountExtensionContactGroups",
      new ApiException(400, "Missing the required parameter 'extensionId' when calling listAccountExtensionContactGroups"));
      }
  

  // create path and map variables
  String path = "/accounts/{account_id}/extensions/{extension_id}/contact-groups".replaceAll("\\{format\\}","json").replaceAll("\\{" + "account_id" + "\\}", apiInvoker.escapeString(accountId.toString())).replaceAll("\\{" + "extension_id" + "\\}", apiInvoker.escapeString(extensionId.toString()));

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
           return (ListGroupsFull) ApiInvoker.deserialize(localVarResponse, "", ListGroupsFull.class);
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
   * Show a list of contact groups belonging to an extension
   * See Account Contact Groups for details on the properties.
   * @param accountId Account ID   * @param extensionId Extension ID   * @param filtersId ID filter   * @param filtersName Name filter   * @param sortId ID sorting   * @param sortName Name sorting   * @param limit Max results   * @param offset Results to skip   * @param fields Field set
  */
  public void listAccountExtensionContactGroups (Integer accountId, Integer extensionId, List<String> filtersId, List<String> filtersName, String sortId, String sortName, Integer limit, Integer offset, String fields, final Response.Listener<ListGroupsFull> responseListener, final Response.ErrorListener errorListener) {
    Object postBody = null;

  
    // verify the required parameter 'accountId' is set
    if (accountId == null) {
       VolleyError error = new VolleyError("Missing the required parameter 'accountId' when calling listAccountExtensionContactGroups",
         new ApiException(400, "Missing the required parameter 'accountId' when calling listAccountExtensionContactGroups"));
    }
    
    // verify the required parameter 'extensionId' is set
    if (extensionId == null) {
       VolleyError error = new VolleyError("Missing the required parameter 'extensionId' when calling listAccountExtensionContactGroups",
         new ApiException(400, "Missing the required parameter 'extensionId' when calling listAccountExtensionContactGroups"));
    }
    

    // create path and map variables
    String path = "/accounts/{account_id}/extensions/{extension_id}/contact-groups".replaceAll("\\{format\\}","json").replaceAll("\\{" + "account_id" + "\\}", apiInvoker.escapeString(accountId.toString())).replaceAll("\\{" + "extension_id" + "\\}", apiInvoker.escapeString(extensionId.toString()));

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
              responseListener.onResponse((ListGroupsFull) ApiInvoker.deserialize(localVarResponse,  "", ListGroupsFull.class));
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
  * 
  * See Account Contact Groups for more info on the properties.
   * @param accountId Account ID
   * @param extensionId Extension ID
   * @param groupId Group ID
   * @return GroupFull
  */
  public GroupFull replaceAccountExtensionContactGroup (Integer accountId, Integer extensionId, Integer groupId) throws TimeoutException, ExecutionException, InterruptedException, ApiException {
     Object postBody = null;
  
      // verify the required parameter 'accountId' is set
      if (accountId == null) {
      VolleyError error = new VolleyError("Missing the required parameter 'accountId' when calling replaceAccountExtensionContactGroup",
      new ApiException(400, "Missing the required parameter 'accountId' when calling replaceAccountExtensionContactGroup"));
      }
  
      // verify the required parameter 'extensionId' is set
      if (extensionId == null) {
      VolleyError error = new VolleyError("Missing the required parameter 'extensionId' when calling replaceAccountExtensionContactGroup",
      new ApiException(400, "Missing the required parameter 'extensionId' when calling replaceAccountExtensionContactGroup"));
      }
  
      // verify the required parameter 'groupId' is set
      if (groupId == null) {
      VolleyError error = new VolleyError("Missing the required parameter 'groupId' when calling replaceAccountExtensionContactGroup",
      new ApiException(400, "Missing the required parameter 'groupId' when calling replaceAccountExtensionContactGroup"));
      }
  

  // create path and map variables
  String path = "/accounts/{account_id}/extensions/{extension_id}/contact-groups/{group_id}".replaceAll("\\{format\\}","json").replaceAll("\\{" + "account_id" + "\\}", apiInvoker.escapeString(accountId.toString())).replaceAll("\\{" + "extension_id" + "\\}", apiInvoker.escapeString(extensionId.toString())).replaceAll("\\{" + "group_id" + "\\}", apiInvoker.escapeString(groupId.toString()));

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
           return (GroupFull) ApiInvoker.deserialize(localVarResponse, "", GroupFull.class);
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
   * 
   * See Account Contact Groups for more info on the properties.
   * @param accountId Account ID   * @param extensionId Extension ID   * @param groupId Group ID
  */
  public void replaceAccountExtensionContactGroup (Integer accountId, Integer extensionId, Integer groupId, final Response.Listener<GroupFull> responseListener, final Response.ErrorListener errorListener) {
    Object postBody = null;

  
    // verify the required parameter 'accountId' is set
    if (accountId == null) {
       VolleyError error = new VolleyError("Missing the required parameter 'accountId' when calling replaceAccountExtensionContactGroup",
         new ApiException(400, "Missing the required parameter 'accountId' when calling replaceAccountExtensionContactGroup"));
    }
    
    // verify the required parameter 'extensionId' is set
    if (extensionId == null) {
       VolleyError error = new VolleyError("Missing the required parameter 'extensionId' when calling replaceAccountExtensionContactGroup",
         new ApiException(400, "Missing the required parameter 'extensionId' when calling replaceAccountExtensionContactGroup"));
    }
    
    // verify the required parameter 'groupId' is set
    if (groupId == null) {
       VolleyError error = new VolleyError("Missing the required parameter 'groupId' when calling replaceAccountExtensionContactGroup",
         new ApiException(400, "Missing the required parameter 'groupId' when calling replaceAccountExtensionContactGroup"));
    }
    

    // create path and map variables
    String path = "/accounts/{account_id}/extensions/{extension_id}/contact-groups/{group_id}".replaceAll("\\{format\\}","json").replaceAll("\\{" + "account_id" + "\\}", apiInvoker.escapeString(accountId.toString())).replaceAll("\\{" + "extension_id" + "\\}", apiInvoker.escapeString(extensionId.toString())).replaceAll("\\{" + "group_id" + "\\}", apiInvoker.escapeString(groupId.toString()));

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
              responseListener.onResponse((GroupFull) ApiInvoker.deserialize(localVarResponse,  "", GroupFull.class));
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
