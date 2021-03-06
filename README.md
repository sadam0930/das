# Dancesport Application System (DAS)
DAS is an open and free competition management system for competitive ballroom
dance. This project (along with [dasdb](https://github.com/yubing24/dasdb) and 
[dasweb](https://github.com/yubing24/dasweb)) aims to provide the community of
dancesport an open and secure implementation of competition management system.

### Goals of DAS ###
* To provide a secure and robust solution to competition organizers and competitors. Currently, the most popular systems do not have the necessary security setup to protect itself from being compromised. Data security is the top priority of this project.

* To provide a free competition management solution to most amateur and collegiate competition organizers. A competition management system should not require a specially-trained technician to be useful for the organizer. Competition organizer’s limited budget should be spent on renting a great venue, inviting unbiased and professional adjudicators, and promote dancesport in society.

* To provide insightful information to competitors, including but not limited to:
  * Dancer and Couple profiles and statistics
  * Dancer and Couple rating, ranking, and progression
  * Adjudicator preference
  
* To help organizers manage competitions efficiently:
  * A more intuitive user interface for all users
  * Small functions that improve quality of life:
    * Set up typical collegiate and amateur competitions quickly
    * Competition finance management: managing entries, tickets, and lessons sales
    * Compatible with major federations’ requirements: WDSF, WDC, USA Dance, NDCA, and Collegiate (North America)
  * TBA partner search and matchmaking

* To provide opportunities for software developers and data enthusiasts:
  * API for developing custom client applications (competition results, statistics, etc.)
  * API for competition data and statistics

## Local Development Setup
DAS is mostly developed for running on Linux platform. Though
it is totally possible to develop and run on Windows, it's not tested
as thoroughly as Linux. We assumes that you are using Ubuntu. Most of
the setup works for Mac as well.

1. Install Go

   You can install Golang SDK through `apt-get`: `$ sudo apt-get install golang-go`
   
   Detailed documentation can be found [here](https://github.com/golang/go/wiki/Ubuntu).
   
2. Go environment

   You need to add `$GOPATH` to your environment (`~/.profile`): 
   
   `export GOPATH=$HOME/go`
   
   Sometimes we will need to use binaries built from other packages:
   
   `export PATH=$GOPATH/bin:$PATH`
   
   Make sure you define `$GOPATH` before adding `$GOPATH/bin` to your `PATH`
   
3. Check out the repository

   First, we need to create necessary directory for DAS:
   
   `$ mkdir -p $GOPATH/src/github.com/yubing24`
   
   Change directory:

   `$ cd $GOPATH/src/github.com/yubing24`
   
   Check out the latest build:
   
   `$ git clone https://github.com/yubing24/das.git`
   
4. Get dependencies

   Most of the dependencies can be get by `go get`: `$ go get ./...`
   This command will download all the dependencies from online repositories.
   
5. Run DAS

   You will need to have the database set up in order to run DAS locally. You can
   visit the [dasdb](https://github.com/yubing24/dasdb) to build the database schema
   for local development.
   
   To run DAS services locally, run `$ go run das.go` in the root of DAS repository.
   
## Necessary Development Tools
* IDE: IntelliJ (with Golang plugin), Goland (requires subscription), and VS Code
* Web Service: Postman, SoapUI

## Design and Development Principle.(Technical)

* **MVC-Based**: The architecture of DAS is inspired by design philosophies of ASP.NET MVC and .NET MVC Core.
If you are familiar with ASP.NET, it should be very easy to get familiar with the architecture of DAS.

* **Clean Architecture**: SOLID principles and Uncle Bob's *Clean Architecture* are heavily applied in DAS. We want the code easy to understand
for developers and easy to maintain. For example, `businesslogic` has no dependency on `dataaccess` or `controller` package. 
`controller` does not directly depend on `dataaccess`. This allows `businesslogic` to be developed
relatively independently without worrying about how it is going to be accessed by users or 
data sources. Similarly, code in `controller` can be changed more freely without concerning with
`dataaccess`. ISP allows the independent changes to different modules. Finally, `config` package
glue everything together and `das.go` gets the entire system started.

* **Test**, but not always driven by it: critical code should be tested as thoroughly as possible. There is
no hard requirement for test coverage, but we do our best to make sure the code executes correctly most of 
the time.