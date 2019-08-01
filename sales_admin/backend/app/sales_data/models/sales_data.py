from app.sales_data import db
from .user import User
from sqlalchemy import func


class SalesData(db.Model):
    """This class defines the sales_data table """

    __tablename__ = 'sales_data'

    # Define the columns of the users table, starting with the primary key
    id = db.Column(db.Integer, primary_key=True)
    customer_name = db.Column(db.String(1000), nullable=True)
    item_description = db.Column(db.String(1000), nullable=True)
    item_price = db.Column(db.DECIMAL(10, 2), nullable=True)
    quantity = db.Column(db.Integer, nullable=True)
    merchant_name = db.Column(db.String(1000), nullable=True)
    merchant_address = db.Column(db.String(1000), nullable=True)
    added_by = db.Column(db.Integer, db.ForeignKey(User.id), nullable=False)
    created_time = db.Column(db.DateTime, server_default=db.func.now())
    modified_time = db.Column(db.DateTime(timezone=True), server_default=db.func.now(), onupdate=db.func.now())

    def __init__(self, **kwargs):
        """Initialize the sales data with details."""
        self.customer_name = kwargs.get("customer_name")
        self.item_description = kwargs.get("item_description")
        self.item_price = kwargs.get("item_price")
        self.quantity = kwargs.get("quantity")
        self.merchant_name = kwargs.get("merchant_name")
        self.merchant_address = kwargs.get("merchant_address")
        self.added_by = kwargs.get("added_by")

    def save(self):
        """Save sales data to the database."""
        db.session.add(self)
        db.session.commit()

    #TODO
    @staticmethod
    def get_total_revenue():
        """Calculate the total revenue"""
        total_revenue = db.session.query(func.sum(SalesData.item_price * SalesData.quantity))
        return {
            "total_revenue": str(round(total_revenue.scalar(), 2))
        }

    @staticmethod
    def calc_distance(num1, num2):
        return num1*num2
