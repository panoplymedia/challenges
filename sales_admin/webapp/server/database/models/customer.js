const loki = require("lokijs");
const { db } = require("../db");

/**
 * LokiJS Collection: "Customer"
 * A storage of all customers from the sales data.
 */

/**
 * ensures the collection exists in the given database
 * @param {object} db LokiJS Database object
 * @returns {object} LokiJS collection
 */
const ensureCustomerCollection = db => {
  let customers = db.getCollection("Customer");
  if (!customers) {
    customers = db.addCollection("Customer");
  }
  return customers;
};
exports.ensureCustomerCollection = ensureCustomerCollection;

/**
 * Create a new customer.
 * If the customer already exists, the existing customer is returned.
 * @param {string} customer.name
 * @returns {object} the created customer document
 */
exports.createCustomer = (db, { name }) => {
  const customers = ensureCustomerCollection(db);
  const created = customers.insert({
    name
  });
  db.saveDatabase();
  return created;
};
