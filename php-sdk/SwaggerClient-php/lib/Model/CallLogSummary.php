<?php
/**
 * CallLogSummary
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
 * CallLogSummary Class Doc Comment
 *
 * @category    Class */
 // @description The Call Log Summary Object is used to briefly represent a call. It consists of properties as below:
/** 
 * @package     Swagger\Client
 * @author      http://github.com/swagger-api/swagger-codegen
 * @license     http://www.apache.org/licenses/LICENSE-2.0 Apache Licene v2
 * @link        https://github.com/swagger-api/swagger-codegen
 */
class CallLogSummary implements ArrayAccess
{
    /**
      * The original name of the model.
      * @var string
      */
    protected static $swaggerModelName = 'CallLogSummary';

    /**
      * Array of property to type mappings. Used for (de)serialization
      * @var string[]
      */
    protected static $swaggerTypes = array(
        'id' => 'string',
        'uuid' => 'string',
        'extension' => '\Swagger\Client\Model\ExtensionSummary',
        'caller_id' => 'string',
        'called_number' => 'string',
        'start_time' => 'string',
        'created_at' => 'string',
        'direction' => 'string',
        'type' => 'string',
        'call_duration' => 'int',
        'is_monitored' => 'string',
        'call_number' => 'string',
        'final_action' => 'string',
        'call_recording' => 'string'
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
        'id' => 'id',
        'uuid' => 'uuid',
        'extension' => 'extension',
        'caller_id' => 'caller_id',
        'called_number' => 'called_number',
        'start_time' => 'start_time',
        'created_at' => 'created_at',
        'direction' => 'direction',
        'type' => 'type',
        'call_duration' => 'call_duration',
        'is_monitored' => 'is_monitored',
        'call_number' => 'call_number',
        'final_action' => 'final_action',
        'call_recording' => 'call_recording'
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
        'id' => 'setId',
        'uuid' => 'setUuid',
        'extension' => 'setExtension',
        'caller_id' => 'setCallerId',
        'called_number' => 'setCalledNumber',
        'start_time' => 'setStartTime',
        'created_at' => 'setCreatedAt',
        'direction' => 'setDirection',
        'type' => 'setType',
        'call_duration' => 'setCallDuration',
        'is_monitored' => 'setIsMonitored',
        'call_number' => 'setCallNumber',
        'final_action' => 'setFinalAction',
        'call_recording' => 'setCallRecording'
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
        'id' => 'getId',
        'uuid' => 'getUuid',
        'extension' => 'getExtension',
        'caller_id' => 'getCallerId',
        'called_number' => 'getCalledNumber',
        'start_time' => 'getStartTime',
        'created_at' => 'getCreatedAt',
        'direction' => 'getDirection',
        'type' => 'getType',
        'call_duration' => 'getCallDuration',
        'is_monitored' => 'getIsMonitored',
        'call_number' => 'getCallNumber',
        'final_action' => 'getFinalAction',
        'call_recording' => 'getCallRecording'
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
        $this->container['id'] = isset($data['id']) ? $data['id'] : null;
        $this->container['uuid'] = isset($data['uuid']) ? $data['uuid'] : null;
        $this->container['extension'] = isset($data['extension']) ? $data['extension'] : null;
        $this->container['caller_id'] = isset($data['caller_id']) ? $data['caller_id'] : null;
        $this->container['called_number'] = isset($data['called_number']) ? $data['called_number'] : null;
        $this->container['start_time'] = isset($data['start_time']) ? $data['start_time'] : null;
        $this->container['created_at'] = isset($data['created_at']) ? $data['created_at'] : null;
        $this->container['direction'] = isset($data['direction']) ? $data['direction'] : null;
        $this->container['type'] = isset($data['type']) ? $data['type'] : null;
        $this->container['call_duration'] = isset($data['call_duration']) ? $data['call_duration'] : null;
        $this->container['is_monitored'] = isset($data['is_monitored']) ? $data['is_monitored'] : null;
        $this->container['call_number'] = isset($data['call_number']) ? $data['call_number'] : null;
        $this->container['final_action'] = isset($data['final_action']) ? $data['final_action'] : null;
        $this->container['call_recording'] = isset($data['call_recording']) ? $data['call_recording'] : null;
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
     * Gets id
     * @return string
     */
    public function getId()
    {
        return $this->container['id'];
    }

    /**
     * Sets id
     * @param string $id ID
     * @return $this
     */
    public function setId($id)
    {
        $this->container['id'] = $id;

        return $this;
    }

    /**
     * Gets uuid
     * @return string
     */
    public function getUuid()
    {
        return $this->container['uuid'];
    }

    /**
     * Sets uuid
     * @param string $uuid Internal system id, may be null
     * @return $this
     */
    public function setUuid($uuid)
    {
        $this->container['uuid'] = $uuid;

        return $this;
    }

    /**
     * Gets extension
     * @return \Swagger\Client\Model\ExtensionSummary
     */
    public function getExtension()
    {
        return $this->container['extension'];
    }

    /**
     * Sets extension
     * @param \Swagger\Client\Model\ExtensionSummary $extension Account extension
     * @return $this
     */
    public function setExtension($extension)
    {
        $this->container['extension'] = $extension;

        return $this;
    }

    /**
     * Gets caller_id
     * @return string
     */
    public function getCallerId()
    {
        return $this->container['caller_id'];
    }

    /**
     * Sets caller_id
     * @param string $caller_id Call made from this phone number
     * @return $this
     */
    public function setCallerId($caller_id)
    {
        $this->container['caller_id'] = $caller_id;

        return $this;
    }

    /**
     * Gets called_number
     * @return string
     */
    public function getCalledNumber()
    {
        return $this->container['called_number'];
    }

    /**
     * Sets called_number
     * @param string $called_number Call made to this phone number
     * @return $this
     */
    public function setCalledNumber($called_number)
    {
        $this->container['called_number'] = $called_number;

        return $this;
    }

    /**
     * Gets start_time
     * @return string
     */
    public function getStartTime()
    {
        return $this->container['start_time'];
    }

    /**
     * Sets start_time
     * @param string $start_time Call start time
     * @return $this
     */
    public function setStartTime($start_time)
    {
        $this->container['start_time'] = $start_time;

        return $this;
    }

    /**
     * Gets created_at
     * @return string
     */
    public function getCreatedAt()
    {
        return $this->container['created_at'];
    }

    /**
     * Sets created_at
     * @param string $created_at Call log creation time. Same as call end time + time needed to create call log
     * @return $this
     */
    public function setCreatedAt($created_at)
    {
        $this->container['created_at'] = $created_at;

        return $this;
    }

    /**
     * Gets direction
     * @return string
     */
    public function getDirection()
    {
        return $this->container['direction'];
    }

    /**
     * Sets direction
     * @param string $direction Call direction: in or out
     * @return $this
     */
    public function setDirection($direction)
    {
        $this->container['direction'] = $direction;

        return $this;
    }

    /**
     * Gets type
     * @return string
     */
    public function getType()
    {
        return $this->container['type'];
    }

    /**
     * Sets type
     * @param string $type Call type: call, fax, audiogram ...
     * @return $this
     */
    public function setType($type)
    {
        $this->container['type'] = $type;

        return $this;
    }

    /**
     * Gets call_duration
     * @return int
     */
    public function getCallDuration()
    {
        return $this->container['call_duration'];
    }

    /**
     * Sets call_duration
     * @param int $call_duration Call duration in seconds
     * @return $this
     */
    public function setCallDuration($call_duration)
    {
        $this->container['call_duration'] = $call_duration;

        return $this;
    }

    /**
     * Gets is_monitored
     * @return string
     */
    public function getIsMonitored()
    {
        return $this->container['is_monitored'];
    }

    /**
     * Sets is_monitored
     * @param string $is_monitored Was call being monitored?
     * @return $this
     */
    public function setIsMonitored($is_monitored)
    {
        $this->container['is_monitored'] = $is_monitored;

        return $this;
    }

    /**
     * Gets call_number
     * @return string
     */
    public function getCallNumber()
    {
        return $this->container['call_number'];
    }

    /**
     * Sets call_number
     * @param string $call_number Internal system call reference number
     * @return $this
     */
    public function setCallNumber($call_number)
    {
        $this->container['call_number'] = $call_number;

        return $this;
    }

    /**
     * Gets final_action
     * @return string
     */
    public function getFinalAction()
    {
        return $this->container['final_action'];
    }

    /**
     * Sets final_action
     * @param string $final_action Last action of call flow
     * @return $this
     */
    public function setFinalAction($final_action)
    {
        $this->container['final_action'] = $final_action;

        return $this;
    }

    /**
     * Gets call_recording
     * @return string
     */
    public function getCallRecording()
    {
        return $this->container['call_recording'];
    }

    /**
     * Sets call_recording
     * @param string $call_recording URL of call recording if available. Empty string if call recording does not exist
     * @return $this
     */
    public function setCallRecording($call_recording)
    {
        $this->container['call_recording'] = $call_recording;

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


