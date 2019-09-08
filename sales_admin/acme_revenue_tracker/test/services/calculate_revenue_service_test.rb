class CalculateRevenueServiceTest < ActiveSupport::TestCase
  test "total revenue calculator sums multiple items" do
    acme_sales = [
          AcmeSale.new(item_price: 10, quantity: 1),
          AcmeSale.new(item_price: 10, quantity: 1)
      ]

    calculated_revenue = CalculateRevenueService.call(acme_sales)

    assert(calculated_revenue.success?, 'Expected a successful response')
    assert_equal(20, calculated_revenue.result)
  end

  test "total revenue calculator multiplies price times quantity for each sale" do
    acme_sales = [
        AcmeSale.new(item_price: 10, quantity: 2),
        AcmeSale.new(item_price: 10, quantity: 3)
    ]

    calculated_revenue = CalculateRevenueService.call(acme_sales)

    assert(calculated_revenue.success?, 'Expected a successful response')
    assert_equal(50, calculated_revenue.result)
  end
end