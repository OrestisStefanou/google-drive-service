import requests

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

headers = {'Cookie': 'ya29.a0ARrdaM8ORNI3bEAGMxekzEEpsKOIZoxCVLcAZr3BUh7ea1cqDOjP19YquhWez8Z312MFfLYWKDP9euyy_SWTG2gbwdwqMsUeWWHIBq-xPA5lnYvhckZHjOUZVpsQ2yA9fniVBRWJvOtVvGks2IntvAQhpdJR'}
r = requests.get('http://127.0.0.1:8080/v1/files/stefanouorestis@gmail.com',headers=headers).json()
print(r)