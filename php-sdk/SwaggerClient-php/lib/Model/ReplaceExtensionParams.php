<?php
/**
 * ReplaceExtensionParams
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
 * ReplaceExtensionParams Class Doc Comment
 *
 * @category    Class */
/** 
 * @package     Swagger\Client
 * @author      http://github.com/swagger-api/swagger-codegen
 * @license     http://www.apache.org/licenses/LICENSE-2.0 Apache Licene v2
 * @link        https://github.com/swagger-api/swagger-codegen
 */
class ReplaceExtensionParams implements ArrayAccess
{
    /**
      * The original name of the model.
      * @var string
      */
    protected static $swaggerModelName = 'ReplaceExtensionParams';

    /**
      * Array of property to type mappings. Used for (de)serialization
      * @var string[]
      */
    protected static $swaggerTypes = array(
        'voicemail_greeting_alternate' => 'object',
        'name_greeting' => 'object',
        'name' => 'string',
        'timezone' => 'string',
        'include_in_directory' => 'bool',
        'extension' => 'int',
        'enable_outbound_calls' => 'bool',
        'usage_type' => 'string',
        'voicemail_password' => 'int',
        'full_name' => 'string',
        'enable_call_waiting' => 'bool',
        'voicemail_greeting_standard' => 'object',
        'voicemail_greeting_type' => 'string',
        'caller_id' => 'string',
        'local_area_code' => 'int',
        'voicemail_enabled' => 'bool',
        'voicemail_greeting_enable_leave_message_prompt' => 'bool',
        'voicemail_transcription' => 'string',
        'voicemail_notifications_emails' => 'string[]',
        'voicemail_notifications_sms' => 'string',
        'call_notifications_emails' => 'string[]',
        'call_notifications_sms' => 'string',
        'route' => 'string[]'
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
        'voicemail_greeting_alternate' => 'voicemail[greeting][alternate]',
        'name_greeting' => 'name_greeting',
        'name' => 'name',
        'timezone' => 'timezone',
        'include_in_directory' => 'include_in_directory',
        'extension' => 'extension',
        'enable_outbound_calls' => 'enable_outbound_calls',
        'usage_type' => 'usage_type',
        'voicemail_password' => 'voicemail[password]',
        'full_name' => 'full_name',
        'enable_call_waiting' => 'enable_call_waiting',
        'voicemail_greeting_standard' => 'voicemail[greeting][standard]',
        'voicemail_greeting_type' => 'voicemail[greeting][type]',
        'caller_id' => 'caller_id',
        'local_area_code' => 'local_area_code',
        'voicemail_enabled' => 'voicemail[enabled]',
        'voicemail_greeting_enable_leave_message_prompt' => 'voicemail[greeting][enable_leave_message_prompt]',
        'voicemail_transcription' => 'voicemail[transcription]',
        'voicemail_notifications_emails' => 'voicemail[notifications][emails]',
        'voicemail_notifications_sms' => 'voicemail[notifications][sms]',
        'call_notifications_emails' => 'call_notifications[emails]',
        'call_notifications_sms' => 'call_notifications[sms]',
        'route' => 'route'
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
        'voicemail_greeting_alternate' => 'setVoicemailGreetingAlternate',
        'name_greeting' => 'setNameGreeting',
        'name' => 'setName',
        'timezone' => 'setTimezone',
        'include_in_directory' => 'setIncludeInDirectory',
        'extension' => 'setExtension',
        'enable_outbound_calls' => 'setEnableOutboundCalls',
        'usage_type' => 'setUsageType',
        'voicemail_password' => 'setVoicemailPassword',
        'full_name' => 'setFullName',
        'enable_call_waiting' => 'setEnableCallWaiting',
        'voicemail_greeting_standard' => 'setVoicemailGreetingStandard',
        'voicemail_greeting_type' => 'setVoicemailGreetingType',
        'caller_id' => 'setCallerId',
        'local_area_code' => 'setLocalAreaCode',
        'voicemail_enabled' => 'setVoicemailEnabled',
        'voicemail_greeting_enable_leave_message_prompt' => 'setVoicemailGreetingEnableLeaveMessagePrompt',
        'voicemail_transcription' => 'setVoicemailTranscription',
        'voicemail_notifications_emails' => 'setVoicemailNotificationsEmails',
        'voicemail_notifications_sms' => 'setVoicemailNotificationsSms',
        'call_notifications_emails' => 'setCallNotificationsEmails',
        'call_notifications_sms' => 'setCallNotificationsSms',
        'route' => 'setRoute'
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
        'voicemail_greeting_alternate' => 'getVoicemailGreetingAlternate',
        'name_greeting' => 'getNameGreeting',
        'name' => 'getName',
        'timezone' => 'getTimezone',
        'include_in_directory' => 'getIncludeInDirectory',
        'extension' => 'getExtension',
        'enable_outbound_calls' => 'getEnableOutboundCalls',
        'usage_type' => 'getUsageType',
        'voicemail_password' => 'getVoicemailPassword',
        'full_name' => 'getFullName',
        'enable_call_waiting' => 'getEnableCallWaiting',
        'voicemail_greeting_standard' => 'getVoicemailGreetingStandard',
        'voicemail_greeting_type' => 'getVoicemailGreetingType',
        'caller_id' => 'getCallerId',
        'local_area_code' => 'getLocalAreaCode',
        'voicemail_enabled' => 'getVoicemailEnabled',
        'voicemail_greeting_enable_leave_message_prompt' => 'getVoicemailGreetingEnableLeaveMessagePrompt',
        'voicemail_transcription' => 'getVoicemailTranscription',
        'voicemail_notifications_emails' => 'getVoicemailNotificationsEmails',
        'voicemail_notifications_sms' => 'getVoicemailNotificationsSms',
        'call_notifications_emails' => 'getCallNotificationsEmails',
        'call_notifications_sms' => 'getCallNotificationsSms',
        'route' => 'getRoute'
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
        $this->container['voicemail_greeting_alternate'] = isset($data['voicemail_greeting_alternate']) ? $data['voicemail_greeting_alternate'] : null;
        $this->container['name_greeting'] = isset($data['name_greeting']) ? $data['name_greeting'] : null;
        $this->container['name'] = isset($data['name']) ? $data['name'] : null;
        $this->container['timezone'] = isset($data['timezone']) ? $data['timezone'] : null;
        $this->container['include_in_directory'] = isset($data['include_in_directory']) ? $data['include_in_directory'] : null;
        $this->container['extension'] = isset($data['extension']) ? $data['extension'] : null;
        $this->container['enable_outbound_calls'] = isset($data['enable_outbound_calls']) ? $data['enable_outbound_calls'] : null;
        $this->container['usage_type'] = isset($data['usage_type']) ? $data['usage_type'] : null;
        $this->container['voicemail_password'] = isset($data['voicemail_password']) ? $data['voicemail_password'] : null;
        $this->container['full_name'] = isset($data['full_name']) ? $data['full_name'] : null;
        $this->container['enable_call_waiting'] = isset($data['enable_call_waiting']) ? $data['enable_call_waiting'] : null;
        $this->container['voicemail_greeting_standard'] = isset($data['voicemail_greeting_standard']) ? $data['voicemail_greeting_standard'] : null;
        $this->container['voicemail_greeting_type'] = isset($data['voicemail_greeting_type']) ? $data['voicemail_greeting_type'] : null;
        $this->container['caller_id'] = isset($data['caller_id']) ? $data['caller_id'] : null;
        $this->container['local_area_code'] = isset($data['local_area_code']) ? $data['local_area_code'] : null;
        $this->container['voicemail_enabled'] = isset($data['voicemail_enabled']) ? $data['voicemail_enabled'] : null;
        $this->container['voicemail_greeting_enable_leave_message_prompt'] = isset($data['voicemail_greeting_enable_leave_message_prompt']) ? $data['voicemail_greeting_enable_leave_message_prompt'] : null;
        $this->container['voicemail_transcription'] = isset($data['voicemail_transcription']) ? $data['voicemail_transcription'] : null;
        $this->container['voicemail_notifications_emails'] = isset($data['voicemail_notifications_emails']) ? $data['voicemail_notifications_emails'] : null;
        $this->container['voicemail_notifications_sms'] = isset($data['voicemail_notifications_sms']) ? $data['voicemail_notifications_sms'] : null;
        $this->container['call_notifications_emails'] = isset($data['call_notifications_emails']) ? $data['call_notifications_emails'] : null;
        $this->container['call_notifications_sms'] = isset($data['call_notifications_sms']) ? $data['call_notifications_sms'] : null;
        $this->container['route'] = isset($data['route']) ? $data['route'] : null;
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
     * Gets voicemail_greeting_alternate
     * @return object
     */
    public function getVoicemailGreetingAlternate()
    {
        return $this->container['voicemail_greeting_alternate'];
    }

    /**
     * Sets voicemail_greeting_alternate
     * @param object $voicemail_greeting_alternate Recording lookup object
     * @return $this
     */
    public function setVoicemailGreetingAlternate($voicemail_greeting_alternate)
    {
        $this->container['voicemail_greeting_alternate'] = $voicemail_greeting_alternate;

        return $this;
    }

    /**
     * Gets name_greeting
     * @return object
     */
    public function getNameGreeting()
    {
        return $this->container['name_greeting'];
    }

    /**
     * Sets name_greeting
     * @param object $name_greeting Recording lookup object
     * @return $this
     */
    public function setNameGreeting($name_greeting)
    {
        $this->container['name_greeting'] = $name_greeting;

        return $this;
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
     * @param string $name Name (required)
     * @return $this
     */
    public function setName($name)
    {
        $this->container['name'] = $name;

        return $this;
    }

    /**
     * Gets timezone
     * @return string
     */
    public function getTimezone()
    {
        return $this->container['timezone'];
    }

    /**
     * Sets timezone
     * @param string $timezone Timezone
     * @return $this
     */
    public function setTimezone($timezone)
    {
        $this->container['timezone'] = $timezone;

        return $this;
    }

    /**
     * Gets include_in_directory
     * @return bool
     */
    public function getIncludeInDirectory()
    {
        return $this->container['include_in_directory'];
    }

    /**
     * Sets include_in_directory
     * @param bool $include_in_directory Include in dial-by-name directory
     * @return $this
     */
    public function setIncludeInDirectory($include_in_directory)
    {
        $this->container['include_in_directory'] = $include_in_directory;

        return $this;
    }

    /**
     * Gets extension
     * @return int
     */
    public function getExtension()
    {
        return $this->container['extension'];
    }

    /**
     * Sets extension
     * @param int $extension Extension number (required)
     * @return $this
     */
    public function setExtension($extension)
    {
        $this->container['extension'] = $extension;

        return $this;
    }

    /**
     * Gets enable_outbound_calls
     * @return bool
     */
    public function getEnableOutboundCalls()
    {
        return $this->container['enable_outbound_calls'];
    }

    /**
     * Sets enable_outbound_calls
     * @param bool $enable_outbound_calls Enable outgoing calls
     * @return $this
     */
    public function setEnableOutboundCalls($enable_outbound_calls)
    {
        $this->container['enable_outbound_calls'] = $enable_outbound_calls;

        return $this;
    }

    /**
     * Gets usage_type
     * @return string
     */
    public function getUsageType()
    {
        return $this->container['usage_type'];
    }

    /**
     * Sets usage_type
     * @param string $usage_type Extension type
     * @return $this
     */
    public function setUsageType($usage_type)
    {
        $this->container['usage_type'] = $usage_type;

        return $this;
    }

    /**
     * Gets voicemail_password
     * @return int
     */
    public function getVoicemailPassword()
    {
        return $this->container['voicemail_password'];
    }

    /**
     * Sets voicemail_password
     * @param int $voicemail_password Voicemail password
     * @return $this
     */
    public function setVoicemailPassword($voicemail_password)
    {
        $this->container['voicemail_password'] = $voicemail_password;

        return $this;
    }

    /**
     * Gets full_name
     * @return string
     */
    public function getFullName()
    {
        return $this->container['full_name'];
    }

    /**
     * Sets full_name
     * @param string $full_name Contact name
     * @return $this
     */
    public function setFullName($full_name)
    {
        $this->container['full_name'] = $full_name;

        return $this;
    }

    /**
     * Gets enable_call_waiting
     * @return bool
     */
    public function getEnableCallWaiting()
    {
        return $this->container['enable_call_waiting'];
    }

    /**
     * Sets enable_call_waiting
     * @param bool $enable_call_waiting Enable Call Waiting
     * @return $this
     */
    public function setEnableCallWaiting($enable_call_waiting)
    {
        $this->container['enable_call_waiting'] = $enable_call_waiting;

        return $this;
    }

    /**
     * Gets voicemail_greeting_standard
     * @return object
     */
    public function getVoicemailGreetingStandard()
    {
        return $this->container['voicemail_greeting_standard'];
    }

    /**
     * Sets voicemail_greeting_standard
     * @param object $voicemail_greeting_standard Recording lookup object
     * @return $this
     */
    public function setVoicemailGreetingStandard($voicemail_greeting_standard)
    {
        $this->container['voicemail_greeting_standard'] = $voicemail_greeting_standard;

        return $this;
    }

    /**
     * Gets voicemail_greeting_type
     * @return string
     */
    public function getVoicemailGreetingType()
    {
        return $this->container['voicemail_greeting_type'];
    }

    /**
     * Sets voicemail_greeting_type
     * @param string $voicemail_greeting_type Voicemail greeting type
     * @return $this
     */
    public function setVoicemailGreetingType($voicemail_greeting_type)
    {
        $this->container['voicemail_greeting_type'] = $voicemail_greeting_type;

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
     * @param string $caller_id Caller ID
     * @return $this
     */
    public function setCallerId($caller_id)
    {
        $this->container['caller_id'] = $caller_id;

        return $this;
    }

    /**
     * Gets local_area_code
     * @return int
     */
    public function getLocalAreaCode()
    {
        return $this->container['local_area_code'];
    }

    /**
     * Sets local_area_code
     * @param int $local_area_code Local area code
     * @return $this
     */
    public function setLocalAreaCode($local_area_code)
    {
        $this->container['local_area_code'] = $local_area_code;

        return $this;
    }

    /**
     * Gets voicemail_enabled
     * @return bool
     */
    public function getVoicemailEnabled()
    {
        return $this->container['voicemail_enabled'];
    }

    /**
     * Sets voicemail_enabled
     * @param bool $voicemail_enabled Voicemail enabled
     * @return $this
     */
    public function setVoicemailEnabled($voicemail_enabled)
    {
        $this->container['voicemail_enabled'] = $voicemail_enabled;

        return $this;
    }

    /**
     * Gets voicemail_greeting_enable_leave_message_prompt
     * @return bool
     */
    public function getVoicemailGreetingEnableLeaveMessagePrompt()
    {
        return $this->container['voicemail_greeting_enable_leave_message_prompt'];
    }

    /**
     * Sets voicemail_greeting_enable_leave_message_prompt
     * @param bool $voicemail_greeting_enable_leave_message_prompt Use leave message prompt after voicemail
     * @return $this
     */
    public function setVoicemailGreetingEnableLeaveMessagePrompt($voicemail_greeting_enable_leave_message_prompt)
    {
        $this->container['voicemail_greeting_enable_leave_message_prompt'] = $voicemail_greeting_enable_leave_message_prompt;

        return $this;
    }

    /**
     * Gets voicemail_transcription
     * @return string
     */
    public function getVoicemailTranscription()
    {
        return $this->container['voicemail_transcription'];
    }

    /**
     * Sets voicemail_transcription
     * @param string $voicemail_transcription Voicemail transcription type
     * @return $this
     */
    public function setVoicemailTranscription($voicemail_transcription)
    {
        $this->container['voicemail_transcription'] = $voicemail_transcription;

        return $this;
    }

    /**
     * Gets voicemail_notifications_emails
     * @return string[]
     */
    public function getVoicemailNotificationsEmails()
    {
        return $this->container['voicemail_notifications_emails'];
    }

    /**
     * Sets voicemail_notifications_emails
     * @param string[] $voicemail_notifications_emails Email notifications for voicemails. Can be a single email or an array of emails
     * @return $this
     */
    public function setVoicemailNotificationsEmails($voicemail_notifications_emails)
    {
        $this->container['voicemail_notifications_emails'] = $voicemail_notifications_emails;

        return $this;
    }

    /**
     * Gets voicemail_notifications_sms
     * @return string
     */
    public function getVoicemailNotificationsSms()
    {
        return $this->container['voicemail_notifications_sms'];
    }

    /**
     * Sets voicemail_notifications_sms
     * @param string $voicemail_notifications_sms SMS notifications for voicemails
     * @return $this
     */
    public function setVoicemailNotificationsSms($voicemail_notifications_sms)
    {
        $this->container['voicemail_notifications_sms'] = $voicemail_notifications_sms;

        return $this;
    }

    /**
     * Gets call_notifications_emails
     * @return string[]
     */
    public function getCallNotificationsEmails()
    {
        return $this->container['call_notifications_emails'];
    }

    /**
     * Sets call_notifications_emails
     * @param string[] $call_notifications_emails Email notifications for calls. Can be a single email or an array of emails
     * @return $this
     */
    public function setCallNotificationsEmails($call_notifications_emails)
    {
        $this->container['call_notifications_emails'] = $call_notifications_emails;

        return $this;
    }

    /**
     * Gets call_notifications_sms
     * @return string
     */
    public function getCallNotificationsSms()
    {
        return $this->container['call_notifications_sms'];
    }

    /**
     * Sets call_notifications_sms
     * @param string $call_notifications_sms SMS notifications for calls
     * @return $this
     */
    public function setCallNotificationsSms($call_notifications_sms)
    {
        $this->container['call_notifications_sms'] = $call_notifications_sms;

        return $this;
    }

    /**
     * Gets route
     * @return string[]
     */
    public function getRoute()
    {
        return $this->container['route'];
    }

    /**
     * Sets route
     * @param string[] $route Route object lookup (must belong to this extension)
     * @return $this
     */
    public function setRoute($route)
    {
        $this->container['route'] = $route;

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

