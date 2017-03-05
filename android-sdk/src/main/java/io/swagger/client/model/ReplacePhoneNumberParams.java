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

package io.swagger.client.model;

import java.util.*;

import io.swagger.annotations.*;
import com.google.gson.annotations.SerializedName;


@ApiModel(description = "")
public class ReplacePhoneNumberParams  {
  
  @SerializedName("route")
  private Object route = null;
  @SerializedName("name")
  private String name = null;
  @SerializedName("block_incoming")
  private Boolean blockIncoming = null;
  @SerializedName("block_anonymous")
  private Boolean blockAnonymous = null;
  @SerializedName("caller_id[name]")
  private String callerIdName = null;
  @SerializedName("caller_id[type]")
  private String callerIdType = null;
  @SerializedName("sms_forwarding[type]")
  private String smsForwardingType = null;
  @SerializedName("sms_forwarding[application]")
  private Object smsForwardingApplication = null;
  @SerializedName("sms_forwarding[extension]")
  private Object smsForwardingExtension = null;
  @SerializedName("pool_item")
  private Object poolItem = null;
  @SerializedName("call_notifications[emails]")
  private List<String> callNotificationsEmails = null;
  @SerializedName("call_notifications[sms]")
  private String callNotificationsSms = null;

  /**
   * Route lookup object
   **/
  @ApiModelProperty(value = "Route lookup object")
  public Object getRoute() {
    return route;
  }
  public void setRoute(Object route) {
    this.route = route;
  }

  /**
   * Phone Name
   **/
  @ApiModelProperty(value = "Phone Name")
  public String getName() {
    return name;
  }
  public void setName(String name) {
    this.name = name;
  }

  /**
   * Block incoming calls
   **/
  @ApiModelProperty(value = "Block incoming calls")
  public Boolean getBlockIncoming() {
    return blockIncoming;
  }
  public void setBlockIncoming(Boolean blockIncoming) {
    this.blockIncoming = blockIncoming;
  }

  /**
   * Block anonymous calls
   **/
  @ApiModelProperty(value = "Block anonymous calls")
  public Boolean getBlockAnonymous() {
    return blockAnonymous;
  }
  public void setBlockAnonymous(Boolean blockAnonymous) {
    this.blockAnonymous = blockAnonymous;
  }

  /**
   * Caller ID name
   **/
  @ApiModelProperty(value = "Caller ID name")
  public String getCallerIdName() {
    return callerIdName;
  }
  public void setCallerIdName(String callerIdName) {
    this.callerIdName = callerIdName;
  }

  /**
   * Caller ID type
   **/
  @ApiModelProperty(value = "Caller ID type")
  public String getCallerIdType() {
    return callerIdType;
  }
  public void setCallerIdType(String callerIdType) {
    this.callerIdType = callerIdType;
  }

  /**
   * 'application' or 'extension'
   **/
  @ApiModelProperty(value = "'application' or 'extension'")
  public String getSmsForwardingType() {
    return smsForwardingType;
  }
  public void setSmsForwardingType(String smsForwardingType) {
    this.smsForwardingType = smsForwardingType;
  }

  /**
   * Application lookup object
   **/
  @ApiModelProperty(value = "Application lookup object")
  public Object getSmsForwardingApplication() {
    return smsForwardingApplication;
  }
  public void setSmsForwardingApplication(Object smsForwardingApplication) {
    this.smsForwardingApplication = smsForwardingApplication;
  }

  /**
   * Extension lookup object
   **/
  @ApiModelProperty(value = "Extension lookup object")
  public Object getSmsForwardingExtension() {
    return smsForwardingExtension;
  }
  public void setSmsForwardingExtension(Object smsForwardingExtension) {
    this.smsForwardingExtension = smsForwardingExtension;
  }

  /**
   * Pool lookup object
   **/
  @ApiModelProperty(value = "Pool lookup object")
  public Object getPoolItem() {
    return poolItem;
  }
  public void setPoolItem(Object poolItem) {
    this.poolItem = poolItem;
  }

  /**
   * Call notifications for emails. Can be a single email or an array of emails
   **/
  @ApiModelProperty(value = "Call notifications for emails. Can be a single email or an array of emails")
  public List<String> getCallNotificationsEmails() {
    return callNotificationsEmails;
  }
  public void setCallNotificationsEmails(List<String> callNotificationsEmails) {
    this.callNotificationsEmails = callNotificationsEmails;
  }

