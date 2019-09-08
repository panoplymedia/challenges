class AcmeSaleTest < ActiveSupport::TestCase
  test "it is not valid if it is missing an item_price" do
    acme_report = AcmeReport.new
    invalid_record = AcmeSale.new(quantity: 1, acme_report: acme_report)

    refute(invalid_record.valid?, 'Expected record to be invalid')
  end

  test "it is not valid if it is missing a quantity" do
    acme_report = AcmeReport.new
    invalid_record = AcmeSale.new(item_price: 100, acme_report: acme_report)

    refute(invalid_record.valid?, 'Expected record to be invalid')
  end

  test "it is valid" do
    acme_report = AcmeReport.new
    invalid_record = AcmeSale.new(item_price: 100, quantity: 1, acme_report: acme_report)

    assert(invalid_record.valid?, 'Expected record to be valid')
  end
end