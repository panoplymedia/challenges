class Item < ApplicationRecord
  has_many :customers, through: :sales
  belongs_to :merchant

  validates :description, presence: true
  validates :price, presence: true
end
