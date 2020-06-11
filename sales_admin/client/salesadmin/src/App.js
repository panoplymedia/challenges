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
        auth: true,
        file: '',
        isUploading: false,
        progress: 0,
        salesData: ''
    };
  }

  componentDidMount(){

    //axios...


  }

  getSalesData = () => {
      axios({
        method: 'GET',
        url: `${BASE_URL}/api/all`,
        // withCredentials: true
      })
  }

  setFile = (event) => {
    this.setState({ file: event.target.files[0]});
  }

  handleUploaderResponse = (res) => {
      console.log(res);
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
        }
          // withCredentials: true
        })

      }

  }

  render() {

    function createData(name, calories, fat, carbs, protein) {
      return { name, calories, fat, carbs, protein };
    }
    
    const rows = [
      createData('Frozen yoghurt', 159, 6.0, 24, 4.0),
      createData('Ice cream sandwich', 237, 9.0, 37, 4.3),
      createData('Eclair', 262, 16.0, 24, 6.0),
      createData('Cupcake', 305, 3.7, 67, 4.3),
      createData('Gingerbread', 356, 16.0, 49, 3.9),
    ];


      if(this.state.auth === true)
        {
          return (
            <div className="App">

              <div className="App_Dashboard">

              <div className="App__Dashboard__Controls">

                <div className="App__Dashboard__Controls__Uploader">
                  <form>
                    <label>CSV Uploader</label> {/*<br /> */}
                    <input type="file" accept="text/csv" name="csv" onChange={this.setFile}/>
                    {/* <br /> */}
                    <button onClick={this.handleUploader}>Submit</button>
                  </form>

                </div>

                <div className="App__Dashboard__Controls__Revenue">Total: $1,595.93</div>

                {/* <div className="App__Dashboard__Controls__ */}



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


                    {rows.map((row) => (
                      <TableRow key={row.name}>
                        <TableCell component="th" scope="row">
                          {row.name}
                        </TableCell>
                        <TableCell align="right">{row.calories}</TableCell>
                        <TableCell align="right">{row.fat}</TableCell>
                        <TableCell align="right">{row.carbs}</TableCell>
                        <TableCell align="right">{row.protein}</TableCell>
                        <TableCell align="right">{row.protein}</TableCell>
                      </TableRow>
                    ))}


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

                <div className="App__AuthBox__Group">
                <label className="">Username:</label>
                <input className="App__AuthBox__Group__Input" type="text" placeholder="demo" />
                </div>
                <div className="App__AuthBox__Group">
                <label className="">Password:</label>
                <input className="App__AuthBox__Group__Input" type="password" placeholder="demo" />
                </div>
                <button className="App__AuthBox__Group__Button">Submit</button>
                </div>
            </div>
          );
      }
    }
}

export default App;
