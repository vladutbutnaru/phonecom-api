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
 * The Region Object may include any of several fields describing the group, as well as the quantity of phone numbers available in that group. Here are the properties:
 **/
@ApiModel(description = "The Region Object may include any of several fields describing the group, as well as the quantity of phone numbers available in that group. Here are the properties:")
public class PhoneNumbersRegionFull  {
  
  @SerializedName("country_code")
  private String countryCode = null;
  @SerializedName("npa")
  private String npa = null;
  @SerializedName("nxx")
  private String nxx = null;
  @SerializedName("is_toll_free")
  private Integer isTollFree = null;
  @SerializedName("city")
  private String city = null;
  @SerializedName("province_postal_code")
  private String provincePostalCode = null;
  @SerializedName("country_postal_code")
  private String countryPostalCode = null;
  @SerializedName("quantity")
  private Integer quantity = null;

  /**
   * Optional. Integer representing the international calling code for the phone number's country.
   **/
  @ApiModelProperty(value = "Optional. Integer representing the international calling code for the phone number's country.")
  public String getCountryCode() {
    return countryCode;
  }
  public void setCountryCode(String countryCode) {
    this.countryCode = countryCode;
  }

  /**
   * Optional. Area Code, for North American phone numbers.
   **/
  @ApiModelProperty(value = "Optional. Area Code, for North American phone numbers.")
  public String getNpa() {
    return npa;
  }
  public void setNpa(String npa) {
    this.npa = npa;
  }

  /**
   * Optional. Second set of 3 digits, for North American phone numbers.
   **/
  @ApiModelProperty(value = "Optional. Second set of 3 digits, for North American phone numbers.")
  public String getNxx() {
    return nxx;
  }
  public void setNxx(String nxx) {
    this.nxx = nxx;
  }

  /**
   * Optional. Boolean describing whether the phone numbers in this group are toll free.
   **/
  @ApiModelProperty(value = "Optional. Boolean describing whether the phone numbers in this group are toll free.")
  public Integer getIsTollFree() {
    return isTollFree;
  }
  public void setIsTollFree(Integer isTollFree) {
    this.isTollFree = isTollFree;
  }

  /**
   * Optional. City name.
   **/
  @ApiModelProperty(value = "Optional. City name.")
  public String getCity() {
    return city;
  }
  public void setCity(String city) {
    this.city = city;
  }

  /**
   * Optional. Two-letter postal code for the state or province.
   **/
  @ApiModelProperty(value = "Optional. Two-letter postal code for the state or province.")
  public String getProvincePostalCode() {
    return provincePostalCode;
  }
  public void setProvincePostalCode(String provincePostalCode) {
    this.provincePostalCode = provincePostalCode;
  }

  /**
   * Optional. Two-letter postal code for the country.
   **/
  @ApiModelProperty(value = "Optional. Two-letter postal code for the country.")
  public String getCountryPostalCode() {
    return countryPostalCode;
  }
  public void setCountryPostalCode(String countryPostalCode) {
    this.countryPostalCode = countryPostalCode;
  }

  /**
   * Integer. Quantity of phone numbers currently available in the given region.
   **/
  @ApiModelProperty(value = "Integer. Quantity of phone numbers currently available in the given region.")
  public Integer getQuantity() {
    return quantity;
  }
  public void setQuantity(Integer quantity) {
    this.quantity = quantity;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    PhoneNumbersRegionFull phoneNumbersRegionFull = (PhoneNumbersRegionFull) o;
    return (countryCode == null ? phoneNumbersRegionFull.countryCode == null : countryCode.equals(phoneNumbersRegionFull.countryCode)) &&
        (npa == null ? phoneNumbersRegionFull.npa == null : npa.equals(phoneNumbersRegionFull.npa)) &&
        (nxx == null ? phoneNumbersRegionFull.nxx == null : nxx.equals(phoneNumbersRegionFull.nxx)) &&
        (isTollFree == null ? phoneNumbersRegionFull.isTollFree == null : isTollFree.equals(phoneNumbersRegionFull.isTollFree)) &&
        (city == null ? phoneNumbersRegionFull.city == null : city.equals(phoneNumbersRegionFull.city)) &&
        (provincePostalCode == null ? phoneNumbersRegionFull.provincePostalCode == null : provincePostalCode.equals(phoneNumbersRegionFull.provincePostalCode)) &&
        (countryPostalCode == null ? phoneNumbersRegionFull.countryPostalCode == null : countryPostalCode.equals(phoneNumbersRegionFull.countryPostalCode)) &&
        (quantity == null ? phoneNumbersRegionFull.quantity == null : quantity.equals(phoneNumbersRegionFull.quantity));
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
    result = 31 * result + (quantity == null ? 0: quantity.hashCode());
    return result;
  }

  @Override
  public String toString()  {
    StringBuilder sb = new StringBuilder();
    sb.append("class PhoneNumbersRegionFull {\n");
    
    sb.append("  countryCode: ").append(countryCode).append("\n");
    sb.append("  npa: ").append(npa).append("\n");
    sb.append("  nxx: ").append(nxx).append("\n");
    sb.append("  isTollFree: ").append(isTollFree).append("\n");
    sb.append("  city: ").append(city).append("\n");
    sb.append("  provincePostalCode: ").append(provincePostalCode).append("\n");
    sb.append("  countryPostalCode: ").append(countryPostalCode).append("\n");
    sb.append("  quantity: ").append(quantity).append("\n");
    sb.append("}\n");
    return sb.toString();
  }
}
