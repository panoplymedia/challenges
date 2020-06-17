import React, { useState } from "react";
import Header from './components/header'
import Login from './components/login'
import RegistrationForm from './components/registrationForm'
import SalesData from './components/salesData'
import { BrowserRouter as Router, Switch, Route } from "react-router-dom";


import './App.css';

function App(props) {
  const [updateErrorMessage, setUser] = useState(null);  
  return (
    <Router>
      <div className="App">
        <Header />
        <Switch>
          <Route path="/" exact={true}>
            <RegistrationForm showError={updateErrorMessage} />
          </Route>

          <Route path="/register">
            <RegistrationForm
              showError={updateErrorMessage}
            />
          </Route>

          <Route path="/login">
            <Login
              showError={updateErrorMessage}
              setUser={setUser}
            />
          </Route>

          <Route path="/home">
            <SalesData />
          </Route>
        </Switch>
      </div>
    </Router>
  );
}

export default App;
