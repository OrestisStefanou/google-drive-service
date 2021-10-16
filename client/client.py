import requests
import json

baseURL = 'http://127.0.0.1:8080/v1'


def ping():
	url = f"{baseURL}/ping"
	try:
		f = open("token.json", "r")
		access_token = f.read()
		f.close()
	except:
		print("Token not found")
		return
	#print("Access token is:",access_token)
	headers = {'Authorization': access_token}
	r = requests.get(url,headers=headers)
	print(r.json())


def get_auth_url():
	url = f"{baseURL}/authenticationURL"
	r = requests.get(url)
	print(r.json())


def create_token(auth_code):
	url = f"{baseURL}/token"
	payload = {'code': auth_code}
	headers = {'Content-type': 'application/x-www-form-urlencoded'}
	r = requests.post(url, data=payload,headers=headers)
	response = r.json()
	print(response)
	if r.status_code == 200:
		json_string = json.dumps(response['AccessToken'])
		f = open("token.json", "w")
		f.write(json_string)
		f.close()


def list_files():
	url = f"{baseURL}/files"
	try:
		f = open("token.json", "r")
		access_token = f.read()
		f.close()
	except:
		print("Token not found")
		return
	#print("Access token is:",access_token)
	headers = {'Authorization': access_token}
	r = requests.get(url,headers=headers)
	if r.status_code == 200:
		files = r.json()['Files']
		for file in files:
			print(json.dumps(file,indent=2))
	else:
		print(r)


def download_file(file_id,filepath):
	url = f"{baseURL}/files/download/{file_id}"
	try:
		f = open("token.json", "r")
		access_token = f.read()
		f.close()
	except:
		print("Token not found")
		return
	headers = {'Authorization': access_token}
	r = requests.get(url,headers=headers)	
	if r.status_code == 200:
		#print(r.content)
		f = open(filepath,"wb")
		f.write(r.content)
		f.close()
	else:
		print(r.json())

list_files()