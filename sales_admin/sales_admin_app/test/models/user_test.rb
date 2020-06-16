require 'test_helper'

class UserTest < ActiveSupport::TestCase
  test "should create a user with a valid email and password" do
    user = User.create(email: "test@mail.com", password: "top_s3cr37")
    assert user.save
  end
end
