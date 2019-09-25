import React, { Component } from 'react';
import { Table } from 'semantic-ui-react'

export default class Sales extends Component {

    showCellData = () => {
        const cellData = this.props.salesData
        console.log( cellData)
        if(cellData !== undefined){
            return cellData.map( (cell, index) => {
                return(
                    <Table.Row key={index}>
                        <Table.Cell>{cell.customer_name}</Table.Cell> 
                        <Table.Cell>{cell.item_description}</Table.Cell> 
                        <Table.Cell>${cell.item_price}.00</Table.Cell> 
                        <Table.Cell>{cell.quantity}</Table.Cell> 
                        <Table.Cell>{cell.merchant_name}</Table.Cell> 
                        <Table.Cell>{cell.merchant_address}</Table.Cell> 
                    </Table.Row>
                )
            })
        }else {
			return (
                <Table.Row>
                <Table.Cell>No Data Upload CSV</Table.Cell> 
                <Table.Cell>No Data Upload CSV</Table.Cell> 
                <Table.Cell>No Data Upload CSV</Table.Cell> 
                <Table.Cell>No Data Upload CSV</Table.Cell> 
                <Table.Cell>No Data Upload CSV</Table.Cell> 
                <Table.Cell>No Data Upload CSV</Table.Cell> 
            </Table.Row>
			);
		}
    }

    render(){
        return (
            <div className="csv-data">
                <Table celled>
                    <Table.Header>
                    <Table.Row>
                        <Table.HeaderCell>Customer Name</Table.HeaderCell>
                        <Table.HeaderCell>Item Description</Table.HeaderCell>
                        <Table.HeaderCell>Item Price</Table.HeaderCell>
                        <Table.HeaderCell>Quantity</Table.HeaderCell>
                        <Table.HeaderCell>Merchant Name</Table.HeaderCell>
                        <Table.HeaderCell>Merchant Address</Table.HeaderCell>
                    </Table.Row>
                    </Table.Header>

                    <Table.Body>
                        {this.showCellData()}
                    </Table.Body>

                    <Table.Footer>
                    <Table.Row>
                        <Table.HeaderCell colSpan='6'>
                        </Table.HeaderCell>
                    </Table.Row>
                    </Table.Footer>
                </Table>
            </div>
        )
      }
};