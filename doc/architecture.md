# Architecture
After some study and research I have decided that the best Architecture for this project should be one based on microservices. There are several reasons for choose this Architecture, being all of them based on the advantages of this architecture:
* Microservices independence: In this service there will be some clearly separated functions, being the most important to give the users information about bike stations. Some other functions are important to, for example, to update the stations occupancy or a possible future function for print tickets when a user takes a bike for its rental. Anyway, any of these functions are independent, not needing any of the other functions to work.

* Scalability: Being independent, these microservices are easier to maintain and easier to scale.

* Availability of future upgrades: With a microservices based architecture, it is easier to create new functions for the app. This easiness came from the fact that microservices are isolated and independent, there are not going to be problems about integration between them.

A first approach to the microservices that will be used is listed below:

1. Bike Stations: This microservice will be used to list and modify the information about bike stations. Clients will connect to this microservice to get information, like free slots, of a specified station.

2. City data: This microservice will allow users to get info about where and how many bike stations are in their city.

3. *Telegram Bot:* This microservice would allow clients to get info about an specified station throw Telegram application. The development of this module could be discarded.

4. Configuration management: This microservice will be the responsible for the storage of the configuration. This will allow us to have the modules configuration in the same place.

5. Data store: In this microservice, all the data will be stored. The other microservices will request the information they need to this module.
