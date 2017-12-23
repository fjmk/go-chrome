package chrome

import (
	"encoding/json"

	database "github.com/mkenney/go-chrome/database"
	"github.com/mkenney/go-chrome/protocol"

	log "github.com/Sirupsen/logrus"
)

/*
Database - https://chromedevtools.github.io/devtools-protocol/tot/Database/
EXPERIMENTAL
*/
type Database struct{}

/*
Disable disables database tracking, prevents database events from being sent to the client.
*/
func (Database) Disable(
	socket *Socket,
) error {
	command := &protocol.Command{
		Method: "Database.disable",
	}
	socket.SendCommand(command)
	return command.Err
}

/*
Enable enables database tracking, database events will now be delivered to the client.
*/
func (Database) Enable(
	socket *Socket,
) error {
	command := &protocol.Command{
		Method: "Database.enable",
	}
	socket.SendCommand(command)
	return command.Err
}

/*
ExecuteSQL executes a SQL query.
*/
func (Database) ExecuteSQL(
	socket *Socket,
	params *database.ExecuteSQLParams,
) (database.ExecuteSQLResult, error) {
	command := &protocol.Command{
		Method: "Database.executeSQL",
		Params: params,
	}
	result := database.ExecuteSQLResult{}
	socket.SendCommand(command)

	if nil != command.Err {
		return result, command.Err
	}

	if nil != command.Result {
		resultData, err := json.Marshal(command.Result)
		if nil != err {
			return result, err
		}

		err = json.Unmarshal(resultData, &result)
		if nil != err {
			return result, err
		}
	}

	return result, command.Err
}

/*
GetDatabaseTableNames gets database table names.
*/
func (Database) GetDatabaseTableNames(
	socket *Socket,
	params *database.GetDatabaseTableNamesParams,
) error {
	command := &protocol.Command{
		Method: "Database.getDatabaseTableNames",
		Params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
OnAddDatabase adds a handler to the Database.addDatabase event. Database.addDatabase fires
whenever a database is added
*/
func (Database) OnAddDatabase(
	socket *Socket,
	callback func(event *database.AddDatabaseEvent),
) {
	handler := protocol.NewEventHandler(
		"Database.addDatabase",
		func(name string, params []byte) {
			event := &database.AddDatabaseEvent{}
			if err := json.Unmarshal(params, event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	socket.AddEventHandler(handler)
}