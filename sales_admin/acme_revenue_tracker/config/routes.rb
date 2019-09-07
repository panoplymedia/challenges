Rails.application.routes.draw do
  root to: 'acme_reports#index'

  resources :acme_reports, only: [:create, :index]
end
