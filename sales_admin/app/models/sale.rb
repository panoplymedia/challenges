class Sale < ApplicationRecord
  belongs_to :item
  belongs_to :customer
  belongs_to :upload

  def revenue
    @revenue ||= quantity.to_i * item.price.to_f
  end
end
