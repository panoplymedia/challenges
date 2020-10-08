class SalesUploadController < ApplicationController
  def index
  end

  def upload

    require 'csv'

    file = params[:csv].tempfile

    table = CSV.parse(File.read(file), headers: true)

    puts table.headers

    ActiveRecord::Base.transaction do

      db_upload = Upload.create! filename: file

      table.each do |sale|

        customer = Customer.find_or_create_by(name: sale['Customer Name'])

        merchant = Merchant.find_or_create_by(name: sale['Merchant Name'], address: sale['Merchant Address'])

        item = Item.find_or_create_by(description: sale['Item Description'], price: sale['Item Price'], merchant: merchant)

        sale = Sale.find_or_create_by(quantity: sale['Quantity'], customer: customer, item: item, upload: db_upload)
      end

      puts "Sales: #{db_upload.total_revenue}"
    end

    redirect_to root_path
  end
end
