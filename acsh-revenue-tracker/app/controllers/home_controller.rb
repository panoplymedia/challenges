require 'roo'

class HomeController < ApplicationController
  skip_before_action :verify_authenticity_token

  def index
    @sales = Sale.all
  end

  def uploaded
    file = params[:csv].tempfile
    software_data = Roo::Spreadsheet.open file
    sheet = software_data.sheet(0)
    ActiveRecord::Base.transaction do
      sheet.each do |row|
        next if row[0].blank? || row[0] == 'Customer Name'
        (first_name, last_name) = row[0].split(' ')
        customer = Customer.find_or_create_by(first_name: first_name, last_name: last_name)
        merchant = Merchant.find_or_create_by(name: row[4], address: row[5])
        item = Item.find_or_create_by(description: row[1], price: row[2], merchant: merchant)
        sale = Sale.find_or_create_by(quantity: 3, customer: customer, item: item)
      end
    end
    redirect_to root_path, notice: "There are now #{Sale.count} sales in the database"
  end
end
