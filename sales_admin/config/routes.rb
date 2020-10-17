Rails.application.routes.draw do

  get 'signup', to: 'users#new', as: 'signup'

  get 'login', to: 'sessions#new', as: 'login'

  get 'logout', to: 'sessions#destroy', as: 'logout'



  root 'sales_upload#index'


  get 'sales_upload/index'

  get 'sales_upload/:id/overview', :controller => 'sales_upload', :action => 'overview', :as => 'upload_overview'

  post 'sales_upload/upload', :upload_csv


  resources :sessions

  resources :users

  # For details on the DSL available within this file, see https://guides.rubyonrails.org/routing.html
end
