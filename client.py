import requests

baseURL = 'http://127.0.0.1:8080/v1/'


def get_auth_url():
	url = f"{baseURL}/authenticationURL"
	r = requests.get(url)
	print(r.json())


def create_token(email,auth_code):
	url = f"{baseURL}/token"
	payload = {'email': email, 'code': auth_code}
	headers = {'Content-type': 'application/x-www-form-urlencoded'}
	r = requests.post(url, data=payload,headers=headers)
	response = r.json()
	print(response)
	if r.status == 200:
		f = open("token", "w")
		f.write(r['AccessToken'])
		f.close()

def list_files(email):
	url = f"{baseURL}/files/{email}"
	f = open("token", "r")
	access_token = f.read()
	print("Access token is:",access_token)
	headers = {'Authorization': access_token}
	r = requests.get(url,headers=headers).json()
	print(r)	

"""
headers = {'Content-type': 'application/x-www-form-urlencoded','Cookie': 'ya29.a0ARrdaM8ORNI3bEAGMxekzEEpsKOIZoxCVLcAZr3BUh7ea1cqDOjP19YquhWez8Z312MFfLYWKDP9euyy_SWTG2gbwdwqMsUeWWHIBq-xPA5lnYvhckZHjOUZVpsQ2yA9fniVBRWJvOtVvGks2IntvAQhpdJR'}
r = requests.get('http://127.0.0.1:8080/v1/ping',headers=headers)
print(r.json())
"""

"""
payload = {'email': 'stefanouorestis@gmail.com', 'code': '4/1AX4XfWgLkjJN3svLxQI-xe-lmBaCqsGVol5HTGZ_OvAVyK2xhZ7_NDtll-o'}
headers = {'Content-type': 'application/x-www-form-urlencoded','AccessToken': 'testingToken'}
r = requests.post("http://127.0.0.1:8080/v1/token", data=payload,headers=headers)
print(r.json())
"""

"""
headers = {'Authorization': 'ya29.a0ARrdaM8ORNI3bEAGMxekzEEpsKOIZoxCVLcAZr3BUh7ea1cqDOjP19YquhWez8Z312MFfLYWKDP9euyy_SWTG2gbwdwqMsUeWWHIBq-xPA5lnYvhckZHjOUZVpsQ2yA9fniVBRWJvOtVvGks2IntvAQhpdJR'}
r = requests.get('http://127.0.0.1:8080/v1/files/stefanouorestis@gmail.com',headers=headers).json()
print(r)
"""