class CreateAcmeReports < ActiveRecord::Migration[5.2]
  def change
    create_table :acme_reports do |t|
      t.string :name

      t.timestamps
    end
  end
end
