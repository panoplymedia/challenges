const multiparty = require("multiparty");
const util = require("util");
const csv = require("fast-csv");
const { db } = require("../../database/db");
const { createOrder } = require("../../database/models/order");

/**
 * Utility: given the CSV row data in object notation,
 * map the data to the normalized property names for readability
 * as well as apply default values if needed
 * @param {object} rowObj
 * @returns {object} formatted csv row
 */
const formatCSVRow = (rowObj = {}) => ({
  //trim whitespace, force to uppercase
  customerName: (rowObj["Customer Name"] || "unknown").trim().toUpperCase(),
  productDescription: (rowObj["Item Description"] || "unknown")
    .trim()
    .toUpperCase(),
  //round to hundredths
  productPrice: (Number.parseFloat(rowObj["Item Price"]) || 0).toFixed(2),
  //round to whole
  quantity: Math.floor(Number.parseInt(rowObj["Quantity"]) || 0),
  //trim whitespace, force to uppercase
  merchantName: (rowObj["Merchant Name"] || "uknown").trim().toUpperCase(),
  merchantAddress: (rowObj["Merchant Address"] || "unkown").trim().toUpperCase()
});
exports.formatCSVRow = formatCSVRow;

/**
 * Given an object whos keys are the sales data headers and
 * values are the order values, this function parses a given
 * order from the CSV reader and normalizes the data and
 * writes to database.
 * @param {object} rowObj a row from the CSV in object notation
 * @returns {undefined} no return value
 */
const handleCSVRow = (rowObj = {}) => {
  const formatted = formatCSVRow(rowObj);
  createOrder(db, formatted);
};
exports.handleCSVRow = handleCSVRow;

/**
 * Orders API post handler for CSV upload
 * This handler is intended for use when a CSV file is
 * uploaded via POST as multi-part form data.
 * @param {object} req
 * @param {object} res
 * @param {function} next
 */
exports.csvPost = (req, res, next) => {
  // create a form to begin parsing
  const form = new multiparty.Form({ uploadDir: "./uploads" });

  form.parse(req, function(err, fields, files) {
    if (err) {
      res.writeHead(400, { "content-type": "text/plain" });
      res.end("invalid request: " + err.message);
      return;
    }

    // ensure received form data of type "file"
    // process the first file received.
    // TODO: add future support for multifile
    if (files && files.file && files.file[0]) {
      csv
        .parseFile(files.file[0].path, { headers: true })
        .on("error", err => {
          res.writeHead(500, { "content-type": "text/plain" });
          res.end(`An error occured ${err.message}`);
        })
        .on("data", handleCSVRow)
        .on("end", () => res.sendStatus(200));
    }
  });
};
