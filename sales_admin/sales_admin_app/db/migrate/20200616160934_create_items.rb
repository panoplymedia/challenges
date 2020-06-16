class CreateItems < ActiveRecord::Migration[5.2]
  def change
    create_table :items do |t|
      t.decimal :price, precision: 10, scale: 2
      t.string :description
      t.timestamps
    end
  end
end
