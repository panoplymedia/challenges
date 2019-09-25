import React, { Component } from 'react';
import './App.css';
import DataPage from './components/DataPage'
import Header from './components/Header'
import Sales from './components/Sales'

export default class App extends Component {
  constructor() {
		super();
		this.state = {
      csv: null,
      salesData: []
		};
  }

  componentDidMount() {
    fetch('http://localhost:3000/api/v1/sales')
      .then(data => data.json())
      .then((data) => { this.setState({ salesData: data }) }); 
  }

  //Method that changes the state of the application to the submitted CSV file
  handleCSV = (event) => {
    this.setState({
      csv: event.target.files[0]
    })
  }

  //Gets sales data after a post is made & then renders component with state change
  newSalesData() {
    fetch('http://localhost:3000/api/v1/sales')
      .then(data => data.json())
      .then((data) => { this.setState({ salesData: data }) }); 
  }

  //Method that uploads file to backend, used axios here instead of fetch
  handleFileUpload = () =>{
    const data = new FormData() 
    data.append('file', this.state.csv)


    let options = {
      method: 'POST',
      body: data
    }
    fetch(`http://localhost:3000/api/v1/csvs`, options)
      .then( () => this.newSalesData())
      .catch(error => {
        console.log("In error logging", error)
    })
     
  }


  render(){
    return (
      <div className="App">
            <Header/>
            <DataPage
            handleCSV={this.handleCSV}
            handleFileUpload={this.handleFileUpload}
            />
            <Sales
            salesData={this.state.salesData}
            calculateTotal={this.calculateTotal}
            />
      </div>
    );
  }
}

