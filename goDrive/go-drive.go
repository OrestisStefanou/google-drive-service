package goDrive

import (
        "context"
        "encoding/json"
        "errors"
        "fmt"
        "io/ioutil"
        "log"
        "os"

        "golang.org/x/oauth2"
        "golang.org/x/oauth2/google"
        "google.golang.org/api/drive/v3"
        "google.golang.org/api/option"
)


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
        tok, err := GetTokenFromFile(tokenPath)
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
