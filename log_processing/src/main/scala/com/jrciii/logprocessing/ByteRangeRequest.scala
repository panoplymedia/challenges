package com.jrciii.logprocessing

/**
  * Represents a byte range request. A status of 200 or 206 is a successful delivery.
  * @param date In yyyy-MM-dd format
  * @param timestamp In HH:mm:ss format
  * @param ipAddress In IPv4 format such as 192.168.1.1
  * @param userAgent User agent of the requester
  * @param request Request path
  * @param status HTTP status code.
  * @param byteRange In the form of '0-100', which is start to finish.
  */
case class ByteRangeRequest(date: String,
                            timestamp: String,
                            ipAddress: String,
                            userAgent: String,
                            request: String,
                            status: Int,
                            byteRange: String)
