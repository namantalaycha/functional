import sys
from slacker import Slacker
slack = Slacker('xoxb-1194530660337-1206300796208-N1jzcyd4zK2fPGQchs55VGw7')
f = open("temp/text.txt", "r")
message= f.read()

slack.chat.post_message('#demo' , message );
