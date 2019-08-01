import os
import tempfile
import pytest
from app.sales_data import create_app, db
from app.sales_data.models.user import User


@pytest.fixture
def app():
    """Create and configure a new app instance for each test."""
    # create a temporary file to isolate the database for each test
    db_fd, db_path = tempfile.mkstemp()
    # create the app with common test config
    app = create_app("testing")
    app.config["SQLALCHEMY_DATABASE_URI"] = "sqlite:///{}".format(db_path)

    # create the database and load test data
    with app.app_context():
        # create all tables
        db.session.close()
        db.drop_all()
        db.create_all()

    yield app

    # close and remove the temporary database
    os.close(db_fd)
    os.unlink(db_path)


@pytest.fixture
def client(app):
    """A test client for the app."""
    return app.test_client()


class AuthActions(object):
    def __init__(self, client):
        self._client = client

    def login(self):
        return self._client.post(
            "/auth/login", json={"body": '{"username": "test", "password": "test"}'}
        )


@pytest.fixture
def auth(client):
    return AuthActions(client)


def populate_users(app):
    with app.app_context():
        # create all tables
        data = [
            {"username": "test", "password": "test"},
        ]
        for user in data:
            user = User(username=user["username"], password=user["password"])
            user.save()

