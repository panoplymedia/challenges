class SalesUploadController < ApplicationController
  def index
  end

  def upload

    require 'csv'

    table = CSV.parse(File.read(params[:csv].tempfile), headers: true)

    puts table.headers

    ActiveRecord::Base.transaction do

      table.each do |sale|

        customer = Customer.find_or_create_by(name: sale['Customer Name'])

        merchant = Merchant.find_or_create_by(name: sale['Merchant Name'], address: sale['Merchant Address'])

        item = Item.find_or_create_by(description: sale['Item Description'], price: sale['Item Price'], merchant: merchant)

        sale = Sale.find_or_create_by(quantity: sale['Quantity'], customer: customer, item: item)

      end

    end

    redirect_to root_path
  end
end
