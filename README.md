# Exercise Readme
### Name: Nicolas Marcelo Ambroa Bernadou
## Use Cases
### Creating An Account
#### Building An Account
- To create an account, first build one using the included [AccountBuilder](./internal/models/builder/builder.go).
- The building process begins with a call to `NewAccountBuilder()`. This will require you to provide only the mandatory info required by the Account API.
####
- Any other additions can be included by calling another builder method. For example, to add an IBAN, just call `WithIban()` after
calling the basic method, like so `NewAccountBuilder(params).WithIban(iban)`.
  - Keep in mind certain fields are not customizable per API docs, for example the field `Type` is always set to `ACCOUNTS` and not changeable.
  - The Builder was designed to be easily extensible so any future changes are quickly implemented. The restriction detailed above
    is fixed immediately with a new `WithAccountType()` method.
####
- To finalize construction, just call the `Build()` method. Keep in mind this method will enforce validation restrictions.
#### Calling the API
- After building an account, just call [Create(account)](./internal/api/accounts/create.go) and the account will be created
for you. This method will also return error information in case anything went wrong.
### Fetching An Account
- Call [Fetch(accountID)](./internal/api/accounts/fetch.go) and the account will be fetched for you. 
This method will also return error information in case anything went wrong (like an invalid ID).
### Deleting An Account
- Call [Delete(accountID, accountVersion)](./internal/api/accounts/delete.go) and the account will be deleted for you.
  This method will also return error information in case anything went wrong (like an invalid ID or Version).

## Considerations
- I am new to Go.
- The tests are made to execute by running `docker-compose up`. There is a dependency to this in the constants file. 
  Specifically in the constant [fake-api](./internal/constants.go). This should be changed if you want to migrate the tests to run locally.
  - This `fake-api` matches the hostname of the `accountapi` service in the [docker-compose.yml](docker-compose.yml) file.

## Improvements
- We could extend the Fake API Service with a Rate Limiter to make sure that an attack on that endpoint doesn't compromise the rest of the service holding the API.
- In this service we could add a Retry Policy with exponential backoff inside a Circuit Breaker to make sure we retry failed requests but not block the entire flow in case of perpetual timeouts returned by the API, for example.
- Logging should be added to this service to Log error details in a logging service (for example SumoLogic) in order to be able to triage potential issues and facilitate RCA concerns in case of incident.
- Another abstraction layer should be added to this service in case the API wants to be extended.
  - For example, if we want to support more Form3 resources instead of only accounts, we could have a `Form3Client` containing an `AccountClient` that can call `Create()`, `Fetch()`, `Delete()`. This allows us to extend resource support while maintaining clean code.
## Libraries used
- [Go UUID](github.com/nu7hatch/gouuid) to generate UUIDs to test the account API.
- [Go Testify](https://github.com/stretchr/testify) for the testing portion of the exercise.
- [Go Validator](https://github.com/go-playground/validator) to add validation for the account creation.

#
# Previous Info: Form3 Take Home Exercise

Engineers at Form3 build highly available distributed systems in a microservices environment. Our take home test is designed to evaluate real world activities that are involved with this role. We recognise that this may not be as mentally challenging and may take longer to implement than some algorithmic tests that are often seen in interview exercises. Our approach however helps ensure that you will be working with a team of engineers with the necessary practical skills for the role (as well as a diverse range of technical wizardry). 

## Instructions
The goal of this exercise is to write a client library in Go to access our fake account API, which is provided as a Docker
container in the file `docker-compose.yaml` of this repository. Please refer to the
[Form3 documentation](https://www.api-docs.form3.tech/api/tutorials/getting-started/create-an-account) for information on how to interact with the API. Please note that the fake account API does not require any authorisation or authentication.

A mapping of account attributes can be found in [models.go](./models.go). Can be used as a starting point, usage of the file is not required.

If you encounter any problems running the fake account API we would encourage you to do some debugging first,
before reaching out for help.

## Submission Guidance

### Shoulds

The finished solution **should:**
- Be written in Go.
- Use the `docker-compose.yaml` of this repository.
- Be a client library suitable for use in another software project.
- Implement the `Create`, `Fetch`, and `Delete` operations on the `accounts` resource.
- Be well tested to the level you would expect in a commercial environment. Note that tests are expected to run against the provided fake account API.
- Be simple and concise.
- Have tests that run from `docker-compose up` - our reviewers will run `docker-compose up` to assess if your tests pass.

### Should Nots

The finished solution **should not:**
- Use a code generator to write the client library.
- Use (copy or otherwise) code from any third party without attribution to complete the exercise, as this will result in the test being rejected.
    - **We will fail tests that plagiarise others' work. This includes (but is not limited to) other past submissions or open-source libraries.**
- Use a library for your client (e.g: go-resty). Anything from the standard library (such as `net/http`) is allowed. Libraries to support testing or types like UUID are also fine.
- Implement client-side validation.
- Implement an authentication scheme.
- Implement support for the fields `data.attributes.private_identification`, `data.attributes.organisation_identification`
  and `data.relationships` or any other fields that are not included in the provided `models.go`, as they are omitted from the provided fake account API implementation.
- Have advanced features, however discussion of anything extra you'd expect a production client to contain would be useful in the documentation.
- Be a command line client or other type of program - the requirement is to write a client library.
- Implement the `List` operation.
> We give no credit for including any of the above in a submitted test, so please only focus on the "Shoulds" above.

## How to submit your exercise

- Include your name in the README. If you are new to Go, please also mention this in the README so that we can consider this when reviewing your exercise
- Create a private [GitHub](https://help.github.com/en/articles/create-a-repo) repository, by copying all files you deem necessary for your submission
- [Invite](https://help.github.com/en/articles/inviting-collaborators-to-a-personal-repository) [@form3tech-interviewer-1](https://github.com/form3tech-interviewer-1) to your private repo
- Let us know you've completed the exercise using the link provided at the bottom of the email from our recruitment team

## License

Copyright 2019-2022 Form3 Financial Cloud

Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with the License.
You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the specific language governing permissions and limitations under the License.
