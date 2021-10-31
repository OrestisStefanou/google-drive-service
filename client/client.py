import requests
import json

baseURL = 'https://radiant-gorge-35067.herokuapp.com/v1'


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


def get_auth_url(scope=""):
	url = f"{baseURL}/authenticationURL"
	if scope:
		payload = {'scope': scope}
		r = requests.get(url,params=payload)
	else:
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


def list_files(query=""):
	url = f"{baseURL}/files"
	try:
		f = open("token.json", "r")
		access_token = f.read()
		f.close()
	except:
		print("Token not found")
		return
	payload = {'query': query}
	headers = {'Authorization': access_token}
	r = requests.get(url,headers=headers,params=payload)
	if r.status_code == 200:
		files = r.json()['Files']
		for file in files:
			print(json.dumps(file,indent=2))
	else:
		print(r.json())


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


def download_exported_file(file_id,mimeType,filepath):
	url = f"{baseURL}/files/download_exported/{file_id}"
	payload = {"mimeType":mimeType}
	try:
		f = open("token.json", "r")
		access_token = f.read()
		f.close()
	except:
		print("Token not found")
		return
	headers = {'Authorization': access_token}
	r = requests.get(url,headers=headers,params=payload)	
	if r.status_code == 200:
		#print(r.content)
		f = open(filepath,"wb")
		f.write(r.content)
		f.close()
	else:
		print(r.json())


def create_folder(folder_name,parent_id=None):
	url = f"{baseURL}/files/folder"
	try:
		f = open("token.json", "r")
		access_token = f.read()
		f.close()
	except:
		print("Token not found")
		return
	#print("Access token is:",access_token)
	headers = {'Authorization': access_token}
	if parent_id:
		payload = {'folder_name': folder_name , "parent_id": parent_id }
	else:
		payload = {'folder_name': folder_name}
	r = requests.post(url, json=payload,headers=headers)
	response = r.json()
	print(response)



def upload_file(filepath,parent_id=None):
	url = f"{baseURL}/files/file"
	try:
		f = open("token.json", "r")
		access_token = f.read()
		f.close()
	except:
		print("Token not found")
		return
	#print("Access token is:",access_token)
	headers = {'Authorization': access_token}
	files = {'file': open(filepath, 'rb')}
	if parent_id:
		payload = {'parent_id': parent_id}
		r = requests.post(url, data=payload,files=files,headers=headers)
	else:
		r = requests.post(url,files=files,headers=headers)
	print(r.json())


def add_permission(file_id,role,permission_type,emails):
	url = f"{baseURL}/permissions/permission"
	try:
		f = open("token.json", "r")
		access_token = f.read()
		f.close()
	except:
		print("Token not found")
		return
	#print("Access token is:",access_token)
	headers = {'Authorization': access_token}
	payload = {
		'file_id': file_id,
		"role":role,
		"type":permission_type,
		"emails":emails
	}
	r = requests.post(url, json=payload,headers=headers)
	print(r.json())


def get_file_permissions(file_id):
	url = f"{baseURL}/permissions/{file_id}"
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
		permissions = r.json()['Permissions']
		for permission in permissions:
			print(json.dumps(permission,indent=2))
	else:
		print(r.json())	

#Run this function first to get an auth url
get_auth_url("DriveScope")

#Run create_token with the auth code that google returned
#create_token('<auth_code>')

# List the files
#list_files()
