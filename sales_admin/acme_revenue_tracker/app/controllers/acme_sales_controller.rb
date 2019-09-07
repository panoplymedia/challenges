class AcmeSalesController < ApplicationController
  def index
    @total_revenue = TotalRevenueCalculator.call(AcmeSale.all)
  end
end
