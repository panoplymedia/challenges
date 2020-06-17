class Customer < ApplicationRecord
  has_many :items
  has_many :merchants, through: :items 
end
