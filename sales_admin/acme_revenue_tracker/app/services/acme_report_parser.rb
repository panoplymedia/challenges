require 'csv'

class AcmeReportParser
  def self.call(file)
    parsed_csv = CSV.parse(file, headers:true)

    parsed_csv.map do |row|
      new_sale = AcmeSale.new
      new_sale.customer_name = row['Customer Name']
      new_sale.item_description = row['Item Description']
      new_sale.item_price = row['Item Price']
      new_sale.quantity = row['Quantity']
      new_sale.merchant_name = row['Merchant Name']
      new_sale.merchant_address = row['Merchant Address']
      new_sale
    end
  end
end