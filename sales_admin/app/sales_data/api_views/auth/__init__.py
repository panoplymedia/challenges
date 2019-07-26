from flask import Blueprint
from .register_view import RegistrationView
from .login_view import LoginView


# This instance of a Blueprint that represents the authentication blueprint
auth_bp = Blueprint('auth', __name__, url_prefix="/auth")


# Define the rule for the registration url --->  /auth/register
registration_view = RegistrationView.as_view('register_view')
# Then add the rule to the blueprint
auth_bp.add_url_rule(
    '/register',
    view_func=registration_view,
    methods=['POST'])


# Define the rule for the registration url --->  /auth/login
login_view = LoginView.as_view('login_view')
# Then add the rule to the blueprint
auth_bp.add_url_rule(
    '/login',
    view_func=login_view,
    methods=['POST']
)
