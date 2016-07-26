#!/usr/bin/env ruby

require 'rubygems'
require 'twitter'

def tweet(haiku)
	client = Twitter::REST::Client.new do |config|
  	config.consumer_key        = "2mNTyLU4k8j2VkIBWzd4dnWIl"
  	config.consumer_secret     = "NXNFxyLOvmP04bCwy3XSyZz6uFzM1kpOmcl2vR3viCEjenDj0x"
  	config.access_token        = "755434618774355968-akRoeduqPcv1uYZ0uHXJi837KuJ5Pl3"
  	config.access_token_secret = "WnTZdqHTJkU7FRIlDvJSz1IYEgWsEun3eerhwL9pPVd1Q"
	end
	client.update(haiku)
end

#tweet("this is the next test\nmany more haikus to come\npoetry is hard")

puts ARGF.read

	
