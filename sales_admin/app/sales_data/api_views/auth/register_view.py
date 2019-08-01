from flask.views import MethodView
from flask import make_response, request, jsonify
from app.sales_data.models.user import User
from app.sales_data.utils.helper import verify_fields, server_error
import json


class RegistrationView(MethodView):
    """This class registers a new user."""

    def post(self):
        """Handle POST request for this view. Url ---> /auth/register"""

        post_data = json.loads(request.get_json().get("body"))
        verify_fields(["username", "password"], post_data)

        try:
            # Query to see if the user already exists
            user = User.query.filter_by(username=post_data['username']).first()

            if not user:
                # There is no user so we'll try to register them
                # Register the user
                user = User(username=post_data["username"], password=post_data["password"])
                user.save()

                response = {
                    'message': 'You registered successfully. Please log in.'
                }
                # return a response notifying the user that they registered successfully
                return make_response(jsonify(response)), 201

            else:
                # There is an existing user. We don't want to register users twice
                # Return a message to the user telling them that they they already exist
                response = {
                    'message': 'User already exists. Please login.'
                }

                return make_response(jsonify(response)), 202
        except Exception as e:
            server_error(e)
