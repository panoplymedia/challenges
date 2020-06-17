class UsersController < ApplicationController
  def register
    user = User.where(email: params[:email]).take
    user = User.create!(email: params[:email], password: params[:password]) if user.nil?
    render json: user
  end
end
