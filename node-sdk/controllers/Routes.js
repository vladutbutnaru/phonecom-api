'use strict';

var url = require('url');


var Routes = require('./RoutesService');


module.exports.createRoute = function createRoute (req, res, next) {
  Routes.createRoute(req.swagger.params, res, next);
};

module.exports.deleteAccountRoute = function deleteAccountRoute (req, res, next) {
  Routes.deleteAccountRoute(req.swagger.params, res, next);
};

module.exports.getAccountRoute = function getAccountRoute (req, res, next) {
  Routes.getAccountRoute(req.swagger.params, res, next);
};

module.exports.listAccountRoutes = function listAccountRoutes (req, res, next) {
  Routes.listAccountRoutes(req.swagger.params, res, next);
};

module.exports.replaceAccountRoute = function replaceAccountRoute (req, res, next) {
  Routes.replaceAccountRoute(req.swagger.params, res, next);
};
