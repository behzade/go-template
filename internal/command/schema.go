package command

import "github.com/behzade/go-template/internal/service"

type SyncEntityCommand struct{}

func (SyncEntityCommand) Run() (string, error) {
	err := service.RenderSchemaTemplate()
	if err != nil {
		return "", err
	}

	msg, err := service.SyncEntityWithSchema()
	return string(msg), err
}

func (SyncEntityCommand) Name() string {
	return "sync_entity"
}

func (SyncEntityCommand) Description() string {
	return "Syncs the entity package with the schema defined in the schema.tmpl.sql"
}

type GetAlterCommand struct{}

func (GetAlterCommand) Run() (string, error) {
	err := service.RenderSchemaTemplate()
	if err != nil {
		return "", err
	}

	res, err := service.SyncDBWithSchema(true)
	if err != nil {
		return "", err
	}

	return res.String(), err
}

func (GetAlterCommand) Name() string {
	return "get_alters"
}

func (GetAlterCommand) Description() string {
	return "dry run schema apply on the db and return pending alters"
}

type ApplyAltersCommand struct{}

func (ApplyAltersCommand) Run() (string, error) {
	err := service.RenderSchemaTemplate()
	if err != nil {
		return "", err
	}

	res, err := service.SyncDBWithSchema(false)
	if err != nil {
		return "", err
	}
	return res.String(), err
}

func (ApplyAltersCommand) Name() string {
	return "apply_alters"
}

func (ApplyAltersCommand) Description() string {
	return "schema apply on the db and return the applied alters"
}
