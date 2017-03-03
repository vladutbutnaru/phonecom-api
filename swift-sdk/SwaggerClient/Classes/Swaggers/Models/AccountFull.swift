//
// AccountFull.swift
//
// Generated by swagger-codegen
// https://github.com/swagger-api/swagger-codegen
//

import Foundation


/** The Full Account Object includes all of the properties in the Account Summary Object, along with the following: */
public class AccountFull: JSONEncodable {
    /** Account ID. Sometimes referred to as \&quot;Voip ID\&quot; or \&quot;voip_id\&quot;. */
    public var id: Int32?
    /** Name on the account. Read-only. */
    public var name: String?
    /** Account user name */
    public var username: String?
    /** Account password. For security reason, this is masked as \&quot;**\&quot; */
    public var password: String?
    /** If this account is a subaccount, this property shows the master account that it belongs to. Otherwise it is NULL. Read-only. Output is an Account Summary Object. */
    public var masterAccount: AccountSummary?
    /** Contact Object. See below for details. */
    public var contact: ContactAccount?
    /** Contact Object for billing purposes. See below for details. */
    public var billingContact: ContactAccount?

    public init() {}

    // MARK: JSONEncodable
    func encodeToJSON() -> AnyObject {
        var nillableDictionary = [String:AnyObject?]()
        nillableDictionary["id"] = self.id?.encodeToJSON()
        nillableDictionary["name"] = self.name
        nillableDictionary["username"] = self.username
        nillableDictionary["password"] = self.password
        nillableDictionary["master_account"] = self.masterAccount?.encodeToJSON()
        nillableDictionary["contact"] = self.contact?.encodeToJSON()
        nillableDictionary["billing_contact"] = self.billingContact?.encodeToJSON()
        let dictionary: [String:AnyObject] = APIHelper.rejectNil(nillableDictionary) ?? [:]
        return dictionary
    }
}
