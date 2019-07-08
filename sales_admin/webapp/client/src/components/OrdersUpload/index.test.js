import Item from "antd/lib/list/Item";
import { onOrderUploadEvent } from ".";

describe("OrdersUpload", () => {
  describe("validateOrderUpload", () => {
    xit("shows an error message if not type text/csv", () => {});
    xit("returns false if not type text/csv", () => {});
    xit("returns true if type text/csv", () => {});
  });

  describe("genOnOrderUploadEvent", () => {
    xit("shows an error message if param structure invalid", () => {});
    xit("shows an error message containing the filename on error events", () => {});
    xit("shows a success message containing the filename on done events", () => {});
    xit("calls onSuccess on done events", () => {});
  });
  describe("renders", () => {
    xit("a button for csv import", () => {});
  });
});
