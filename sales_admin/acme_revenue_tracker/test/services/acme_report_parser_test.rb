class AcmeReportParserTest < ActiveSupport::TestCase
  include ActionDispatch::TestProcess::FixtureFile

  test "acme report parser returns an AcmeSale list" do
    upload = fixture_file_upload "#{Rails.root}/test/fixtures/files/salesdata.csv"
    parsed_upload = AcmeReportParser.call(upload)

    expected_number_of_acme_sales = 5 # rows in example csv

    assert(parsed_upload.success?, 'Expected a successful response')

    result = parsed_upload.result
    assert_equal(expected_number_of_acme_sales, result.count)
    assert_equal(AcmeSale, result.first.class)
    assert_equal('Jack Burton', result.first.customer_name)
    assert_equal('Premium Cowboy Boots', result.first.item_description)
    assert_equal(149.99, result.first.item_price)
    assert_equal(1, result.first.quantity)
    assert_equal('Carpenter Outfitters', result.first.merchant_name)
    assert_equal('99 Factory Drive', result.first.merchant_address)
  end

  test "acme report parser returns an error if file is missing" do
    parsed_upload = AcmeReportParser.call(nil)

    refute(parsed_upload.success?, 'Expected an unsuccessful response')
    assert_equal(AcmeReportParser::DEFAULT_ERROR_MESSAGE, parsed_upload.error_message)
  end
end