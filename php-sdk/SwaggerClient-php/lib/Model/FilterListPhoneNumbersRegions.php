<?php
/**
 * FilterListPhoneNumbersRegions
 *
 * PHP version 5
 *
 * @category Class
 * @package  Swagger\Client
 * @author   http://github.com/swagger-api/swagger-codegen
 * @license  http://www.apache.org/licenses/LICENSE-2.0 Apache Licene v2
 * @link     https://github.com/swagger-api/swagger-codegen
 */

/**
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

/**
 * NOTE: This class is auto generated by the swagger code generator program.
 * https://github.com/swagger-api/swagger-codegen
 * Do not edit the class manually.
 */

namespace Swagger\Client\Model;

use \ArrayAccess;

/**
 * FilterListPhoneNumbersRegions Class Doc Comment
 *
 * @category    Class */
/** 
 * @package     Swagger\Client
 * @author      http://github.com/swagger-api/swagger-codegen
 * @license     http://www.apache.org/licenses/LICENSE-2.0 Apache Licene v2
 * @link        https://github.com/swagger-api/swagger-codegen
 */
class FilterListPhoneNumbersRegions implements ArrayAccess
{
    /**
      * The original name of the model.
      * @var string
      */
    protected static $swaggerModelName = 'FilterListPhoneNumbersRegions';

    /**
      * Array of property to type mappings. Used for (de)serialization
      * @var string[]
      */
    protected static $swaggerTypes = array(
        'country_code' => 'string',
        'npa' => 'int[]',
        'nxx' => 'string',
        'is_toll_free' => 'string',
        'city' => 'string',
        'province_postal_code' => 'string',
        'country_postal_code' => 'string'
    );

    public static function swaggerTypes()
    {
        return self::$swaggerTypes;
    }

    /**
     * Array of attributes where the key is the local name, and the value is the original name
     * @var string[]
     */
    protected static $attributeMap = array(
        'country_code' => 'country_code',
        'npa' => 'npa',
        'nxx' => 'nxx',
        'is_toll_free' => 'is_toll_free',
        'city' => 'city',
        'province_postal_code' => 'province_postal_code',
        'country_postal_code' => 'country_postal_code'
    );

    public static function attributeMap()
    {
        return self::$attributeMap;
    }

    /**
     * Array of attributes to setter functions (for deserialization of responses)
     * @var string[]
     */
    protected static $setters = array(
        'country_code' => 'setCountryCode',
        'npa' => 'setNpa',
        'nxx' => 'setNxx',
        'is_toll_free' => 'setIsTollFree',
        'city' => 'setCity',
        'province_postal_code' => 'setProvincePostalCode',
        'country_postal_code' => 'setCountryPostalCode'
    );

    public static function setters()
    {
        return self::$setters;
    }

    /**
     * Array of attributes to getter functions (for serialization of requests)
     * @var string[]
     */
    protected static $getters = array(
        'country_code' => 'getCountryCode',
        'npa' => 'getNpa',
        'nxx' => 'getNxx',
        'is_toll_free' => 'getIsTollFree',
        'city' => 'getCity',
        'province_postal_code' => 'getProvincePostalCode',
        'country_postal_code' => 'getCountryPostalCode'
    );

    public static function getters()
    {
        return self::$getters;
    }

    

    

    /**
     * Associative array for storing property values
     * @var mixed[]
     */
    protected $container = array();

    /**
     * Constructor
     * @param mixed[] $data Associated array of property value initalizing the model
     */
    public function __construct(array $data = null)
    {
        $this->container['country_code'] = isset($data['country_code']) ? $data['country_code'] : null;
        $this->container['npa'] = isset($data['npa']) ? $data['npa'] : null;
        $this->container['nxx'] = isset($data['nxx']) ? $data['nxx'] : null;
        $this->container['is_toll_free'] = isset($data['is_toll_free']) ? $data['is_toll_free'] : null;
        $this->container['city'] = isset($data['city']) ? $data['city'] : null;
        $this->container['province_postal_code'] = isset($data['province_postal_code']) ? $data['province_postal_code'] : null;
        $this->container['country_postal_code'] = isset($data['country_postal_code']) ? $data['country_postal_code'] : null;
    }

