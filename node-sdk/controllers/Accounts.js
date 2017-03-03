'use strict';

var url = require('url');


var Accounts = require('./AccountsService');


module.exports.getAccount = function getAccount (req, res, next) {
  Accounts.getAccount(req.swagger.params, res, next);
};

module.exports.listAccounts = function listAccounts (req, res, next) {
  Accounts.listAccounts(req.swagger.params, res, next);
};
