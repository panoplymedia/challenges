const loki = require("lokijs");
/**
 * LokiJS: Database instantiation
 * The database is persisted in the given json file.
 */
const db = new loki("./db.json", {
  autoload: true,
  autoupdate: true
});
exports.db = db;
