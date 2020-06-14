import requests
import json
url = 'https://hooks.slack.com/services/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'

msg = { "test" : "sending data successfull"}
print("successfull")
requests.post( url , data=json.dumps(msg) )
print("successfull")
