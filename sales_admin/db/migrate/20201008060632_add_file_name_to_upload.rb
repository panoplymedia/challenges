class AddFileNameToUpload < ActiveRecord::Migration[6.0]
  def change
    add_column :uploads, :filename, :string
  end
end
