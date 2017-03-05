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

import io.swagger.client.model.CallDetails;
import io.swagger.client.model.ExtensionSummary;
import java.util.*;

import io.swagger.annotations.*;
import com.google.gson.annotations.SerializedName;


/**
 * The Full Call Log Object includes the properties in the Call Log Summary Object, along with the following:
 **/
@ApiModel(description = "The Full Call Log Object includes the properties in the Call Log Summary Object, along with the following:")
public class CallLogFull  {
  
  @SerializedName("id")
  private String id = null;
  @SerializedName("uuid")
  private String uuid = null;
  @SerializedName("extension")
  private ExtensionSummary extension = null;
  @SerializedName("caller_id")
  private String callerId = null;
  @SerializedName("called_number")
  private String calledNumber = null;
  @SerializedName("start_time")
  private String startTime = null;
  @SerializedName("created_at")
  private String createdAt = null;
  @SerializedName("direction")
  private String direction = null;
  @SerializedName("type")
  private String type = null;
  @SerializedName("call_duration")
  private Integer callDuration = null;
  @SerializedName("is_monitored")
  private String isMonitored = null;
  @SerializedName("call_number")
  private String callNumber = null;
  @SerializedName("final_action")
  private String finalAction = null;
  @SerializedName("call_recording")
  private String callRecording = null;
  @SerializedName("details")
  private List<CallDetails> details = null;
  @SerializedName("caller_cnam")
  private String callerCnam = null;

  /**
   * ID
   **/
  @ApiModelProperty(value = "ID")
  public String getId() {
    return id;
  }
  public void setId(String id) {
    this.id = id;
  }

  /**
   * Internal system id, may be null
   **/
  @ApiModelProperty(value = "Internal system id, may be null")
  public String getUuid() {
    return uuid;
  }
  public void setUuid(String uuid) {
    this.uuid = uuid;
  }

  /**
   * Account extension
   **/
  @ApiModelProperty(value = "Account extension")
  public ExtensionSummary getExtension() {
    return extension;
  }
  public void setExtension(ExtensionSummary extension) {
    this.extension = extension;
  }

  /**
   * Call made from this phone number
   **/
  @ApiModelProperty(value = "Call made from this phone number")
  public String getCallerId() {
    return callerId;
  }
  public void setCallerId(String callerId) {
    this.callerId = callerId;
  }

  /**
   * Call made to this phone number
   **/
  @ApiModelProperty(value = "Call made to this phone number")
  public String getCalledNumber() {
    return calledNumber;
  }
  public void setCalledNumber(String calledNumber) {
    this.calledNumber = calledNumber;
  }

  /**
   * Call start time
   **/
  @ApiModelProperty(value = "Call start time")
  public String getStartTime() {
    return startTime;
  }
  public void setStartTime(String startTime) {
    this.startTime = startTime;
  }

  /**
   * Call log creation time. Same as call end time + time needed to create call log
   **/
  @ApiModelProperty(value = "Call log creation time. Same as call end time + time needed to create call log")
  public String getCreatedAt() {
    return createdAt;
  }
  public void setCreatedAt(String createdAt) {
    this.createdAt = createdAt;
  }

  /**
   * Call direction: in or out
   **/
  @ApiModelProperty(value = "Call direction: in or out")
  public String getDirection() {
    return direction;
  }
  public void setDirection(String direction) {
    this.direction = direction;
  }

  /**
   * Call type: call, fax, audiogram ...
   **/
  @ApiModelProperty(value = "Call type: call, fax, audiogram ...")
  public String getType() {
    return type;
  }
  public void setType(String type) {
    this.type = type;
  }

  /**
   * Call duration in seconds
   **/
  @ApiModelProperty(value = "Call duration in seconds")
  public Integer getCallDuration() {
    return callDuration;
  }
  public void setCallDuration(Integer callDuration) {
    this.callDuration = callDuration;
  }

  /**
   * Was call being monitored?
   **/
  @ApiModelProperty(value = "Was call being monitored?")
  public String getIsMonitored() {
    return isMonitored;
  }
  public void setIsMonitored(String isMonitored) {
    this.isMonitored = isMonitored;
  }

  /**
   * Internal system call reference number
   **/
  @ApiModelProperty(value = "Internal system call reference number")
  public String getCallNumber() {
    return callNumber;
  }
  public void setCallNumber(String callNumber) {
    this.callNumber = callNumber;
  }

  /**
   * Last action of call flow
   **/
  @ApiModelProperty(value = "Last action of call flow")
  public String getFinalAction() {
    return finalAction;
  }
  public void setFinalAction(String finalAction) {
    this.finalAction = finalAction;
  }

  /**
   * URL of call recording if available. Empty string if call recording does not exist
   **/
  @ApiModelProperty(value = "URL of call recording if available. Empty string if call recording does not exist")
  public String getCallRecording() {
    return callRecording;
  }
  public void setCallRecording(String callRecording) {
    this.callRecording = callRecording;
  }

