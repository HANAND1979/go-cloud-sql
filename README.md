"# go-cloud-sql" 
Useful references
https://github.com/go-sql-driver/mysql/wiki/Examples
https://cloud.google.com/sql/docs/access-control#instanceaccess
https://cloud.google.com/appengine/docs/go/cloud-sql/reference

Important note: 
//This is the working code blocks for connecting to CloudSQL using Go from a Standard AppEngine environment
//Important reference - https://cloud.google.com/sql/docs/access-control#instanceaccess
//Make sure you choose the correct combination
//This code works for Standard AppEngine environment with First Generation CloudSQL instance
//Open issue: Insert query is still not very reliable - found to be inserting the same row multiple times
//Yet to try: Connecting AppEngine Managed VM with 2nd gen Cloud SQL instance
