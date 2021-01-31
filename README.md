# Gian Carlo Pace
It's the first time I approach golang.

## My first steps in go. Random thoughts (31/01/2021)
### A new language
When I approach a new programming language, I try immediately to write some simple 
program that works: it is so rewarding to get a software running! 
But after a few hours, I start immediately to hit the hard ceiling of my limited knowledge.

That happened also approaching golang. 
I knew nothing about the ecosystem and libraries, about know how to build 
a more complex application, about how to create modules, about how to test the code.

This is where the path get steeper and it corresponds with the moment when I need to 
put some grounds where to build my knowledge before continuing on coding. So I started 
reading an the introductory book http://www.golang-book.com/books/intro. 
It's very basic and it is addressed to a very unexperienced software developer but,
skipping the really introductory information on what a function and a method are,
it puts all the starting concepts of the language one after the other. 

After a quick read I have more instruments to create the design I prefer and to 
be a bit more idiomatic in the usage of the language (avoiding some java-ish drifts).

## Growing Object-Oriented Software, Guided by Tests
One of my favourites books that shaped the way I write software is Growing Object-Oriented 
Software, Guided by Tests [GOOS01].
I really value this book as a great tool to build software at a sustainable pace.

In a nutshell the steps to build the software should be:
- Create a walking skeleton
- Write an acceptance test (fully integrated) that demonstrate the featrure the customer needs
- Write an incremental set of unit tests to make the acceptance test pass 
- Loop creating a new acceptance test

So I started from the first point.

[GOOS01] http://www.growing-object-oriented-software.com/

# Creating a walking skeleton
To create a walking skeleton I needed to know:
- How structure the code in the repository 
- how to build the code (with all the complications of the naming conventions)
- how to use the module system (confusion for the env var setup)
- how to create an launch tests for the whole application
- how to use mocks (that is really complicated confronted to Java)

To do so I read a lot of blogpost and also some github code to better understand how the 
major projects are behaving (e.g. https://www.wolfe.id.au/2020/03/10/how-do-i-structure-my-go-project/) 
and to have an idea on how an idiomatic golang codebase looks like.

After studying golang and digging in the details of the ecosystem, I needed a CI/CD system to code properly.
Therefore, I configured the github actions to:
- launch the `docker-compose-tests.yml` from the pipeline
- build the software and launch the tests

The `docker-compose-tests.yml` file is the docker compose file that is not launching 
the client part I developed. 
----
# Form3 Take Home Exercise

## Instructions
The goal of this exercise is to write a client library in Go to access our fake account API, which is provided as a Docker
container in the file `docker-compose.yaml` of this repository. Please refer to the
[Form3 documentation](http://api-docs.form3.tech/api.html#organisation-accounts) for information on how to interact with the API.

If you encounter any problems running the fake account API we would encourage you to do some debugging first,
before reaching out for help.

### The solution is expected to
- Be written in Go
- Contain documentation of your technical decisions
- Implement the `Create`, `Fetch`, `List` and `Delete` operations on the `accounts` resource. Note that filtering of the List operation is not required, but you should support paging
- Be well tested to the level you would expect in a commercial environment. Make sure your tests are easy to read.

#### Docker-compose
 - Add your solution to the provided docker-compose file
 - We should be able to run `docker-compose up` and see your tests run against the provided account API service 

### Please don't
- Use a code generator to write the client library
- Use (copy or otherwise) code from any third party without attribution to complete the exercise, as this will result in the test being rejected
- Use a library for your client (e.g: go-resty). Only test libraries are allowed.
- Implement an authentication scheme
- Implement support for the fields `data.attributes.private_identification`, `data.attributes.organisation_identification`
  and `data.relationships`, as they are omitted in the provided fake account API implementation
  
## How to submit your exercise
- Include your name in the README. If you are new to Go, please also mention this in the README so that we can consider this when reviewing your exercise
- Create a private [GitHub](https://help.github.com/en/articles/create-a-repo) repository, copy the `docker-compose` from this repository
- [Invite](https://help.github.com/en/articles/inviting-collaborators-to-a-personal-repository) @form3tech-interviewer-1 to your private repo
- Let us know you've completed the exercise using the link provided at the bottom of the email from our recruitment team

## License
Copyright 2019-2021 Form3 Financial Cloud

Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with the License.
You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the specific language governing permissions and limitations under the License.