  /**
   * A list of call flows from beginning of call to end of call.
   **/
  @ApiModelProperty(value = "A list of call flows from beginning of call to end of call.")
  public List<CallDetails> getDetails() {
    return details;
  }
  public void setDetails(List<CallDetails> details) {
    this.details = details;
  }

  /**
   * Internal system caller id / name
   **/
  @ApiModelProperty(value = "Internal system caller id / name")
  public String getCallerCnam() {
    return callerCnam;
  }
  public void setCallerCnam(String callerCnam) {
    this.callerCnam = callerCnam;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    CallLogFull callLogFull = (CallLogFull) o;
    return (id == null ? callLogFull.id == null : id.equals(callLogFull.id)) &&
        (uuid == null ? callLogFull.uuid == null : uuid.equals(callLogFull.uuid)) &&
        (extension == null ? callLogFull.extension == null : extension.equals(callLogFull.extension)) &&
        (callerId == null ? callLogFull.callerId == null : callerId.equals(callLogFull.callerId)) &&
        (calledNumber == null ? callLogFull.calledNumber == null : calledNumber.equals(callLogFull.calledNumber)) &&
        (startTime == null ? callLogFull.startTime == null : startTime.equals(callLogFull.startTime)) &&
        (createdAt == null ? callLogFull.createdAt == null : createdAt.equals(callLogFull.createdAt)) &&
        (direction == null ? callLogFull.direction == null : direction.equals(callLogFull.direction)) &&
        (type == null ? callLogFull.type == null : type.equals(callLogFull.type)) &&
        (callDuration == null ? callLogFull.callDuration == null : callDuration.equals(callLogFull.callDuration)) &&
        (isMonitored == null ? callLogFull.isMonitored == null : isMonitored.equals(callLogFull.isMonitored)) &&
        (callNumber == null ? callLogFull.callNumber == null : callNumber.equals(callLogFull.callNumber)) &&
        (finalAction == null ? callLogFull.finalAction == null : finalAction.equals(callLogFull.finalAction)) &&
        (callRecording == null ? callLogFull.callRecording == null : callRecording.equals(callLogFull.callRecording)) &&
        (details == null ? callLogFull.details == null : details.equals(callLogFull.details)) &&
        (callerCnam == null ? callLogFull.callerCnam == null : callerCnam.equals(callLogFull.callerCnam));
  }

  @Override
  public int hashCode() {
    int result = 17;
    result = 31 * result + (id == null ? 0: id.hashCode());
    result = 31 * result + (uuid == null ? 0: uuid.hashCode());
    result = 31 * result + (extension == null ? 0: extension.hashCode());
    result = 31 * result + (callerId == null ? 0: callerId.hashCode());
    result = 31 * result + (calledNumber == null ? 0: calledNumber.hashCode());
    result = 31 * result + (startTime == null ? 0: startTime.hashCode());
    result = 31 * result + (createdAt == null ? 0: createdAt.hashCode());
    result = 31 * result + (direction == null ? 0: direction.hashCode());
    result = 31 * result + (type == null ? 0: type.hashCode());
    result = 31 * result + (callDuration == null ? 0: callDuration.hashCode());
    result = 31 * result + (isMonitored == null ? 0: isMonitored.hashCode());
    result = 31 * result + (callNumber == null ? 0: callNumber.hashCode());
    result = 31 * result + (finalAction == null ? 0: finalAction.hashCode());
    result = 31 * result + (callRecording == null ? 0: callRecording.hashCode());
    result = 31 * result + (details == null ? 0: details.hashCode());
    result = 31 * result + (callerCnam == null ? 0: callerCnam.hashCode());
    return result;
  }

  @Override
  public String toString()  {
    StringBuilder sb = new StringBuilder();
    sb.append("class CallLogFull {\n");
    
    sb.append("  id: ").append(id).append("\n");
    sb.append("  uuid: ").append(uuid).append("\n");
    sb.append("  extension: ").append(extension).append("\n");
    sb.append("  callerId: ").append(callerId).append("\n");
    sb.append("  calledNumber: ").append(calledNumber).append("\n");
    sb.append("  startTime: ").append(startTime).append("\n");
    sb.append("  createdAt: ").append(createdAt).append("\n");
    sb.append("  direction: ").append(direction).append("\n");
    sb.append("  type: ").append(type).append("\n");
    sb.append("  callDuration: ").append(callDuration).append("\n");
    sb.append("  isMonitored: ").append(isMonitored).append("\n");
    sb.append("  callNumber: ").append(callNumber).append("\n");
    sb.append("  finalAction: ").append(finalAction).append("\n");
    sb.append("  callRecording: ").append(callRecording).append("\n");
    sb.append("  details: ").append(details).append("\n");
    sb.append("  callerCnam: ").append(callerCnam).append("\n");
    sb.append("}\n");
    return sb.toString();
  }
}