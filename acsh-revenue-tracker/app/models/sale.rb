class Sale < ApplicationRecord
  belongs_to :customer
  belongs_to :item

  validates :quantity, presence: true
end
