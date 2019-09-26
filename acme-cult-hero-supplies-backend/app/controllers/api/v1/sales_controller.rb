class Api::V1::SalesController < ApplicationController
    def index
        response = Sale.all

        render json: response, status: :ok
    end
end
