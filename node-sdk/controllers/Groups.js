'use strict';

var url = require('url');


var Groups = require('./GroupsService');


module.exports.createAccountExtensionContactGroup = function createAccountExtensionContactGroup (req, res, next) {
  Groups.createAccountExtensionContactGroup(req.swagger.params, res, next);
};

module.exports.deleteAccountExtensionContactGroup = function deleteAccountExtensionContactGroup (req, res, next) {
  Groups.deleteAccountExtensionContactGroup(req.swagger.params, res, next);
};

module.exports.getAccountExtensionContactGroup = function getAccountExtensionContactGroup (req, res, next) {
  Groups.getAccountExtensionContactGroup(req.swagger.params, res, next);
};

module.exports.listAccountExtensionContactGroups = function listAccountExtensionContactGroups (req, res, next) {
  Groups.listAccountExtensionContactGroups(req.swagger.params, res, next);
};

module.exports.replaceAccountExtensionContactGroup = function replaceAccountExtensionContactGroup (req, res, next) {
  Groups.replaceAccountExtensionContactGroup(req.swagger.params, res, next);
};
