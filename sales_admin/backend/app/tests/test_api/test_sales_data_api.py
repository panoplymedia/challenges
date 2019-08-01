import json
from app.sales_data.models.sales_data import SalesData
from ..conftest import populate_users
import io
import os
import csv


def test_sales_data_upload(auth, client, app):
    with app.app_context():
        populate_users(app)

        # get auth_token
        res = auth.login()
        token = json.loads(res.data.decode())

        headers = {
            "Authorization": "Bearer {}".format(token["access_token"])
        }

        file_path = os.path.dirname(os.path.realpath(__file__)) + "/salesdata.csv"
        with open(file_path, "rb") as f:
            data = dict(file=(io.BytesIO(f.read()), 'salesdata.csv'))
            response = client.post(
                '/sales_data/upload', data=data,
                headers=headers,
                follow_redirects=True,
                content_type='multipart/form-data'
            )
            assert response.data == b"Sales data uploaded"

        # Query sales data is added
        with app.app_context():
            assert (
                SalesData.query.all()
                is not None
            )


def test_sales_get_total_revenue(auth, client, app):
    with app.app_context():
        populate_users(app)

        # get auth_token
        res = auth.login()
        token = json.loads(res.data.decode())

        headers = {
            "Authorization": "Bearer {}".format(token["access_token"])
        }

        file_path = os.path.dirname(os.path.realpath(__file__)) + "/salesdata.csv"
        with app.app_context():
            with open(file_path) as csv_file:
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
                            "added_by": 1
                        }
                        sales_data = SalesData(**sd)
                        sales_data.save()
                    line_count += 1

            response = client.get('/sales_data/total_revenue', headers=headers)
            assert response.json == {'total_revenue': '526.45'}
