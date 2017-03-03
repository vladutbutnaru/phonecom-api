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
import java.util.*;

import io.swagger.annotations.*;
import com.google.gson.annotations.SerializedName;


/**
 * The Full Trunk Object is identical to the Trunk Summary Object, along with the following:
 **/
@ApiModel(description = "The Full Trunk Object is identical to the Trunk Summary Object, along with the following:")
public class TrunkFull  {
  
  @SerializedName("id")
  private Integer id = null;
  @SerializedName("name")
  private String name = null;
  @SerializedName("uri")
  private String uri = null;
  @SerializedName("max_concurrent_calls")
  private Integer maxConcurrentCalls = null;
  @SerializedName("max_minutes_per_month")
  private Integer maxMinutesPerMonth = null;
  @SerializedName("greeting")
  private MediaSummary greeting = null;
  @SerializedName("error_message")
  private MediaSummary errorMessage = null;
  @SerializedName("codecs")
  private List<String> codecs = null;

  /**
   * Integer Trunk ID. Read-only.
   **/
  @ApiModelProperty(required = true, value = "Integer Trunk ID. Read-only.")
  public Integer getId() {
    return id;
  }
  public void setId(Integer id) {
    this.id = id;
  }

  /**
   * Name. Required.
   **/
  @ApiModelProperty(required = true, value = "Name. Required.")
  public String getName() {
    return name;
  }
  public void setName(String name) {
    this.name = name;
  }

  /**
   * Fully-qualified SIP URI. Required.
   **/
  @ApiModelProperty(required = true, value = "Fully-qualified SIP URI. Required.")
  public String getUri() {
    return uri;
  }
  public void setUri(String uri) {
    this.uri = uri;
  }

  /**
   * Max concurrent calls. Default is 10.
   **/
  @ApiModelProperty(required = true, value = "Max concurrent calls. Default is 10.")
  public Integer getMaxConcurrentCalls() {
    return maxConcurrentCalls;
  }
  public void setMaxConcurrentCalls(Integer maxConcurrentCalls) {
    this.maxConcurrentCalls = maxConcurrentCalls;
  }

  /**
   * Max minutes per month. Default is 750.
   **/
  @ApiModelProperty(required = true, value = "Max minutes per month. Default is 750.")
  public Integer getMaxMinutesPerMonth() {
    return maxMinutesPerMonth;
  }
  public void setMaxMinutesPerMonth(Integer maxMinutesPerMonth) {
    this.maxMinutesPerMonth = maxMinutesPerMonth;
  }

  /**
   * Greeting. Output is a Media Summary Object. Input must be a Media Lookup Object. Must refer to a media recording that has is_hold_music set to FALSE.
   **/
  @ApiModelProperty(required = true, value = "Greeting. Output is a Media Summary Object. Input must be a Media Lookup Object. Must refer to a media recording that has is_hold_music set to FALSE.")
  public MediaSummary getGreeting() {
    return greeting;
  }
  public void setGreeting(MediaSummary greeting) {
    this.greeting = greeting;
  }

  /**
   * Error Message. Output is a Media Summary Object. Input must be a Media Lookup Object. Must refer to a media recording that has is_hold_music set to FALSE.
   **/
  @ApiModelProperty(required = true, value = "Error Message. Output is a Media Summary Object. Input must be a Media Lookup Object. Must refer to a media recording that has is_hold_music set to FALSE.")
  public MediaSummary getErrorMessage() {
    return errorMessage;
  }
  public void setErrorMessage(MediaSummary errorMessage) {
    this.errorMessage = errorMessage;
  }

  /**
   * Custom audio codec configuration, if any is needed. If provided, must be a simple array containing the prioritized list of desired codecs. Supported codecs are: g711u 64k, g711u 56k, g711a 64k, g711a 56k, g7231, g728, g729, g729A, g729B, g729AB, gms full, rfc2833, t38, ilbc, h263, g722, g722_1, g729D, g729E, amr, amr_wb, efr, evrc, h264, mpeg4, red, cng, SIP Info to 2833
   **/
  @ApiModelProperty(required = true, value = "Custom audio codec configuration, if any is needed. If provided, must be a simple array containing the prioritized list of desired codecs. Supported codecs are: g711u 64k, g711u 56k, g711a 64k, g711a 56k, g7231, g728, g729, g729A, g729B, g729AB, gms full, rfc2833, t38, ilbc, h263, g722, g722_1, g729D, g729E, amr, amr_wb, efr, evrc, h264, mpeg4, red, cng, SIP Info to 2833")
  public List<String> getCodecs() {
    return codecs;
  }
  public void setCodecs(List<String> codecs) {
    this.codecs = codecs;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    TrunkFull trunkFull = (TrunkFull) o;
    return (id == null ? trunkFull.id == null : id.equals(trunkFull.id)) &&
        (name == null ? trunkFull.name == null : name.equals(trunkFull.name)) &&
        (uri == null ? trunkFull.uri == null : uri.equals(trunkFull.uri)) &&
        (maxConcurrentCalls == null ? trunkFull.maxConcurrentCalls == null : maxConcurrentCalls.equals(trunkFull.maxConcurrentCalls)) &&
        (maxMinutesPerMonth == null ? trunkFull.maxMinutesPerMonth == null : maxMinutesPerMonth.equals(trunkFull.maxMinutesPerMonth)) &&
        (greeting == null ? trunkFull.greeting == null : greeting.equals(trunkFull.greeting)) &&
        (errorMessage == null ? trunkFull.errorMessage == null : errorMessage.equals(trunkFull.errorMessage)) &&
        (codecs == null ? trunkFull.codecs == null : codecs.equals(trunkFull.codecs));
  }

  @Override
  public int hashCode() {
    int result = 17;
    result = 31 * result + (id == null ? 0: id.hashCode());
    result = 31 * result + (name == null ? 0: name.hashCode());
    result = 31 * result + (uri == null ? 0: uri.hashCode());
    result = 31 * result + (maxConcurrentCalls == null ? 0: maxConcurrentCalls.hashCode());
    result = 31 * result + (maxMinutesPerMonth == null ? 0: maxMinutesPerMonth.hashCode());
    result = 31 * result + (greeting == null ? 0: greeting.hashCode());
    result = 31 * result + (errorMessage == null ? 0: errorMessage.hashCode());
    result = 31 * result + (codecs == null ? 0: codecs.hashCode());
    return result;
  }

  @Override
  public String toString()  {
    StringBuilder sb = new StringBuilder();
    sb.append("class TrunkFull {\n");
    
    sb.append("  id: ").append(id).append("\n");
    sb.append("  name: ").append(name).append("\n");
    sb.append("  uri: ").append(uri).append("\n");
    sb.append("  maxConcurrentCalls: ").append(maxConcurrentCalls).append("\n");
    sb.append("  maxMinutesPerMonth: ").append(maxMinutesPerMonth).append("\n");
    sb.append("  greeting: ").append(greeting).append("\n");
    sb.append("  errorMessage: ").append(errorMessage).append("\n");
    sb.append("  codecs: ").append(codecs).append("\n");
    sb.append("}\n");
    return sb.toString();
  }
}
