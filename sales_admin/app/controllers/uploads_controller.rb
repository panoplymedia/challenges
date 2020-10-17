class UploadsController < ApplicationController
  before_action :set_upload, only: [:show, :destroy]

  # GET /uploads
  # GET /uploads.json
  def index
    @uploads = Upload.last(7)
  end


  def from_csv

    require 'csv'

    csv = params[:csv]

    table = CSV.parse(File.read(csv.tempfile), headers: true)

    puts table.headers

    ActiveRecord::Base.transaction do

      db_upload = Upload.create! filename: csv.original_filename

      table.each do |sale|

        customer = Customer.find_or_create_by(name: sale['Customer Name'])

        merchant = Merchant.find_or_create_by(name: sale['Merchant Name'], address: sale['Merchant Address'])

        item = Item.find_or_create_by(description: sale['Item Description'], price: sale['Item Price'], merchant: merchant)

        sale = Sale.find_or_create_by(quantity: sale['Quantity'], customer: customer, item: item, upload: db_upload)
      end

      puts "Sales: #{db_upload.total_revenue}"
    end

    redirect_to root_path
  end


  def show
  end

  def edit
  end


  def destroy
    @upload.destroy
    respond_to do |format|
      format.html { redirect_to uploads_url, notice: 'Upload was successfully destroyed.' }
      format.json { head :no_content }
    end
  end

  private
    # Use callbacks to share common setup or constraints between actions.
    def set_upload
      @upload = Upload.find(params[:id])
    end

    # Only allow a list of trusted parameters through.
    def upload_params
      params.require(:upload).permit(:filename)
    end
end
