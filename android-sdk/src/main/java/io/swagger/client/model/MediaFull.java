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


import io.swagger.annotations.*;
import com.google.gson.annotations.SerializedName;


/**
 * The Full Recording Object includes all of the properties from the Recording Summary Object, along with the following:
 **/
@ApiModel(description = "The Full Recording Object includes all of the properties from the Recording Summary Object, along with the following:")
public class MediaFull  {
  
  @SerializedName("id")
  private Integer id = null;
  @SerializedName("name")
  private String name = null;
  @SerializedName("type")
  private String type = null;

  /**
   * Recording ID. Read-only.
   **/
  @ApiModelProperty(value = "Recording ID. Read-only.")
  public Integer getId() {
    return id;
  }
  public void setId(Integer id) {
    this.id = id;
  }

  /**
   * Name of recording
   **/
  @ApiModelProperty(value = "Name of recording")
  public String getName() {
    return name;
  }
  public void setName(String name) {
    this.name = name;
  }

  /**
   * Can be hold_music or greeting. Indicates the purpose of this recording and where it can be used.
   **/
  @ApiModelProperty(value = "Can be hold_music or greeting. Indicates the purpose of this recording and where it can be used.")
  public String getType() {
    return type;
  }
  public void setType(String type) {
    this.type = type;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    MediaFull mediaFull = (MediaFull) o;
    return (id == null ? mediaFull.id == null : id.equals(mediaFull.id)) &&
        (name == null ? mediaFull.name == null : name.equals(mediaFull.name)) &&
        (type == null ? mediaFull.type == null : type.equals(mediaFull.type));
  }

  @Override
  public int hashCode() {
    int result = 17;
    result = 31 * result + (id == null ? 0: id.hashCode());
    result = 31 * result + (name == null ? 0: name.hashCode());
    result = 31 * result + (type == null ? 0: type.hashCode());
    return result;
  }

  @Override
  public String toString()  {
    StringBuilder sb = new StringBuilder();
    sb.append("class MediaFull {\n");
    
    sb.append("  id: ").append(id).append("\n");
    sb.append("  name: ").append(name).append("\n");
    sb.append("  type: ").append(type).append("\n");
    sb.append("}\n");
    return sb.toString();
  }
}
