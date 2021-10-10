package goDrive

import (
        "context"
        "encoding/json"
        "errors"
        "fmt"
        "io/ioutil"
        "io"
        "log"
        "net/http"
        "os"

        "golang.org/x/oauth2"
        "golang.org/x/oauth2/google"
        "google.golang.org/api/drive/v3"
        "google.golang.org/api/option"
)

//Service global variable
var srv *drive.Service

// Retrieve a token, saves the token, then returns the generated client.
func getClient(config *oauth2.Config) *http.Client {
        // The file token.json stores the user's access and refresh tokens, and is
        // created automatically when the authorization flow completes for the first
        // time.
        tokFile := "token.json"
        tok, err := tokenFromFile(tokFile)
        if err != nil {
                tok = getTokenFromWeb(config)
                saveToken(tokFile, tok)
        }
        return config.Client(context.Background(), tok)
}


// Request a token from the web, then returns the retrieved token.
func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
        authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
        fmt.Printf("Go to the following link in your browser then type the "+
                "authorization code: \n%v\n", authURL)

        var authCode string
        if _, err := fmt.Scan(&authCode); err != nil {
                log.Fatalf("Unable to read authorization code %v", err)
        }

        tok, err := config.Exchange(context.TODO(), authCode)
        if err != nil {
                log.Fatalf("Unable to retrieve token from web %v", err)
        }
        return tok
}


// Retrieves a token from a local file.
func tokenFromFile(file string) (*oauth2.Token, error) {
        f, err := os.Open(file)
        if err != nil {
                return nil, err
        }
        defer f.Close()
        tok := &oauth2.Token{}
        err = json.NewDecoder(f).Decode(tok)
        return tok, err
}


// Saves a token to a file path.
func saveToken(path string, token *oauth2.Token) {
        fmt.Printf("Saving credential file to: %s\n", path)
        f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
        if err != nil {
                log.Fatalf("Unable to cache oauth token: %v", err)
        }
        defer f.Close()
        json.NewEncoder(f).Encode(token)
}


//TEST AUTHENTICATION FUNCTIONS
// Retrieves a token from a local file.
func GetTokenFromFile(file string) (*oauth2.Token, error) {
        f, err := os.Open(file)
        if err != nil {
                return nil, err
        }
        defer f.Close()
        tok := &oauth2.Token{}
        err = json.NewDecoder(f).Decode(tok)
        return tok, err
}


func Get_user_auth_url() string {
        b, err := ioutil.ReadFile("credentials.json")
        if err != nil {
                log.Fatalf("Unable to read client secret file: %v", err)
        }

        // If modifying these scopes, delete your previously saved token.json.
        config, err := google.ConfigFromJSON(b, drive.DriveMetadataScope, drive.DriveFileScope, drive.DriveScope)
        if err != nil {
                log.Fatalf("Unable to parse client secret file to config: %v", err)
        }
        authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
        return authURL        
}

func get_user_auth_code(config *oauth2.Config) *oauth2.Token {
        var authCode string
        fmt.Println("Enter the authentication code")
        if _, err := fmt.Scan(&authCode); err != nil {
                log.Fatalf("Unable to read authorization code %v", err)
        }

        tok, err := config.Exchange(context.TODO(), authCode)
        if err != nil {
                log.Fatalf("Unable to retrieve token from web %v", err)
        }
        return tok        
}

func getClientTest(config *oauth2.Config) *http.Client {
        tok := get_user_auth_code(config)
        return config.Client(context.Background(), tok)
}


func getOauth2Config() (*oauth2.Config,error) {
        b, err := ioutil.ReadFile("credentials.json")
        if err != nil {
                fmt.Println("Couldn't read credentials")
                return nil,err
        }

        // If modifying these scopes, delete your previously saved token.json.
        config, err := google.ConfigFromJSON(b, drive.DriveMetadataScope, drive.DriveFileScope, drive.DriveScope)
        if err != nil {
                fmt.Println("google.ConfigFromJSON failed:",err)
                return nil,err
        }
        return config,nil        
}


func CreateUserToken(authCode,tokenPath string) (string,error) {
        config,err := getOauth2Config()
        if err != nil {
                return "",err
        }
        //Get the token from google
        tok, err := config.Exchange(context.TODO(), authCode)
        if err != nil {
                fmt.Println("config.Exchange failed:",err)
                return "",err
        }
        //Check if the token is valid
        _,err = getClientService(tok)
        if err != nil {
                return "",err
        }
        //Save the token 
        saveToken(tokenPath,tok)
        return tok.AccessToken,nil     
}


