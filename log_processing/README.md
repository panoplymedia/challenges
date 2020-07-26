# Log Processing

## Background

You are managing a content delivery network, and a customer is questioning whether all of an asset is being delivered to a particular IP address and user agent.

The assets in question are primarily delivered via range requests, so you'll need to efficiently process the log lines and the byte range in each line to determine which bytes are actually being delivered.

## Functional requirements

Your solution will need to be able to calculate whether every byte within a specific byte range for an asset is delivered to a given IP address and user agent. The assets will be identified by the request path in the log line. Requests for different assets will be interspersed within the log file. Only requests with HTTP status of 200 or 206 should be counted.

For example, given a 1000 byte file, the range requests may be something like 0-200 bytes, 200-400 bytes, etc. In this scenario, if the customer asks whether bytes 0-150 were delivered, the answer is yes.

However, in some cases the range requests may skip a segment of the asset, going from byte 400 to request bytes 500-600. In this scenario, if the customer asks whether bytes 401 to 550 were delivered, the answer is no, because not all bytes in that range were delivered.

A sample log file is provided in the current directory.

## Delivery requirements

Please provide code for processing the log files, along with documentation or discussion of any platforms, technologies, etc. that you may use in a production solution. An ideal solution will provide an output format where a user can easily query specific byte ranges for specific assets.

Think about things like:

* Testing
* How to store the data
* How would your solution differ if it had to scale?

Please submit your solution as a pull request, or package it up and send it to doug.ramsay@megaphone.fm.
