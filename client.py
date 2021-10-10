import requests

"""
r = requests.get('http://127.0.0.1:8080/authenticationURL')
print(r.json())
"""


payload = {'email': 'stefanouorestis@gmail.com', 'code': '4/1AX4XfWgLkjJN3svLxQI-xe-lmBaCqsGVol5HTGZ_OvAVyK2xhZ7_NDtll-o'}
headers = {'Content-type': 'application/x-www-form-urlencoded'}
r = requests.post("http://127.0.0.1:8080/token", data=payload,headers=headers)
print(r.json())


"""
r = requests.get('http://127.0.0.1:8080/files/stefanouorestis@gmail.com').json()
files = r["Files"]
for file in files:
	print(file)
"""