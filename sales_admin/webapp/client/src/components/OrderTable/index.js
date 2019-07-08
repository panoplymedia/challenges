import React from "react";
import PropTypes from "prop-types";
import { Table } from "antd";

export const strSort = (a = "", b = "") => a.localeCompare(b);
export const numSort = (a = 0, b = 0) => a - b;
/**
 * Order table column definitions
 */
export const ORDER_COLS = [
  {
    title: "Import Date",
    dataIndex: "importTimeEpoch",
    sorter: (a, b) => numSort(a.importTimeEpoch, b.importTimeEpoch),
    render: timestamp => {
      console.log(timestamp);
      debugger;
      return new Date(timestamp).toLocaleString();
    }
  },
  {
    title: "Customer Name",
    dataIndex: "customer.name",
    sorter: (a, b) => strSort(a.customer.name, b.customer.name),
    defaultSortOrder: "descend"
  },
  {
    title: "Item Description",
    dataIndex: "product.description",
    sorter: (a, b) => strSort(a.product.description, b.product.description),
    defaultSortOrder: "descend"
  },
  {
    title: "Merchant Name",
    sorter: (a, b) => strSort(a.product.erchant.name, b.product.merchant.name),
    dataIndex: "product.merchant.name"
  },
  {
    title: "Merchant Address",
    sorter: (a, b) =>
      strSort(a.product.merchant.address, b.product.merchant.address),
    dataIndex: "product.merchant.address"
  },
  {
    title: "Item Price ($)",
    dataIndex: "product.price",
    align: "right",
    sorter: (a, b) => numSort(a.product.price, b.product.price),
    defaultSortOrder: "descend"
  },
  {
    title: "Quantity",
    dataIndex: "quantity",
    align: "right",
    sorter: (a, b) => numSort(a.quantity, b.quantity),
    defaultSortOrder: "descend"
  },
  {
    title: "Total ($)",
    dataIndex: "total",
    align: "right",
    sorter: (a, b) => numSort(a.total, b.total)
  }
];

/**
 * An ant design table for representing Order records
 * @param {object} props
 */
export const OrdersTable = ({ orders, ...other }) => {
  return (
    <Table
      size={"middle"}
      pagination={false}
      columns={ORDER_COLS}
      // TODo: typically using index as key is bad
      // but without a unique ID, and without row manipulation
      // this is ok for the time being.
      dataSource={orders.map((order, index) => ({
        ...order,
        total: (order.quantity * order.product.price).toFixed(2),
        key: index
      }))}
      {...other}
    />
  );
};

OrdersTable.propTypes = {
  orders: PropTypes.arrayOf(
    PropTypes.shape({
      customer: {
        name: PropTypes.string
      },
      product: {
        name: PropTypes.string,
        price: PropTypes.number,
        merchant: {
          name: PropTypes.string,
          address: PropTypes.string
        }
      },
      quantity: PropTypes.number
    })
  )
};
