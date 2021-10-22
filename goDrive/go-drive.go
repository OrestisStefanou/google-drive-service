package goDrive

import (
        "context"
        "io"
        "io/ioutil"
        "log"

        "golang.org/x/oauth2"
        "golang.org/x/oauth2/google"
        "google.golang.org/api/drive/v3"
        "google.golang.org/api/option"
)

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


func getOauth2Config() (*oauth2.Config,error) {
        b, err := ioutil.ReadFile("credentials.json")
        if err != nil {
                log.Println("Couldn't read credentials")
                return nil,err
        }

        // If modifying these scopes, delete your previously saved token.json.
        config, err := google.ConfigFromJSON(b, drive.DriveMetadataScope, drive.DriveFileScope, drive.DriveScope)
        if err != nil {
                log.Println("google.ConfigFromJSON failed:",err)
                return nil,err
        }
        return config,nil        
}


func GetUserToken(authCode string) (*oauth2.Token,error) {
        config,err := getOauth2Config()
        if err != nil {
                return nil,err
        }
        //Get the token from google
        tok, err := config.Exchange(context.TODO(), authCode)
        if err != nil {
                log.Println("config.Exchange failed:",err)
                return nil,err
        }
        return tok,nil     
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
                log.Println("drive.NewService failed:",err)
                return nil,err
        }
        return service,nil           
}


//Function that returns a slice with pointers to drive.File 
func GetFileList(tok *oauth2.Token,query string) ([]*drive.File,error) {
        //Get client's service
        service,err := getClientService(tok)
        if err != nil {
                log.Println("getClientService failed:",err)
                return nil,err 
        }
        var filesList *drive.FileList
        if query == ""{
                filesList,err = service.Files.List().Fields("nextPageToken, files(id, name, mimeType,parents,size,webContentLink,webViewLink)").Do()
        } else {
                filesList,err = service.Files.List().Q(query).Fields("nextPageToken, files(id, name, mimeType,parents,size,webContentLink,webViewLink)").Do()
        }
        if err != nil {
                log.Println("Files list call failed:",err)
                return nil,err
        }
        return filesList.Files,nil
}


//Function that downloads a file with id = {fileId} from client's drive
func DownloadFile(tok *oauth2.Token,fileId string) ([]byte,error) {
        //Get client's service
        service,err := getClientService(tok)
        if err != nil {
                log.Println("getClientService failed:",err)
                return nil,err 
        }
        http_response,err := service.Files.Get(fileId).Download()
        if err != nil {
                log.Println(err)
                return nil,err 
        }
        file_data,err := ioutil.ReadAll(http_response.Body)
        if err != nil {
                log.Println(err)
                return nil,err
        }
        return file_data,nil
}


//Function that exports the file with id = {fileId} to mimeType = {mimeType}
//and then download's it from client's drive
func DownloadExportedFile(tok *oauth2.Token,fileId,mimeType string) ([]byte,error) {
        //Get client's service
        service,err := getClientService(tok)
        if err != nil {
                log.Println("getClientService failed:",err)
                return nil,err 
        }
        http_response,err := service.Files.Export(fileId,mimeType).Download()
        if err != nil {
                log.Println(err)
                return nil,err 
        }
        file_data,err := ioutil.ReadAll(http_response.Body)
        if err != nil {
                log.Println(err)
                return nil,err
        }
        return file_data,nil
}


//Function to create a new folder in client's drive
func CreateFolder(tok *oauth2.Token,folderName,parentId string) (*drive.File,error) {
        folder_metadata := new(drive.File)
        //Set the name of the new folder
        folder_metadata.Name = folderName
        //Set the parent of the new folder
        if len(parentId) > 0 {
                folder_metadata.Parents = append(folder_metadata.Parents,parentId)
        }
        //Set the mimeType to folder
        folder_metadata.MimeType = "application/vnd.google-apps.folder"
        //Get client's service
        service,err := getClientService(tok)
        if err != nil {
                log.Println("getClientService failed:",err)
                return nil,err 
        }
        //Make the call to create the new folder
        f,err := service.Files.Create(folder_metadata).Do()
        if err != nil {
                log.Println("Files.Create call failed:",err)
                return nil,err
        }
        return f,nil
}

//Function to upload a file in client's drive
func UploadFile(tok *oauth2.Token,file_to_upload io.Reader,parentId,filename string) (*drive.File,error) {
        file_metadata := new(drive.File)
        //Set the name of the uploaded file
        file_metadata.Name = filename
        //Set the parent folder if given
        if len(parentId) > 0 {
                file_metadata.Parents = append(file_metadata.Parents,parentId)
        }
        service,err := getClientService(tok)
        if err != nil {
                log.Println("getClientService failed:",err)
                return nil,err 
        }
        f,err := service.Files.Create(file_metadata).Media(file_to_upload).Do()
        if err != nil {
                log.Println("Files.Create call failed")
                return nil,err 
        }
        return f,nil

}