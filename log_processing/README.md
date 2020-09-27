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
5. Allow the user to query the merged byte ranges.

## Installation And Running
To install, sbt is required. cd into log_processing and run `sbt assembly`. Then run the jar, supplying the directory to your CSV with byte range requests
```
cd log_processing
sbt assembly
java -jar target\scala-2.12\log_processing-assembly-0.0.1.jar "C:\Users\You\code\challenges\log_processing\src\test\resources\sample"
```
Some Spark output will be displayed. Wait for the input prompt to appear. It looks like "> ".
You can query the dataset by entering Spark SQL at the prompt and hitting enter.
To exit, type 'quit'

## Querying the byte ranges
After starting the application, wait for the "> " prompt to appear. Then enter your query. Type 'quit' to exit.

To check if a file was completely delivered, check if that combination of ipAddress, userAgent, request has only one
byte range starting at 0 and ending with the file size.
```
> select * from delivered where ipAddress='183.3.129.45' AND userAgent='Mozilla/5.0 (iPhone; CPU iPhone OS 13_6 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) CriOS/83.0.4147.71 Mobile/15E148 Safari/604.1' AND request='/32668757-95c0-4ded-9b81-71607a644e92' AND byteRanges=array(named_struct('start', 0, 'end', 1741))
+------------+---------------------------------------------------------------------------------------------------------------------------------------------+-------------------------------------+-----------+
|ipAddress   |userAgent                                                                                                                                    |request                              |byteRanges |
+------------+---------------------------------------------------------------------------------------------------------------------------------------------+-------------------------------------+-----------+
|183.3.129.45|Mozilla/5.0 (iPhone; CPU iPhone OS 13_6 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) CriOS/83.0.4147.71 Mobile/15E148 Safari/604.1|/32668757-95c0-4ded-9b81-71607a644e92|[[0, 1741]]|
+------------+---------------------------------------------------------------------------------------------------------------------------------------------+-------------------------------------+-----------+
```

## How Would This App Look And Scale In Production
In production, this app would be changed to stream from a distributed source such as S3 or Kafka, and write to one or more 
sinks such as ElasticSearch and S3. The Spark app itself could run on EMR.

To improve scaling, determining how far back in time we need to keep requests to merge to determine successful delivery would help.
Old requests would get dropped from memory. This demo app does not consider the timestamp when aggregating.
