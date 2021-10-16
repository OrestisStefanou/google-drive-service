package goDrive

import (
        "context"
        "fmt"
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


func GetUserToken(authCode string) (*oauth2.Token,error) {
        config,err := getOauth2Config()
        if err != nil {
                return nil,err
        }
        //Get the token from google
        tok, err := config.Exchange(context.TODO(), authCode)
        if err != nil {
                fmt.Println("config.Exchange failed:",err)
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
                fmt.Println("drive.NewService failed:",err)
                return nil,err
        }
        return service,nil           
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

//Function that downloads a file with id = {fileId} from client's drive to specified filepath
func DownloadFile(tok *oauth2.Token,fileId string) ([]byte,error) {
        //Get client's service
        service,err := getClientService(tok)
        if err != nil {
                fmt.Println("getClientService failed:",err)
                return nil,err 
        }
        http_response,err := service.Files.Get(fileId).Download()
        if err != nil {
                fmt.Println(err)
                return nil,err 
        }
        file_data,err := ioutil.ReadAll(http_response.Body)
        if err != nil {
                fmt.Println(err)
                return nil,err
        }
        return file_data,nil
}