# UrbanBike
This repository will be used to develop the final project for the subject cloud computing.

# Index
* [Github and repository configuration](./doc/repository_preparation.md).
* [Problem and proposed solution](./doc/problem_definition.md)
* Tools
* Roadmap

# Architecture

After some study and research I have decided that the best Architecture for this project should be one based on microservices. There are several reasons for choose this Architecture, being all of them based on the advantages of this architecture:
* Microservices independence: In this service there will be some clearly separated functions, bei  ng the most important to give the users information about bike stations. Some other functions are important to, for example, to update the stations occupancy or a possible future function for print tickets when a user takes a bike for its rental. Anyway, any of these functions are independent, not needing any of the other functions to work.

* Scalability: Being independent, these microservices are easier to maintain and easier to scale.

* Availability of future upgrades: With a microservices based architecture, it is easier to create new functions for the app. This easiness came from the fact that microservices are isolated and independent, there are not going to be problems about integration between them.
