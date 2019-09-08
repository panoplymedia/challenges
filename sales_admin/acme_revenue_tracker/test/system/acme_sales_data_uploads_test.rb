require "application_system_test_case"

class AcmeSalesDataUploadsTest < ApplicationSystemTestCase
  test "when a user uploads a new Acme Report, the total revenue is displayed" do
    visit root_url

    total_revenue = 526.45 # calculated from example csv
    refute_text total_revenue

    attach_file('acme_sales_csv', 'test/fixtures/files/salesdata.csv')
    click_on 'Upload'

    assert_text total_revenue
    assert_text 'Upload successful'
  end

  test "when a user fails to select a report, they see an error" do
    visit root_url

    click_on 'Upload'

    assert_text ParseAcmeReport::DEFAULT_ERROR_MESSAGE
  end
end
