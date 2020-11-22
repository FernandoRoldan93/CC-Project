# UrbanBike
This repository will be used to develop the final project for the subject cloud computing.

# Index
* [Github and repository configuration](./doc/repository_preparation.md).
* [Problem and proposed solution](./doc/problem_definition.md)
* [Architecture](./doc/architecture.md)
* [Tools](./doc/tools.md)
* [Roadmap](./doc/roadmap.md)

## Architecture

After some study and research I have decided that the best architecture for this project should be one based on microservices. The reasons to choose this architecture comes from its advantages. For example, with this type of architecture we have different microservices, taking each one control of one feature of the system. This allow us to take some of the features as critical, for example, the management of the bike stations. This feature is one of the main features of the system and has to be operative most part of the time. Being an isolated microservice, it would be easier to recover it in case of failure or to scale it.

Looking at the rest of possible architectures, they does not fit for our project. Talking more specifically, about the multilayer architecture, the main problem in this kind of architectures, is that they are difficult to scale, being the scalability a requirement of the project: this kind of city transport has a lot of users and depending to the cities that we will give service, the number of users and the amount of information to handle can grow very fast.
About the event driven architecture, the problem is that testing and development are harder in this architecture and we need to have in mind that testing is a key part of the development: we cannot deploy something that has not been tested properly. Microkernel architecture does not fit here neither: even if we have a main feature (station management, for example), it is not the *core* of the system. And the last architecture that was taken in consideration was a monolithic one. Nevertheless this architecture provides a lot of disadvantages to the system, as for example, the difficulty to scale the project or even to maintain it.


## Tools

A lot of different tools will be used in the development of this project. These tools have to fit with the project in some way, this means that every tool should be choose with a specific purpose. In this section I will give a first look at these tools and the reasons to choose them. This section can be found on [this document](./doc/tools.md).

## Roadmap

In order to have a structure to follow or a general guide of the project, you will find in the next lines a roadmap of the project. This roadmap will give information about what are going to be the steps of the system development. Almost every point or step will be related to one [milestone](https://github.com/FernandoRoldan93/UrbanBike/milestones) of the project.

* [Basic Structure](https://github.com/FernandoRoldan93/UrbanBike/milestone/9): The main objective of this milestone will be to have a data structure defined that allows an administrator to create add new stations and bikes
  * [[US1] Create a station](https://github.com/FernandoRoldan93/UrbanBike/issues/7)
  * [[US2] Create a new bike in the system](https://github.com/FernandoRoldan93/UrbanBike/issues/8)


* [Basic operations](https://github.com/FernandoRoldan93/UrbanBike/milestone/10): At this point, having a data structure and some bikes and stations in the system, we want to be able of operate over that data structures, listing the objects we created and making operations over them, for example, rent or park a bike.
  * [[US3] Get all the stations created in the system](https://github.com/FernandoRoldan93/UrbanBike/issues/9)
  * [[US4] Get all the bikes created in the system](https://github.com/FernandoRoldan93/UrbanBike/issues/17)
  * [[US5] Park a bike in a station](https://github.com/FernandoRoldan93/UrbanBike/issues/18)
  * [[US6] Rent a bike](https://github.com/FernandoRoldan93/UrbanBike/issues/11)



* [Usage Statistics](https://github.com/FernandoRoldan93/UrbanBike/milestone/12): Having a way to operate over the objects we have in the system, we want some statistics over that usage, for example knowing the average occupation and the average number of rents in a certain station.
  * [[US7] Get the average occupancy of a station](https://github.com/FernandoRoldan93/UrbanBike/issues/12)
  * [[US8] Get the average of rented bikes in a station](https://github.com/FernandoRoldan93/UrbanBike/issues/14)

## Code structure

In this application, all the code will be placed, as usual, in the ['src' folder](./src) and at the moment there are only 2 classes implemented partially. Those classes are [Station](./src/station/station.go) and [Bike](./src/bike/bike.go).

The different modules will be developed and placed in folders representing their packages, an example can be found in the classes listed above. That code has been tested syntactically using ```gofmt -e <.go file>``` and ```go build``` trying to check for possible errors and syntax errors. The `gofmt` command will run automatically in a Github Action. This action will also run a linter on the YAML files trying to look for syntax errors.
