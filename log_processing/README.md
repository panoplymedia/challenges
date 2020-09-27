# Log Processing

## Background

This application demonstrates the use of Spark Structured Streaming to process streaming byte range requests in order
to determine complete or incomplete delivery of bytes. The demo process consists of:

1. Stream from a folder containing CSV files of byte range requests
2. Reject any requests without a status of 200 or 206
3. Group requests by ip, user agent and request path
4. Merge these successful byte ranges. For example this group of ranges: (0-100) and (50-200) will become: (0-200)
    If the file is 200 bytes, this represents a complete delivery. This group of ranges: (0-100), (50-200), (300-400), (350-500)
    will become (0-200) and (300-500). Since there are multiple ranges, this represents a gap, or an incomplete delivery.
5. Allow the user to query the merged byte ranges stored in the Memory sink.

## Installation And Running
To build from the source, sbt is required. Alternatively you can download a prebuilt jar here https://github.com/jrciii/challenges/releases/download/my-solution/log_processing-assembly-0.0.1.jar

### Building with sbt
cd into log_processing and run `sbt assembly`. Then run the jar, supplying the directory to your CSV with byte range requests
```
cd log_processing
sbt assembly
```

### Running
To run the jar, use the java command, specify the jar and the full path to the directory containing the CSV files with 
the byte range requests.
```
java -jar target\scala-2.12\log_processing-assembly-0.0.1.jar "C:\Users\You\code\challenges\log_processing\src\test\resources\sample"
```
Some Spark output will be displayed. Wait for the input prompt to appear. It looks like "> ".
You can query the dataset by entering Spark SQL queries at the prompt and hitting enter. The table to select from is 'delivered'. 
To exit, type 'quit'

## Querying the byte ranges
After starting the application, wait for the "> " prompt to appear. Then enter your Spark SQL query. Type 'quit' to exit.

To check if a file was completely delivered, check if that combination of ipAddress, userAgent, request has only one
byte range starting at 0 and ending with the file size. If a row is returned, the file was delivered successfully.
```
> select * from delivered where ipAddress='183.3.129.45' AND userAgent='Mozilla/5.0 (iPhone; CPU iPhone OS 13_6 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) CriOS/83.0.4147.71 Mobile/15E148 Safari/604.1' AND request='/32668757-95c0-4ded-9b81-71607a644e92' AND byteRanges=array(named_struct('start', 0, 'end', 1741))
+------------+---------------------------------------------------------------------------------------------------------------------------------------------+-------------------------------------+-----------+
|ipAddress   |userAgent                                                                                                                                    |request                              |byteRanges |
+------------+---------------------------------------------------------------------------------------------------------------------------------------------+-------------------------------------+-----------+
|183.3.129.45|Mozilla/5.0 (iPhone; CPU iPhone OS 13_6 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) CriOS/83.0.4147.71 Mobile/15E148 Safari/604.1|/32668757-95c0-4ded-9b81-71607a644e92|[[0, 1741]]|
+------------+---------------------------------------------------------------------------------------------------------------------------------------------+-------------------------------------+-----------+
```

