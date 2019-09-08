class AddAcmeReportToAcmeSales < ActiveRecord::Migration[5.2]
  def change
    add_reference :acme_sales, :acme_report, foreign_key: true
  end
end
