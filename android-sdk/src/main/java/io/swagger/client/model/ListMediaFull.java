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

import io.swagger.client.model.FilterIdNameArray;
import io.swagger.client.model.MediaFull;
import io.swagger.client.model.SortIdName;
import java.util.*;

import io.swagger.annotations.*;
import com.google.gson.annotations.SerializedName;


@ApiModel(description = "")
public class ListMediaFull  {
  
  @SerializedName("filters")
  private FilterIdNameArray filters = null;
  @SerializedName("sort")
  private SortIdName sort = null;
  @SerializedName("total")
  private Integer total = null;
  @SerializedName("offset")
  private Integer offset = null;
  @SerializedName("limit")
  private Integer limit = null;
  @SerializedName("items")
  private List<MediaFull> items = null;

  /**
   **/
  @ApiModelProperty(value = "")
  public FilterIdNameArray getFilters() {
    return filters;
  }
  public void setFilters(FilterIdNameArray filters) {
    this.filters = filters;
  }

  /**
   **/
  @ApiModelProperty(value = "")
  public SortIdName getSort() {
    return sort;
  }
  public void setSort(SortIdName sort) {
    this.sort = sort;
  }

  /**
   **/
  @ApiModelProperty(value = "")
  public Integer getTotal() {
    return total;
  }
  public void setTotal(Integer total) {
    this.total = total;
  }

  /**
   **/
  @ApiModelProperty(value = "")
  public Integer getOffset() {
    return offset;
  }
  public void setOffset(Integer offset) {
    this.offset = offset;
  }

  /**
   **/
  @ApiModelProperty(value = "")
  public Integer getLimit() {
    return limit;
  }
  public void setLimit(Integer limit) {
    this.limit = limit;
  }

  /**
   **/
  @ApiModelProperty(value = "")
  public List<MediaFull> getItems() {
    return items;
  }
  public void setItems(List<MediaFull> items) {
    this.items = items;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    ListMediaFull listMediaFull = (ListMediaFull) o;
    return (filters == null ? listMediaFull.filters == null : filters.equals(listMediaFull.filters)) &&
        (sort == null ? listMediaFull.sort == null : sort.equals(listMediaFull.sort)) &&
        (total == null ? listMediaFull.total == null : total.equals(listMediaFull.total)) &&
        (offset == null ? listMediaFull.offset == null : offset.equals(listMediaFull.offset)) &&
        (limit == null ? listMediaFull.limit == null : limit.equals(listMediaFull.limit)) &&
        (items == null ? listMediaFull.items == null : items.equals(listMediaFull.items));
  }

  @Override
  public int hashCode() {
    int result = 17;
    result = 31 * result + (filters == null ? 0: filters.hashCode());
    result = 31 * result + (sort == null ? 0: sort.hashCode());
    result = 31 * result + (total == null ? 0: total.hashCode());
    result = 31 * result + (offset == null ? 0: offset.hashCode());
    result = 31 * result + (limit == null ? 0: limit.hashCode());
    result = 31 * result + (items == null ? 0: items.hashCode());
    return result;
  }

  @Override
  public String toString()  {
    StringBuilder sb = new StringBuilder();
    sb.append("class ListMediaFull {\n");
    
    sb.append("  filters: ").append(filters).append("\n");
    sb.append("  sort: ").append(sort).append("\n");
    sb.append("  total: ").append(total).append("\n");
    sb.append("  offset: ").append(offset).append("\n");
    sb.append("  limit: ").append(limit).append("\n");
    sb.append("  items: ").append(items).append("\n");
    sb.append("}\n");
    return sb.toString();
  }
}
