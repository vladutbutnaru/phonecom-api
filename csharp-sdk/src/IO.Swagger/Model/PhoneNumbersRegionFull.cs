/* 
 * Phone.com API
 *
 * This is a Phone.com api Swagger definition
 *
 * OpenAPI spec version: 1.0.0
 * Contact: apisupport@phone.com
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

using System;
using System.Linq;
using System.IO;
using System.Text;
using System.Collections;
using System.Collections.Generic;
using System.Collections.ObjectModel;
using System.Runtime.Serialization;
using Newtonsoft.Json;
using Newtonsoft.Json.Converters;

namespace IO.Swagger.Model
{
    /// <summary>
    /// The Region Object may include any of several fields describing the group, as well as the quantity of phone numbers available in that group. Here are the properties:
    /// </summary>
    [DataContract]
    public partial class PhoneNumbersRegionFull :  IEquatable<PhoneNumbersRegionFull>
    {
        /// <summary>
        /// Initializes a new instance of the <see cref="PhoneNumbersRegionFull" /> class.
        /// </summary>
        /// <param name="CountryCode">Optional. Integer representing the international calling code for the phone number&#39;s country..</param>
        /// <param name="Npa">Optional. Area Code, for North American phone numbers..</param>
        /// <param name="Nxx">Optional. Second set of 3 digits, for North American phone numbers..</param>
        /// <param name="IsTollFree">Optional. Boolean describing whether the phone numbers in this group are toll free..</param>
        /// <param name="City">Optional. City name..</param>
        /// <param name="ProvincePostalCode">Optional. Two-letter postal code for the state or province..</param>
        /// <param name="CountryPostalCode">Optional. Two-letter postal code for the country..</param>
        /// <param name="Quantity">Integer. Quantity of phone numbers currently available in the given region..</param>
        public PhoneNumbersRegionFull(string CountryCode = null, string Npa = null, string Nxx = null, int? IsTollFree = null, string City = null, string ProvincePostalCode = null, string CountryPostalCode = null, int? Quantity = null)
        {
            this.CountryCode = CountryCode;
            this.Npa = Npa;
            this.Nxx = Nxx;
            this.IsTollFree = IsTollFree;
            this.City = City;
            this.ProvincePostalCode = ProvincePostalCode;
            this.CountryPostalCode = CountryPostalCode;
            this.Quantity = Quantity;
        }
        
        /// <summary>
        /// Optional. Integer representing the international calling code for the phone number&#39;s country.
        /// </summary>
        /// <value>Optional. Integer representing the international calling code for the phone number&#39;s country.</value>
        [DataMember(Name="country_code", EmitDefaultValue=false)]
        public string CountryCode { get; set; }
        /// <summary>
        /// Optional. Area Code, for North American phone numbers.
        /// </summary>
        /// <value>Optional. Area Code, for North American phone numbers.</value>
        [DataMember(Name="npa", EmitDefaultValue=false)]
        public string Npa { get; set; }
        /// <summary>
        /// Optional. Second set of 3 digits, for North American phone numbers.
        /// </summary>
        /// <value>Optional. Second set of 3 digits, for North American phone numbers.</value>
        [DataMember(Name="nxx", EmitDefaultValue=false)]
        public string Nxx { get; set; }
        /// <summary>
        /// Optional. Boolean describing whether the phone numbers in this group are toll free.
        /// </summary>
        /// <value>Optional. Boolean describing whether the phone numbers in this group are toll free.</value>
        [DataMember(Name="is_toll_free", EmitDefaultValue=false)]
        public int? IsTollFree { get; set; }
        /// <summary>
        /// Optional. City name.
        /// </summary>
        /// <value>Optional. City name.</value>
        [DataMember(Name="city", EmitDefaultValue=false)]
        public string City { get; set; }
        /// <summary>
        /// Optional. Two-letter postal code for the state or province.
        /// </summary>
        /// <value>Optional. Two-letter postal code for the state or province.</value>
        [DataMember(Name="province_postal_code", EmitDefaultValue=false)]
        public string ProvincePostalCode { get; set; }
        /// <summary>
        /// Optional. Two-letter postal code for the country.
        /// </summary>
        /// <value>Optional. Two-letter postal code for the country.</value>
        [DataMember(Name="country_postal_code", EmitDefaultValue=false)]
        public string CountryPostalCode { get; set; }
        /// <summary>
        /// Integer. Quantity of phone numbers currently available in the given region.
        /// </summary>
        /// <value>Integer. Quantity of phone numbers currently available in the given region.</value>
        [DataMember(Name="quantity", EmitDefaultValue=false)]
        public int? Quantity { get; set; }
        /// <summary>
        /// Returns the string presentation of the object
        /// </summary>
        /// <returns>String presentation of the object</returns>
        public override string ToString()
        {
            var sb = new StringBuilder();
            sb.Append("class PhoneNumbersRegionFull {\n");
            sb.Append("  CountryCode: ").Append(CountryCode).Append("\n");
            sb.Append("  Npa: ").Append(Npa).Append("\n");
            sb.Append("  Nxx: ").Append(Nxx).Append("\n");
            sb.Append("  IsTollFree: ").Append(IsTollFree).Append("\n");
            sb.Append("  City: ").Append(City).Append("\n");
            sb.Append("  ProvincePostalCode: ").Append(ProvincePostalCode).Append("\n");
            sb.Append("  CountryPostalCode: ").Append(CountryPostalCode).Append("\n");
            sb.Append("  Quantity: ").Append(Quantity).Append("\n");
            sb.Append("}\n");
            return sb.ToString();
        }
  
        /// <summary>
        /// Returns the JSON string presentation of the object
        /// </summary>
        /// <returns>JSON string presentation of the object</returns>
        public string ToJson()
        {
            return JsonConvert.SerializeObject(this, Formatting.Indented);
        }

        /// <summary>
        /// Returns true if objects are equal
        /// </summary>
        /// <param name="obj">Object to be compared</param>
        /// <returns>Boolean</returns>
        public override bool Equals(object obj)
        {
            // credit: http://stackoverflow.com/a/10454552/677735
            return this.Equals(obj as PhoneNumbersRegionFull);
        }

        /// <summary>
        /// Returns true if PhoneNumbersRegionFull instances are equal
        /// </summary>
        /// <param name="other">Instance of PhoneNumbersRegionFull to be compared</param>
        /// <returns>Boolean</returns>
        public bool Equals(PhoneNumbersRegionFull other)
        {
            // credit: http://stackoverflow.com/a/10454552/677735
            if (other == null)
                return false;

            return 
                (
                    this.CountryCode == other.CountryCode ||
                    this.CountryCode != null &&
                    this.CountryCode.Equals(other.CountryCode)
                ) && 
                (
                    this.Npa == other.Npa ||
                    this.Npa != null &&
                    this.Npa.Equals(other.Npa)
                ) && 
                (
                    this.Nxx == other.Nxx ||
                    this.Nxx != null &&
                    this.Nxx.Equals(other.Nxx)
                ) && 
                (
                    this.IsTollFree == other.IsTollFree ||
                    this.IsTollFree != null &&
                    this.IsTollFree.Equals(other.IsTollFree)
                ) && 
                (
                    this.City == other.City ||
                    this.City != null &&
                    this.City.Equals(other.City)
                ) && 
                (
                    this.ProvincePostalCode == other.ProvincePostalCode ||
                    this.ProvincePostalCode != null &&
                    this.ProvincePostalCode.Equals(other.ProvincePostalCode)
                ) && 
                (
                    this.CountryPostalCode == other.CountryPostalCode ||
                    this.CountryPostalCode != null &&
                    this.CountryPostalCode.Equals(other.CountryPostalCode)
                ) && 
                (
                    this.Quantity == other.Quantity ||
                    this.Quantity != null &&
                    this.Quantity.Equals(other.Quantity)
                );
        }

        /// <summary>
        /// Gets the hash code
        /// </summary>
        /// <returns>Hash code</returns>
        public override int GetHashCode()
        {
            // credit: http://stackoverflow.com/a/263416/677735
            unchecked // Overflow is fine, just wrap
            {
                int hash = 41;
                // Suitable nullity checks etc, of course :)
                if (this.CountryCode != null)
                    hash = hash * 59 + this.CountryCode.GetHashCode();
                if (this.Npa != null)
                    hash = hash * 59 + this.Npa.GetHashCode();
                if (this.Nxx != null)
                    hash = hash * 59 + this.Nxx.GetHashCode();
                if (this.IsTollFree != null)
                    hash = hash * 59 + this.IsTollFree.GetHashCode();
                if (this.City != null)
                    hash = hash * 59 + this.City.GetHashCode();
                if (this.ProvincePostalCode != null)
                    hash = hash * 59 + this.ProvincePostalCode.GetHashCode();
                if (this.CountryPostalCode != null)
                    hash = hash * 59 + this.CountryPostalCode.GetHashCode();
                if (this.Quantity != null)
                    hash = hash * 59 + this.Quantity.GetHashCode();
                return hash;
            }
        }
    }

}
