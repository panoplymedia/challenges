class AcmeReportsController < ApplicationController
  def index
    @total_revenue = TotalRevenueCalculator.call(AcmeSale.all)
  end

  def create
    if params['acme_sales_csv']
      acme_sales = AcmeReportParser.call(params['acme_sales_csv'].open)

      acme_sales.each do |acme_sale|
        acme_sale.save!
      end
    else
      flash[:error] = 'Please select a file to upload'
    end

    redirect_to '/'
  end
end
