class AcmeReportParserTest < ActiveSupport::TestCase
  include ActionDispatch::TestProcess::FixtureFile

  test "acme report parser returns an AcmeSale list" do
    upload = fixture_file_upload "#{Rails.root}/test/fixtures/files/salesdata.csv"
    parsed_upload = AcmeReportParser.call(upload)

    expected_number_of_acme_sales = 5 # rows in example csv
    assert(parsed_upload.success?, 'Expected a successful response')
    assert_equal(parsed_upload.result.count, expected_number_of_acme_sales)
    assert_equal(parsed_upload.result.first.class, AcmeSale)
    assert_equal(parsed_upload.result.first.customer_name, 'Jack Burton')
    assert_equal(parsed_upload.result.first.item_description, 'Premium Cowboy Boots')
    assert_equal(parsed_upload.result.first.item_price, 149.99)
    assert_equal(parsed_upload.result.first.quantity, 1)
    assert_equal(parsed_upload.result.first.merchant_name, 'Carpenter Outfitters')
    assert_equal(parsed_upload.result.first.merchant_address, '99 Factory Drive')
  end

  test "acme report parser returns an error if file is missing" do
    parsed_upload = AcmeReportParser.call(nil)

    refute(parsed_upload.success?, 'Expected an unsuccessful response')
    assert_equal(parsed_upload.error_message, AcmeReportParser::DEFAULT_ERROR_MESSAGE)
  end
end