require 'test_helper'

class AuthenticationControllerTest < ActionDispatch::IntegrationTest
  test "should login with a valid user email and password" do
    user = User.create(email: 'test@mail.com', password: 'tops3cr37')
    post('/auth/login', params: { email: user.email, password: user.password }) 
    assert_response :success
    assert_equal "test@mail.com", JSON.parse(response.body)["user"]
  end

  test "should not login without a valid email and password" do
    post('/auth/login', params: { email: "invalid", password: "is_wrong" })
    assert_response :unauthorized
    assert_equal "unauthorized", JSON.parse(response.body)["error"]
  end
end
