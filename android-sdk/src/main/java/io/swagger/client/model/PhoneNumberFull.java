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

import io.swagger.client.model.CallNotifications;
import io.swagger.client.model.CallerIdPhoneNumber;
import io.swagger.client.model.RouteSummary;
import io.swagger.client.model.SmsForwarding;

import io.swagger.annotations.*;
import com.google.gson.annotations.SerializedName;


/**
 * The Full Phone Number Object has all of the properties of the Phone Number Summary Object, along with several more, as shown below:
 **/
@ApiModel(description = "The Full Phone Number Object has all of the properties of the Phone Number Summary Object, along with several more, as shown below:")
public class PhoneNumberFull  {
  
  @SerializedName("id")
  private Integer id = null;
  @SerializedName("name")
  private String name = null;
  @SerializedName("phone_number")
  private String phoneNumber = null;
  @SerializedName("block_incoming")
  private Boolean blockIncoming = null;
  @SerializedName("block_anonymous")
  private Boolean blockAnonymous = null;
  @SerializedName("route")
  private RouteSummary route = null;
  @SerializedName("caller_id")
  private CallerIdPhoneNumber callerId = null;
  @SerializedName("sms_forwarding")
  private SmsForwarding smsForwarding = null;
  @SerializedName("call_notifications")
  private CallNotifications callNotifications = null;

  /**
   * Integer Phone number ID. This is the internal Phone.com ID for this number, not the phone number itself. Read-only.
   **/
  @ApiModelProperty(value = "Integer Phone number ID. This is the internal Phone.com ID for this number, not the phone number itself. Read-only.")
  public Integer getId() {
    return id;
  }
  public void setId(Integer id) {
    this.id = id;
  }

  /**
   * Name
   **/
  @ApiModelProperty(value = "Name")
  public String getName() {
    return name;
  }
  public void setName(String name) {
    this.name = name;
  }

  /**
   * Phone number, in E.164 format
   **/
  @ApiModelProperty(value = "Phone number, in E.164 format")
  public String getPhoneNumber() {
    return phoneNumber;
  }
  public void setPhoneNumber(String phoneNumber) {
    this.phoneNumber = phoneNumber;
  }

  /**
   * Whether to block incoming calls. Boolean.
   **/
  @ApiModelProperty(value = "Whether to block incoming calls. Boolean.")
  public Boolean getBlockIncoming() {
    return blockIncoming;
  }
  public void setBlockIncoming(Boolean blockIncoming) {
    this.blockIncoming = blockIncoming;
  }

  /**
   * Whether to block anonymous calls. Boolean.
   **/
  @ApiModelProperty(value = "Whether to block anonymous calls. Boolean.")
  public Boolean getBlockAnonymous() {
    return blockAnonymous;
  }
  public void setBlockAnonymous(Boolean blockAnonymous) {
    this.blockAnonymous = blockAnonymous;
  }

  /**
   * The Route assigned to handle incoming calls for this number, if any. Output is a Route Summary Object, or NULL if not set. Input can be a Route Lookup Object or NULL to unset.
   **/
  @ApiModelProperty(value = "The Route assigned to handle incoming calls for this number, if any. Output is a Route Summary Object, or NULL if not set. Input can be a Route Lookup Object or NULL to unset.")
  public RouteSummary getRoute() {
    return route;
  }
  public void setRoute(RouteSummary route) {
    this.route = route;
  }

  /**
   * Caller ID Object, or NULL
   **/
  @ApiModelProperty(value = "Caller ID Object, or NULL")
  public CallerIdPhoneNumber getCallerId() {
    return callerId;
  }
  public void setCallerId(CallerIdPhoneNumber callerId) {
    this.callerId = callerId;
  }

  /**
   * SMS Forwarding Object, or NULL
   **/
  @ApiModelProperty(value = "SMS Forwarding Object, or NULL")
  public SmsForwarding getSmsForwarding() {
    return smsForwarding;
  }
  public void setSmsForwarding(SmsForwarding smsForwarding) {
    this.smsForwarding = smsForwarding;
  }

  /**
   **/
  @ApiModelProperty(value = "")
  public CallNotifications getCallNotifications() {
    return callNotifications;
  }
  public void setCallNotifications(CallNotifications callNotifications) {
    this.callNotifications = callNotifications;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    PhoneNumberFull phoneNumberFull = (PhoneNumberFull) o;
    return (id == null ? phoneNumberFull.id == null : id.equals(phoneNumberFull.id)) &&
        (name == null ? phoneNumberFull.name == null : name.equals(phoneNumberFull.name)) &&
        (phoneNumber == null ? phoneNumberFull.phoneNumber == null : phoneNumber.equals(phoneNumberFull.phoneNumber)) &&
        (blockIncoming == null ? phoneNumberFull.blockIncoming == null : blockIncoming.equals(phoneNumberFull.blockIncoming)) &&
        (blockAnonymous == null ? phoneNumberFull.blockAnonymous == null : blockAnonymous.equals(phoneNumberFull.blockAnonymous)) &&
        (route == null ? phoneNumberFull.route == null : route.equals(phoneNumberFull.route)) &&
        (callerId == null ? phoneNumberFull.callerId == null : callerId.equals(phoneNumberFull.callerId)) &&
        (smsForwarding == null ? phoneNumberFull.smsForwarding == null : smsForwarding.equals(phoneNumberFull.smsForwarding)) &&
        (callNotifications == null ? phoneNumberFull.callNotifications == null : callNotifications.equals(phoneNumberFull.callNotifications));
  }

  @Override
  public int hashCode() {
    int result = 17;
    result = 31 * result + (id == null ? 0: id.hashCode());
    result = 31 * result + (name == null ? 0: name.hashCode());
    result = 31 * result + (phoneNumber == null ? 0: phoneNumber.hashCode());
    result = 31 * result + (blockIncoming == null ? 0: blockIncoming.hashCode());
    result = 31 * result + (blockAnonymous == null ? 0: blockAnonymous.hashCode());
    result = 31 * result + (route == null ? 0: route.hashCode());
    result = 31 * result + (callerId == null ? 0: callerId.hashCode());
    result = 31 * result + (smsForwarding == null ? 0: smsForwarding.hashCode());
    result = 31 * result + (callNotifications == null ? 0: callNotifications.hashCode());
    return result;
  }

  @Override
  public String toString()  {
    StringBuilder sb = new StringBuilder();
    sb.append("class PhoneNumberFull {\n");
    
    sb.append("  id: ").append(id).append("\n");
    sb.append("  name: ").append(name).append("\n");
    sb.append("  phoneNumber: ").append(phoneNumber).append("\n");
    sb.append("  blockIncoming: ").append(blockIncoming).append("\n");
    sb.append("  blockAnonymous: ").append(blockAnonymous).append("\n");
    sb.append("  route: ").append(route).append("\n");
    sb.append("  callerId: ").append(callerId).append("\n");
    sb.append("  smsForwarding: ").append(smsForwarding).append("\n");
    sb.append("  callNotifications: ").append(callNotifications).append("\n");
    sb.append("}\n");
    return sb.toString();
  }
}