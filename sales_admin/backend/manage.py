# class for handling migrations
from flask_script import Manager
from flask_migrate import Migrate, MigrateCommand
from app.sales_data import db, create_app


# initialize the app with all its configurations
app = create_app("development")  # config_name=os.getenv('APP_SETTINGS'))
migrate = Migrate(app, db)
# create an instance of class that will handle our commands
manager = Manager(app)

# Define the migration command to always be preceded by the word "db"
# Example usage: python manage.py db init
manager.add_command('db', MigrateCommand)


if __name__ == '__main__':
    manager.run()
