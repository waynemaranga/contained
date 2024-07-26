# `contained`

![go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![docker](https://img.shields.io/badge/Docker-2496ED?style=for-the-badge&logo=docker&logoColor=white)
![mongodb](https://img.shields.io/badge/MongoDB-47A248?style=for-the-badge&logo=mongodb&logoColor=white)
![postgresql](https://img.shields.io/badge/PostgreSQL-4169E1?style=for-the-badge&logo=postgresql&logoColor=white)
![aws](https://img.shields.io/badge/Amazon_AWS-FF9900?style=for-the-badge&logo=amazonaws&logoColor=white)
![azure](https://img.shields.io/badge/Azure-0089D6?style=for-the-badge&logo=microsoft-azure&logoColor=white)
![gcp](https://img.shields.io/badge/GCP-4285F4?style=for-the-badge&logo=google-cloud&logoColor=white)
![cloudflare](https://img.shields.io/badge/Cloudflare-F38020?style=for-the-badge&logo=cloudflare&logoColor=white)
![circleci](https://img.shields.io/badge/CircleCI-343434?style=for-the-badge&logo=circleci&logoColor=white)
![github_actions](https://img.shields.io/badge/GitHub_Actions-2088FF?style=for-the-badge&logo=github-actions&logoColor=white)

<!-- ![nginx](https://img.shields.io/badge/NGINX-009639?style=for-the-badge&logo=nginx&logoColor=white) -->
A simple containerised Go program.

`go build -o contained ./cmd/contained/main.go`

or...

`docker build --tag contained -f ./.docker/Dockerfile .`

 `docker run --detach --publish 8080:8080 contained`