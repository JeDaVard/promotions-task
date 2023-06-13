# Requirement

_note: the below requirement is copy pasted from the email_

We receive some records in a CSV file (example promotions.csv attached) every 30 minutes. We would like to store these objects in a way to be accessed by an endpoint.
Given an ID the endpoint should return the object, otherwise, return not found. 

Eg:
curl https://localhost:1321/promotions/1 {"id":"172FFC14-D229-4C93-B06B-F48B8C095512", "price":9.68, "expiration_date": "2022-06-04 06:01:20"}

Additionally, consider:
1. The .csv file could be very big (billions of entries) - how would your application perform?
2. Every new file is immutable, that is, you should erase and write the whole storage;
3. How would your application perform in peak periods (millions of requests per minute)?
4. How would you operate this app in production (e.g. deployment, scaling, monitoring)? 
5. The application should be written in golang;
6. Main deliverable is the code for the app including usage instructions, ideally in a repo/github gist.

# Solution

### Get the app running

Run mysql server
```shell
docker-compose up -d
```
Copy the .env.example file to .env
```shell
cp .env.example .env
```

Run the application
```shell
go run main.go
```

### Services
It exposes 2 API endpoints:
1. For bulk recording promotions (an example p.csv file with 200K records is on root dir)
2. For getting a promotion by id

For more information open the swagger documentation at http://localhost:3001/docs/index.html
(You can use Swagger UI or any API client e.g. Postman to test the endpoints)

### System document
The application is written in golang using gin and gorm. As a storage MySQL is selected. But depending on business problem it can be any sql/noSql or object storage. For writing CSV files I assumed that it's an http request for simplicity, but it can by an event or a cron job.

### Answers to additional
1. The .csv file could be very big (billions of entries) - how would your application perform?
    - The application breaks the csv down into chunks and uses bulk insert to insert the records in the database chunk by chunk. It is very fast and can handle billions of records
2. Every new file is immutable, that is, you should erase and write the whole storage;
    - Didn't quite understand the relationship between file immutability and erasing the storage. But if this is about the csv file that we receive is not a new list of promotions but the updated list of all promotions, then a flow can be developed to solve this problem. Simplest: truncate the table and before insert the new records (a transaction can be used to make it an atomic op)
3. How would your application perform in peak periods (millions of requests per minute)?
    - In peak periods we can have more replicas. Additionally we can separate customer API service from the "worker" service that writes CSV data to DB as it has consistent load.
4. How would you operate this app in production (e.g. deployment, scaling, monitoring)?
    - A CI/CD pipeline can be used for testing and dockerizing the app and initiating the deployment to the cloud provider. It's scaling policies can be easily configured in a Kubernetes cluster. And for monitoring we can use Prometheus and Grafana. Additionally for performance monitoring elastic APM can be used.
5. The application should be written in golang;
    - Yes, it is written in golang
6. Main deliverable is the code for the app including usage instructions, ideally in a repo/github gist.
    - Done

Thank you.