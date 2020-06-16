class AuthenticationController < ApplicationController
  before_action :authorize_request, except: :login

  def login
    @user = User.find_by_email(params[:email])
    if @user&.authenticate(params[:password])
      token = JsonWebToken.encode(user_id: @user.id)
      cookies.signed[:jwt] = { value: token, httponly: true, expires: 1.hour.from_now }
      render json: { user: @user.email, id: @user.id }, status: :ok
    else
      render json: { error: 'unauthorized' }, status: :unauthorized
    end
  end

  def destroy
    cookies.delete(:jwt)
    render json: { message: "logged out", status: :ok }
  end

  private

  def login_params
    params.permit(:email, :password)
  end

end
