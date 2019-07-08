import React from "react";
import "./App.css";
import { OrdersView } from "./views/OrdersView";
import { HashRouter, Route } from "react-router-dom";
import { RootLayout } from "./components/RootLayout";
function App() {
  return (
    <HashRouter>
      <RootLayout>
        <Route path="/" component={OrdersView} />
      </RootLayout>
    </HashRouter>
  );
}

export default App;
