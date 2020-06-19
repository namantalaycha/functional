import sys
from slacker import Slacker
slack = Slacker('xoxb-1194530660337-1206300796208-To4noH0m82VZaauXLoLTixuy')
f = open("temp/text.txt", "r")
message= f.read()

slack.chat.post_message('#demo' , message );