func TestInitializeService() {
        ctx := context.Background()
        b, err := ioutil.ReadFile("credentials.json")
        if err != nil {
                log.Fatalf("Unable to read client secret file: %v", err)
        }

        // If modifying these scopes, delete your previously saved token.json.
        config, err := google.ConfigFromJSON(b, drive.DriveMetadataScope, drive.DriveFileScope, drive.DriveScope)
        if err != nil {
                log.Fatalf("Unable to parse client secret file to config: %v", err)
        }
        //send_user_auth_url(config)
        client := getClientTest(config)

        srv, err = drive.NewService(ctx, option.WithHTTPClient(client))
        if err != nil {
                log.Fatalf("Unable to retrieve Drive client: %v", err)
        }        
}


//Function to check if tokenExists
func TokenExists(tokenPath string) (bool,error) {
    _, err := os.Stat(tokenPath)
    if err == nil {
        return true, nil
    }
    if errors.Is(err, os.ErrNotExist) {
        return false, nil
    }
    return false, err
}


//Function to get client's service
func getClientService(token *oauth2.Token) (*drive.Service,error) {
        config,err := getOauth2Config()
        if err != nil {
                return nil,err
        }
       //Get client
        client := config.Client(context.Background(), token)
        //Get client's google drive service
        ctx := context.Background()
        service, err := drive.NewService(ctx, option.WithHTTPClient(client))
        if err != nil {
                fmt.Println("drive.NewService failed:",err)
                return nil,err
        }
        return service,nil           
}

//Function to check if token is valid
func TokenIsValid(tokenPath string) (bool,error) {
        //Get the token
        tok, err := tokenFromFile(tokenPath)
        if err != nil {
                fmt.Println("tokenFromFile failed:",err)
                return false,err 
        }
        //Check if token works
        _,err = getClientService(tok)
        if err != nil {
                return false,err
        }
        return true,nil
}

//Function that returns a slice with pointers to drive.File 
func GetFileList(tok *oauth2.Token) ([]*drive.File,error) {
        //Get client's service
        service,err := getClientService(tok)
        if err != nil {
                fmt.Println("getClientService failed:",err)
                return nil,err 
        }
        filesList,err := service.Files.List().Fields("nextPageToken, files(id, name, mimeType,parents,size,webContentLink,webViewLink)").Do()
        if err != nil {
                fmt.Println("Files list call failed:",err)
                return nil,err
        }
        return filesList.Files,nil
}


//TEST AUTHENTICATION FUNCTIONS FINISHED

//Upload a file in the drive
func UploadFile(filepath, filename, parentFolderId string) error {
        file_metadata := new(drive.File)
        //Set the name of the uploaded file
        file_metadata.Name = filename
        //Set the parent folder
        file_metadata.Parents = append(file_metadata.Parents,parentFolderId)
        //Get the content of the file we want to upload
        f,err := os.Open(filepath)
        defer f.Close()
        if err != nil{
                return  err
        }
        _,err = srv.Files.Create(file_metadata).Media(f).Do()
        if err != nil {
                return err 
        }
        return nil
}


//Create a new folder in the drive
func CreateFolder(folderName,parentId string) error {
        folder_metadata := new(drive.File)
        //Set the name of the new folder
        folder_metadata.Name = folderName
        //Set the parent of the folder
        folder_metadata.Parents = append(folder_metadata.Parents,parentId)
        //Set the mimeType to a folder
        folder_metadata.MimeType = "application/vnd.google-apps.folder"
        //Create the folder
        _,err := srv.Files.Create(folder_metadata).Do()
        if err != nil {
                return err 
        }
        fmt.Println("Folder Created")
        return nil
}


//Create a local file to store a file we download from our drive
func createLocalFile(filepath string,filedata io.Reader) error {
        buf,err := ioutil.ReadAll(filedata)
        if err != nil {
                return err 
        }
        //Create and write the content to the new file
        err = ioutil.WriteFile(filepath,buf,0644)
        if err != nil {
                return err 
        }
        return nil        
}


