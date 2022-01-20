**Gophercises**

In this repo there are 4 Different mini-applications, packages, and tools that are each designed to teach something different. </br>
Through the course I will learn about and practice using: </br>
Channels </br>
Functional Options </br>
Chaining Interfaces </br>
Various parts of the standard library (io, time, and many other packages) </br>
Reading input from the command line </br>
And much much more! </br>

**Build and Insatll Go Program**

1.go env GOPATH </br>
2. cd .. </br>
3. go install  </br>
4. ls $GOPATH/bin </br>
5. go build </br>
6. go run main.go </br>


**Exercise details** </br>
**1. caeserCipher**

     Caesar cipher, involves encrypting a string using a fairly basic cipher where you are changing the ascii value of each character.
     
     Input:- Add any character or string in casar.in file and in the terminal just command "go run main.go < camel.in"
              e.g:- middle-Outz😉
              
     OutPut:- okffng-Qwvb😉
     
     Note:- This change the ascii value only for charecter/string not for numbers, symbols etc.
     
**2. camelCase**              
       Camel Case goal is to find out how many words were used to create a camelcase string.
       
       Input:- Add any string with capital small and run "go run main.go < camel.in"
              e.g:- saveChangesInTheEditor
              
       Output:- index: 0 Run: 115 
                index: 1 Run: 97
                index: 2 Run: 118
                index: 3 Run: 101
                index: 4 Run: 67
                index: 5 Run: 104
                index: 6 Run: 97
                index: 7 Run: 110
                index: 8 Run: 103
                index: 9 Run: 101
                index: 10 Run: 115
                index: 11 Run: 73
                index: 12 Run: 110
                index: 13 Run: 84
                index: 14 Run: 104
                index: 15 Run: 101
                index: 16 Run: 69
                index: 17 Run: 100
                index: 18 Run: 105
                index: 19 Run: 116
                index: 20 Run: 111
                index: 21 Run: 114
                9
                
   **3. quizGame**
            In this exercixe, will read in a quiz provided through a CSV file and will then give the quiz to keeping track of how many questions get right and how many get                   incorrect.
            
           The CSV file should default to **'problems.csv'** 
           The default time limit should be 30 seconds, but should also be customizable. Quiz should stop as soon as the time limit has exceeded. That is, you shouldn’t wait for            the user to answer one final questions but should ideally stop the quiz entirely even if you are currently waiting on an answer from the end user.
           
           Output:- 
                Problem #1: 5+5 = 2
                Problem #2: 7+3 = 2
                Problem #3: 1+1 = 23
                Problem #4: 8+3 = 11
                Problem #5: 1+2 = 3
                Problem #6: 8+6 = 15
                Problem #7: 3+1 = 4
                Problem #8: 1+4 = 5
                Problem #9: 5+1 = 6
                Problem #10: 2+3 = 5
                Problem #11: 3+3 = 6
                Problem #12: 2+4 = 6
                Problem #13: 5+2 = 7

                Scored 9 out of 13.
                
   **4.urlShortener**            
        In urlShortener exercise, to create an http.Handler that will look at the path of any incoming web request and determine if it should redirect the user to a new                 page,much like URL shortener would.
        
        On command line run this "go run main/main.go" command and ou will see server will start on port e.g:- 8081.
        
        
        
        
        
       
