'use strict';

var url = require('url');


var Applications = require('./ApplicationsService');


module.exports.getAccountApplication = function getAccountApplication (req, res, next) {
  Applications.getAccountApplication(req.swagger.params, res, next);
};

module.exports.listAccountApplications = function listAccountApplications (req, res, next) {
  Applications.listAccountApplications(req.swagger.params, res, next);
};
