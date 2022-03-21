const { Parser } = require("acorn");

const MyParser = Parser.extend(require("acorn-jsx")());
