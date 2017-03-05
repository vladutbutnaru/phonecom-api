//
// FilterListAvailableNumbers.swift
//
// Generated by swagger-codegen
// https://github.com/swagger-api/swagger-codegen
//

import Foundation


public class FilterListAvailableNumbers: JSONEncodable {
    public var phoneNumber: String?
    public var countryCode: String?
    public var npa: [Int32]?
    public var nxx: String?
    public var xxxx: String?
    public var city: String?
    public var province: String?
    public var country: String?
    public var price: String?
    public var category: String?

    public init() {}

    // MARK: JSONEncodable
    func encodeToJSON() -> AnyObject {
        var nillableDictionary = [String:AnyObject?]()
        nillableDictionary["phone_number"] = self.phoneNumber
        nillableDictionary["country_code"] = self.countryCode
        nillableDictionary["npa"] = self.npa?.encodeToJSON()
        nillableDictionary["nxx"] = self.nxx
        nillableDictionary["xxxx"] = self.xxxx
        nillableDictionary["city"] = self.city
        nillableDictionary["province"] = self.province
        nillableDictionary["country"] = self.country
        nillableDictionary["price"] = self.price
        nillableDictionary["category"] = self.category
        let dictionary: [String:AnyObject] = APIHelper.rejectNil(nillableDictionary) ?? [:]
        return dictionary
    }
}