package main

import (
	"bufio"
    "fmt"
    "os"
    "os/user"
    "strings"
    "github.com/plar/go-adaptive-radix-tree"
    "io/ioutil"
    "encoding/json"
)



func main() {
	

	usr, _ := user.Current()
    uhdir := usr.HomeDir
	IPASN,_ := os.Open(uhdir+"/asn/IPASN.DAT")
	defer IPASN.Close()

	t  := art.New()

	scanner := bufio.NewScanner(IPASN)
	
	
	for scanner.Scan() {
		if strings.HasPrefix(scanner.Text(), ";") == false {
			words := strings.Fields(scanner.Text())

			value, found := t.Search(art.Key(words[1])) 
            
            if found {
            	newVal := fmt.Sprintf("%v,%s", value,words[0])
            	t.Insert(art.Key(words[1]),newVal )
			} else {
				t.Insert(art.Key(words[1]), words[0])

			}
		}
	}
	

	AsnJson, _ := os.Open(uhdir+"/asn/asn.json")
	defer AsnJson.Close()

   	byteAsnJson, _ := ioutil.ReadAll(AsnJson)

    org := os.Args[1]

    var resAsnJson map[string]interface{}
   	json.Unmarshal([]byte(byteAsnJson), &resAsnJson)

   	for k,w :=range resAsnJson {
   		switch v := w.(type) {
		
		case string:
			fndwords := strings.Fields(strings.ToLower(v))

            for _,x := range fndwords {

				if strings.HasPrefix(x, strings.ToLower(org)) {

					value, found := t.Search(art.Key(string(k)))
				    if found {
				        fmt.Printf("\n# [AS%s] %s\n%v",string(k),v,value)
				    }
									
				}
			}

		
		}
   	
        
        	    
}
}
