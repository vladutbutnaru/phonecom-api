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

import java.util.*;

import io.swagger.annotations.*;
import com.google.gson.annotations.SerializedName;


@ApiModel(description = "")
public class FilterListAvailableNumbers  {
  
  @SerializedName("phone_number")
  private String phoneNumber = null;
  @SerializedName("country_code")
  private String countryCode = null;
  @SerializedName("npa")
  private List<Integer> npa = null;
  @SerializedName("nxx")
  private String nxx = null;
  @SerializedName("xxxx")
  private String xxxx = null;
  @SerializedName("city")
  private String city = null;
  @SerializedName("province")
  private String province = null;
  @SerializedName("country")
  private String country = null;
  @SerializedName("price")
  private String price = null;
  @SerializedName("category")
  private String category = null;

  /**
   **/
  @ApiModelProperty(value = "")
  public String getPhoneNumber() {
    return phoneNumber;
  }
  public void setPhoneNumber(String phoneNumber) {
    this.phoneNumber = phoneNumber;
  }

  /**
   **/
  @ApiModelProperty(value = "")
  public String getCountryCode() {
    return countryCode;
  }
  public void setCountryCode(String countryCode) {
    this.countryCode = countryCode;
  }

  /**
   **/
  @ApiModelProperty(value = "")
  public List<Integer> getNpa() {
    return npa;
  }
  public void setNpa(List<Integer> npa) {
    this.npa = npa;
  }

  /**
   **/
  @ApiModelProperty(value = "")
  public String getNxx() {
    return nxx;
  }
  public void setNxx(String nxx) {
    this.nxx = nxx;
  }

  /**
   **/
  @ApiModelProperty(value = "")
  public String getXxxx() {
    return xxxx;
  }
  public void setXxxx(String xxxx) {
    this.xxxx = xxxx;
  }

  /**
   **/
  @ApiModelProperty(value = "")
  public String getCity() {
    return city;
  }
  public void setCity(String city) {
    this.city = city;
  }

  /**
   **/
  @ApiModelProperty(value = "")
  public String getProvince() {
    return province;
  }
  public void setProvince(String province) {
    this.province = province;
  }

  /**
   **/
  @ApiModelProperty(value = "")
  public String getCountry() {
    return country;
  }
  public void setCountry(String country) {
    this.country = country;
  }

  /**
   **/
  @ApiModelProperty(value = "")
  public String getPrice() {
    return price;
  }
  public void setPrice(String price) {
    this.price = price;
  }

  /**
   **/
  @ApiModelProperty(value = "")
  public String getCategory() {
    return category;
  }
  public void setCategory(String category) {
    this.category = category;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    FilterListAvailableNumbers filterListAvailableNumbers = (FilterListAvailableNumbers) o;
    return (phoneNumber == null ? filterListAvailableNumbers.phoneNumber == null : phoneNumber.equals(filterListAvailableNumbers.phoneNumber)) &&
        (countryCode == null ? filterListAvailableNumbers.countryCode == null : countryCode.equals(filterListAvailableNumbers.countryCode)) &&
        (npa == null ? filterListAvailableNumbers.npa == null : npa.equals(filterListAvailableNumbers.npa)) &&
        (nxx == null ? filterListAvailableNumbers.nxx == null : nxx.equals(filterListAvailableNumbers.nxx)) &&
        (xxxx == null ? filterListAvailableNumbers.xxxx == null : xxxx.equals(filterListAvailableNumbers.xxxx)) &&
        (city == null ? filterListAvailableNumbers.city == null : city.equals(filterListAvailableNumbers.city)) &&
        (province == null ? filterListAvailableNumbers.province == null : province.equals(filterListAvailableNumbers.province)) &&
        (country == null ? filterListAvailableNumbers.country == null : country.equals(filterListAvailableNumbers.country)) &&
        (price == null ? filterListAvailableNumbers.price == null : price.equals(filterListAvailableNumbers.price)) &&
        (category == null ? filterListAvailableNumbers.category == null : category.equals(filterListAvailableNumbers.category));
  }

  @Override
  public int hashCode() {
    int result = 17;
    result = 31 * result + (phoneNumber == null ? 0: phoneNumber.hashCode());
    result = 31 * result + (countryCode == null ? 0: countryCode.hashCode());
    result = 31 * result + (npa == null ? 0: npa.hashCode());
    result = 31 * result + (nxx == null ? 0: nxx.hashCode());
    result = 31 * result + (xxxx == null ? 0: xxxx.hashCode());
    result = 31 * result + (city == null ? 0: city.hashCode());
    result = 31 * result + (province == null ? 0: province.hashCode());
    result = 31 * result + (country == null ? 0: country.hashCode());
    result = 31 * result + (price == null ? 0: price.hashCode());
    result = 31 * result + (category == null ? 0: category.hashCode());
    return result;
  }

  @Override
  public String toString()  {
    StringBuilder sb = new StringBuilder();
    sb.append("class FilterListAvailableNumbers {\n");
    
    sb.append("  phoneNumber: ").append(phoneNumber).append("\n");
    sb.append("  countryCode: ").append(countryCode).append("\n");
    sb.append("  npa: ").append(npa).append("\n");
    sb.append("  nxx: ").append(nxx).append("\n");
    sb.append("  xxxx: ").append(xxxx).append("\n");
    sb.append("  city: ").append(city).append("\n");
    sb.append("  province: ").append(province).append("\n");
    sb.append("  country: ").append(country).append("\n");
    sb.append("  price: ").append(price).append("\n");
    sb.append("  category: ").append(category).append("\n");
    sb.append("}\n");
    return sb.toString();
  }
}