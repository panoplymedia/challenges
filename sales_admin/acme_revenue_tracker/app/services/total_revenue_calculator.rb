class TotalRevenueCalculator
  def self.call(sales)
    total_revenue = 0

    sales.each do |sale|
      total_revenue += sale.item_price * sale.quantity
    end

    total_revenue
  end
end