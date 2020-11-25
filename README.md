# jarvisESB

An Enterprise Service Bus {ESB} with multi-tenancy built in.

## [Heroku App](https://Notyetdeployed.herokuapp.com)

### [GitHub](https://github.com/deye9/ESB.git)

## Table of Contents

1. [**Project overview**](#project-overview)
2. [**UX**](#ux)

- [**User Stories**](#user-stories)
- [**Design**](#design)
  - [**Libraries**](#libraries)
  - [**Color Scheme**](#color-scheme)
  - [**Typography**](#typography)
- [**Wireframes**](#wireframes)

3. [**Features**](#features)

   - [**Existing Features**](#existing-features)
   - [**Features Left to Implement**](#features-left-to-implemement)

4. [**Technologies Used**](#technologies-used)

5. [**Databases Used**](#databases-used)

   - [**API - Spoonacular**](#api-spoonacular)
   - [**PostgreSQL**](#PostgreSQL)

6. [**Testing**](#testing)

   - [**Validators**](#validators)
   - [**Manual Testing**](#manual-testing)

7. [**Deployment**](#deployment)

8. [**Credits**](#credits)

   - [**Content**](#content)
   - [**Media**](#media)
   - [**Acknowledgements**](#acknowledgements)

---

## Project overview

An Enterprise Service Bus "ESB" using the Domain-Driven Design Principles. This application was built using Golang and uses the following rules

<ol>
    <li><b>Domain:</b> <br />
    This is where the domain and business logic of the application is defined.</li>
    <li><b>Infrastructure:</b> <br />
    This layer consists of everything that exists independently of our application: external libraries, database engines, routes and so on.</li>
    <li><b>Application:</b> <br />
    This layer serves as a passage between the domain and the interface layer. The sends the requests from the interface layer to the domain layer, which processes it and returns a response.</li>
    <li><b>Interface:</b> <br />
    This layer holds everything that interacts with other systems, such as web services, RMI interfaces or web applications, and batch processing frontends.</li>
</ol>

## Caveats

Middleware's should be arranged as needed e.g ["auth", "cors"] PostgresSQL is the database of choice. To switch it out kinda update your .env

To run your test use the command below clear && environment=test go test ./... -coverprofile cover.out && go tool cover -func cover.out && go tool cover -html cover.out
