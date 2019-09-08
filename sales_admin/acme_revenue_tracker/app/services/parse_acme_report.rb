require 'csv'

class ParseAcmeReport
  DEFAULT_ERROR_MESSAGE = 'Select a properly formatted CSV'

  def self.call(file)
    begin
      parsed_csv = CSV.parse(file.open, headers:true)
      original_filename = file.original_filename

      acme_sales = parsed_csv.map do |row|
        new_sale = AcmeSale.new
        new_sale.customer_name = row['Customer Name']
        new_sale.item_description = row['Item Description']
        new_sale.item_price = row['Item Price']
        new_sale.quantity = row['Quantity']
        new_sale.merchant_name = row['Merchant Name']
        new_sale.merchant_address = row['Merchant Address']
        new_sale
      end

      acme_report = AcmeReport.new(acme_sales: acme_sales, name: original_filename)

      OpenStruct.new(success?: true, result: acme_report, error_message: nil)
    rescue
      OpenStruct.new(success?: false, result: nil, error_message: DEFAULT_ERROR_MESSAGE)
    end
  end
end