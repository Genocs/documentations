Welcome to Microshop!
===================

This document describe how to setup a new project from scratch!


----------

Setup Idea
-------------


Project management
-------------
- VSTS



Architecture
-------------
- Clean Architecte
- Microservise Architecture


DevOps
-------------
Source code Versionig (The code repository)
- VSTS Github



CI/CD (Continuos Integration / Continuous Deploy) (How build and run tests)
- VSTS Pipeline


Package manager (How store the libraries and container images)
- VSTS
- Nuget
- DockerHub


The platforms (Where execute your solution)

- Public Cloud
	- Azure		
	- Google Cloud


Software development

- Backend services
- Frontend Services
- Native Mobile App

- Data analysis and processing


### Backens services

Programming languages
- .NET C#
- Typescript
- html/css
- Java
- Kotlin
- Nodejs
- Python
 


Setup Components
-------------
## Main Software components setup

- Azure Infrastructure Services
- Google Firebase Infrastructure
- FE BackOffice Angular Project
- BE Web API Microservice .NET Core Project
- Android Native Mobile App Project
- Infrastructure External Services
    - Mongo DB cloud Project
    - Enterprise Service Bus Clud Project
	
	
### Setup timeline


| Setup Activity                               | time spend   |
| :------------------------------------------- | :------------|
| Azure Infrastructure Services                |  2h          |
| Google Firebase Infrastructure               |  2h          |
| FE BackOffice Angular Project                |  2h          |
| BE Web API Microservice .NET Core Project    |  2h          |
| Infrastructure MongoDb            		   |  1h          |
| Infrastructure Enterprise Service Bus        |  1h          |





## Azure Infrastructure Services

Following are the step required to setup the cloud infrastructure needed to deploy the backend services.

The backend services are deveoloped using .NET core 3.1 on Linux.


1. Connect to [Azure portal](https://portal.azure.com/)
2. Create a resource group: RG-Genocs
3. Create the app service plan: ASP-RGGenocs-core
4. Create the app service: microshop-anagraph-api

NOTE: .NET core app service do not allows to use appinsight. Evaluate to use on OpenTrace framework like Jagger



## Google Firebase Infrastructure

Following are the step required to setup the cloud infrastructure needed to deploy the front-end portal and to provide social auth login support.


1. Connect to [Firebase Console](https://firebase.google.com/)




## FE BackOffice Angular Project

This section describe how to setup the Angular front-end portal, confifure it using the Auth token coming from Google firebase platform and pass it to authenticate API.





## BE Web API Microservice .NET Core Project

The backend infrastructure ecosystem is compose by a set of microservices buid using Clean architecture patterns tu build the a bounch of Microservices. 



The template