To get some requests with gaps, try this query:
```
> select * from delivered where size(byteRanges) > 1
+---------------+---------------------------------------------------------------------------------------------------------------------------------------------+-------------------------------------+-------------------------+
|ipAddress      |userAgent                                                                                                                                    |request                              |byteRanges               |
+---------------+---------------------------------------------------------------------------------------------------------------------------------------------+-------------------------------------+-------------------------+
|106.220.65.54  |Mozilla/5.0 (iPhone; CPU iPhone OS 13_6 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) CriOS/83.0.4147.71 Mobile/15E148 Safari/604.1|/c66c0276-3cc2-47f9-a5ae-42954421a4b1|[[0, 1077], [1436, 1795]]|
|228.122.181.108|Mozilla/5.0 (iPhone; CPU iPhone OS 13_6 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) CriOS/83.0.4147.71 Mobile/15E148 Safari/604.1|/566e5232-72f1-4a70-aa89-6e8adfb5a73f|[[0, 504], [1008, 1260]] |
|14.98.169.176  |Mozilla/5.0 (iPhone; CPU iPhone OS 13_6 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) CriOS/83.0.4147.71 Mobile/15E148 Safari/604.1|/92e2666d-fc50-457a-95ba-58ac268eac48|[[0, 400], [800, 1600]]  |
|71.95.75.219   |Mozilla/5.0 (iPhone; CPU iPhone OS 13_6 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) CriOS/83.0.4147.71 Mobile/15E148 Safari/604.1|/32668757-95c0-4ded-9b81-71607a644e92|[[0, 696], [1044, 1741]] |
|228.93.19.159  |Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:78.0) Gecko/20100101 Firefox/78.0                                                               |/968687c0-3efd-4ae5-aeec-9d4ac5e1598e|[[0, 872], [1090, 1092]] |
|249.94.28.246  |Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:78.0) Gecko/20100101 Firefox/78.0                                                               |/c66c0276-3cc2-47f9-a5ae-42954421a4b1|[[0, 718], [1077, 1795]] |
|94.230.87.124  |Mozilla/5.0 (iPhone; CPU iPhone OS 13_6 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) CriOS/83.0.4147.71 Mobile/15E148 Safari/604.1|/cd6d50fd-ada7-45da-9d5f-d59694f52be7|[[0, 556], [834, 1112]]  |
|133.15.230.36  |Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.89 Safari/537.36                           |/6d8a9754-e8c7-4193-8491-58b2122c1c10|[[0, 578], [867, 1156]]  |
|141.239.177.203|Mozilla/5.0 (iPhone; CPU iPhone OS 13_6 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) CriOS/83.0.4147.71 Mobile/15E148 Safari/604.1|/ea65704e-6829-4871-a92a-5f0c3b80addf|[[0, 633], [844, 1057]]  |
|4.254.102.152  |Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:78.0) Gecko/20100101 Firefox/78.0                                                               |/921fa9d7-79c3-4758-bb7c-76110cb21187|[[305, 610], [915, 1525]]|
|55.28.147.17   |Mozilla/5.0 (iPhone; CPU iPhone OS 13_6 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) CriOS/83.0.4147.71 Mobile/15E148 Safari/604.1|/968687c0-3efd-4ae5-aeec-9d4ac5e1598e|[[0, 436], [654, 872]]   |
|109.132.83.87  |Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:78.0) Gecko/20100101 Firefox/78.0                                                               |/6d8a9754-e8c7-4193-8491-58b2122c1c10|[[0, 289], [578, 867]]   |
|109.72.52.215  |Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:78.0) Gecko/20100101 Firefox/78.0                                                               |/c66c0276-3cc2-47f9-a5ae-42954421a4b1|[[0, 718], [1077, 1795]] |
|97.75.45.253   |Mozilla/5.0 (iPhone; CPU iPhone OS 13_6 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) CriOS/83.0.4147.71 Mobile/15E148 Safari/604.1|/32668757-95c0-4ded-9b81-71607a644e92|[[0, 1044], [1392, 1741]]|
|196.204.201.103|Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:78.0) Gecko/20100101 Firefox/78.0                                                               |/968687c0-3efd-4ae5-aeec-9d4ac5e1598e|[[0, 654], [872, 1090]]  |
|26.216.213.231 |Mozilla/5.0 (iPhone; CPU iPhone OS 13_6 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) CriOS/83.0.4147.71 Mobile/15E148 Safari/604.1|/92e2666d-fc50-457a-95ba-58ac268eac48|[[0, 400], [800, 1600]]  |
|182.122.91.54  |Mozilla/5.0 (iPhone; CPU iPhone OS 13_6 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) CriOS/83.0.4147.71 Mobile/15E148 Safari/604.1|/566e5232-72f1-4a70-aa89-6e8adfb5a73f|[[0, 504], [756, 1008]]  |
|71.95.75.219   |Mozilla/5.0 (iPhone; CPU iPhone OS 13_6 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) CriOS/83.0.4147.71 Mobile/15E148 Safari/604.1|/cd6d50fd-ada7-45da-9d5f-d59694f52be7|[[0, 556], [834, 1390]]  |
|153.116.196.14 |Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.89 Safari/537.36                           |/6d8a9754-e8c7-4193-8491-58b2122c1c10|[[0, 867], [1156, 1446]] |
|182.122.91.54  |Mozilla/5.0 (iPhone; CPU iPhone OS 13_6 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) CriOS/83.0.4147.71 Mobile/15E148 Safari/604.1|/ea65704e-6829-4871-a92a-5f0c3b80addf|[[0, 311], [522, 1057]]  |
+---------------+---------------------------------------------------------------------------------------------------------------------------------------------+-------------------------------------+-------------------------+
only showing top 20 rows
```

## How Would This App Look And Scale In Production
In production, this app would be changed to stream from a distributed source such as S3 or Kafka, and write to one or more 
sinks such as ElasticSearch and S3.
There would be a separate application to query the output sink for delivery statuses, such as a microservice that queries ElasticSearch.
The Spark app itself could run on EMR.

To improve scaling, determining how far back in time we need to keep requests to merge to determine successful delivery would help.
Old requests would get dropped from memory. This demo app does not consider the timestamp when aggregating.
