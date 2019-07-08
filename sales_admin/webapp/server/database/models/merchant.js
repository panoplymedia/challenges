const loki = require("lokijs");
const { db } = require("../db");

/**
 * LokiJS Collection: "Merchant"
 * A storage of all merchants from the sales data.
 */

/**
 * ensures the collection exists in the given database
 * @param {object} db LokiJS Database object
 * @returns {object} LokiJS collection
 */
const ensureMerchantCollection = db => {
  let merchants = db.getCollection("Merchant");
  if (!merchants) {
    merchants = db.addCollection("Merchant");
  }
  return merchants;
};
exports.ensureMerchantCollection = ensureMerchantCollection;

/**
 * Create a new merchant.
 * If the merchant already exists, no operation occurs.
 * @param {string} merchant.name
 * @param {string} merchant.address
 * @returns {object} the created merchant document
 */
exports.createMerchant = (db, { name, address }) => {
  const merchants = ensureMerchantCollection(db);
  const created = merchants.insert({
    name,
    address
  });
  db.saveDatabase();
  return created;
};
