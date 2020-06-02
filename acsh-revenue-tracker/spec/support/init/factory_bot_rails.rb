require 'factory_bot_rails'

RSpec.configure do |config|
  # Allows the use of "create :user" as a shorthand for "FactoryBot.create :user".
  config.include FactoryBot::Syntax::Methods
end

