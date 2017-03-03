'use strict';

var url = require('url');


var Contacts = require('./ContactsService');


module.exports.createAccountExtensionContact = function createAccountExtensionContact (req, res, next) {
  Contacts.createAccountExtensionContact(req.swagger.params, res, next);
};

module.exports.deleteAccountExtensionContact = function deleteAccountExtensionContact (req, res, next) {
  Contacts.deleteAccountExtensionContact(req.swagger.params, res, next);
};

module.exports.getAccountExtensionContact = function getAccountExtensionContact (req, res, next) {
  Contacts.getAccountExtensionContact(req.swagger.params, res, next);
};

module.exports.listAccountExtensionContacts = function listAccountExtensionContacts (req, res, next) {
  Contacts.listAccountExtensionContacts(req.swagger.params, res, next);
};

module.exports.replaceAccountExtensionContact = function replaceAccountExtensionContact (req, res, next) {
  Contacts.replaceAccountExtensionContact(req.swagger.params, res, next);
};
