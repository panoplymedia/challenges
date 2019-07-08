import { orderCollection } from "./order";

describe("Order Collection", () => {
  describe("ensureOrderCollection", () => {
    xit("returns an order collection", () ={});
  });
  describe("createOrder", () => {
    xit("adds a customer to the database if not exists", () => {});
    xit("adds a merchant to the database if not exists", () => {});
    xit("adds a product to the database if not exists", () => {});
    xit("adds a order to the database", () => {});
    xit("throws an error on invalid entry", () => {});
  });
  describe("getAllOrders", () => {
    xit("returns an array of orders", () => {});
    xit("returns each order's customer", () => {});
    xit("returns each order's product", () => {});
    xit("returns each order's quantity", () => {});
    xit("returns an empty array if no orders", () => {});
  });
  describe("getRevenue", () => {
    xit("returns the total revenue of all orders", () => {});
    xit("returns 0 if no orders");
  });
});
