Models 
customer - has_many :items
item - belongs_to :customer belongs_to :merchant
merchant - has_many :items
user

Features
- interface to upload CSV data to DB
- calculate sales data for user, total sales revenue
- auth

Gems added - 
- bcrypt 3.1.7
- jwt
- figaro
- rack-cors

to load app: 
to run the app locally:
- clone this repository
- `cd sales_admin_app`
- install dependencies with `bundle install`
- setup db 
    - `rails db:setup`
    - `rails db:migrate`
    - `rails db:seed`
- `rake start`

to test: 
- `rails t`
