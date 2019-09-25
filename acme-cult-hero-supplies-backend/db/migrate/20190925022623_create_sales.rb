class CreateSales < ActiveRecord::Migration[5.2]
  def change
    create_table :sales do |t|
      t.string :customer_name
      t.string :item_description
      t.integer :item_price
      t.integer :quantity
      t.string :merchant_name
      t.string :merchant_address

      t.timestamps
    end
  end
end
