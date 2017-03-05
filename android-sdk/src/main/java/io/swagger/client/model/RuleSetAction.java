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

import io.swagger.client.model.ExtensionSummary;
import io.swagger.client.model.MediaSummary;
import io.swagger.client.model.MenuSummary;
import io.swagger.client.model.QueueSummary;
import io.swagger.client.model.RuleSetForwardItem;
import io.swagger.client.model.TrunkSummary;
import java.util.*;

import io.swagger.annotations.*;
import com.google.gson.annotations.SerializedName;


/**
 * Filter Object. Optional. See below for details.
 **/
@ApiModel(description = "Filter Object. Optional. See below for details.")
public class RuleSetAction  {
  
  @SerializedName("action")
  private String action = null;
  @SerializedName("extension")
  private ExtensionSummary extension = null;
  @SerializedName("items")
  private List<RuleSetForwardItem> items = null;
  @SerializedName("timeout")
  private Integer timeout = null;
  @SerializedName("hold_music")
  private MediaSummary holdMusic = null;
  @SerializedName("greeting")
  private MediaSummary greeting = null;
  @SerializedName("duration")
  private Integer duration = null;
  @SerializedName("menu")
  private MenuSummary menu = null;
  @SerializedName("queue")
  private QueueSummary queue = null;
  @SerializedName("trunk")
  private TrunkSummary trunk = null;

  /**
   * Required.
   **/
  @ApiModelProperty(value = "Required.")
  public String getAction() {
    return action;
  }
  public void setAction(String action) {
    this.action = action;
  }

  /**
   * Extension that this action refers to. Output is an Extension Summary Object. Input must be an Extension Lookup Object. Required.
   **/
  @ApiModelProperty(value = "Extension that this action refers to. Output is an Extension Summary Object. Input must be an Extension Lookup Object. Required.")
  public ExtensionSummary getExtension() {
    return extension;
  }
  public void setExtension(ExtensionSummary extension) {
    this.extension = extension;
  }

  /**
   * This action is for forwarding calls to any number of extensions or phone numbers. The forwarding is handled in parallel, meaning that all phone numbers and/or extensions will ring simultaneously. When the call is answered by any single phone number or extension, ringing will stop for all of them. Subsequent actions in this rule set will be performed if the call is not answered before the timeout period is reached, or if it is forwarded to an extension that has its own route and that route does not result in any actions that disconnect the call or take over call handling.
   **/
  @ApiModelProperty(value = "This action is for forwarding calls to any number of extensions or phone numbers. The forwarding is handled in parallel, meaning that all phone numbers and/or extensions will ring simultaneously. When the call is answered by any single phone number or extension, ringing will stop for all of them. Subsequent actions in this rule set will be performed if the call is not answered before the timeout period is reached, or if it is forwarded to an extension that has its own route and that route does not result in any actions that disconnect the call or take over call handling.")
  public List<RuleSetForwardItem> getItems() {
    return items;
  }
  public void setItems(List<RuleSetForwardItem> items) {
    this.items = items;
  }

  /**
   * Seconds that our routing engine should wait until moving on. Optional. Must be an integer between 5 and 90. Default is 5 seconds.
   **/
  @ApiModelProperty(value = "Seconds that our routing engine should wait until moving on. Optional. Must be an integer between 5 and 90. Default is 5 seconds.")
  public Integer getTimeout() {
    return timeout;
  }
  public void setTimeout(Integer timeout) {
    this.timeout = timeout;
  }

  /**
   * Hold Music to be played while callers are waiting. Output is a Media Summary Object. Input must be a Media Lookup Object. Optional. Must refer to a media recording that has is_hold_music set to TRUE. Default is to play a standard ring tone.
   **/
  @ApiModelProperty(value = "Hold Music to be played while callers are waiting. Output is a Media Summary Object. Input must be a Media Lookup Object. Optional. Must refer to a media recording that has is_hold_music set to TRUE. Default is to play a standard ring tone.")
  public MediaSummary getHoldMusic() {
    return holdMusic;
  }
  public void setHoldMusic(MediaSummary holdMusic) {
    this.holdMusic = holdMusic;
  }

  /**
   * Greeting that this action refers to. Output is a Media Summary Object. Input must be a Media Lookup Object. Required. Must refer to a media recording that has is_hold_music set to FALSE.
   **/
  @ApiModelProperty(value = "Greeting that this action refers to. Output is a Media Summary Object. Input must be a Media Lookup Object. Required. Must refer to a media recording that has is_hold_music set to FALSE.")
  public MediaSummary getGreeting() {
    return greeting;
  }
  public void setGreeting(MediaSummary greeting) {
    this.greeting = greeting;
  }

