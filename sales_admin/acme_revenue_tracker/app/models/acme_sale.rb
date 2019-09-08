class AcmeSale < ApplicationRecord
  belongs_to :acme_report

  validates :item_price, presence: true
  validates :quantity, presence: true
end