  /**
   * Call notification for SMS
   **/
  @ApiModelProperty(value = "Call notification for SMS")
  public String getCallNotificationsSms() {
    return callNotificationsSms;
  }
  public void setCallNotificationsSms(String callNotificationsSms) {
    this.callNotificationsSms = callNotificationsSms;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    ReplacePhoneNumberParams replacePhoneNumberParams = (ReplacePhoneNumberParams) o;
    return (route == null ? replacePhoneNumberParams.route == null : route.equals(replacePhoneNumberParams.route)) &&
        (name == null ? replacePhoneNumberParams.name == null : name.equals(replacePhoneNumberParams.name)) &&
        (blockIncoming == null ? replacePhoneNumberParams.blockIncoming == null : blockIncoming.equals(replacePhoneNumberParams.blockIncoming)) &&
        (blockAnonymous == null ? replacePhoneNumberParams.blockAnonymous == null : blockAnonymous.equals(replacePhoneNumberParams.blockAnonymous)) &&
        (callerIdName == null ? replacePhoneNumberParams.callerIdName == null : callerIdName.equals(replacePhoneNumberParams.callerIdName)) &&
        (callerIdType == null ? replacePhoneNumberParams.callerIdType == null : callerIdType.equals(replacePhoneNumberParams.callerIdType)) &&
        (smsForwardingType == null ? replacePhoneNumberParams.smsForwardingType == null : smsForwardingType.equals(replacePhoneNumberParams.smsForwardingType)) &&
        (smsForwardingApplication == null ? replacePhoneNumberParams.smsForwardingApplication == null : smsForwardingApplication.equals(replacePhoneNumberParams.smsForwardingApplication)) &&
        (smsForwardingExtension == null ? replacePhoneNumberParams.smsForwardingExtension == null : smsForwardingExtension.equals(replacePhoneNumberParams.smsForwardingExtension)) &&
        (poolItem == null ? replacePhoneNumberParams.poolItem == null : poolItem.equals(replacePhoneNumberParams.poolItem)) &&
        (callNotificationsEmails == null ? replacePhoneNumberParams.callNotificationsEmails == null : callNotificationsEmails.equals(replacePhoneNumberParams.callNotificationsEmails)) &&
        (callNotificationsSms == null ? replacePhoneNumberParams.callNotificationsSms == null : callNotificationsSms.equals(replacePhoneNumberParams.callNotificationsSms));
  }

  @Override
  public int hashCode() {
    int result = 17;
    result = 31 * result + (route == null ? 0: route.hashCode());
    result = 31 * result + (name == null ? 0: name.hashCode());
    result = 31 * result + (blockIncoming == null ? 0: blockIncoming.hashCode());
    result = 31 * result + (blockAnonymous == null ? 0: blockAnonymous.hashCode());
    result = 31 * result + (callerIdName == null ? 0: callerIdName.hashCode());
    result = 31 * result + (callerIdType == null ? 0: callerIdType.hashCode());
    result = 31 * result + (smsForwardingType == null ? 0: smsForwardingType.hashCode());
    result = 31 * result + (smsForwardingApplication == null ? 0: smsForwardingApplication.hashCode());
    result = 31 * result + (smsForwardingExtension == null ? 0: smsForwardingExtension.hashCode());
    result = 31 * result + (poolItem == null ? 0: poolItem.hashCode());
    result = 31 * result + (callNotificationsEmails == null ? 0: callNotificationsEmails.hashCode());
    result = 31 * result + (callNotificationsSms == null ? 0: callNotificationsSms.hashCode());
    return result;
  }

  @Override
  public String toString()  {
    StringBuilder sb = new StringBuilder();
    sb.append("class ReplacePhoneNumberParams {\n");
    
    sb.append("  route: ").append(route).append("\n");
    sb.append("  name: ").append(name).append("\n");
    sb.append("  blockIncoming: ").append(blockIncoming).append("\n");
    sb.append("  blockAnonymous: ").append(blockAnonymous).append("\n");
    sb.append("  callerIdName: ").append(callerIdName).append("\n");
    sb.append("  callerIdType: ").append(callerIdType).append("\n");
    sb.append("  smsForwardingType: ").append(smsForwardingType).append("\n");
    sb.append("  smsForwardingApplication: ").append(smsForwardingApplication).append("\n");
    sb.append("  smsForwardingExtension: ").append(smsForwardingExtension).append("\n");
    sb.append("  poolItem: ").append(poolItem).append("\n");
    sb.append("  callNotificationsEmails: ").append(callNotificationsEmails).append("\n");
    sb.append("  callNotificationsSms: ").append(callNotificationsSms).append("\n");
    sb.append("}\n");
    return sb.toString();
  }
}