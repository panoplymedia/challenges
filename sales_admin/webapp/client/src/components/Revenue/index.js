import React from "react";
import PropTypes from "prop-types";

export const Revenue = ({ revenue = 0 }) => (
  <div>
    <span>
      <b>Total Revenue: </b>${revenue.toFixed(2)}
    </span>
  </div>
);

Revenue.propTypes = {
  revenue: PropTypes.number
};
