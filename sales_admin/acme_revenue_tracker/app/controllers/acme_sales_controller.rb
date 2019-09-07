class AcmeSalesController < ApplicationController
  def index
    @total_revenue = TotalRevenueCalculator.call
  end
end
