require 'test_helper'
require 'rack/test'
class ItemsControllerTest < ActionDispatch::IntegrationTest
  setup do
    @user = User.create!(email: 'test@email.com', password: 'pwpwpw')
    @token = JsonWebToken.encode(user_id: @user.id)
  end
  test "should return a list of Items and their associated customer and merchant" do
    merchant = Merchant.create(name: 'CVS', address: 'Providence, RI')
    customer = Customer.create(name: 'Bob')
    item = Item.create({ merchant: merchant, customer: customer, description: 'Bandage', price: 1.50, quantity: 10 })
    get('/items.json', headers: { Authorization: @token })
    assert_response :success
  end

  test "should upload a CSV and import it into the database" do
    file = Rack::Test::UploadedFile.new("../salesdata.csv", "mime/type")
    post('/api/upload_csv', params: { file: file }, headers: { Authorization: @token }, as: 'multipart/form-data' )
    assert_response :success
    assert_equal "successfully uploaded data", JSON.parse(response.body)['message']
  end
end
