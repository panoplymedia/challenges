class Merchant < ApplicationRecord
  has_many :items
  has_many :customers, through: :items
end
