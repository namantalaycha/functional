import sys
from slacker import Slacker
slack = Slacker('xoxb-1194530660337-1206300796208-gpwoVts9WPOYXkq5zqoNtGtZ')
message="hello everyone"
slack.chat.post_message('#demo' , message );
