import React, { Component } from 'react';
import './App.css';
import axios from 'axios';

import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Paper from '@material-ui/core/Paper';

import { BASE_URL } from './constants';

class App extends Component {
  constructor(){
    super();
    this.state = {
        auth: false,
        file: '',
        isUploading: false,
        progress: 0,
        salesData: '',
        inputKey: Date.now(),
        revenue: 0,
        JSX: ''
    };
  }

  componentDidMount(){


    this.initializePage();
    this.authCheck();


  }

  authCheck = () => {
    axios({
      url: `${BASE_URL}/authcheck`,
      method: 'GET',
      withCredentials: true
    }).then(res => { 
      console.log('authcheck', res) 
      if(res.data.auth === true) this.setState({ auth: true, username: res.data.username });
    });
  }
  
  initializePage = () => {
    this.getSalesData(this.getTableJSX);
  }
  getTotalRevenue = () => {
    axios({
      method: 'GET',
      url: `${BASE_URL}/api/salesdata/revenue`,
      withCredentials: true
    }).then(res => {
        // console.log(res.data);
        this.setState({ revenue : res.data.revenue});
        this.initializePage();
    }).catch(err => console.log(err))


  }

  

  getSalesData = (cb = '') => {
      setTimeout(() => {
              
            axios({
              method: 'GET',
              url: `${BASE_URL}/api/salesdata`,
              withCredentials: true
            }).then(res => {
                if(res.data.rowCount === 0) {
                  this.setState({ salesData: '' })
                  console.log('set salesData to nothing');
                }
                else {
                  console.log('getSalesData', res.data);
                  if(!cb)
                    this.setState({ salesData: res.data })
                  else {
                    this.setState({ salesData: res.data }, cb);
                  }
                  
                }
            }).catch(err => console.log('error', err));

          }, 200);
  }

  setFile = (event) => {
    this.setState({ file: event.target.files[0]});
  }

  handleUploaderResponse = (res) => {
      console.log(res);
      this.setState({
        progress: 0,
        file: '',
        inputKey: Date.now()
      });
        this.getSalesData(this.getTableJSX);

  }

  handleUploader = (event) => {
      event.preventDefault();
      let errs = [];
      let response = '';
      if(!this.state.file) 
          errs.push('File is required');

      if(errs.length > 0) {
          response = JSON.stringify(errs);
          this.handleUploaderResponse(response);

      } else {

        let fd = new FormData();
        fd.append("csv", this.state.file);

        console.log(this.state.file);

        axios({
          method: 'POST',
          url: `${BASE_URL}/api/upload`,
          onUploadProgress: (ProgressEvent) => {
            let percentage = Math.round( (ProgressEvent.loaded * 100) / ProgressEvent.total );
            this.setState({ progress: percentage }, this.logPercent)
          },
          data: fd,
          
          withCredentials: true
        }).then(res => this.handleUploaderResponse(res))
        .catch(err => this.handleUploaderResponse(JSON.stringify(err)));

      }

  }

  getTableJSX = () => {
      let tableData  = [...this.state.salesData];
      console.log('td', tableData);
      // console.log('tabledata', tableData)
      let JSX = tableData.map(sale => {
        return(
        <TableRow key={sale.id}>
        <TableCell component="th" scope="row">{sale.customer_name}</TableCell>
        {/* <TableCell align="right">{sale.customer_name}</TableCell> */}
        <TableCell align="right">{sale.description}</TableCell>
        <TableCell align="right">{sale.price}</TableCell>
        <TableCell align="right">{sale.quantity}</TableCell>
        <TableCell align="right">{sale.merchant_name}</TableCell>
        <TableCell align="right">{sale.merchant_address}</TableCell>
      </TableRow>
        );
      });
      this.setState({ JSX: JSX});
  }

  logPercent = () => console.log(this.state.progress);

  clearDB = () => {
    window.confirm('This will erase all data in the database and is irreversible. Are you sure?');
    axios({
      method: 'GET',
      url: `${BASE_URL}/api/deleteall`,
      withCredentials: true
    }).then(res => { 
        if(res.data.Success==="deleted")
        { 
           this.initializePage();
           this.setState({ JSX: '', revenue: 0 });
        }
    });
  }
  
  render() {

      if(this.state.auth === true)
        {
          return (
            <div className="App">
              <div className="App__Navbar">
                <div className="Company">ACHS</div>
              <div className="Log">Logged in as {this.state.username}</div>

              </div>

              <div className="App_Dashboard">

              <div className="App__Dashboard__Controls">

                <div className="App__Dashboard__Controls__Uploader">
                  <form>
                    <label>CSV Uploader</label> {/*<br /> */}
                    <input type="file" key={this.state.inputKey} accept="text/csv" name="csv" onChange={this.setFile}/>
                    {/* <br /> */}
                    <button onClick={this.handleUploader}>Submit</button>
                  </form>

                </div>
                <div className="App__Dashboard__Controls__Revenue">Total: ${this.state.revenue}</div>


              </div>
              <div className="App__Dashboard__Buttons">
              <button onClick={this.getTotalRevenue}>Calculate Revenue</button>
                <button onClick={this.clearDB}>Clear Database</button>
                


              </div>



              <div className="App__Dashboard__Table">

              <TableContainer component={Paper}>
                <Table className="App__Dashboard__Table__MUI" aria-label="simple table">
                  <TableHead >
                    <TableRow >
                      <TableCell>Customer Name</TableCell>
                      <TableCell align="right">Description</TableCell>
                      <TableCell align="right">Price</TableCell>
                      <TableCell align="right">Quantity</TableCell>
                      <TableCell align="right">Merchant Name</TableCell>
                      <TableCell align="right">Merchant Address</TableCell>
                    </TableRow>
                  </TableHead>
                  <TableBody>


                  {this.state.JSX}

                  </TableBody>
                </Table>
              </TableContainer>


              </div>

              </div>

            </div>
          );
      }
      else {
          return(
            <div className="App">

            
            <div className="App__AuthBox">
                <div className="App__AuthBox__Header">
                <h4>Please Login:</h4>
                </div>

                    <a className="GitAuth" href={`${BASE_URL}/login`}>Login with GitHub</a>
            </div>
            </div>
          );
      }
    }
}

export default App;