  /**
   * Required. Seconds that the caller should be placed on hold before being moved onto the next action. Must be an integer between 1 and 60 seconds.
   **/
  @ApiModelProperty(value = "Required. Seconds that the caller should be placed on hold before being moved onto the next action. Must be an integer between 1 and 60 seconds.")
  public Integer getDuration() {
    return duration;
  }
  public void setDuration(Integer duration) {
    this.duration = duration;
  }

  /**
   * Menu that this action refers to. Required. Output is a Menu Summary Object. Input must be a Menu Lookup Object.
   **/
  @ApiModelProperty(value = "Menu that this action refers to. Required. Output is a Menu Summary Object. Input must be a Menu Lookup Object.")
  public MenuSummary getMenu() {
    return menu;
  }
  public void setMenu(MenuSummary menu) {
    this.menu = menu;
  }

  /**
   * Queue that this action refers to. Required. Output is a Queue Summary Object. Input must be a Queue Lookup Object.
   **/
  @ApiModelProperty(value = "Queue that this action refers to. Required. Output is a Queue Summary Object. Input must be a Queue Lookup Object.")
  public QueueSummary getQueue() {
    return queue;
  }
  public void setQueue(QueueSummary queue) {
    this.queue = queue;
  }

  /**
   * Trunk that this action refers to. Required. Output is a Trunk Summary Object. Input must be a Trunk Lookup Object.
   **/
  @ApiModelProperty(value = "Trunk that this action refers to. Required. Output is a Trunk Summary Object. Input must be a Trunk Lookup Object.")
  public TrunkSummary getTrunk() {
    return trunk;
  }
  public void setTrunk(TrunkSummary trunk) {
    this.trunk = trunk;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    RuleSetAction ruleSetAction = (RuleSetAction) o;
    return (action == null ? ruleSetAction.action == null : action.equals(ruleSetAction.action)) &&
        (extension == null ? ruleSetAction.extension == null : extension.equals(ruleSetAction.extension)) &&
        (items == null ? ruleSetAction.items == null : items.equals(ruleSetAction.items)) &&
        (timeout == null ? ruleSetAction.timeout == null : timeout.equals(ruleSetAction.timeout)) &&
        (holdMusic == null ? ruleSetAction.holdMusic == null : holdMusic.equals(ruleSetAction.holdMusic)) &&
        (greeting == null ? ruleSetAction.greeting == null : greeting.equals(ruleSetAction.greeting)) &&
        (duration == null ? ruleSetAction.duration == null : duration.equals(ruleSetAction.duration)) &&
        (menu == null ? ruleSetAction.menu == null : menu.equals(ruleSetAction.menu)) &&
        (queue == null ? ruleSetAction.queue == null : queue.equals(ruleSetAction.queue)) &&
        (trunk == null ? ruleSetAction.trunk == null : trunk.equals(ruleSetAction.trunk));
  }

  @Override
  public int hashCode() {
    int result = 17;
    result = 31 * result + (action == null ? 0: action.hashCode());
    result = 31 * result + (extension == null ? 0: extension.hashCode());
    result = 31 * result + (items == null ? 0: items.hashCode());
    result = 31 * result + (timeout == null ? 0: timeout.hashCode());
    result = 31 * result + (holdMusic == null ? 0: holdMusic.hashCode());
    result = 31 * result + (greeting == null ? 0: greeting.hashCode());
    result = 31 * result + (duration == null ? 0: duration.hashCode());
    result = 31 * result + (menu == null ? 0: menu.hashCode());
    result = 31 * result + (queue == null ? 0: queue.hashCode());
    result = 31 * result + (trunk == null ? 0: trunk.hashCode());
    return result;
  }

  @Override
  public String toString()  {
    StringBuilder sb = new StringBuilder();
    sb.append("class RuleSetAction {\n");
    
    sb.append("  action: ").append(action).append("\n");
    sb.append("  extension: ").append(extension).append("\n");
    sb.append("  items: ").append(items).append("\n");
    sb.append("  timeout: ").append(timeout).append("\n");
    sb.append("  holdMusic: ").append(holdMusic).append("\n");
    sb.append("  greeting: ").append(greeting).append("\n");
    sb.append("  duration: ").append(duration).append("\n");
    sb.append("  menu: ").append(menu).append("\n");
    sb.append("  queue: ").append(queue).append("\n");
    sb.append("  trunk: ").append(trunk).append("\n");
    sb.append("}\n");
    return sb.toString();
  }
}