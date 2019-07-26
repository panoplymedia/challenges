from flask.views import MethodView
from flask import make_response, request, jsonify, g
from app.sales_data.models.sales_data import SalesData
from app.sales_data.utils.helper import server_error, create_tmp_file, remove_tmp_directory
from werkzeug.exceptions import HTTPException
import csv


class SalesDataView(MethodView):
    """This class handles all event endpoints."""

    def post(self):
        """Handle POST request for this view. Url ---> /sales_data/upload"""

        file_contents = request.files["file"].stream.read()
        temp_d, report_file = create_tmp_file(file_contents)

        try:
            with open(report_file) as csv_file:
                csv_reader = csv.reader(csv_file, delimiter=',')
                line_count = 0
                column_mapping_dict = dict()
                for row in csv_reader:
                    if line_count == 0:
                        # Saving column position so the data can be uploaded in any order
                        column_count = 0
                        for column_name in row:
                            column_mapping_dict[column_name] = column_count
                            column_count += 1

                    else:
                        # Add sales data
                        sd = {
                            "customer_name": row[column_mapping_dict["Customer Name"]],
                            "item_description": row[column_mapping_dict["Item Description"]],
                            "item_price": row[column_mapping_dict["Item Price"]],
                            "quantity": row[column_mapping_dict["Quantity"]],
                            "merchant_name": row[column_mapping_dict["Merchant Name"]],
                            "merchant_address": row[column_mapping_dict["Merchant Address"]],
                            "added_by": g.user_id
                        }
                        sales_data = SalesData(**sd)
                        sales_data.save()
                    line_count += 1

            return make_response("Sales data uploaded"), 201
        except Exception as e:
            server_error(e)
        finally:
            remove_tmp_directory(temp_d)

    def get(self):
        """Handle GET request for this view. Url ---> /sales_data/total_revenue"""

        try:
            response = SalesData.get_total_revenue()

            return make_response(jsonify(response)), 200
        except HTTPException:
            raise
        except Exception as e:
            server_error(e)
