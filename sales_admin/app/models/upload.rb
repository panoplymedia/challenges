class Upload < ApplicationRecord
  has_many :sales

  def total_revenue
    @total_revenue ||= sales.map{ |s| s.revenue }.sum
  end

  def user_email
    return nil unless user_id

    return User.find(user_id).email
  end
end
