class Api::V1::SalesController < ApplicationController
    def index
        render json: Sale.all
    end
end
