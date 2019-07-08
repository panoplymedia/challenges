import React, { Component } from "react";
import { OrdersTable } from "../../components/OrderTable";
import { PageHeader, message } from "antd";
import { OrdersUpload } from "../../components/OrdersUpload";
import { Revenue } from "../../components/Revenue";

export class OrdersView extends Component {
  state = {
    orders: [],
    revenue: 0,
    loading: true
  };

  fetchData() {
    Promise.all([
      // get orders
      fetch("/orders")
        .then(resp => resp.json())
        .then(orders => {
          this.setState({ orders });
        }),
      // get revenue
      fetch("/orders/revenue")
        .then(resp => resp.json())
        .then(({ revenue }) => this.setState({ revenue }))
    ])
      .then(() => this.setState({ loading: false }))
      .catch(() => message.error("An error occured."));
  }

  componentDidMount() {
    this.fetchData();
  }

  render() {
    const { loading, orders, revenue } = this.state;
    return (
      <div>
        <PageHeader
          title={`Sales Admin`}
          subTitle={<Revenue revenue={revenue} />}
          extra={<OrdersUpload onSuccess={() => this.fetchData()} />}
        />
        <OrdersTable orders={orders} loading={loading} />
      </div>
    );
  }
}
