require "application_system_test_case"

class AcmeSalesDataUploadsTest < ApplicationSystemTestCase
  test "uploading the acme sales data csv" do
    visit root_url

    attach_file('acme_sales_csv', 'test/system/salesdata.csv')
    click_on 'Upload'

    total_revenue = 526.45 # calculated from example csv
    assert_text total_revenue
  end
end
