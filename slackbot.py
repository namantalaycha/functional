import sys
from slacker import Slacker
slack = Slacker(sys.argv[1])
f = open("temp/text.txt", "r")
message= f.read()
slack.chat.post_message('#demo' , message );
