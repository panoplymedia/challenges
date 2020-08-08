import json,re,sys

# global variables to map each log field in an array to a label
DATE = 0
TIMESTAMP = 1
IP_ADDRESS = 2
USER_AGENT = 3
REQUEST = 4
STATUS = 5
BYTE_RANGE = 6

lines = []
file_path = None
output_file_path = None

#
# method that detects and pulls out a valid log line from being concatenated with the header field names
#
def strip_header_fields(line):
    search_res = re.search(r'\d{4}-\d{2}-\d{2}', line)
    if search_res is not None:
        split_pos = search_res.start()
        return line[split_pos:]
    return None

#
# method that formats a log into a json object
#
def format(split_line):
    byte_split = split_line[BYTE_RANGE].split('-')

    return {
        'date': split_line[DATE],
        'timestamp': split_line[TIMESTAMP],
        'ip_address': split_line[IP_ADDRESS],
        'user_agent': split_line[USER_AGENT],
        'request': split_line[REQUEST],
        'status': int(split_line[STATUS]),
        'min_byte': int(byte_split[0]),
        'max_byte': int(byte_split[1])
    }

#
# method to split a log line
#
def split(line):
    return re.split(r'\t+', line)

#
# method that tests if a request was successful (HTTP OK response)
#
def is_OK(line):
    status = line[STATUS]
    if status == '200' or status == '206':
        return True
    return False

#
# method that orchestrates the running of the parser  
#
def main():
    # reading log file
    with open(file_path, 'r') as fh:
        log_line = strip_header_fields(fh.readline())

        # skip first line if only header info in first line
        if log_line is None:
            log_line = fh.readline()

        # loop through the rest of the file's lines
        while log_line is not None and log_line != '':
            log_line = log_line.replace('\n', '')
            split_line = split(log_line)

            try:
                if is_OK(split_line):
                    lines.append(format(split_line))
            except Exception as e:
                print('\nAn error occured parsing the log line [{}], error = [{}]\n'.format(log_line, str(e)))

            log_line = fh.readline()

    # output writing
    if output_file_path is None:
        for l in lines:
            print('{}\n'.format(l))
    else:
        with open(output_file_path, 'w+') as fh:
            for l in lines:
                fh.write('{}\n'.format(json.dumps({ "index" : { "_index" : "logs" } })))
                fh.write('{}\n'.format(json.dumps(l)))

#
# entry point
#
if __name__ == '__main__':
    if len(sys.argv) == 2:
        file_path = sys.argv[1]
    elif len(sys.argv) == 3:
        file_path = sys.argv[1]
        output_file_path = sys.argv[2]
    else:
        raise RuntimeError('No filepath provided.')

    main()
