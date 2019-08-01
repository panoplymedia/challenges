import json
from app.sales_data.models.user import User


def add_user_data():
    data = [
            {"username": "test", "password": "test"}
        ]
    for user in data:
        user = User(username=user["username"], password=user["password"])
        user.save()


def test_register(client, app):
    response = client.post("/auth/register", json={"body": '{"username": "a", "password": "a"}'})
    # get the results returned in json format
    result = json.loads(response.data.decode())
    # assert that the request contains a success message and a 201 status code
    assert result['message'] == "You registered successfully. Please log in."
    assert response.status_code == 201

    # test that the user was inserted into the database
    with app.app_context():
        assert (
            User.query.filter_by(username="a").first()
            is not None
        )


def test_login(app, auth):
    # Add user data to the database
    with app.app_context():
        add_user_data()

    response = auth.login()

    # get the results returned in json format
    result = json.loads(response.data.decode())
    # assert that the request contains a success message and a 201 status code
    assert result['message'] == "Successfully logged in."
    assert len(result['access_token']) == 139
    assert response.status_code == 200
