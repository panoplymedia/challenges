class CreateItems < ActiveRecord::Migration[6.0]
  def change
    create_table :items do |t|
      t.references :merchant, null: false, foreign_key: true
      t.string :description, null: false
      t.float :price, null: false
      t.timestamps
    end
  end
end
