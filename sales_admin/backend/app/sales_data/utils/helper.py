from flask import current_app, make_response, jsonify, abort
import tempfile
import shutil


# Verify all required fields are passed and not empty strings in the body otherwise throw a 400 error
def verify_fields(required, passed):
    for field in required:
        if field in passed:
            if passed[field] is not None and passed[field] != "":
                continue
        # If either condition is met, return 400 error
        abort(make_response(jsonify({'message': '{} must be non-empty.'.format(field)}), 400))


def server_error(error):
    # log error
    current_app.logger.error(error)

    # Create a response containing an string error message
    response = {
        'message': "Internal Server Error"
    }
    # Return a server error using the HTTP Error Code 500 (Internal Server Error)
    abort(make_response(jsonify(response), 500))


def create_tmp_file(contents):
        temp_d = tempfile.mkdtemp()
        temp_f = tempfile.mkstemp(dir=temp_d)
        report_file = temp_f[1]

        with open(report_file, 'wb') as f:
            f.write(contents)

        return temp_d, report_file


def remove_tmp_directory(directory):
    shutil.rmtree(directory)
