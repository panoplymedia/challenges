import React from "react";
import PropTypes from "prop-types";
import { Upload, message, Button } from "antd";

/**
 * Called when a file is selected for upload.
 * Return value determines if upload proceeds
 * @param {object} file
 * @returns {boolean} true if valid
 */
export const validateOrderUpload = file => {
  const isCSV = file.type === "text/csv";
  if (!isCSV) {
    message.error("Only CSV files are supported");
  }
  return isCSV;
};

/**
 * Event handler generator for the uploader.
 * Builds a handler used to trigger based off of upload state events
 * @param {function} onSuccess a callback called when upload succeeds
 * @returns {undefined} no return value
 */
export const genOnOrderUploadEvent = onSuccess => info => {
  if (info && info.file) {
    switch (info.file.status) {
      case "done":
        message.success(`${info.file.name} uploaded succesfully`);
        onSuccess();
        break;
      case "error":
        message.error(`Failed to upload ${info.file.name}`);
        break;
      default:
      //no-op
    }
  }
};

/**
 * An ant design upload component purposed for Order CSV data
 * @param {object} props
 */
export const OrdersUpload = ({ url = "/orders", onSuccess = () => {} }) => (
  <Upload
    action={url}
    showUploadList={false}
    beforeUpload={validateOrderUpload}
    onChange={genOnOrderUploadEvent(onSuccess)}
  >
    <Button type="primary">Import Sales CSV</Button>
  </Upload>
);

OrdersUpload.propTypes = {
  url: PropTypes.string
};
