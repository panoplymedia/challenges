import React from "react";
import PropTypes from "prop-types";
import { Layout, Menu, Card } from "antd";
import { Link } from "react-router-dom";
const { Header, Content, Footer } = Layout;

export const RootLayout = ({ children }) => {
  return (
    <Layout className="layout">
      <Header>
        <Menu
          theme="dark"
          mode="horizontal"
          defaultSelectedKeys={["1"]}
          style={{ lineHeight: "64px" }}
        >
          {/* Todo: router links */}
          <Menu.Item key="1">
            <Link to="/" className="nav-text">
              Orders
            </Link>
          </Menu.Item>
        </Menu>
      </Header>
      <Content>
        <Card style={{ margin: 24 }}>{children}</Card>
      </Content>
      <Footer style={{ textAlign: "center" }}>{/* TODO: footer */}</Footer>
    </Layout>
  );
};

RootLayout.propTypes = {
  children: PropTypes.oneOfType([
    PropTypes.arrayOf(PropTypes.node),
    PropTypes.node
  ]).isRequired
};
