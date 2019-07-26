from flask import Blueprint
from .sales_data_view import SalesDataView
from app.sales_data.utils.authorization import get_user_id


# This instance of a Blueprint that represents the sales_data blueprint
sales_data_bp = Blueprint('sales_data', __name__, url_prefix="/")
sales_data_bp.before_request(get_user_id)

# Define the rule for the registration url --->  /sales_data
sales_data_view = SalesDataView.as_view('sales_data_view')
# Then add the rule to the blueprint
sales_data_bp.add_url_rule(
    '/sales_data/upload',
    view_func=sales_data_view,
    methods=['POST'])

sales_data_bp.add_url_rule(
    '/sales_data/total_revenue',
    view_func=sales_data_view,
    methods=['GET'])
