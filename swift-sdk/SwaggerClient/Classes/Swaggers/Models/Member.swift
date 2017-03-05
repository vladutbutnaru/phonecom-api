//
// Member.swift
//
// Generated by swagger-codegen
// https://github.com/swagger-api/swagger-codegen
//

import Foundation


public class Member: JSONEncodable {
    /** Extension that this member refers to. Output is an Extension Summary Object. Input must be an Extension Lookup Object. */
    public var _extension: ExtensionSummary?
    /** Phone number */
    public var phoneNumber: String?

    public init() {}

    // MARK: JSONEncodable
    func encodeToJSON() -> AnyObject {
        var nillableDictionary = [String:AnyObject?]()
        nillableDictionary["extension"] = self._extension?.encodeToJSON()
        nillableDictionary["phone_number"] = self.phoneNumber
        let dictionary: [String:AnyObject] = APIHelper.rejectNil(nillableDictionary) ?? [:]
        return dictionary
    }
}