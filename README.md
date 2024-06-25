# CI/CD - Course [ Notely ]
![tests status badge](http://github.com/adamhu714/cicd-course/actions/workflows/ci.yml/badge.svg)

Within this project, I used github actions to implement both a continuous integration and a continuous deployment workflow.
The yml files created for the CI/CD are found in the .github folder.

The CI workflow I set up automatically runs unit tests, security tests, formatting and linting tests upon a pull request to the main branch.

For deployment, I used Google Cloud Platform to containerize and automatically deploy/update the application serverlessly whenever the main branch is pushed to.
Additionally, I set up a remote sqlite database using Turso and automated migrations in the github actions workflow,
as well as connected the database to the GCP deployment of the application.

Image showcasing online deployed application with persistent data: 

![image](https://github.com/adamhu714/cicd-course/assets/105497355/f13e78c1-d4a4-4940-993f-3935864f03e0)


---

This repo contains starter golang code for the "Notely" application for the "Learn CICD" course on [Boot.dev](https://boot.dev).

## Local Development

Make sure you're on Go version 1.22+.

Create a `.env` file in the root of the project with the following contents:

```bash
PORT="8080"
```

Run the server:

```bash
go build -o notely && ./notely
```

*This starts the server in non-database mode.* It will serve a simple webpage at `http://localhost:8080`.


