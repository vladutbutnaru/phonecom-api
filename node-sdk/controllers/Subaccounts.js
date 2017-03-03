'use strict';

var url = require('url');


var Subaccounts = require('./SubaccountsService');


module.exports.createAccountSubaccount = function createAccountSubaccount (req, res, next) {
  Subaccounts.createAccountSubaccount(req.swagger.params, res, next);
};

module.exports.listAccountSubaccounts = function listAccountSubaccounts (req, res, next) {
  Subaccounts.listAccountSubaccounts(req.swagger.params, res, next);
};
