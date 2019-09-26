FactoryBot.define do
  factory :random_sale, class: Sale do
      customer_name { Faker::Name.name  }
      item_description { Faker::Lorem.sentence }
      item_price { Faker::Number.decimal(l_digits: 2) }
      quantity { Faker::Number.within(range: 1..10) }
      merchant_name { Faker::Company.name }
      merchant_address { Faker::Address.street_name }
    end
  end