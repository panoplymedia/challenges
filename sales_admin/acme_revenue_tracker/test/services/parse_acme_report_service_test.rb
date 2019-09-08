class ParseAcmeReportServiceTest < ActiveSupport::TestCase
  include ActionDispatch::TestProcess::FixtureFile

  test "acme report parser returns an AcmeSale list" do
    upload = fixture_file_upload "#{Rails.root}/test/fixtures/files/salesdata.csv"
    parsed_upload = ParseAcmeReportService.call(upload)

    expected_number_of_acme_sales = 5 # rows in example csv

    assert(parsed_upload.success?, 'Expected a successful response')
    assert_equal('salesdata.csv', parsed_upload.result.name)

    acme_sales = parsed_upload.result.acme_sales
    assert_equal(expected_number_of_acme_sales, acme_sales.size)
    assert_equal(AcmeSale, acme_sales.first.class)
    assert_equal('Jack Burton', acme_sales.first.customer_name)
    assert_equal('Premium Cowboy Boots', acme_sales.first.item_description)
    assert_equal(149.99, acme_sales.first.item_price)
    assert_equal(1, acme_sales.first.quantity)
    assert_equal('Carpenter Outfitters', acme_sales.first.merchant_name)
    assert_equal('99 Factory Drive', acme_sales.first.merchant_address)
  end

  test "acme report parser returns an error if file is missing" do
    parsed_upload = ParseAcmeReportService.call(nil)

    refute(parsed_upload.success?, 'Expected an unsuccessful response')
    assert_equal(ParseAcmeReportService::DEFAULT_ERROR_MESSAGE, parsed_upload.error_message)
  end
end