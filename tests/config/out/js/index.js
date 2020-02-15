"use strict";

exports.AllParameter = require("./allparameter.js")

exports.AllRepeatedParameter = require("./allrepeatedparameter.js")

exports.EmptyMessage = require("./emptymessage.js")

/**
 *  UpperCamelCase comment
 */
exports.UpperCamelCase = require("./uppercamelcase.js")

/**
 *  lowerCamelCase comment
 */
exports.LowerCamelCase = require("./lowercamelcase.js")

exports.SnakeCase = require("./snakecase.js")

exports.DependTest = require("./dependtest.js")

/**
 *  comment
 */
exports.TestEnum = require("./testenum.js")

exports.DependMessage = require("./dependmessage.js")

exports.PackageMessage = require("./packagemessage.js")

exports.DepDep = {}

exports.DepDep.DependTestMessage = require("./depdep.dependtestmessage.js")

exports.MyPackage = {}

exports.MyPackage.MyMessage = require("./mypackage.mymessage.js")

exports.MyPackage.MyEnum = require("./mypackage.myenum.js")

