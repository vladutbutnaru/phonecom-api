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
import io.swagger.client.model.RuleSet;
import java.util.*;

import io.swagger.annotations.*;
import com.google.gson.annotations.SerializedName;


/**
 * The root level of the Full Route Object includes all of the properties in the Route Summary Object, along with two more:
 **/
@ApiModel(description = "The root level of the Full Route Object includes all of the properties in the Route Summary Object, along with two more:")
public class RouteFull  {
  
  @SerializedName("id")
  private Integer id = null;
  @SerializedName("name")
  private String name = null;
  @SerializedName("extension")
  private ExtensionSummary extension = null;
  @SerializedName("rules")
  private List<RuleSet> rules = null;

  /**
   * Integer ID. Read-only.
   **/
  @ApiModelProperty(value = "Integer ID. Read-only.")
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
   * Extension to which this route belongs. Output is an Extension Summary Object. Input must be an Extension Lookup Object. Optional. Cannot be changed after a route is created.
   **/
  @ApiModelProperty(value = "Extension to which this route belongs. Output is an Extension Summary Object. Input must be an Extension Lookup Object. Optional. Cannot be changed after a route is created.")
  public ExtensionSummary getExtension() {
    return extension;
  }
  public void setExtension(ExtensionSummary extension) {
    this.extension = extension;
  }

  /**
   * Array of Rule Set Objects. Required. See below for details. When processing incoming calls, the first matching rule set will be used, and all others will be ignored.
   **/
  @ApiModelProperty(value = "Array of Rule Set Objects. Required. See below for details. When processing incoming calls, the first matching rule set will be used, and all others will be ignored.")
  public List<RuleSet> getRules() {
    return rules;
  }
  public void setRules(List<RuleSet> rules) {
    this.rules = rules;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    RouteFull routeFull = (RouteFull) o;
    return (id == null ? routeFull.id == null : id.equals(routeFull.id)) &&
        (name == null ? routeFull.name == null : name.equals(routeFull.name)) &&
        (extension == null ? routeFull.extension == null : extension.equals(routeFull.extension)) &&
        (rules == null ? routeFull.rules == null : rules.equals(routeFull.rules));
  }

  @Override
  public int hashCode() {
    int result = 17;
    result = 31 * result + (id == null ? 0: id.hashCode());
    result = 31 * result + (name == null ? 0: name.hashCode());
    result = 31 * result + (extension == null ? 0: extension.hashCode());
    result = 31 * result + (rules == null ? 0: rules.hashCode());
    return result;
  }

  @Override
  public String toString()  {
    StringBuilder sb = new StringBuilder();
    sb.append("class RouteFull {\n");
    
    sb.append("  id: ").append(id).append("\n");
    sb.append("  name: ").append(name).append("\n");
    sb.append("  extension: ").append(extension).append("\n");
    sb.append("  rules: ").append(rules).append("\n");
    sb.append("}\n");
    return sb.toString();
  }
}