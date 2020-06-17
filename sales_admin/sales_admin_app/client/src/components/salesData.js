import React, { useState, useEffect } from "react";
import { withRouter } from "react-router-dom";
import FileUpload from './fileUpload'

function SalesData(props) {
    const [state, setState] = useState([]);
    useEffect(() => {
      fetchSalesData().then((res) => {
        if(res.error) {
          setState({data: res})
        } else {
          setState({data: res})
        }
      })
    }, []);

    if (window.localStorage.userAuth) {
        if(state.data) {
          return (
            <div>
              <FileUpload prop={"test"} />
              <table className="table">
                <thead>
                  <tr>
                    <th scope="col">Id</th>
                    <th scope="col">Customer Name</th>
                    <th scope="col">Item Description</th>
                    <th scope="col">Item Price</th>
                    <th scope="col">Quantity</th>
                    <th scope="col">Merchant Name</th>
                    <th scope="col">Merchant Address</th>
                  </tr>
                </thead>
                <tbody>
                  {state.data.length > 1 ? state.data.map((row) => {
                    return (
                    <tr key={row.id}>
                      <th scope="row">{row.id}</th>
                      <th>{row.customer.name}</th>
                      <th>{row.description}</th>
                      <th>{row.price}</th>
                      <th>{row.quantity}</th>
                      <th>{row.merchant.name}</th>
                      <th>{row.merchant.address}</th>
                    </tr>
                    );
                  }) : null }
                </tbody>
              </table>
            </div>
          );
        } else {
          return <div>Loading...</div>
        }
    } else {
      props.history.push("/login");
      return null
    }
}

async function fetchSalesData() {
  try {
    const response = await fetch("http://localhost:3001/items", {
      method: "GET",
      mode: "cors",
      headers: {
        Authorization: JSON.parse(window.localStorage.userAuth).token,
        type: "bearer",
      },
    });
    const json = await response.json();
    if (response.status === 401) {
      // reset local storage
      window.localStorage.clear()
    }
      return json
  } catch(e) {
    console.log("error:", e)
    return e
  }
}
export default withRouter(SalesData);
