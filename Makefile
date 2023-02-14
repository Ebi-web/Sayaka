DBUSER:=postgres
DBPASSWORD:=password
DBPORT:=5432
DBNAME:=sayaka
DOCKER_DNS:=db
DOCKER_COMPOSE_IMPL:=docker-compose
FLYWAY_CONF?=-url=jdbc:postgresql://$(DOCKER_DNS):$(DBPORT)/$(DBNAME) -user=$(DBUSER) -password=$(DBPASSWORD)
MIGRATION_SERVICE:=migration
.PHONY: flyway/info
flyway/info:
	$(DOCKER_COMPOSE_IMPL) run --rm $(MIGRATION_SERVICE) $(FLYWAY_CONF) info

.PHONY: flyway/validate
flyway/validate:
	$(DOCKER_COMPOSE_IMPL) run --rm $(MIGRATION_SERVICE) $(FLYWAY_CONF) validate

.PHONY: flyway/migrate
flyway/migrate:
	$(DOCKER_COMPOSE_IMPL) run --rm $(MIGRATION_SERVICE) $(FLYWAY_CONF) migrate

.PHONY: flyway/repair
flyway/repair:
	$(DOCKER_COMPOSE_IMPL) run --rm $(MIGRATION_SERVICE) $(FLYWAY_CONF) repair

.PHONY: flyway/baseline
flyway/baseline:
	$(DOCKER_COMPOSE_IMPL) run --rm $(MIGRATION_SERVICE) $(FLYWAY_CONF) baseline