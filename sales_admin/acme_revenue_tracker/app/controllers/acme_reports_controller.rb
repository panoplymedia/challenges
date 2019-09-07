class AcmeReportsController < ApplicationController
  def create
    acme_sales = AcmeReportParser.call(params['acme_sales_csv'].open)

    acme_sales.each do |acme_sale|
      acme_sale.save!
    end

    redirect_to '/'
  end
end
