class AcmeReportsController < ApplicationController
  def index
    calculated_revenue = CalculateRevenue.call(AcmeSale.all)
    @total_revenue = calculated_revenue.result
  end

  def create
    parsed_upload = ParseAcmeReport.call(params['acme_sales_csv'])

    if parsed_upload.success?
      if !parsed_upload.result.save
        flash[:error] = ParseAcmeReport::DEFAULT_ERROR_MESSAGE
      else
        flash[:success] = 'Upload successful'
      end

    else
      flash[:error] = parsed_upload.error_message
    end

    redirect_to '/'
  end
end
