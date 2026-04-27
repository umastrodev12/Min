package updater

import (
	"encoding/json"
	"os"
)

type updateInfo struct {
    UpdateName string `json:"name"`  
    Version    string `json:"version"`
}


func WriteUpdateJson() {
    info := updateInfo{UpdateName: "BETA Mighty", Version: "0.0.1"}
    
    
    file, err := os.Create("update.json") // verify the error on creating the file
    if err != nil {
        panic(err)
    }
    defer file.Close() 
    
    enc := json.NewEncoder(file)
    if err := enc.Encode(info); err != nil {
        panic(err)
    }
}
