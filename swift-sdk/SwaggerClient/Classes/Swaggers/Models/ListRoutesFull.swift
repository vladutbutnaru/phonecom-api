//
// ListRoutesFull.swift
//
// Generated by swagger-codegen
// https://github.com/swagger-api/swagger-codegen
//

import Foundation


public class ListRoutesFull: JSONEncodable {
    public var filters: FilterIdNameArray?
    public var sort: SortIdName?
    public var total: Int32?
    public var offset: Int32?
    public var limit: Int32?
    public var items: RoutesFull?

    public init() {}

    // MARK: JSONEncodable
    func encodeToJSON() -> AnyObject {
        var nillableDictionary = [String:AnyObject?]()
        nillableDictionary["filters"] = self.filters?.encodeToJSON()
        nillableDictionary["sort"] = self.sort?.encodeToJSON()
        nillableDictionary["total"] = self.total?.encodeToJSON()
        nillableDictionary["offset"] = self.offset?.encodeToJSON()
        nillableDictionary["limit"] = self.limit?.encodeToJSON()
        nillableDictionary["items"] = self.items?.encodeToJSON()
        let dictionary: [String:AnyObject] = APIHelper.rejectNil(nillableDictionary) ?? [:]
        return dictionary
    }
}
