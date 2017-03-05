//
// CreateContactParams.swift
//
// Generated by swagger-codegen
// https://github.com/swagger-api/swagger-codegen
//

import Foundation


public class CreateContactParams: JSONEncodable {
    /** First Name */
    public var firstName: String?
    /** Middle Name */
    public var middleName: String?
    /** Last Name */
    public var lastName: String?
    /** Prefix */
    public var _prefix: String?
    /** Phonetic First Name */
    public var phoneticFirstName: String?
    /** Phonetic Middle Name */
    public var phoneticMiddleName: String?
    /** Phonetic Last Name */
    public var phoneticLastName: String?
    /** Suffix */
    public var suffix: String?
    /** Nickname */
    public var nickname: String?
    /** Company Name */
    public var company: String?
    /** Department Name */
    public var department: String?
    /** Job Title */
    public var jobTitle: String?
    /** Email Addresses */
    public var emails: [AnyObject]?
    /** Phone Numbers */
    public var phoneNumbers: [AnyObject]?
    /** Addresses */
    public var addresses: [AnyObject]?
    /** Contact Group */
    public var group: AnyObject?

    public init() {}

    // MARK: JSONEncodable
    func encodeToJSON() -> AnyObject {
        var nillableDictionary = [String:AnyObject?]()
        nillableDictionary["first_name"] = self.firstName
        nillableDictionary["middle_name"] = self.middleName
        nillableDictionary["last_name"] = self.lastName
        nillableDictionary["prefix"] = self._prefix
        nillableDictionary["phonetic_first_name"] = self.phoneticFirstName
        nillableDictionary["phonetic_middle_name"] = self.phoneticMiddleName
        nillableDictionary["phonetic_last_name"] = self.phoneticLastName
        nillableDictionary["suffix"] = self.suffix
        nillableDictionary["nickname"] = self.nickname
        nillableDictionary["company"] = self.company
        nillableDictionary["department"] = self.department
        nillableDictionary["job_title"] = self.jobTitle
        nillableDictionary["emails"] = self.emails?.encodeToJSON()
        nillableDictionary["phone_numbers"] = self.phoneNumbers?.encodeToJSON()
        nillableDictionary["addresses"] = self.addresses?.encodeToJSON()
        nillableDictionary["group"] = self.group
        let dictionary: [String:AnyObject] = APIHelper.rejectNil(nillableDictionary) ?? [:]
        return dictionary
    }
}