class AcmeReportsController < ApplicationController
  def index
    @total_revenue = TotalRevenueCalculator.call(AcmeSale.all)
  end

  def create
    parsed_upload = AcmeReportParser.call(params['acme_sales_csv'])

    if parsed_upload.success?
      parsed_upload.result.each do |acme_sale|
        acme_sale.save!
      end

      flash[:success] = 'Upload successful'
    else
      flash[:error] = parsed_upload.error_message
    end

    redirect_to '/'
  end
end
