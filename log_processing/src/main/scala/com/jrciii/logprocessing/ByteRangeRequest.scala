package com.jrciii.logprocessing

case class ByteRangeRequest(date: String,
                            timestamp: String,
                            ipAddress: String,
                            userAgent: String,
                            request: String,
                            status: Int,
                            byteRange: String)
