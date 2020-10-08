class Upload < ApplicationRecord
  has_many :sales

  def total_revenue
    @total_revenue ||= sales.map{ |s| s.revenue }.sum
  end
end
