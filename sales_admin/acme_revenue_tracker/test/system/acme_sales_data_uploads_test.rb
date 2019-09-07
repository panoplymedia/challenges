require "application_system_test_case"

class AcmeSalesDataUploadsTest < ApplicationSystemTestCase
  test "uploading the acme sales data csv" do
    visit root_url

    total_revenue = 526.45 # calculated from example csv
    refute_text total_revenue

    attach_file('acme_sales_csv', 'test/fixtures/files/salesdata.csv')
    click_on 'Upload'

    assert_text total_revenue
  end
end
