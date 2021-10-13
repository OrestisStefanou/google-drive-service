import requests
import json

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
	if r.status_code == 200:
		f = open("token", "w")
		f.write(r['AccessToken'])
		f.close()

def list_files(email):
	url = f"{baseURL}files/{email}"
	f = open("token", "r")
	access_token = f.read()
	#print("Access token is:",access_token)
	headers = {'Authorization': access_token}
	r = requests.get(url,headers=headers)
	if r.status_code == 200:
		files = r.json()['Files']
		for file in files:
			print(json.dumps(file,indent=2))
	else:
		print(r)

def download_file(email,file_id):
	url = f"{baseURL}files/download/{email}/{file_id}"
	f = open("token", "r")
	access_token = f.read()
	#print("Access token is:",access_token)
	headers = {'Authorization': access_token}
	r = requests.get(url,headers=headers)	
	if r.status_code == 200:
		#print(r.content)
		f = open("test.docx","wb")
		f.write(r.content)
		f.close()
	else:
		print(r.json())

#list_files('stefanouorestis@gmail.com')
download_file('stefanouorestis@gmail.com','1eqTY8ce0tCSjzfENhMBg3-4rR5HnOEPr')