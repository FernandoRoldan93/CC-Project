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

In order to have a structure to follow or a general guide of the project, you will find in the next lines a roadmap of the project. This roadmap will give information about what are going to be the steps of the project development. Almost every point or step will be related to one [milestone](https://github.com/FernandoRoldan93/UrbanBike/milestones) of the project. You can find further information and a list of issues on each one of the milestone pages.

1. [Architecture definition](https://github.com/FernandoRoldan93/UrbanBike/milestone/2): As a first step we have to choose what architecture will have the application.

2. [Development of the different modules](https://github.com/FernandoRoldan93/UrbanBike/milestone/3): This step will be parallel to the other ones, the reason is that we can always add more modules or functionalities to the project. This also include the development of the different tests on each of the modules that requires it. But in any case this step will include, at least, the next parts:
  * Station class
  * City class
  * [External API reader](https://github.com/FernandoRoldan93/UrbanBike/milestone/5)
  * Data storage
  * Configuration module

3. [Continuous Integration](https://github.com/FernandoRoldan93/UrbanBike/milestone/8): Adding tools that will allow to check all the tests after every change on the system.

4. [Dockerize the project](https://github.com/FernandoRoldan93/UrbanBike/milestone/6): Compression of the system in containers that will allow to deploy it easily on any machine.

5. [Deploy](https://github.com/FernandoRoldan93/UrbanBike/milestone/7): Once we have the system in a container we can deploy it in any cloud provider of our choice.

## Code structure

In this application, all the code will be placed, as usual, in the ['src' folder](./src) and at the moment there are only 2 classes implemented partially. Those classes are [Station](./src/station/station.go) and [City](./src/station/city.go).

The different modules will be developed and placed in folders representing their packages, an example can be found in the classes listed above. That code has been tested syntactically using 'gofmt -e <.go file>' and 'go build' trying to check for possible errors and syntax errors.
