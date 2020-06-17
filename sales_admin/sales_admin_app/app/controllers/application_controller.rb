class ApplicationController < ActionController::API
  include ActionController::ImplicitRender
  include ActionController::MimeResponds
  include ActionView::Layouts

  def not_found
    render json: { error: 'not_found' }
  end

  def authorize_request
    jwt = request.headers['Authorization']
    begin
      @decoded = JsonWebToken.decode(jwt)
      @current_user = User.find(@decoded[:user_id])
    rescue ActiveRecord::RecordNotFound => e
      render json: { errors: e.message }, status: :unauthorized
    rescue JWT::DecodeError => e
      render json: { errors: e.message }, status: :unauthorized
    end
  end
end
