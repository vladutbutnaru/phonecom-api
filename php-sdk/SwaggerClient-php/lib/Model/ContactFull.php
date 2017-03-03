<?php
/**
 * ContactFull
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
 * ContactFull Class Doc Comment
 *
 * @category    Class */
 // @description The Full Contact Object includes all of the properties in the Contact Summary Object, along with several more:
/** 
 * @package     Swagger\Client
 * @author      http://github.com/swagger-api/swagger-codegen
 * @license     http://www.apache.org/licenses/LICENSE-2.0 Apache Licene v2
 * @link        https://github.com/swagger-api/swagger-codegen
 */
class ContactFull implements ArrayAccess
{
    /**
      * The original name of the model.
      * @var string
      */
    protected static $swaggerModelName = 'ContactFull';

    /**
      * Array of property to type mappings. Used for (de)serialization
      * @var string[]
      */
    protected static $swaggerTypes = array(
        'id' => 'int',
        'prefix' => 'string',
        'first_name' => 'string',
        'middle_name' => 'string',
        'last_name' => 'string',
        'suffix' => 'string',
        'nickname' => 'string',
        'company' => 'string',
        'phonetic_first_name' => 'string',
        'phonetic_middle_name' => 'string',
        'phonetic_last_name' => 'string',
        'department' => 'string',
        'job_title' => 'string',
        'emails' => '\Swagger\Client\Model\Emails',
        'phone_numbers' => '\Swagger\Client\Model\PhoneNumberContact[]',
        'addresses' => '\Swagger\Client\Model\Addresses',
        'group' => '\Swagger\Client\Model\GroupListContacts',
        'created_at' => 'int',
        'updated_at' => 'int'
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
        'prefix' => 'prefix',
        'first_name' => 'first_name',
        'middle_name' => 'middle_name',
        'last_name' => 'last_name',
        'suffix' => 'suffix',
        'nickname' => 'nickname',
        'company' => 'company',
        'phonetic_first_name' => 'phonetic_first_name',
        'phonetic_middle_name' => 'phonetic_middle_name',
        'phonetic_last_name' => 'phonetic_last_name',
        'department' => 'department',
        'job_title' => 'job_title',
        'emails' => 'emails',
        'phone_numbers' => 'phone_numbers',
        'addresses' => 'addresses',
        'group' => 'group',
        'created_at' => 'created_at',
        'updated_at' => 'updated_at'
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
        'prefix' => 'setPrefix',
        'first_name' => 'setFirstName',
        'middle_name' => 'setMiddleName',
        'last_name' => 'setLastName',
        'suffix' => 'setSuffix',
        'nickname' => 'setNickname',
        'company' => 'setCompany',
        'phonetic_first_name' => 'setPhoneticFirstName',
        'phonetic_middle_name' => 'setPhoneticMiddleName',
        'phonetic_last_name' => 'setPhoneticLastName',
        'department' => 'setDepartment',
        'job_title' => 'setJobTitle',
        'emails' => 'setEmails',
        'phone_numbers' => 'setPhoneNumbers',
        'addresses' => 'setAddresses',
        'group' => 'setGroup',
        'created_at' => 'setCreatedAt',
        'updated_at' => 'setUpdatedAt'
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
        'prefix' => 'getPrefix',
        'first_name' => 'getFirstName',
        'middle_name' => 'getMiddleName',
        'last_name' => 'getLastName',
        'suffix' => 'getSuffix',
        'nickname' => 'getNickname',
        'company' => 'getCompany',
        'phonetic_first_name' => 'getPhoneticFirstName',
        'phonetic_middle_name' => 'getPhoneticMiddleName',
        'phonetic_last_name' => 'getPhoneticLastName',
        'department' => 'getDepartment',
        'job_title' => 'getJobTitle',
        'emails' => 'getEmails',
        'phone_numbers' => 'getPhoneNumbers',
        'addresses' => 'getAddresses',
        'group' => 'getGroup',
        'created_at' => 'getCreatedAt',
        'updated_at' => 'getUpdatedAt'
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
        $this->container['prefix'] = isset($data['prefix']) ? $data['prefix'] : null;
        $this->container['first_name'] = isset($data['first_name']) ? $data['first_name'] : null;
        $this->container['middle_name'] = isset($data['middle_name']) ? $data['middle_name'] : null;
        $this->container['last_name'] = isset($data['last_name']) ? $data['last_name'] : null;
        $this->container['suffix'] = isset($data['suffix']) ? $data['suffix'] : null;
        $this->container['nickname'] = isset($data['nickname']) ? $data['nickname'] : null;
        $this->container['company'] = isset($data['company']) ? $data['company'] : null;
        $this->container['phonetic_first_name'] = isset($data['phonetic_first_name']) ? $data['phonetic_first_name'] : null;
        $this->container['phonetic_middle_name'] = isset($data['phonetic_middle_name']) ? $data['phonetic_middle_name'] : null;
        $this->container['phonetic_last_name'] = isset($data['phonetic_last_name']) ? $data['phonetic_last_name'] : null;
        $this->container['department'] = isset($data['department']) ? $data['department'] : null;
        $this->container['job_title'] = isset($data['job_title']) ? $data['job_title'] : null;
        $this->container['emails'] = isset($data['emails']) ? $data['emails'] : null;
        $this->container['phone_numbers'] = isset($data['phone_numbers']) ? $data['phone_numbers'] : null;
        $this->container['addresses'] = isset($data['addresses']) ? $data['addresses'] : null;
        $this->container['group'] = isset($data['group']) ? $data['group'] : null;
        $this->container['created_at'] = isset($data['created_at']) ? $data['created_at'] : null;
        $this->container['updated_at'] = isset($data['updated_at']) ? $data['updated_at'] : null;
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
     * @return int
     */
    public function getId()
    {
        return $this->container['id'];
    }

    /**
     * Sets id
     * @param int $id Integer ID. Read-only.
     * @return $this
     */
    public function setId($id)
    {
        $this->container['id'] = $id;

        return $this;
    }

    /**
     * Gets prefix
     * @return string
     */
    public function getPrefix()
    {
        return $this->container['prefix'];
    }

    /**
     * Sets prefix
     * @param string $prefix Salutation, such as Mr, Mrs, or Dr
     * @return $this
     */
    public function setPrefix($prefix)
    {
        $this->container['prefix'] = $prefix;

        return $this;
    }

    /**
     * Gets first_name
     * @return string
     */
    public function getFirstName()
    {
        return $this->container['first_name'];
    }

    /**
     * Sets first_name
     * @param string $first_name First name or given name
     * @return $this
     */
    public function setFirstName($first_name)
    {
        $this->container['first_name'] = $first_name;

        return $this;
    }

    /**
     * Gets middle_name
     * @return string
     */
    public function getMiddleName()
    {
        return $this->container['middle_name'];
    }

    /**
     * Sets middle_name
     * @param string $middle_name Middle or second name
     * @return $this
     */
    public function setMiddleName($middle_name)
    {
        $this->container['middle_name'] = $middle_name;

        return $this;
    }

    /**
     * Gets last_name
     * @return string
     */
    public function getLastName()
    {
        return $this->container['last_name'];
    }

    /**
     * Sets last_name
     * @param string $last_name Last name or surname
     * @return $this
     */
    public function setLastName($last_name)
    {
        $this->container['last_name'] = $last_name;

        return $this;
    }

    /**
     * Gets suffix
     * @return string
     */
    public function getSuffix()
    {
        return $this->container['suffix'];
    }

    /**
     * Sets suffix
     * @param string $suffix Suffix, such as \"Jr.\", \"Sr.\", \"II\", or \"III\"
     * @return $this
     */
    public function setSuffix($suffix)
    {
        $this->container['suffix'] = $suffix;

        return $this;
    }

    /**
     * Gets nickname
     * @return string
     */
    public function getNickname()
    {
        return $this->container['nickname'];
    }

    /**
     * Sets nickname
     * @param string $nickname Nickname, or a shortened informal version of the contact's name
     * @return $this
     */
    public function setNickname($nickname)
    {
        $this->container['nickname'] = $nickname;

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
     * @param string $company Name of the contact's company
     * @return $this
     */
    public function setCompany($company)
    {
        $this->container['company'] = $company;

        return $this;
    }

    /**
     * Gets phonetic_first_name
     * @return string
     */
    public function getPhoneticFirstName()
    {
        return $this->container['phonetic_first_name'];
    }

    /**
     * Sets phonetic_first_name
     * @param string $phonetic_first_name Phonetic first name. Useful for remembering how to pronounce the contact's name.
     * @return $this
     */
    public function setPhoneticFirstName($phonetic_first_name)
    {
        $this->container['phonetic_first_name'] = $phonetic_first_name;

        return $this;
    }

    /**
     * Gets phonetic_middle_name
     * @return string
     */
    public function getPhoneticMiddleName()
    {
        return $this->container['phonetic_middle_name'];
    }

    /**
     * Sets phonetic_middle_name
     * @param string $phonetic_middle_name Phonetic middle name. Useful for remembering how to pronounce the contact's name.
     * @return $this
     */
    public function setPhoneticMiddleName($phonetic_middle_name)
    {
        $this->container['phonetic_middle_name'] = $phonetic_middle_name;

        return $this;
    }

    /**
     * Gets phonetic_last_name
     * @return string
     */
    public function getPhoneticLastName()
    {
        return $this->container['phonetic_last_name'];
    }

    /**
     * Sets phonetic_last_name
     * @param string $phonetic_last_name Phonetic last name. Useful for remembering how to pronounce the contact's name.
     * @return $this
     */
    public function setPhoneticLastName($phonetic_last_name)
    {
        $this->container['phonetic_last_name'] = $phonetic_last_name;

        return $this;
    }

    /**
     * Gets department
     * @return string
     */
    public function getDepartment()
    {
        return $this->container['department'];
    }

    /**
     * Sets department
     * @param string $department Name of the contact's department
     * @return $this
     */
    public function setDepartment($department)
    {
        $this->container['department'] = $department;

        return $this;
    }

    /**
     * Gets job_title
     * @return string
     */
    public function getJobTitle()
    {
        return $this->container['job_title'];
    }

    /**
     * Sets job_title
     * @param string $job_title Contact's job title
     * @return $this
     */
    public function setJobTitle($job_title)
    {
        $this->container['job_title'] = $job_title;

        return $this;
    }

    /**
     * Gets emails
     * @return \Swagger\Client\Model\Emails
     */
    public function getEmails()
    {
        return $this->container['emails'];
    }

    /**
     * Sets emails
     * @param \Swagger\Client\Model\Emails $emails
     * @return $this
     */
    public function setEmails($emails)
    {
        $this->container['emails'] = $emails;

        return $this;
    }

    /**
     * Gets phone_numbers
     * @return \Swagger\Client\Model\PhoneNumberContact[]
     */
    public function getPhoneNumbers()
    {
        return $this->container['phone_numbers'];
    }

    /**
     * Sets phone_numbers
     * @param \Swagger\Client\Model\PhoneNumberContact[] $phone_numbers
     * @return $this
     */
    public function setPhoneNumbers($phone_numbers)
    {
        $this->container['phone_numbers'] = $phone_numbers;

        return $this;
    }

    /**
     * Gets addresses
     * @return \Swagger\Client\Model\Addresses
     */
    public function getAddresses()
    {
        return $this->container['addresses'];
    }

    /**
     * Sets addresses
     * @param \Swagger\Client\Model\Addresses $addresses
     * @return $this
     */
    public function setAddresses($addresses)
    {
        $this->container['addresses'] = $addresses;

        return $this;
    }

    /**
     * Gets group
     * @return \Swagger\Client\Model\GroupListContacts
     */
    public function getGroup()
    {
        return $this->container['group'];
    }

    /**
     * Sets group
     * @param \Swagger\Client\Model\GroupListContacts $group
     * @return $this
     */
    public function setGroup($group)
    {
        $this->container['group'] = $group;

        return $this;
    }

    /**
     * Gets created_at
     * @return int
     */
    public function getCreatedAt()
    {
        return $this->container['created_at'];
    }

    /**
     * Sets created_at
     * @param int $created_at Integer UNIX timestamp when the contact was created. Read-only.
     * @return $this
     */
    public function setCreatedAt($created_at)
    {
        $this->container['created_at'] = $created_at;

        return $this;
    }

    /**
     * Gets updated_at
     * @return int
     */
    public function getUpdatedAt()
    {
        return $this->container['updated_at'];
    }

    /**
     * Sets updated_at
     * @param int $updated_at Integer UNIX timestamp when the contact was created. Read-only.
     * @return $this
     */
    public function setUpdatedAt($updated_at)
    {
        $this->container['updated_at'] = $updated_at;

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


