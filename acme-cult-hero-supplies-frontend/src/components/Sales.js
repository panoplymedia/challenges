import React, { Component } from 'react';
import { Table } from 'semantic-ui-react'

export default class Sales extends Component {

    showCellData = () => {
        const cellData = this.props.salesData
        if(cellData.length > 0 ){
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
    
  calculateTotal(){
    const salesData = this.props.salesData
    let totalSales = 0;
    let saleAmount =0;
    if(salesData.length > 0){
      salesData.map( sale => {
        saleAmount = sale.quantity * sale.item_price
        totalSales += saleAmount
        return totalSales
      })
    }
    return totalSales
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

                    <Table.Footer className="total">
                    <Table.Row>
                        <Table.HeaderCell colSpan='6'>
                            Total Sales: ${this.calculateTotal() }
                        </Table.HeaderCell>
                    </Table.Row>
                    </Table.Footer>
                </Table>
            </div>
        )
      }
};