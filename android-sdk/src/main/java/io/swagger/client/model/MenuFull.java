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
import io.swagger.client.model.OptionsListMenus;
import io.swagger.client.model.RouteSummary;

import io.swagger.annotations.*;
import com.google.gson.annotations.SerializedName;


/**
 * The Full Menu Object contains the same properties as the Menu Summary Object, along with the following:
 **/
@ApiModel(description = "The Full Menu Object contains the same properties as the Menu Summary Object, along with the following:")
public class MenuFull  {
  
  @SerializedName("id")
  private Integer id = null;
  @SerializedName("name")
  private String name = null;
  @SerializedName("allow_extension_dial")
  private Boolean allowExtensionDial = null;
  @SerializedName("keypress_wait_time")
  private Integer keypressWaitTime = null;
  @SerializedName("greeting")
  private MediaSummary greeting = null;
  @SerializedName("keypress_error")
  private MediaSummary keypressError = null;
  @SerializedName("timeout_handler")
  private RouteSummary timeoutHandler = null;
  @SerializedName("options")
  private OptionsListMenus options = null;

  /**
   * Integer Menu ID. Read-only.
   **/
  @ApiModelProperty(value = "Integer Menu ID. Read-only.")
  public Integer getId() {
    return id;
  }
  public void setId(Integer id) {
    this.id = id;
  }

  /**
   * Name. Required. Unique.
   **/
  @ApiModelProperty(value = "Name. Required. Unique.")
  public String getName() {
    return name;
  }
  public void setName(String name) {
    this.name = name;
  }

  /**
   * Boolean. Determines whether a caller can enter an extension number to bypass the menu.
   **/
  @ApiModelProperty(value = "Boolean. Determines whether a caller can enter an extension number to bypass the menu.")
  public Boolean getAllowExtensionDial() {
    return allowExtensionDial;
  }
  public void setAllowExtensionDial(Boolean allowExtensionDial) {
    this.allowExtensionDial = allowExtensionDial;
  }

  /**
   * Boolean. Determines whether a caller can enter an extension number to bypass the menu.
   **/
  @ApiModelProperty(value = "Boolean. Determines whether a caller can enter an extension number to bypass the menu.")
  public Integer getKeypressWaitTime() {
    return keypressWaitTime;
  }
  public void setKeypressWaitTime(Integer keypressWaitTime) {
    this.keypressWaitTime = keypressWaitTime;
  }

  /**
   * Greeting that is played when a caller enters a menu. Output is a Media Summary Object. Input must be a Media Lookup Object. Must refer to a media recording that has is_hold_music set to FALSE.
   **/
  @ApiModelProperty(value = "Greeting that is played when a caller enters a menu. Output is a Media Summary Object. Input must be a Media Lookup Object. Must refer to a media recording that has is_hold_music set to FALSE.")
  public MediaSummary getGreeting() {
    return greeting;
  }
  public void setGreeting(MediaSummary greeting) {
    this.greeting = greeting;
  }

  /**
   * Message that is played when the caller makes a keypress error. Output is a Media Summary Object. Input must be a Media Lookup Object. Must refer to a media recording that has is_hold_music set to FALSE.
   **/
  @ApiModelProperty(value = "Message that is played when the caller makes a keypress error. Output is a Media Summary Object. Input must be a Media Lookup Object. Must refer to a media recording that has is_hold_music set to FALSE.")
  public MediaSummary getKeypressError() {
    return keypressError;
  }
  public void setKeypressError(MediaSummary keypressError) {
    this.keypressError = keypressError;
  }

  /**
   * Route that will be entered when the caller fails to choose a menu option within the allotted time. Output is a Route Summary Object if the route is named, otherwise the Full Route Object will be shown. Input must be a Route Lookup Object pointing to a named route.
   **/
  @ApiModelProperty(value = "Route that will be entered when the caller fails to choose a menu option within the allotted time. Output is a Route Summary Object if the route is named, otherwise the Full Route Object will be shown. Input must be a Route Lookup Object pointing to a named route.")
  public RouteSummary getTimeoutHandler() {
    return timeoutHandler;
  }
  public void setTimeoutHandler(RouteSummary timeoutHandler) {
    this.timeoutHandler = timeoutHandler;
  }

  /**
   **/
  @ApiModelProperty(value = "")
  public OptionsListMenus getOptions() {
    return options;
  }
  public void setOptions(OptionsListMenus options) {
    this.options = options;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    MenuFull menuFull = (MenuFull) o;
    return (id == null ? menuFull.id == null : id.equals(menuFull.id)) &&
        (name == null ? menuFull.name == null : name.equals(menuFull.name)) &&
        (allowExtensionDial == null ? menuFull.allowExtensionDial == null : allowExtensionDial.equals(menuFull.allowExtensionDial)) &&
        (keypressWaitTime == null ? menuFull.keypressWaitTime == null : keypressWaitTime.equals(menuFull.keypressWaitTime)) &&
        (greeting == null ? menuFull.greeting == null : greeting.equals(menuFull.greeting)) &&
        (keypressError == null ? menuFull.keypressError == null : keypressError.equals(menuFull.keypressError)) &&
        (timeoutHandler == null ? menuFull.timeoutHandler == null : timeoutHandler.equals(menuFull.timeoutHandler)) &&
        (options == null ? menuFull.options == null : options.equals(menuFull.options));
  }

  @Override
  public int hashCode() {
    int result = 17;
    result = 31 * result + (id == null ? 0: id.hashCode());
    result = 31 * result + (name == null ? 0: name.hashCode());
    result = 31 * result + (allowExtensionDial == null ? 0: allowExtensionDial.hashCode());
    result = 31 * result + (keypressWaitTime == null ? 0: keypressWaitTime.hashCode());
    result = 31 * result + (greeting == null ? 0: greeting.hashCode());
    result = 31 * result + (keypressError == null ? 0: keypressError.hashCode());
    result = 31 * result + (timeoutHandler == null ? 0: timeoutHandler.hashCode());
    result = 31 * result + (options == null ? 0: options.hashCode());
    return result;
  }

  @Override
  public String toString()  {
    StringBuilder sb = new StringBuilder();
    sb.append("class MenuFull {\n");
    
    sb.append("  id: ").append(id).append("\n");
    sb.append("  name: ").append(name).append("\n");
    sb.append("  allowExtensionDial: ").append(allowExtensionDial).append("\n");
    sb.append("  keypressWaitTime: ").append(keypressWaitTime).append("\n");
    sb.append("  greeting: ").append(greeting).append("\n");
    sb.append("  keypressError: ").append(keypressError).append("\n");
    sb.append("  timeoutHandler: ").append(timeoutHandler).append("\n");
    sb.append("  options: ").append(options).append("\n");
    sb.append("}\n");
    return sb.toString();
  }
}