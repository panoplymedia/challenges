class ItemsController < ApplicationController
  before_action :authorize_request
  require 'csv'

  def index
    @items = Item.all.includes(:merchant, :customer)
  end

  def post_csv
    file = params[:file]
    CSV.foreach(file.path, headers: true, header_converters: :symbol) do |row|
      # parse row
      # create models with relevant params
      # update if exists
      row = row.to_hash

      item = Item.where(description: row[:item_description]).take
      customer = Customer.where(name: row[:customer_name]).take
      merchant = Merchant.where(name: row[:merchant_name]).take

      if customer.nil?
        customer = Customer.create!(name: row[:customer_name])
      else
        customer.update_attributes(name: customer.name)
      end

      if merchant.nil?
        merchant = Merchant.create!(name: row[:merchant_name], address: row[:merchant_address])
      else
        merchant.update_attributes(name: row[:merchant_name], address: row[:merchant_address])
      end

      if item.nil?
        Item.create!(
          description: row[:item_description],
          price: row[:item_price],
          quantity: row[:quantity],
          merchant_id: merchant.id,
          customer_id: customer.id
        )
      else
        item.update_attributes(
          description: row[:item_description],
          price: row[:item_price],
          quantity: row[:quantity],
          merchant_id: merchant.id,
          customer_id: customer.id
        )
      end
    end
    render json: { message: "successfully uploaded data" }, status: :ok 
  end
end
