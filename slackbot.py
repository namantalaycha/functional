import sys
from slacker import Slacker
aoi_token = input("Enter your API TOKEN")
slack = Slacker('aoi_token')
f = open("temp/text.txt", "r")
message= f.read()

slack.chat.post_message('#demo' , message );
