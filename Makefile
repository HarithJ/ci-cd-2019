DOCKER_DEV_COMPOSE_FILE := docker/dev/docker-compose.yml
DOCKER_PROD_FILE := docker/prod/Dockerfile

start-build:
	@ docker-compose -f $(DOCKER_DEV_COMPOSE_FILE) build
	@ docker-compose -f $(DOCKER_DEV_COMPOSE_FILE) run app /bin/ash

start-nobuild:
	@ docker-compose -f $(DOCKER_DEV_COMPOSE_FILE) run app /bin/ash

test-prod:
	@ docker build -f $(DOCKER_PROD_FILE) -t harithj/ci-cd-2019 .
	@ docker push harithj/ci-cd-2019

production:
	@ docker build -f $(DOCKER_PROD_FILE) -t gcr.io/apprenticeship-projects-1/ci-cd-2019 .
	@ docker push gcr.io/apprenticeship-projects-1/ci-cd-2019

ssh-production:
	@ docker run -it gcr.io/apprenticeship-projects-1/ci-cd-2019 /bin/ash
