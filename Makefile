TIME := $(shell date +%s)
MIGRATIONS_FOLDER := migrations

.PHONY: migration
migration:
ifndef NAME 
	$(error NAME is not set)
endif
	@touch $(MIGRATIONS_FOLDER)/$(TIME)_$(NAME).up.sql
	@touch $(MIGRATIONS_FOLDER)/$(TIME)_$(NAME).down.sql

GENERATED := GameType HeroClass Result

$(GENERATED): %:
	cd models && stringer -type=$@
	cd models && jsonenums -type=$@

.PHONY: generate
generate: $(GENERATED)