//Download a file from drive and store it to 'filepath'
func DownloadFile(filepath,fileId string) error {
        http_response,err := srv.Files.Get(fileId).Download()
        if err != nil {
                return err 
        }
        err = createLocalFile(filepath,http_response.Body)
        return err
}


//Export and download a file
func ExportDownloadFile(filepath,fileId,mimeType string) error {
        http_response,err := srv.Files.Export(fileId,mimeType).Download()
        if err != nil {
                return err 
        }
        err = createLocalFile(filepath,http_response.Body)
        return err
}


func ListFiles() {
        r, err := srv.Files.List().Fields("nextPageToken, files(id, name, mimeType,parents)").Do()        
        if err != nil {
                log.Fatalf("Unable to retrieve files: %v", err)
        }
        fmt.Println("Files:")
        if len(r.Files) == 0 {
                fmt.Println("No files found.")
        } else {
                for _, i := range r.Files {
                        fmt.Printf("%s %s (%s) ", i.Name,i.MimeType,i.Id )
                        fmt.Printf("Parent:%v\n",i.Parents)
                }
        }
}


//Create a function to add permissions(access to other google users) to a file
func AddFilePermission(fileId,role,permissionType string,emails []string) error {
        // Role: The role granted by this permission
        // - owner
        // - organizer
        // - fileOrganizer
        // - writer
        // - commenter
        // - reader

        // permissionType: The type of the grantee. Valid values are:
        // - user
        // - group
        // - anyone  When creating a permission, if type is user or group, you
        // must provide an emailAddress for the user or group
        if permissionType == "anyone" {
                new_permission := new(drive.Permission)
                new_permission.Role = role 
                new_permission.Type = permissionType
                _,err := srv.Permissions.Create(fileId,new_permission).Do()
                if err != nil {
                        return err
                }
                return nil                
        }
        for _,email := range emails {
                new_permission := new(drive.Permission)
                new_permission.EmailAddress = email
                new_permission.Role = role 
                new_permission.Type = permissionType
                //Create a new permission for the file
                _,err := srv.Permissions.Create(fileId,new_permission).Do()
                if err != nil {
                        return err
                }
        }
        fmt.Println("Permissions added!")
        return nil
}


//Function to get information for a specific file
func GetFileMetadata(fileId string) (*drive.File,error) {
        file_metadata,err := srv.Files.Get(fileId).Fields("id, name, mimeType,webViewLink").Do()
        if err != nil {
                return nil,err 
        }
        return file_metadata,nil
}


//Query filesService 
func SearchFilesService(query string) error {
        filesList,err := srv.Files.List().Q(query).Fields("nextPageToken, files(id, name, mimeType,parents)").Do()
        if err != nil {
                return err 
        }
        fmt.Println("Search results")
        if len(filesList.Files) == 0 {
                fmt.Println("No results")
                return nil
        }
        for _, file_metadata := range filesList.Files {
                fmt.Printf("Filename: %s, Id: %s ,MimeType:%s ,Parents:%v\n", file_metadata.Name, file_metadata.Id, file_metadata.MimeType,file_metadata.Parents)
        }
        return nil
}


//Funtion to list the permissions of a file
func GetFilePermissions(fileId string) error {
        permissionsList,err := srv.Permissions.List(fileId).Fields("nextPageToken, permissions(emailAddress, role, type)").Do()
        if err != nil {
                return err 
        }
        if len(permissionsList.Permissions) == 0 {
                fmt.Println("No permissions found for this file")
                return nil
        }
        for _,permission := range permissionsList.Permissions {
                fmt.Printf("%s ,Role:%s ,Type:%s\n",permission.EmailAddress,permission.Role,permission.Type)
        }
        return nil
}


//Function to initialize the google drive service
func InitializeService() {
        ctx := context.Background()
        b, err := ioutil.ReadFile("credentials.json")
        if err != nil {
                log.Fatalf("Unable to read client secret file: %v", err)
        }

        // If modifying these scopes, delete your previously saved token.json.
        config, err := google.ConfigFromJSON(b, drive.DriveMetadataScope, drive.DriveFileScope, drive.DriveScope)
        if err != nil {
                log.Fatalf("Unable to parse client secret file to config: %v", err)
        }
        client := getClient(config)

        srv, err = drive.NewService(ctx, option.WithHTTPClient(client))
        if err != nil {
                log.Fatalf("Unable to retrieve Drive client: %v", err)
        }        
}
