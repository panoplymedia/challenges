/**
 * Utility method originating from the ExpressJS example "route-map"
 * @param {object} app the ExpressJS app server
 * @param {object} routeObj an object whos keys are route paths and values are either
 * ExpressJS route handlers (get,delete,post,put) or sub route paths (ex: /:id)
 * @param {string|undefined} routePath The current route path, used to build nested routes
 * @returns {undefined} no return value
 */
const mapRoutes = function(app, routeObj, routePath) {
  const route = routePath || "";
  for (const key in routeObj) {
    switch (typeof routeObj[key]) {
      // { '/path': { ... }}
      case "object":
        mapRoutes(app, routeObj[key], route + key);
        break;
      // get: function(){ ... }
      case "function":
        app[key](route, routeObj[key]);
        break;
    }
  }
};

exports.mapRoutes = mapRoutes;
