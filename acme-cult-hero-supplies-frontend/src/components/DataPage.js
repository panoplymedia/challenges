import React, { Component } from 'react';
import { Input } from 'semantic-ui-react'
import { Button } from 'semantic-ui-react'

export default class DataPage extends Component {


    render(){
        return (
            <div className="csv-data">
                <Input
                    type="file"
                    ref={(input) => { this.filesInput = input }}
                    name="file"
                    icon='file text outline'
                    iconPosition='left'
                    onChange={this.props.handleCSV}
                    />
                    <Button className="btn-upload" content='Upload' onClick={this.props.handleFileUpload} />
            </div>
        )
      }
};