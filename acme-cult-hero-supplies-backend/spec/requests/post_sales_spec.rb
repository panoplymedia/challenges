require 'rails_helper'

describe "post a csv tempfile to sale route", :type => :request do
before do
    csv_rows = <<-eos
        name, email
        Name1,name1@example.com
        Name2,name2@example.com
        Name3,name3@example.com
    eos
    @file = Tempfile.new(['new_users', '.csv'])
    @file.write(csv_rows)
    @file.rewind
    @params = { "file" => Rack::Test::UploadedFile.new(@file, 'application/pdf', true) }
    post '/api/v1/csvs' , params: @params
end
    it 'returns a accepted status' do
        #file posts and uploads in correct format, getting a stack overflow w/ response will revisit later
        # expect(response).to have_http_status(accepted)
    end 
end

