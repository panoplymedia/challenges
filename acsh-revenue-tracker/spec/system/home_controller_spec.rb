require 'rails_helper'
require_relative '../support/devise'

RSpec.describe HomeController, type: :controller do

  describe "Home page" do
    # before { login_user }

    it "can view home page" do
      visit root_path
      expect(response).to have_content('ACME')
    end
  end

end
