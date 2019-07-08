const loki = require("lokijs");
const { createCustomer } = require("./customer");
const { createMerchant } = require("./merchant");
const { createProduct } = require("./product");
/**
 * LokiJS Collection: "Order"
 * A storage of all orders from the sales data.
 * An Order comprises of:
 *  - 1 customer
 *  - 1 product
 *  - 1 quantity
 */

/**
 * ensures the collection exists in the given database
 * @param {object} db LokiJS Database object
 * @returns {object} LokiJS collection
 */
const ensureOrderCollection = db => {
  let order = db.getCollection("Order");
  if (!order) {
    order = db.addCollection("Order");
  }
  return order;
};
exports.ensureOrderCollection = ensureOrderCollection;

/**
 * Creates a new order in the Order collection
 * @param {string} order.customerName
 * @param {string} order.productDescription
 * @param {number} order.productPrice
 * @param {number} order.quantity
 * @param {string} order.merchantName
 * @param {string} order.merchantAddress
 * @returns {object} the created order document
 */
exports.createOrder = (
  db,
  {
    customerName,
    productDescription,
    productPrice,
    quantity,
    merchantName,
    merchantAddress
  }
) => {
  const orders = ensureOrderCollection(db);
  const customer = createCustomer(db, {
    name: customerName
  });
  const merchant = createMerchant(db, {
    name: merchantName,
    address: merchantAddress
  });
  const product = createProduct(db, {
    description: productDescription,
    price: productPrice,
    merchant
  });

  const resp = orders.insert({
    customer,
    product,
    quantity
  });

  // persists the database after insert (since its a json file)
  db.saveDatabase();
  return resp;
};

/**
 * Returns all orders, eliminates metadata
 * @returns {object[]} array of orders
 */
exports.getAllOrders = db => {
  const orders = ensureOrderCollection(db);
  return orders.find().map(order => ({
    customer: order.customer,
    product: order.product,
    quantity: order.quantity,
    importTimeEpoch: order.meta.created
  }));
};

/**
 * Returns the total revenue.
 * @returns {number} revenue
 */
exports.getRevenue = db => {
  const orders = ensureOrderCollection(db);
  if (orders.data.length) {
    return orders.data
      .map(order => order.product.price * order.quantity)
      .reduce((prev = 0, curr = 0) => prev + curr);
  } else {
    return 0;
  }
};
