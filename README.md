# Weather-Monster


Dependencies

Golang 1.13 or higher

Docker for Mac/Windows

github.com/shyam81992/Weather-Monster-job // Which is used to notify the subscribers on temperature record creation

Steps to run the project 

    1. Create base folder and clone the projects github.com/shyam81992/Weather-Monster and github.com/shyam81992/Weather-Monster-job
    2. Run go mod tidy to install the dependencies.
    3. Open a new terminal and go to the folder base_folder/Weather-Monster/postgres and run the command docker-compose up (To start the postgres db).
    4. Open a new terminal and go to the folder base_folder/Weather-Monster/rabbitmq and run the command docker-compose up (To start the Rabbitmq).
    5. Open a new terminal and go to the folder base_folder/Weather-Monster-job and run the command 
        . config/dev.sh && go run main.go (To start the Weather-Monster-job)
    6. Open a new terminal and go to the folder base_folder/Weather-Monster and run the command 
        . config/dev.sh && go run main.go (To start the Weather-Monster)
    7. To run the testcases Open a new terminal and go to the folder base_folder/Weather-Monster/tescases and run the command . ../config/dev.sh && go test -v

Note: If you are running on windows please install git bash so that you can able to run the shell script  ". config/dev.sh"
    
Area of improvements:

PATCH /cities/:id route expects all the three parameters to be present in the patch body that can be improved to accept if any one of the  parameter is present in the patch body.

Delete /cities/:id  Now this route deletes the record from the city and adds it to the city_deleted table. I went with this approach because when we try to delete the historic data form temperature table it will result in performance degrade. So that now we can delete or archive the data form temperature table and webhook table lazily.
Another way(use cascade and reference constraints)


Get /forecasts/:city_id Query for this route can be optimized by creating compound index on id and created_at

Test cases: Now the test cases just verify only response status which can improved to validate the response data.

Weather-Monster-Job
Now it only listens to the temperature creation event and notify the subscriber which can be improved to have filter (ex : filter option can be provided at the webhook level). If notification rate per second is high we can improve the performance of it by scaling through region or by any other factor.(Kinesis, sqs, sns, rabbitmq)

 