    /**
     * show all the invalid properties with reasons.
     *
     * @return array invalid properties with reasons
     */
    public function listInvalidProperties()
    {
        $invalid_properties = array();
        return $invalid_properties;
    }

    /**
     * validate all the properties in the model
     * return true if all passed
     *
     * @return bool True if all properteis are valid
     */
    public function valid()
    {
        return true;
    }


    /**
     * Gets country_code
     * @return string
     */
    public function getCountryCode()
    {
        return $this->container['country_code'];
    }

    /**
     * Sets country_code
     * @param string $country_code
     * @return $this
     */
    public function setCountryCode($country_code)
    {
        $this->container['country_code'] = $country_code;

        return $this;
    }

    /**
     * Gets npa
     * @return int[]
     */
    public function getNpa()
    {
        return $this->container['npa'];
    }

    /**
     * Sets npa
     * @param int[] $npa
     * @return $this
     */
    public function setNpa($npa)
    {
        $this->container['npa'] = $npa;

        return $this;
    }

    /**
     * Gets nxx
     * @return string
     */
    public function getNxx()
    {
        return $this->container['nxx'];
    }

    /**
     * Sets nxx
     * @param string $nxx
     * @return $this
     */
    public function setNxx($nxx)
    {
        $this->container['nxx'] = $nxx;

        return $this;
    }

    /**
     * Gets is_toll_free
     * @return string
     */
    public function getIsTollFree()
    {
        return $this->container['is_toll_free'];
    }

    /**
     * Sets is_toll_free
     * @param string $is_toll_free
     * @return $this
     */
    public function setIsTollFree($is_toll_free)
    {
        $this->container['is_toll_free'] = $is_toll_free;

        return $this;
    }

    /**
     * Gets city
     * @return string
     */
    public function getCity()
    {
        return $this->container['city'];
    }

    /**
     * Sets city
     * @param string $city
     * @return $this
     */
    public function setCity($city)
    {
        $this->container['city'] = $city;

        return $this;
    }

    /**
     * Gets province_postal_code
     * @return string
     */
    public function getProvincePostalCode()
    {
        return $this->container['province_postal_code'];
    }

    /**
     * Sets province_postal_code
     * @param string $province_postal_code
     * @return $this
     */
    public function setProvincePostalCode($province_postal_code)
    {
        $this->container['province_postal_code'] = $province_postal_code;

        return $this;
    }

    /**
     * Gets country_postal_code
     * @return string
     */
    public function getCountryPostalCode()
    {
        return $this->container['country_postal_code'];
    }

    /**
     * Sets country_postal_code
     * @param string $country_postal_code
     * @return $this
     */
    public function setCountryPostalCode($country_postal_code)
    {
        $this->container['country_postal_code'] = $country_postal_code;

        return $this;
    }
    /**
     * Returns true if offset exists. False otherwise.
     * @param  integer $offset Offset
     * @return boolean
     */
    public function offsetExists($offset)
    {
        return isset($this->container[$offset]);
    }

    /**
     * Gets offset.
     * @param  integer $offset Offset
     * @return mixed
     */
    public function offsetGet($offset)
    {
        return isset($this->container[$offset]) ? $this->container[$offset] : null;
    }

    /**
     * Sets value based on offset.
     * @param  integer $offset Offset
     * @param  mixed   $value  Value to be set
     * @return void
     */
    public function offsetSet($offset, $value)
    {
        if (is_null($offset)) {
            $this->container[] = $value;
        } else {
            $this->container[$offset] = $value;
        }
    }

    /**
     * Unsets offset.
     * @param  integer $offset Offset
     * @return void
     */
    public function offsetUnset($offset)
    {
        unset($this->container[$offset]);
    }

    /**
     * Gets the string presentation of the object
     * @return string
     */
    public function __toString()
    {
        if (defined('JSON_PRETTY_PRINT')) { // use JSON pretty print
            return json_encode(\Swagger\Client\ObjectSerializer::sanitizeForSerialization($this), JSON_PRETTY_PRINT);
        }

        return json_encode(\Swagger\Client\ObjectSerializer::sanitizeForSerialization($this));
    }
}

