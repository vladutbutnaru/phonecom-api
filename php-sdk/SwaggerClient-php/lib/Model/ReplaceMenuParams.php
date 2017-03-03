<?php
/**
 * ReplaceMenuParams
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
 * ReplaceMenuParams Class Doc Comment
 *
 * @category    Class */
/** 
 * @package     Swagger\Client
 * @author      http://github.com/swagger-api/swagger-codegen
 * @license     http://www.apache.org/licenses/LICENSE-2.0 Apache Licene v2
 * @link        https://github.com/swagger-api/swagger-codegen
 */
class ReplaceMenuParams implements ArrayAccess
{
    /**
      * The original name of the model.
      * @var string
      */
    protected static $swaggerModelName = 'ReplaceMenuParams';

    /**
      * Array of property to type mappings. Used for (de)serialization
      * @var string[]
      */
    protected static $swaggerTypes = array(
        'name' => 'string',
        'greeting' => 'object',
        'invalid_keypress' => 'object',
        'allow_extension_dial' => 'bool',
        'keypress_wait_time' => 'int',
        'timeout_handler' => 'object',
        'options' => 'object[]'
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
        'greeting' => 'greeting',
        'invalid_keypress' => 'invalid_keypress',
        'allow_extension_dial' => 'allow_extension_dial',
        'keypress_wait_time' => 'keypress_wait_time',
        'timeout_handler' => 'timeout_handler',
        'options' => 'options'
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
        'greeting' => 'setGreeting',
        'invalid_keypress' => 'setInvalidKeypress',
        'allow_extension_dial' => 'setAllowExtensionDial',
        'keypress_wait_time' => 'setKeypressWaitTime',
        'timeout_handler' => 'setTimeoutHandler',
        'options' => 'setOptions'
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
        'greeting' => 'getGreeting',
        'invalid_keypress' => 'getInvalidKeypress',
        'allow_extension_dial' => 'getAllowExtensionDial',
        'keypress_wait_time' => 'getKeypressWaitTime',
        'timeout_handler' => 'getTimeoutHandler',
        'options' => 'getOptions'
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
        $this->container['greeting'] = isset($data['greeting']) ? $data['greeting'] : null;
        $this->container['invalid_keypress'] = isset($data['invalid_keypress']) ? $data['invalid_keypress'] : null;
        $this->container['allow_extension_dial'] = isset($data['allow_extension_dial']) ? $data['allow_extension_dial'] : null;
        $this->container['keypress_wait_time'] = isset($data['keypress_wait_time']) ? $data['keypress_wait_time'] : null;
        $this->container['timeout_handler'] = isset($data['timeout_handler']) ? $data['timeout_handler'] : null;
        $this->container['options'] = isset($data['options']) ? $data['options'] : null;
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
     * @param string $name
     * @return $this
     */
    public function setName($name)
    {
        $this->container['name'] = $name;

        return $this;
    }

    /**
     * Gets greeting
     * @return object
     */
    public function getGreeting()
    {
        return $this->container['greeting'];
    }

    /**
     * Sets greeting
     * @param object $greeting
     * @return $this
     */
    public function setGreeting($greeting)
    {
        $this->container['greeting'] = $greeting;

        return $this;
    }

    /**
     * Gets invalid_keypress
     * @return object
     */
    public function getInvalidKeypress()
    {
        return $this->container['invalid_keypress'];
    }

    /**
     * Sets invalid_keypress
     * @param object $invalid_keypress
     * @return $this
     */
    public function setInvalidKeypress($invalid_keypress)
    {
        $this->container['invalid_keypress'] = $invalid_keypress;

        return $this;
    }

    /**
     * Gets allow_extension_dial
     * @return bool
     */
    public function getAllowExtensionDial()
    {
        return $this->container['allow_extension_dial'];
    }

    /**
     * Sets allow_extension_dial
     * @param bool $allow_extension_dial
     * @return $this
     */
    public function setAllowExtensionDial($allow_extension_dial)
    {
        $this->container['allow_extension_dial'] = $allow_extension_dial;

        return $this;
    }

    /**
     * Gets keypress_wait_time
     * @return int
     */
    public function getKeypressWaitTime()
    {
        return $this->container['keypress_wait_time'];
    }

    /**
     * Sets keypress_wait_time
     * @param int $keypress_wait_time
     * @return $this
     */
    public function setKeypressWaitTime($keypress_wait_time)
    {
        $this->container['keypress_wait_time'] = $keypress_wait_time;

        return $this;
    }

    /**
     * Gets timeout_handler
     * @return object
     */
    public function getTimeoutHandler()
    {
        return $this->container['timeout_handler'];
    }

    /**
     * Sets timeout_handler
     * @param object $timeout_handler
     * @return $this
     */
    public function setTimeoutHandler($timeout_handler)
    {
        $this->container['timeout_handler'] = $timeout_handler;

        return $this;
    }

    /**
     * Gets options
     * @return object[]
     */
    public function getOptions()
    {
        return $this->container['options'];
    }

    /**
     * Sets options
     * @param object[] $options
     * @return $this
     */
    public function setOptions($options)
    {
        $this->container['options'] = $options;

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


