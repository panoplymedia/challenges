class UsersController < ApplicationController

  skip_before_action :authorize, only: [:new, :create]


  def new
    @user = User.new
  end

  def create
    @user = User.new(params[:user].permit!)
    if @user.save
      session[:user_id] = @user.id
      redirect_to root_url, notice: "Thank you for signing up!"
    else
      render "new"
    end
  end

end
