# Weather-Monster


Dependencies

Golang 1.13 or higher

Docker for Mac/Windows

github.com/shyam81992/Weather-Monster-job // Which is used to notify the subscribers on temperature record creation

Steps to run the project 

    1. Run go mod tidy to install the dependencies.
    2. Open a new terminal and go to the folder postgres 
    3. Run the command docker-compose up (To start the postgres db)
    4. Open a new terminal and go to the folder rabbitmq and run the command docker-compose up (To start the Rabbitmq)
    5. clone the github.com/shyam81992/Weather-Monster-job in a separate folder and run the command 
        . config/dev.sh && go run main.go
    6. open a new terminal and got to the Weather-Monster project root folder and run the command 
        . config/dev.sh && go run main.go

Note: If you are running on windows please install git bash so that you can the shell script  ". config/dev.sh"
