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
    
Area of improvements :
PATCH /cities/:id route expects all the three parameters to be present in the patch  body that can be improved to accept if any one of the  parameter is present in the patch body.

Delete  /cities/:id   Now this route deletes the record from the city and adds it to the delete_city table. I went with this approach  because we try to delete the historic data form temperature table will result in performance issue. So that now we can delete or archive  the data form temperature table and webhook table lazily.
Another way(use cascade and reference constraints)

Test cases:  Now the test cases just verify only response status which can improved to validate the response data.

Weather-Monster-Job
Now it only listens to the temperature creation event and notify the subscriber which can be improved to have filter (ex : filter option can be provided at the webhook level). If notification per second is high we can improve the performance of it by scaling through region or by any other factor.(Kinesis,sqs, rabbitmq)

 
