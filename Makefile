TIME := $(shell date +%s)
MIGRATIONS_FOLDER := migrations

.PHONY: migration
migration:
ifndef NAME 
	$(error NAME is not set)
endif
	@touch $(MIGRATIONS_FOLDER)/$(TIME)_$(NAME).up.sql
	@touch $(MIGRATIONS_FOLDER)/$(TIME)_$(NAME).down.sql

.PHONY: foo
foo:
	echo "foo"