Rails.application.routes.draw do

  root 'sales_upload#index'

  get 'sales_upload/index'

  get 'sales_upload/:id/overview', :controller => 'sales_upload', :action => 'overview', :as => 'upload_overview'

  post 'sales_upload/upload', :upload_csv
  # For details on the DSL available within this file, see https://guides.rubyonrails.org/routing.html
end
