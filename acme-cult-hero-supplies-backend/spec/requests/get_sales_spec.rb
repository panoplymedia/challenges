require 'rails_helper'

describe "get all sales route", :type => :request do
  let!(:sales) {FactoryBot.create_list(:random_sale, 3)}
before {get '/api/v1/sales'}
it 'returns all sales' do
    expect(JSON.parse(response.body).size).to eq(3)
  end
it 'returns status code 200' do
    expect(response).to have_http_status(:success)
  end

it 'rspec returns a generic success' do
    expect(3).to eq(3)
  end

end 