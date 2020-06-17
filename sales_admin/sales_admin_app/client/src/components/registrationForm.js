import React, { useState } from "react";
import { withRouter } from "react-router-dom";

function RegistrationForm(props) {
  const [state, setState] = useState({
    email: "",
    password: "",
    passwordConfirm: "",
    success: null,
  });

  const handleInputChange = (e) => {
    const { id, value } = e.target;
    setState((prevState) => ({
      ...prevState,
      [id]: value,
    }));
  };

  const handleSubmit = (e) => {
    e.preventDefault();
    if (state.password === state.passwordConfirm) {
      return postRegistration(state, redirectToLogin);
    } else {
        console.log("try again")
    }
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
        <div className="form-group text-left">
          <label htmlFor="exampleInputPassword1">Confirm Password</label>
          <input
            type="password"
            className="form-control"
            id="passwordConfirm"
            placeholder="Confirm Password"
            value={state.passwordConfirm}
            onChange={handleInputChange}
          />
        </div>
        <button
          type="submit"
          className="btn btn-primary"
          onClick={handleSubmit}
        >
          Register
        </button>
      </form>
    </div>
  );
}

async function postRegistration(state, redirectToLogin) {
  try {
    const response = await fetch("http://localhost:3001/users/register", {
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
    redirectToLogin();
    return await response.json();
  } catch (error) {
    redirectToLogin();
  }
}

export default withRouter(RegistrationForm);