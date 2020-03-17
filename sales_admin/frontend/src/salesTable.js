import React, { Component } from 'react'

class SalesTable extends Component {

   constructor(props) {
      super(props)
      this.state = {
         sales: []
      }
   }

   uploadSalesData = event => {
    const data = new FormData()
    data.append('file', event.target.files[0])
    fetch("http://localhost:8080/sales/", {
          body: data,
          method: "POST"
      }).then(res => res.json())
      .then((data) => {
        if (data == null) {
          this.setState({ sales: [] })
        } else {
          this.setState({ sales: data })
        }
      })
    }

   componentDidMount() {
     fetch("http://localhost:8080/sales/")
            .then(res => res.json())
            .then((data) => {
              if (data == null) {
                this.setState({ sales: [] })
              } else {
                this.setState({ sales: data })
              }
            })
            .catch(console.log)
   }

   renderTableHeader() {
      let header = Object.keys(this.state.sales[0])
      return header.map((key, index) => {
        var columnName = key.replace("_", " ")
        columnName = columnName.toLowerCase()
            .split(' ')
            .map((s) => s.charAt(0).toUpperCase() + s.substring(1))
            .join(' ')
        return <th key={index}>{columnName}</th>
      })
   }

   renderSalesData() {
      return this.state.sales.map((sale, index) => {
         const {customer_name, item_description, item_price, quantity,  merchant_name, merchant_address} = sale
         return (
            <tr key={index}>
               <td>{customer_name}</td>
               <td>{item_description}</td>
               <td>{item_price}</td>
               <td>{quantity}</td>
               <td>{merchant_name}</td>
               <td>{merchant_address}</td>
            </tr>
         )
      })
   }

   render() {
     if (this.state.sales.length > 0) {
       return (
          <div>
          <h4 id='title'>Update Sales Data (must be in .csv form)</h4>
          <input type="file" name="file" onChange={this.uploadSalesData}/>
             <h1 id='title'>Sales Data</h1>
             <table id='sales'>
                <tbody>
                <tr>{this.renderTableHeader()}</tr>
                {this.renderSalesData()}
                </tbody>
             </table>
          </div>
       )
     } else {
       return (
          <div>
          <h4 id='title'>Add Sales Data (must be in .csv form)</h4>
          <input type="file" name="file" onChange={this.uploadSalesData}/>
          </div>
       )
     }
   }
}

export default SalesTable
