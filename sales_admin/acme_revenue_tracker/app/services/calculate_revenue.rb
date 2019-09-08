class CalculateRevenue
  def self.call(sales)
    revenue = 0

    sales.each do |sale|
      revenue += sale.item_price * sale.quantity
    end

    OpenStruct.new(success?: true, result: revenue, error_message: nil)
  end
end