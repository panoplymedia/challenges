import React, { useState } from "react";
import { withRouter } from "react-router-dom";

function Login(props) {
  const [state, setState] = useState({
    email: '',
    password: '',
    successMessage: null
  })

  const handleInputChange = e => {
    const { id, value } = e.target;
    setState((prevStates) => ({
      ...prevStates,
      [id]: value,
    }));
  }

  const handleSubmit = e => {
    e.preventDefault();
    postLogin(state, redirectToHome, redirectToRegister, redirectToLogin);
  }

  const redirectToHome = (json) => {
    window.localStorage.setItem("userAuth", JSON.stringify(json));
    props.history.push("/home");
  };
  const redirectToRegister = () => {
    window.localStorage.clear()
    props.history.push("/register");
  };
    
  const redirectToLogin = () => {
    window.localStorage.clear();
    props.history.push("/login");
  };

  return (
    <div className="card col-12 col-lg-4 login-card mt-2 hv-center">
      <form>
        <div className="form-group text-left">
          <label htmlFor="exampleInputEmail1">Email address</label>
          <input
            type="email"
            className="form-control"
            id="email"
            aria-describedby="emailHelp"
            placeholder="Enter email"
            value={state.email}
            onChange={handleInputChange}
          />
          <small id="emailHelp" className="form-text text-muted">
            We'll never share your email with anyone else.
          </small>
        </div>
        <div className="form-group text-left">
          <label htmlFor="exampleInputPassword1">Password</label>
          <input
            type="password"
            className="form-control"
            id="password"
            placeholder="Password"
            value={state.password}
            onChange={handleInputChange}
          />
        </div>
        <div className="form-check"></div>
        <button
          type="submit"
          className="btn btn-primary"
          onClick={handleSubmit}
        >
          Submit
        </button>
      </form>
      <div
        className="alert alert-success mt-2"
        style={{ display: state.successMessage ? "block" : "none" }}
        role="alert"
      >
        {state.successMessage}
      </div>
      <div className="registerMessage">
        <span>Dont have an account? </span>
        <span className="loginText" onClick={() => redirectToRegister()}>
          Register
        </span>
      </div>
    </div>
  );
}

async function postLogin(state, redirectToHome, redirectToRegister, redirectToLogin) {
  try {
    const response = await fetch("http://localhost:3001/auth/login", {
      method: "POST",
      mode: "cors",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        email: state.email,
        password: state.password,
      }),
    });
    const json = await response.json();

    if (json.error) {
      alert('Please Try again, or create a new account')
      redirectToLogin();
    } else {
       redirectToHome(json);
    }
  } catch (error) {
    redirectToRegister();
  }
}

export default withRouter(Login);