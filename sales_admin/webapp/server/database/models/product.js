const loki = require("lokijs");
const { db } = require("../db");

/**
 * LokiJS Collection: "Product"
 * A storage of all products from the sales data.
 * A Product comprises of:
 *  - a description
 *  - a price
 *  - 1 merchant
 *
 */

/**
 * ensures the collection exists in the given database
 * @param {object} db LokiJS Database object
 * @returns {object} LokiJS collection
 */
const ensureProductCollection = db => {
  let products = db.getCollection("Product");
  if (!products) {
    products = db.addCollection("Product");
  }
  return products;
};
exports.ensureProductCollection = ensureProductCollection;

/**
 * Create a new product.
 * If the product already exists, no operation occurs.
 * @param {string} product.description
 * @param {number} product.price
 * @param {object} product.merchant
 * @returns {object} the created product document
 */
exports.createProduct = (db, { description, price, merchant }) => {
  const products = ensureProductCollection(db);
  const created = products.insert({
    description,
    price,
    merchant
  });
  db.saveDatabase();
  return created;
};
