require 'csv'

class Api::V1::CsvsController < ApplicationController
    def create
        csv_text = params["file"].tempfile
        CSV.foreach(csv_text).with_index do |row, i|
            next if i == 0
            Sale.create(customer_name: row[0], item_description: row[1], item_price: row[2], quantity: row[3], 
                merchant_name: row[4], merchant_address: row[5])
          end
    end
    
    def show
    
    end

    private
   
end