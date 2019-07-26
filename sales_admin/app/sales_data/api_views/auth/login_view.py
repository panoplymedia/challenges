from flask.views import MethodView
from flask import make_response, request, jsonify
from app.sales_data.models.user import User
from app.sales_data.utils.helper import verify_fields, server_error


class LoginView(MethodView):
    """This class-based view handles user login and access token generation."""

    def post(self):
        """Handle POST request for this view. Url ---> /auth/login"""

        # Get the json post data
        post_data = request.get_json()
        verify_fields(["username", "password"], post_data)

        try:
            # Get the user object using their email (unique to every user)
            user = User.query.filter_by(username=post_data["username"]).first()

            # Try to authenticate the found user using their password
            if user and user.password_is_valid(post_data["password"]):
                # Generate the access token. This will be used as the authorization header
                access_token = User.generate_token(user.id)
                if access_token:
                    response = {
                        'message': 'Successfully logged in.',
                        'access_token': access_token.decode()
                    }
                    return make_response(jsonify(response)), 200
            else:
                # User does not exist. Therefore, we return an error message
                response = {
                    'message': 'Invalid email or password, Please try again'
                }
                return make_response(jsonify(response)), 401

        except Exception as e:
            server_error(e)
