require 'cinch'
require 'require_all'

require_all 'plugins'

BOTNAME = 'diegobot'
REALNAME = 'Diego the african bot'

bot = Cinch::Bot.new do
  configure do |c|
    c.server = "irc.freenode.org"
    c.channels = ["#mcgillecetest"]
    c.plugins.plugins = [Ping, Messenger, UrbanDictionary, WolframAlpha]
    c.nick = BOTNAME
    c.user = BOTNAME
    c.realname = REALNAME
  end
end

bot.start
