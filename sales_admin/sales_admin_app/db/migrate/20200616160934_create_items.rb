class CreateItems < ActiveRecord::Migration[5.2]
  def change
    create_table :items do |t|
      t.belongs_to :merchant
      t.belongs_to :customer
      t.decimal :price, precision: 10, scale: 2
      t.integer :quantity
      t.string :description
      t.timestamps
    end
  end
end
