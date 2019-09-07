class TotalRevenueCalculatorTest < ActiveSupport::TestCase
  test "total revenue calculator sums multiple items" do
    acme_sales = [
          AcmeSale.new(item_price: 10, quantity: 1),
          AcmeSale.new(item_price: 10, quantity: 1)
      ]

    result = TotalRevenueCalculator.call(acme_sales)

    assert_equal(20, result)
  end

  test "total revenue calculator multiplies price times quantity for each sale" do
    acme_sales = [
        AcmeSale.new(item_price: 10, quantity: 2),
        AcmeSale.new(item_price: 10, quantity: 3)
    ]

    result = TotalRevenueCalculator.call(acme_sales)

    assert_equal(50, result)
  end
end