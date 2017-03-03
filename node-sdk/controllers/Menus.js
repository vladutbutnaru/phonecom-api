'use strict';

var url = require('url');


var Menus = require('./MenusService');


module.exports.createAccountMenu = function createAccountMenu (req, res, next) {
  Menus.createAccountMenu(req.swagger.params, res, next);
};

module.exports.deleteAccountMenu = function deleteAccountMenu (req, res, next) {
  Menus.deleteAccountMenu(req.swagger.params, res, next);
};

module.exports.getAccountMenu = function getAccountMenu (req, res, next) {
  Menus.getAccountMenu(req.swagger.params, res, next);
};

module.exports.listAccountMenus = function listAccountMenus (req, res, next) {
  Menus.listAccountMenus(req.swagger.params, res, next);
};

module.exports.replaceAccountMenu = function replaceAccountMenu (req, res, next) {
  Menus.replaceAccountMenu(req.swagger.params, res, next);
};
