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
public class FilterListPhoneNumbersRegions  {
  
  @SerializedName("country_code")
  private String countryCode = null;
  @SerializedName("npa")
  private List<Integer> npa = null;
  @SerializedName("nxx")
  private String nxx = null;
  @SerializedName("is_toll_free")
  private String isTollFree = null;
  @SerializedName("city")
  private String city = null;
  @SerializedName("province_postal_code")
  private String provincePostalCode = null;
  @SerializedName("country_postal_code")
  private String countryPostalCode = null;

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
  public String getIsTollFree() {
    return isTollFree;
  }
  public void setIsTollFree(String isTollFree) {
    this.isTollFree = isTollFree;
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
  public String getProvincePostalCode() {
    return provincePostalCode;
  }
  public void setProvincePostalCode(String provincePostalCode) {
    this.provincePostalCode = provincePostalCode;
  }

  /**
   **/
  @ApiModelProperty(value = "")
  public String getCountryPostalCode() {
    return countryPostalCode;
  }
  public void setCountryPostalCode(String countryPostalCode) {
    this.countryPostalCode = countryPostalCode;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    FilterListPhoneNumbersRegions filterListPhoneNumbersRegions = (FilterListPhoneNumbersRegions) o;
    return (countryCode == null ? filterListPhoneNumbersRegions.countryCode == null : countryCode.equals(filterListPhoneNumbersRegions.countryCode)) &&
        (npa == null ? filterListPhoneNumbersRegions.npa == null : npa.equals(filterListPhoneNumbersRegions.npa)) &&
        (nxx == null ? filterListPhoneNumbersRegions.nxx == null : nxx.equals(filterListPhoneNumbersRegions.nxx)) &&
        (isTollFree == null ? filterListPhoneNumbersRegions.isTollFree == null : isTollFree.equals(filterListPhoneNumbersRegions.isTollFree)) &&
        (city == null ? filterListPhoneNumbersRegions.city == null : city.equals(filterListPhoneNumbersRegions.city)) &&
        (provincePostalCode == null ? filterListPhoneNumbersRegions.provincePostalCode == null : provincePostalCode.equals(filterListPhoneNumbersRegions.provincePostalCode)) &&
        (countryPostalCode == null ? filterListPhoneNumbersRegions.countryPostalCode == null : countryPostalCode.equals(filterListPhoneNumbersRegions.countryPostalCode));
  }

  @Override
  public int hashCode() {
    int result = 17;
    result = 31 * result + (countryCode == null ? 0: countryCode.hashCode());
    result = 31 * result + (npa == null ? 0: npa.hashCode());
    result = 31 * result + (nxx == null ? 0: nxx.hashCode());
    result = 31 * result + (isTollFree == null ? 0: isTollFree.hashCode());
    result = 31 * result + (city == null ? 0: city.hashCode());
    result = 31 * result + (provincePostalCode == null ? 0: provincePostalCode.hashCode());
    result = 31 * result + (countryPostalCode == null ? 0: countryPostalCode.hashCode());
    return result;
  }

  @Override
  public String toString()  {
    StringBuilder sb = new StringBuilder();
    sb.append("class FilterListPhoneNumbersRegions {\n");
    
    sb.append("  countryCode: ").append(countryCode).append("\n");
    sb.append("  npa: ").append(npa).append("\n");
    sb.append("  nxx: ").append(nxx).append("\n");
    sb.append("  isTollFree: ").append(isTollFree).append("\n");
    sb.append("  city: ").append(city).append("\n");
    sb.append("  provincePostalCode: ").append(provincePostalCode).append("\n");
    sb.append("  countryPostalCode: ").append(countryPostalCode).append("\n");
    sb.append("}\n");
    return sb.toString();
  }
}
