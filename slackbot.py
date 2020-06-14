import sys
from slacker import Slacker
slack = Slacker('xoxb-1194530660337-1206300796208-XzURRjkJB5a1caEXUqfT13ui')
message="hello everyone"

slack.chat.post_message('#demo' , message );
