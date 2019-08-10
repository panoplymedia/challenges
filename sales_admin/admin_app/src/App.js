import React from 'react';
import { Form, Button, FormGroup, Card, CardBody, CardText } from 'reactstrap';
import ReactFileReader from 'react-file-reader';
import Papa from 'papaparse';
import './App.css';

export default class App extends React.Component {

  constructor(props) {
    super(props);
    this.state = {
      fileContent: {},
      fileInfo: {},
      revenueTotal: 0
    }
    this.handleSubmit = this.handleSubmit.bind(this);
    this.handleFiles = this.handleFiles.bind(this); 
    this.onGenerate = this.onGenerate.bind(this);
    this.updateData = this.updateData.bind(this);
  }

  handleSubmit(e) {
    e.preventDefault();
  }

  handleFiles(files) {
    this.setState({fileInfo: files[0]});
    var reader = new FileReader();
    reader.onload = (e) => {
      this.setState({fileContent: reader.result});
    }
    reader.readAsText(files[0]);
  }

  onGenerate() {
    const { fileInfo } = this.state;
    Papa.parse(fileInfo, {
      complete: this.updateData,
      header: true
    });
  }

  updateData(result) {
    var data = result.data;
    console.log(data);
    let revenueTotal = 0;
    for(var entry of data) {
      if(entry["Item Price"] && entry["Quantity"]) {
        revenueTotal += entry["Item Price"]*entry["Quantity"];
      }
      else {
        alert("Invalid CSV");
        break;
      }
    };
    this.setState({revenueTotal: revenueTotal});
  }

  render() {
    return (
      <div className="App">
        <div className="formContainer">
          <h3>Generate Total Revenue</h3>
          <Form onSubmit={this.handleSubmit}>
            <FormGroup className="formPadding">
              {/*<Input type="file" id="inputFile" className="form-control" multiple=""/>*/}
              <ReactFileReader handleFiles={this.handleFiles} fileTypes={'.csv'} className="setInlineBlock">
                  <Button color="secondary" className='btn'>Upload CSV</Button> 
              </ReactFileReader>
              <div className="setInlineBlock fileName">{this.state.fileInfo.name}</div>
            </FormGroup>
            <div className="submitBtnContainer">
              <Button className="customBtn" size="lg" block onClick={this.onGenerate}>Generate</Button>
            </div>
          </Form> 
          <Card>
            <CardBody>
              <CardText>Total Revenue: ${this.state.revenueTotal.toFixed(2)}</CardText>
            </CardBody>
          </Card>
        </div>
      </div>
    );
  }
}

