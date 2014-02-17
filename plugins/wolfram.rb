require 'cinch'
require 'wolfram'

class WolframAlpha
  include Cinch::Plugin

  Wolfram.appid = '6KVGE3-JHPTXY6WPK'
  set :prefix, /^/

  match /diegobot: (.*)/
  def execute(m, query)
    result = get_wolfram_result(query)
    result = 'No result found!' if result.nil? || result.strip.empty?

    m.reply "#{m.user.nick}: #{result}"
  end

  def get_wolfram_result(query)
    result = Wolfram.fetch(query)
    hash = Wolfram::HashPresenter.new(result).to_hash
    # hash has one key, :pods
    # the key holds a hash of key string, and as value a list with one item in
    # it
    hash_clean = remove_empty_line_in_hash(hash[:pods])
    generate_string_from_hash(hash_clean)
  end

  def remove_empty_line_in_hash(hash)
    hash.select { |_, v| !v[0].empty? }
  end

  def generate_string_from_hash(hash)
    query_result = ''
    count = 0

    hash_clean.each do |_, v|
      count += 1
      query_result << v[0] << "\n"
      break if count == 3 # only show the first 3 data
    end
    query_result.strip
  end
end
