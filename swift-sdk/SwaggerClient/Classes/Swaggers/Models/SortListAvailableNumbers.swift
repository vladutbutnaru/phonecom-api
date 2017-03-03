//
// SortListAvailableNumbers.swift
//
// Generated by swagger-codegen
// https://github.com/swagger-api/swagger-codegen
//

import Foundation


public class SortListAvailableNumbers: JSONEncodable {
    public var _internal: String?
    public var price: String?
    public var phoneNumber: String?

    public init() {}

    // MARK: JSONEncodable
    func encodeToJSON() -> AnyObject {
        var nillableDictionary = [String:AnyObject?]()
        nillableDictionary["internal"] = self._internal
        nillableDictionary["price"] = self.price
        nillableDictionary["phone_number"] = self.phoneNumber
        let dictionary: [String:AnyObject] = APIHelper.rejectNil(nillableDictionary) ?? [:]
        return dictionary
    }
}
