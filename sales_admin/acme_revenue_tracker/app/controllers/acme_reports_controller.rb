class AcmeReportsController < ApplicationController
  def index
    calculated_revenue = CalculateRevenueService.call(AcmeSale.all)
    @total_revenue = calculated_revenue.result
  end

  def create
    parsed_acme_report = ParseAcmeReportService.call(params['acme_sales_csv'])

    if parsed_acme_report.success? && parsed_acme_report.result.save
      flash[:success] = 'Upload successful'
    else
      flash[:error] = parsed_acme_report.error_message || ParseAcmeReportService::DEFAULT_ERROR_MESSAGE
    end

    redirect_to '/'
  end
end
