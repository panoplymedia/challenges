const { db } = require("../../database/db");
const { getAllOrders, getRevenue } = require("../../database/models/order");
const { csvPost } = require("./post");

/**
 * Orders API: Express route handlers
 */
const orders = {
  /**
   * return all orders
   */
  list: (req, res) => {
    const orders = getAllOrders(db);
    res.send(orders);
  },
  getRevenue: (req, res) => {
    res.send({
      revenue: getRevenue(db)
    });
  },
  /**
   * Return one order
   */
  get: (req, res) => {
    // not implemented
    res.writeHead(400, { "content-type": "text/plain" });
    res.end("not implemented");
  },
  /**
   * Delete orders
   */
  delete: (req, res) => {
    // not implemented
    res.writeHead(400, { "content-type": "text/plain" });
    res.end("not implemented");
  },
  /**
   * Add order(s)
   */
  post: (req, res, next) => {
    const contentType = req.headers["content-type"] || "";
    if (contentType.indexOf("multipart/form-data") === 0) {
      csvPost(req, res, next);
    } else {
      res.writeHead(400, { "content-type": "text/plain" });
      res.end("content type not supported.");
    }
  }
};
exports.orders = orders;
