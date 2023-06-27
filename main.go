
package main

// importing the required packages for the code.

import (
    "fmt"
    "os"
    "time"
    "gopkg.in/src-d/go-git.v4"

    "gopkg.in/src-d/go-git.v4/plumbing"

    "path/filepath"

    "gopkg.in/src-d/go-git.v4/plumbing/transport/http"

)

var (

    username        = "muzakkirsaifi123"

    repositoryName  = "terratest-for-terraform"

    destinationPath = "../" // Global variable for destination path

    folderName      = "module-terraform"                               // Global variable for folder name

)




func ClonePrivateRepo(token string, done chan bool) error {

    // Create the destination directory if it doesn't exist

    err := os.MkdirAll(destinationPath, os.ModePerm)

    if err != nil {

        return err

    }




    repoPath := filepath.Join(destinationPath, folderName)




    repoURL := fmt.Sprintf("https://github.com/%s/%s.git", username, repositoryName)



    _, err = git.PlainClone(repoPath, false, &git.CloneOptions{

        URL:           repoURL,

        Auth:          &http.BasicAuth{Username: "token", Password: token},

        ReferenceName: plumbing.ReferenceName("refs/heads/232323"), // or specify the branch you want to clone

    })



    done <- true




    return err

}




func DeleteClonedRepo() error {

    repoPath := filepath.Join(destinationPath, folderName)


    err := os.RemoveAll(repoPath)

    if err != nil {

        return err

    }

    return nil

}

func main() {

    token := "< token >"




    // Create a channel to signal the completion of the cloning process

    done := make(chan bool)




    // Start the cloning process

    go ClonePrivateRepo(token, done)




    // Wait for the cloning process to finish

    <-done




    fmt.Println("Repository cloned successfully!")




    // Continue with other operations if needed

    time.Sleep(10 * time.Second) // Sleep for an additional 2 seconds to allow any pending operations to complete




    // Delete the cloned repository

    err := DeleteClonedRepo()

    if err != nil {

        fmt.Printf("Error deleting cloned repository: %v\n", err)

    } else {

        fmt.Println("Cloned repository deleted.")

    }




    fmt.Println("Finished.")

}










