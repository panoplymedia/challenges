Rails.application.routes.draw do
  # For details on the DSL available within this file, see http://guides.rubyonrails.org/routing.html
  resources :users, param: :_username
  post '/users/register', to: 'users#register'
  post '/auth/login', to: 'authentication#login'
  delete '/auth/logout', to: 'authentication#destroy'
  resources :items
  scope '/api' do
    post 'upload_csv', to: 'items#post_csv'
  end
  #last...
  get '/*a', to: 'application#not_found'
end
