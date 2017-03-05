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

import io.swagger.client.model.ContactSubaccount;

import io.swagger.annotations.*;
import com.google.gson.annotations.SerializedName;


@ApiModel(description = "")
public class CreateSubaccountParams  {
  
  @SerializedName("username")
  private String username = null;
  @SerializedName("password")
  private String password = null;
  @SerializedName("contact")
  private ContactSubaccount contact = null;
  @SerializedName("billing_contact")
  private ContactSubaccount billingContact = null;

  /**
   * Sub account user name
   **/
  @ApiModelProperty(required = true, value = "Sub account user name")
  public String getUsername() {
    return username;
  }
  public void setUsername(String username) {
    this.username = username;
  }

  /**
   * Sub account password
   **/
  @ApiModelProperty(required = true, value = "Sub account password")
  public String getPassword() {
    return password;
  }
  public void setPassword(String password) {
    this.password = password;
  }

  /**
   * Contact Object. See below for details.
   **/
  @ApiModelProperty(value = "Contact Object. See below for details.")
  public ContactSubaccount getContact() {
    return contact;
  }
  public void setContact(ContactSubaccount contact) {
    this.contact = contact;
  }

  /**
   * Contact Object for billing purposes. See below for details.
   **/
  @ApiModelProperty(value = "Contact Object for billing purposes. See below for details.")
  public ContactSubaccount getBillingContact() {
    return billingContact;
  }
  public void setBillingContact(ContactSubaccount billingContact) {
    this.billingContact = billingContact;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    CreateSubaccountParams createSubaccountParams = (CreateSubaccountParams) o;
    return (username == null ? createSubaccountParams.username == null : username.equals(createSubaccountParams.username)) &&
        (password == null ? createSubaccountParams.password == null : password.equals(createSubaccountParams.password)) &&
        (contact == null ? createSubaccountParams.contact == null : contact.equals(createSubaccountParams.contact)) &&
        (billingContact == null ? createSubaccountParams.billingContact == null : billingContact.equals(createSubaccountParams.billingContact));
  }

  @Override
  public int hashCode() {
    int result = 17;
    result = 31 * result + (username == null ? 0: username.hashCode());
    result = 31 * result + (password == null ? 0: password.hashCode());
    result = 31 * result + (contact == null ? 0: contact.hashCode());
    result = 31 * result + (billingContact == null ? 0: billingContact.hashCode());
    return result;
  }

  @Override
  public String toString()  {
    StringBuilder sb = new StringBuilder();
    sb.append("class CreateSubaccountParams {\n");
    
    sb.append("  username: ").append(username).append("\n");
    sb.append("  password: ").append(password).append("\n");
    sb.append("  contact: ").append(contact).append("\n");
    sb.append("  billingContact: ").append(billingContact).append("\n");
    sb.append("}\n");
    return sb.toString();
  }
}