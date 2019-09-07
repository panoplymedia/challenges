class AcmeReportParserTest < ActiveSupport::TestCase
  test "acme report parser returns an AcmeSale list" do
    csv = File.read('test/fixtures/files/salesdata.csv')
    result = AcmeReportParser.call(csv)

    expected_number_of_acme_sales = 5 # rows is example csv
    assert_equal(result.count, expected_number_of_acme_sales)
    assert_equal(result.first.class, AcmeSale)
    assert_equal(result.first.customer_name, 'Jack Burton')
    assert_equal(result.first.item_description, 'Premium Cowboy Boots')
    assert_equal(result.first.item_price, 149.99)
    assert_equal(result.first.quantity, 1)
    assert_equal(result.first.merchant_name, 'Carpenter Outfitters')
    assert_equal(result.first.merchant_address, '99 Factory Drive')
  end
end