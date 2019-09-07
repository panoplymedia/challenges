Rails.application.routes.draw do
  root to: 'acme_sales#index'

  resources :acme_reports, only: [:create]
end
