class AddUploadToSale < ActiveRecord::Migration[6.0]
  def change
    add_reference :sales, :upload, foreign_key: true
  end
end
