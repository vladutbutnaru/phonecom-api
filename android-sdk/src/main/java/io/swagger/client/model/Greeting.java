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

import io.swagger.client.model.MediaSummary;

import io.swagger.annotations.*;
import com.google.gson.annotations.SerializedName;


/**
 * Voicemail Greeting Object. See below for details. Can be set to NULL to reset greeting options.
 **/
@ApiModel(description = "Voicemail Greeting Object. See below for details. Can be set to NULL to reset greeting options.")
public class Greeting  {
  
  @SerializedName("type")
  private String type = null;
  @SerializedName("alternate")
  private MediaSummary alternate = null;
  @SerializedName("standard")
  private MediaSummary standard = null;
  @SerializedName("enable_leave_message_prompt")
  private Boolean enableLeaveMessagePrompt = null;

  /**
   * The greeting to play. Can be \"name\" for the name_greeting as described above, \"standard\" for the standard greeting, or \"alternate\" for an alternate greeting. See below for details.
   **/
  @ApiModelProperty(value = "The greeting to play. Can be \"name\" for the name_greeting as described above, \"standard\" for the standard greeting, or \"alternate\" for an alternate greeting. See below for details.")
  public String getType() {
    return type;
  }
  public void setType(String type) {
    this.type = type;
  }

  /**
   * Greeting to be played when type=\"alternate\". Output is a Greeting Summary Object. Input must be a Greeting Lookup Object.
   **/
  @ApiModelProperty(value = "Greeting to be played when type=\"alternate\". Output is a Greeting Summary Object. Input must be a Greeting Lookup Object.")
  public MediaSummary getAlternate() {
    return alternate;
  }
  public void setAlternate(MediaSummary alternate) {
    this.alternate = alternate;
  }

  /**
   * Greeting to be played when type=\"standard\". Output is a Greeting Summary Object. Input must be a Greeting Lookup Object.
   **/
  @ApiModelProperty(value = "Greeting to be played when type=\"standard\". Output is a Greeting Summary Object. Input must be a Greeting Lookup Object.")
  public MediaSummary getStandard() {
    return standard;
  }
  public void setStandard(MediaSummary standard) {
    this.standard = standard;
  }

  /**
   * Whether to prompt the caller with the following words after the voicemail greeting has been played: \"Please leave your message after the tone. When finished, hang up or press the pound key.\" Boolean.
   **/
  @ApiModelProperty(value = "Whether to prompt the caller with the following words after the voicemail greeting has been played: \"Please leave your message after the tone. When finished, hang up or press the pound key.\" Boolean.")
  public Boolean getEnableLeaveMessagePrompt() {
    return enableLeaveMessagePrompt;
  }
  public void setEnableLeaveMessagePrompt(Boolean enableLeaveMessagePrompt) {
    this.enableLeaveMessagePrompt = enableLeaveMessagePrompt;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    Greeting greeting = (Greeting) o;
    return (type == null ? greeting.type == null : type.equals(greeting.type)) &&
        (alternate == null ? greeting.alternate == null : alternate.equals(greeting.alternate)) &&
        (standard == null ? greeting.standard == null : standard.equals(greeting.standard)) &&
        (enableLeaveMessagePrompt == null ? greeting.enableLeaveMessagePrompt == null : enableLeaveMessagePrompt.equals(greeting.enableLeaveMessagePrompt));
  }

  @Override
  public int hashCode() {
    int result = 17;
    result = 31 * result + (type == null ? 0: type.hashCode());
    result = 31 * result + (alternate == null ? 0: alternate.hashCode());
    result = 31 * result + (standard == null ? 0: standard.hashCode());
    result = 31 * result + (enableLeaveMessagePrompt == null ? 0: enableLeaveMessagePrompt.hashCode());
    return result;
  }

  @Override
  public String toString()  {
    StringBuilder sb = new StringBuilder();
    sb.append("class Greeting {\n");
    
    sb.append("  type: ").append(type).append("\n");
    sb.append("  alternate: ").append(alternate).append("\n");
    sb.append("  standard: ").append(standard).append("\n");
    sb.append("  enableLeaveMessagePrompt: ").append(enableLeaveMessagePrompt).append("\n");
    sb.append("}\n");
    return sb.toString();
  }
}