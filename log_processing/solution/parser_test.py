import unittest
import parser

class TestMethods(unittest.TestCase):

    HEADER_LINE_RAW = '# date	timestamp	ip_address	user_agent	request	status	byte_range2020-07-31	16:46:07	63.110.194.22	Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:78.0) Gecko/20100101 Firefox/78.0	/6d8a9754-e8c7-4193-8491-58b2122c1c10	206	0-289'
    HEADER_LINE_PARSED = '2020-07-31	16:46:07	63.110.194.22	Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:78.0) Gecko/20100101 Firefox/78.0	/6d8a9754-e8c7-4193-8491-58b2122c1c10	206	0-289'
    SAMPLE_LINE = '2020-07-31	12:46:27	228.93.19.159	Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:78.0) Gecko/20100101 Firefox/78.0	/92e2666d-fc50-457a-95ba-58ac268eac48	206	0-400'
    SAMPLE_LINE_FORMATTED = {
        'date': '2020-07-31',
        'timestamp': '12:46:27',
        'ip_address': '228.93.19.159',
        'user_agent': 'Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:78.0) Gecko/20100101 Firefox/78.0',
        'request': '/92e2666d-fc50-457a-95ba-58ac268eac48',
        'status': 206,
        'min_byte': 0,
        'max_byte': 400
    }

    def test_strip_header_fields_1(self):
        out = parser.strip_header_fields(self.HEADER_LINE_RAW)
        self.assertEqual(self.HEADER_LINE_PARSED, out)

    def test_transform_1(self):
        split_line = parser.split(self.SAMPLE_LINE)
        formatted_line = parser.format(split_line)
        self.assertEqual(self.SAMPLE_LINE_FORMATTED, formatted_line)

    def test_is_OK_1(self):
        split_line = parser.split(self.SAMPLE_LINE)
        out = parser.is_OK(split_line)
        self.assertTrue(out)

if __name__ == '__main__':
    unittest.main()