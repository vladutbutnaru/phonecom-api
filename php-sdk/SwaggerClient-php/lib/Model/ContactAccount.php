<?php
/**
 * ContactAccount
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
 * ContactAccount Class Doc Comment
 *
 * @category    Class */
/** 
 * @package     Swagger\Client
 * @author      http://github.com/swagger-api/swagger-codegen
 * @license     http://www.apache.org/licenses/LICENSE-2.0 Apache Licene v2
 * @link        https://github.com/swagger-api/swagger-codegen
 */
class ContactAccount implements ArrayAccess
{
    /**
      * The original name of the model.
      * @var string
      */
    protected static $swaggerModelName = 'ContactAccount';

    /**
      * Array of property to type mappings. Used for (de)serialization
      * @var string[]
      */
    protected static $swaggerTypes = array(
        'name' => 'string',
        'company' => 'string',
        'address' => '\Swagger\Client\Model\Address',
        'phone' => 'string',
        'fax' => 'string',
        'primary_email' => 'string',
        'alternate_email' => 'string'
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
        'name' => 'name',
        'company' => 'company',
        'address' => 'address',
        'phone' => 'phone',
        'fax' => 'fax',
        'primary_email' => 'primary_email',
        'alternate_email' => 'alternate_email'
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
        'name' => 'setName',
        'company' => 'setCompany',
        'address' => 'setAddress',
        'phone' => 'setPhone',
        'fax' => 'setFax',
        'primary_email' => 'setPrimaryEmail',
        'alternate_email' => 'setAlternateEmail'
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
        'name' => 'getName',
        'company' => 'getCompany',
        'address' => 'getAddress',
        'phone' => 'getPhone',
        'fax' => 'getFax',
        'primary_email' => 'getPrimaryEmail',
        'alternate_email' => 'getAlternateEmail'
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
        $this->container['name'] = isset($data['name']) ? $data['name'] : null;
        $this->container['company'] = isset($data['company']) ? $data['company'] : null;
        $this->container['address'] = isset($data['address']) ? $data['address'] : null;
        $this->container['phone'] = isset($data['phone']) ? $data['phone'] : null;
        $this->container['fax'] = isset($data['fax']) ? $data['fax'] : null;
        $this->container['primary_email'] = isset($data['primary_email']) ? $data['primary_email'] : null;
        $this->container['alternate_email'] = isset($data['alternate_email']) ? $data['alternate_email'] : null;
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
     * Gets name
     * @return string
     */
    public function getName()
    {
        return $this->container['name'];
    }

    /**
     * Sets name
     * @param string $name Name. Required.
     * @return $this
     */
    public function setName($name)
    {
        $this->container['name'] = $name;

        return $this;
    }

    /**
     * Gets company
     * @return string
     */
    public function getCompany()
    {
        return $this->container['company'];
    }

    /**
     * Sets company
     * @param string $company Company name
     * @return $this
     */
    public function setCompany($company)
    {
        $this->container['company'] = $company;

        return $this;
    }

    /**
     * Gets address
     * @return \Swagger\Client\Model\Address
     */
    public function getAddress()
    {
        return $this->container['address'];
    }

    /**
     * Sets address
     * @param \Swagger\Client\Model\Address $address
     * @return $this
     */
    public function setAddress($address)
    {
        $this->container['address'] = $address;

        return $this;
    }

    /**
     * Gets phone
     * @return string
     */
    public function getPhone()
    {
        return $this->container['phone'];
    }

    /**
     * Sets phone
     * @param string $phone Phone number. Required.
     * @return $this
     */
    public function setPhone($phone)
    {
        $this->container['phone'] = $phone;

        return $this;
    }

    /**
     * Gets fax
     * @return string
     */
    public function getFax()
    {
        return $this->container['fax'];
    }

    /**
     * Sets fax
     * @param string $fax Fax number
     * @return $this
     */
    public function setFax($fax)
    {
        $this->container['fax'] = $fax;

        return $this;
    }

    /**
     * Gets primary_email
     * @return string
     */
    public function getPrimaryEmail()
    {
        return $this->container['primary_email'];
    }

    /**
     * Sets primary_email
     * @param string $primary_email Primary email address. Required.
     * @return $this
     */
    public function setPrimaryEmail($primary_email)
    {
        $this->container['primary_email'] = $primary_email;

        return $this;
    }

    /**
     * Gets alternate_email
     * @return string
     */
    public function getAlternateEmail()
    {
        return $this->container['alternate_email'];
    }

    /**
     * Sets alternate_email
     * @param string $alternate_email Alternate email address
     * @return $this
     */
    public function setAlternateEmail($alternate_email)
    {
        $this->container['alternate_email'] = $alternate_email;

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

