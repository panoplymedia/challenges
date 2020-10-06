require 'test_helper'

class SalesUploadControllerTest < ActionDispatch::IntegrationTest
  test "should get index" do
    get sales_upload_index_url
    assert_response :success
  end

  test "should get upload" do
    get sales_upload_upload_url
    assert_response :success
  end

end
