import os
import logging
from logging.handlers import RotatingFileHandler
from flask import Flask
from flask_cors import CORS
from flask_sqlalchemy import SQLAlchemy
from app.instance.config import app_config


db = SQLAlchemy()


def create_app(config_name="development"):
    """Create and configure an instance of the Flask application."""
    app = Flask(__name__, instance_relative_config=True)
    cors = CORS(app, resources={r"/*": {"origins": "*", "methods": ["GET", "HEAD", "POST", "OPTIONS", "PUT", "PATCH", "DELETE"], "expose_headers": "Authorization"}})
    app.config.from_mapping(
        # a default secret that should be overridden by instance config
        SECRET_KEY="dev",
        # store the database in the instance folder
        DATABASE=os.path.join(app.instance_path, "sqlite.db"),
    )

    app.config.from_object(app_config[config_name])
    app.config['SQLALCHEMY_TRACK_MODIFICATIONS'] = False

    # Set logging
    if app.config.get("LOG_LEVEL") is not None and app.config.get("LOG_LOCATION") is not None:
        formatter = logging.Formatter(
            "[%(asctime)s] {%(pathname)s:%(lineno)d} %(levelname)s - %(message)s")
        handler = RotatingFileHandler(app.config['LOG_LOCATION'], maxBytes=10000, backupCount=1)
        handler.setLevel(app.config["LOG_LEVEL"])
        handler.setFormatter(formatter)
        app.logger.addHandler(handler)

    # register the database commands with SQLAlchemy
    db.init_app(app)

    # set strict_slashes to False, endpoints like /sales_data/upload and /sales_data/upload/ will map to the same thing
    app.url_map.strict_slashes = False

    # apply the blueprints to the app
    from .api_views.auth import auth_bp
    from .api_views.sales_data import sales_data_bp

    app.register_blueprint(auth_bp)
    app.register_blueprint(sales_data_bp)

    return app
