from app.sales_data.models.user import User
from flask import request, abort, make_response, jsonify, g


def get_user_id():
    # Headers aren't passed in OPTIONS call
    if request.method in ["OPTIONS"]:
        return
    auth_header = request.headers.get('Authorization')
    access_token = auth_header.split(" ")[1]

    unauthorized_response = {
        'message': "Unauthorized: The server could not verify that you are authorized to access the URL "
                   "requested. You either supplied the wrong credentials (e.g. a bad password), or your browser "
                   "doesn't understand how to supply the credentials required."
    }

    if access_token:
        # Attempt to decode the token and get the User ID
        user_id = User.decode_token(access_token)

        if not isinstance(user_id, str):
            g.user_id = user_id
        else:
            abort(make_response(jsonify(unauthorized_response), 401))

    else:
        abort(make_response(jsonify(unauthorized_response), 401))
