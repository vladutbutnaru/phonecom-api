//
// ContactAccount.swift
//
// Generated by swagger-codegen
// https://github.com/swagger-api/swagger-codegen
//

import Foundation


public class ContactAccount: JSONEncodable {
    /** Name. Required. */
    public var name: String?
    /** Company name */
    public var company: String?
    public var address: Address?
    /** Phone number. Required. */
    public var phone: String?
    /** Fax number */
    public var fax: String?
    /** Primary email address. Required. */
    public var primaryEmail: String?
    /** Alternate email address */
    public var alternateEmail: String?

    public init() {}

    // MARK: JSONEncodable
    func encodeToJSON() -> AnyObject {
        var nillableDictionary = [String:AnyObject?]()
        nillableDictionary["name"] = self.name
        nillableDictionary["company"] = self.company
        nillableDictionary["address"] = self.address?.encodeToJSON()
        nillableDictionary["phone"] = self.phone
        nillableDictionary["fax"] = self.fax
        nillableDictionary["primary_email"] = self.primaryEmail
        nillableDictionary["alternate_email"] = self.alternateEmail
        let dictionary: [String:AnyObject] = APIHelper.rejectNil(nillableDictionary) ?? [:]
        return dictionary
    }
}
