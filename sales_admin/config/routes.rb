Rails.application.routes.draw do

  root 'sales_upload#index'

  get 'sales_upload/index'
  get 'sales_upload/upload'
  # For details on the DSL available within this file, see https://guides.rubyonrails.org/routing.html
end
