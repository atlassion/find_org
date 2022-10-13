package main

import (
    "fmt"
    "github.com/reinhrst/fzf-lib"
    "encoding/json"
    "os"
    "path/filepath"
    "os/user"
    "io/ioutil"
    "strings"
)

func getAsnDir() (string, error) {
    usr, err := user.Current()
    if err != nil {
        return "", err
    }
    
    return filepath.Join(usr.HomeDir, "asn"), nil
}


func findasn(findit,v string) {
    dirAsn,_ := getAsnDir()

    ipasnBytes, err := ioutil.ReadFile(dirAsn+"/IPASN.DAT")
    if err != nil {
        panic(err)
    }

    quotes := strings.Split(string(ipasnBytes), "\n")
    
    var options = fzf.DefaultOptions()
  

    var myFzf = fzf.New(quotes, options)
    var result fzf.SearchResult
    

    myFzf.Search(fmt.Sprintf(" %s$",findit))
    result = <- myFzf.GetResultChannel()
        
    for _, match := range result.Matches{
        fmt.Println(match.Key,v)

    }
    
    myFzf.End()
    }

func main() {
    dirAsn,_ := getAsnDir()
    AsnJson, _ := os.Open(dirAsn+"/asn.json")
    defer AsnJson.Close()

    byteAsnJson, _ := ioutil.ReadAll(AsnJson)
    org := os.Args[1]

    var resAsnJson map[string]interface{}
    json.Unmarshal([]byte(byteAsnJson), &resAsnJson)

    for findit, w := range resAsnJson {
        switch v := w.(type) {
        case string:
            fndwords := strings.Fields(strings.ToLower(v))

            for _, x := range fndwords {

                if strings.HasPrefix(x, strings.ToLower(org)) {

                    findasn(findit,v)

                }
            }

        }

    }
 
}
