//
// ExpressServiceCodeSummary.swift
//
// Generated by swagger-codegen
// https://github.com/swagger-api/swagger-codegen
//

import Foundation


/** The Express Service Code Summary Object is used to briefly represent a Express Service Code. It consists of the ID and code: */
public class ExpressServiceCodeSummary: JSONEncodable {
    /** ID */
    public var id: Int32?
    /** An 8-digit number in string format */
    public var expressServiceCode: String?

    public init() {}

    // MARK: JSONEncodable
    func encodeToJSON() -> AnyObject {
        var nillableDictionary = [String:AnyObject?]()
        nillableDictionary["id"] = self.id?.encodeToJSON()
        nillableDictionary["express_service_code"] = self.expressServiceCode
        let dictionary: [String:AnyObject] = APIHelper.rejectNil(nillableDictionary) ?? [:]
        return dictionary
    }
}
