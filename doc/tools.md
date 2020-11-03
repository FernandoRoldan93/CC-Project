# Tools

A lot of different tools will be used in the development of this project. These tools have to fit with the project in some way, this means that every tool should be choose with a specific purpose. In this section I will give a first look at these tools and the reasons to choose them. In further updates more tools will be added.

### Language
Talking about programming languages we can find a lot of options like Java, Python, C++, etc.. But for this project I have decided to use [Golang](https://golang.org/). The reason to choose this language is that Go is having a huge rise in popularity and I do not have any experience in this language, being this project the perfect opportunity to learn more about it.

### Framework for API
For this purpose I think that the best solution would be [Gin](https://github.com/gin-gonic/gin). The reasons to choose it came from its huge popularity, having 40k+ stars on Github. It is also ideal for Golang starters as it is said in this [article](https://medium.com/devtechtoday/top-7-golang-web-frameworks-in-2020-and-beyond-9ca2a89eb904), being perfect for us.

### Distributed configuration
About the distributed configuration, I decided to use [etcd](https://etcd.io/) due to its easiness and popularity.

### Data storage
After some consideration, the problem data objects, does not have any relation between them. Because of this, a relational database is not necessary being a non-relational DB the other option. Talking about non-relational DB, the option that I prefer is MongoDB. The reasons to choose it are that it is familiar to me, easy to use while having a huge popularity.
