Rails.application.routes.draw do

  root 'home#index'

  devise_for :users
  
  post '/upload', to: 'home#uploaded'
end
