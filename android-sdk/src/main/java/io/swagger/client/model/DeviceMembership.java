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

import io.swagger.client.model.DeviceSummary;

import io.swagger.annotations.*;
import com.google.gson.annotations.SerializedName;


/**
 * Device Membership Object, or NULL. Read-only. See below for details.
 **/
@ApiModel(description = "Device Membership Object, or NULL. Read-only. See below for details.")
public class DeviceMembership  {
  
  @SerializedName("line")
  private Integer line = null;
  @SerializedName("device")
  private DeviceSummary device = null;

  /**
   * Line number to which this extension is assigned. Integer.
   **/
  @ApiModelProperty(value = "Line number to which this extension is assigned. Integer.")
  public Integer getLine() {
    return line;
  }
  public void setLine(Integer line) {
    this.line = line;
  }

  /**
   * Device that this extension belongs to. Output is an Device Summary Object.
   **/
  @ApiModelProperty(value = "Device that this extension belongs to. Output is an Device Summary Object.")
  public DeviceSummary getDevice() {
    return device;
  }
  public void setDevice(DeviceSummary device) {
    this.device = device;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    DeviceMembership deviceMembership = (DeviceMembership) o;
    return (line == null ? deviceMembership.line == null : line.equals(deviceMembership.line)) &&
        (device == null ? deviceMembership.device == null : device.equals(deviceMembership.device));
  }

  @Override
  public int hashCode() {
    int result = 17;
    result = 31 * result + (line == null ? 0: line.hashCode());
    result = 31 * result + (device == null ? 0: device.hashCode());
    return result;
  }

  @Override
  public String toString()  {
    StringBuilder sb = new StringBuilder();
    sb.append("class DeviceMembership {\n");
    
    sb.append("  line: ").append(line).append("\n");
    sb.append("  device: ").append(device).append("\n");
    sb.append("}\n");
    return sb.toString();
  }
}
