//
// QueueSummary.swift
//
// Generated by swagger-codegen
// https://github.com/swagger-api/swagger-codegen
//

import Foundation


/** The Queue Summary Object is used to briefly represent a calling queue. It can be seen in several places throughout this API. Here are the properties: */
public class QueueSummary: JSONEncodable {
    /** Integer ID. Read-only. */
    public var id: Int32?
    /** Name. Required. */
    public var name: String?

    public init() {}

    // MARK: JSONEncodable
    func encodeToJSON() -> AnyObject {
        var nillableDictionary = [String:AnyObject?]()
        nillableDictionary["id"] = self.id?.encodeToJSON()
        nillableDictionary["name"] = self.name
        let dictionary: [String:AnyObject] = APIHelper.rejectNil(nillableDictionary) ?? [:]
        return dictionary
    }
}
