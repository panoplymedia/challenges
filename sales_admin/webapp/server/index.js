/**
 * Sales Admin API Server.
 * A RESTful API server for the ACME Cult Hero Supplies aquisition.
 *
 * Technology:
 *   - ExpressJS for the API routing
 *   - LokiJS for a persistent data storage.
 *
 * Runs on port 4000.
 */

const app = require("express")();
const { mapRoutes } = require("./util/mapRoutes");
const { orders } = require("./routes/orders");
const fs = require("fs");

// setup: need a uploads dir, create if not exists
const dir = "./uploads";
if (!fs.existsSync(dir)) {
  fs.mkdirSync(dir);
}

// Takes the given route structure and applies it to the api server
mapRoutes(app, {
  "/orders": {
    get: orders.list,
    post: orders.post,
    "/revenue": {
      get: orders.getRevenue
    }
  }
});

//todo: env
app.listen(4000);
console.log("Sales Admin server started on port 4000");
