# Log Processing Solution

## TLDR - How do I run this?
To ensure compatibility with whatever OS reviewers are using I'm running everything in Docker containers. Docker cli commands can get pretty verbose, so I've wrapped the commands in shell scripts.

### Requirements
- A working install of Docker ([Instructions Here](https://docs.docker.com/get-docker/))

### Steps
Execute the following commands from `./log_processing` to walk through my solution:
- `./0_run_tests.sh`
    - Runs all the tests via pytest inside a Spark Docker container
- `./1_process_logs.sh`
    - Processes the sample log file into the datastore (parquet)
- `./2_check_delivery.sh`
    - This is an example script to demonstrate how to check if an ip/ua/asset has been delivered

### Additional info
All data is stored in `./data`, should you wish to inspect the output.

If you wish to re-run the processing you can `rm -rf ./data/parquet/delivery.parquet` to get a clean slate.

## My thought process around this solution

### Tool Choice
#### ETL
I chose Spark to process the logs because it is an all around good ETL tool to work with. It scales extremely well, and it can run natively on most platforms that you'd likely be deploying to (EMR, GKE, DataProc, Hadoop, Spark Standalone, Spark Local).

#### Data Store
I chose Parquet because it works really well with Spark out of the box, and it is a really good columnar storage format. I also considered using Vertica here, but I decided against it for this challenge as it would add a lot of complexity to the setup for the reviewers.

#### Front End
In my example script I used Spark again to demonstrate how the data could be accessed. The disadvantage here is that it takes Spark a few seconds to get started. This is not ideal for a large number of small queries. If this was a production deployment I'd likely use something like Vertica, Presto, or a Spark Thrift Server to access the data, which would allow instant access.

### Data Model
To take full advantage of Parquet's columnar datastore I modeled the data such that each byte delivered is exploded out into it's own row. This makes the ETL a little more compute intensive, but with the proper encoding it speeds up the query times. I like this tradeoff because it is easier to scale the Spark ETL than it is to scale your DataStore.

This model takes advantage of the [Delta Encoding](https://github.com/apache/parquet-format/blob/master/Encodings.md#delta-encoding-delta_binary_packed--5) built in to Parquet makes it really efficient to store (and query) sequential data in a given column.

I have experience tuning similar data models using Vertica's [Common Delta](https://www.vertica.com/docs/9.2.x/HTML/Content/Authoring/SQLReferenceManual/Statements/encoding-type.htm#COMMONDE) encoding as well. This is likely what I'd use in a production deployment (mainly due to my familiarity).
