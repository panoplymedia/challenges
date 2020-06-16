Rails.application.routes.draw do
  # For details on the DSL available within this file, see http://guides.rubyonrails.org/routing.html
  resources :users, param: :_username
  post '/auth/login', to: 'authentication#login'
  delete '/auth/logout', to: 'authentication#destroy'
  scope '/api' do
    resources :merchants
    resources :customers
    resources :items
  end
  #last...
  get '/*a', to: 'application#not_found'
end
