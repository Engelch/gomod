package debugerrorce

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"regexp"
	"strings"
)

// @title               PostgreSQL helper routines
// @version             0.0.1
// @description         helper routines to check for existing permissions and tables.
// @contact.name        Christian Engel
// @contact.email       engel-ch@outlook.com
// @license.name        MIT

// / foreach export function
// PsqlVerifyTablePermissions
// @Summary             check for existence of table and permissions for the specified DB user
// @Description         PostgreSQL-specific code
// @Produce             error | nil
// @Success             nil "no error ⤳ tables existing in the database, user existing, user has required permissions"
// @Failure             error "error message returned in error message"

type PsqlTablePermission string

const (
	SELECT     PsqlTablePermission = "r"
	INSERT     PsqlTablePermission = "a"
	UPDATE     PsqlTablePermission = "w"
	DELETE     PsqlTablePermission = "d"
	TRUNCATE   PsqlTablePermission = "D"
	TRIGGER    PsqlTablePermission = "t"
	REFERENCES PsqlTablePermission = "x"
	ALL_PERMS  PsqlTablePermission = "arwdDxt"
)

func ArrayContains[T comparable](s []T, e T) bool {
	for _, v := range s {
		if v == e {
			return true
		}
	}
	return false
}

// PsqlListTablesInCurrentDatabase delivers all existing tables in the current database
func PsqlListTablesInCurrentDatabase(dbPool *pgxpool.Pool) ([]string, error) {
	var err error
	var dbmsAnswer string
	var returnTable []string

	if dbPool == nil {
		return nil, errors.New(CurrentFunctionName() + ":supplied dbPool is nil")
	}
	sqlRequest := "SELECT table_name FROM information_schema.tables WHERE table_schema='public'"
	rows, err := dbPool.Query(context.Background(), sqlRequest)
	if err != nil {
		return nil, errors.New("DBMS_ERROR request:" + sqlRequest + ":" + err.Error())
	}
	for rows.Next() {
		err := rows.Scan(&dbmsAnswer)
		if err != nil {
			log.Fatal(err)
		}
		returnTable = append(returnTable, dbmsAnswer)
	}
	return returnTable, nil
}

// / foreach export function
// PsqlVerifyTablePermissions
// @Summary             check for existence of table and permissions for the specified DB user
// @Description         PostgreSQL-specific code
// @Produce             error | nil
// @Success             nil "no error ⤳ tables existing in the database, user existing, user has required permissions"
// @Failure             error "USAGE_ERROR: wrong arguments to function"
// @Failure             error "DBMS_ERROR: DBMS returned an error"
// @Failure             error "TABLE_ERROR: permissions not existing for a table"
func PsqlVerifyTablePermissions(dbPool *pgxpool.Pool, user string, perms PsqlTablePermission, tables ...string) error {
	var sqlRequest string
	var dbmsAnswer string
	var err error
	var tablesExisting []string
	isAlNum := regexp.MustCompile(`^[A-Za-z][A-Za-z0-9]+$`).MatchString
	if user == "" {
		return errors.New("USAGE_ERROR:cannot check perms for unknown user")
	}
	user = strings.ToLower(user)
	if !isAlNum(user) {
		return errors.New("USAGE_ERROR:user name is not alNum and/or does not start with a letter")
	}
	if dbPool == nil {
		return errors.New("USAGE_ERROR:dbPool is nil")
	}
	CondDebugSet(true)
	tablesExisting, err = PsqlListTablesInCurrentDatabase(dbPool)
	if err != nil {
		return errors.New(CurrentFunctionName() + ":" + err.Error())
	}
	fmt.Printf("table names existing in DB are:\n\t%s\n", strings.Join(tablesExisting, "\n\t"))
	for _, table := range tables {
		if !isAlNum(table) {
			return errors.New("USAGE_ERROR:table name" + table + " is not allowed")
		}
		if !ArrayContains(tablesExisting, table) {
			return errors.New("DBMS_ERROR:table " + table + " is not in the current database")
		}
		for _, perm := range perms {
			sqlRequest = "select relacl from pg_class where relname = '" + table + "'"
			err = dbPool.QueryRow(context.Background(), sqlRequest).Scan(&dbmsAnswer)
			if err != nil {
				return errors.New("DBMS_ERROR request:" + sqlRequest + ":" + err.Error())
			}
			CondDebugln("DB returned:" + dbmsAnswer)
			if !strings.Contains(dbmsAnswer, user+"=") {
				return errors.New("DBMS_ERROR:answer for " + table + " does not contain:" + user + "=")
			}
			dbmsAnswer = dbmsAnswer[strings.Index(dbmsAnswer, user+"=")+1:]
			fmt.Println("dbmsAnswer is now:" + dbmsAnswer)
			if !strings.Contains(dbmsAnswer, "/") {
				return errors.New("DBMS_ERROR:answer for " + table + " does not contain /")
			}
			dbmsAnswer = dbmsAnswer[:strings.Index(dbmsAnswer, "/")]
			CondDebugln("DB answer processed:" + dbmsAnswer)
			if !strings.Contains(dbmsAnswer, string(perm)) {
				return errors.New("TABLE_ERROR:table:" + table + ":does not have permission:" + string(perm) + ":for user:" + user)
			}
		}
	}
	return nil
}

// PsqlGetUser returns the user name from a postgresql connection string.
func PsqlGetUser(constr string) (string, error) {
	if !strings.HasPrefix(constr, "postgresql://") {
		return "", errors.New("connection string does not start with postgresql://")
	}
	dbUser := constr[13:]
	index := strings.Index(dbUser, ":")
	if index == -1 {
		return "", errors.New("Cannot find user separation in:" + dbUser)
	}
	return dbUser[:index], nil
}
