class TotalRevenueCalculator
  def self.call(sales)
    total_revenue = 0

    sales.each do |sale|
      total_revenue += sale.item_price * sale.quantity
    end

    OpenStruct.new(success?: true, result: total_revenue, error_message: nil)
  end
end