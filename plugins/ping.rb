require 'cinch'

class Ping
  include Cinch::Plugin

  match /ping/
  def execute(m)
    m.reply "#{m.user.nick}: pong!"
  end
end
