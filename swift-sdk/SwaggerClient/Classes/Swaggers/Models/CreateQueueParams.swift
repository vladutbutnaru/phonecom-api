//
// CreateQueueParams.swift
//
// Generated by swagger-codegen
// https://github.com/swagger-api/swagger-codegen
//

import Foundation


public class CreateQueueParams: JSONEncodable {
    /** Name of queue */
    public var name: String?
    /** Recording lookup object */
    public var greeting: AnyObject?
    /** Recording lookup object */
    public var holdMusic: AnyObject?
    /** Max seconds for hold */
    public var maxHoldTime: Int32?
    /** Type of caller id */
    public var callerIdType: String?
    /** Number of seconds to ring each member */
    public var ringTime: Int32?
    /** Extensions or phone numbers */
    public var members: [AnyObject]?

    public init() {}

    // MARK: JSONEncodable
    func encodeToJSON() -> AnyObject {
        var nillableDictionary = [String:AnyObject?]()
        nillableDictionary["name"] = self.name
        nillableDictionary["greeting"] = self.greeting
        nillableDictionary["hold_music"] = self.holdMusic
        nillableDictionary["max_hold_time"] = self.maxHoldTime?.encodeToJSON()
        nillableDictionary["caller_id_type"] = self.callerIdType
        nillableDictionary["ring_time"] = self.ringTime?.encodeToJSON()
        nillableDictionary["members"] = self.members?.encodeToJSON()
        let dictionary: [String:AnyObject] = APIHelper.rejectNil(nillableDictionary) ?? [:]
        return dictionary
    }
}
