class AcmeReport < ApplicationRecord
  has_many :acme_sales
  accepts_nested_attributes_for :acme_sales
  validates_associated :acme_sales
end
