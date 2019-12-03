package main

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/stretchr/testify/mock"
)

type mockDB struct {
	*mock.Mock
}

func (d *mockDB) Query(query string, args ...interface{}) (*sql.Rows, error) {
	d.Called(query, args)
	return nil, nil
}

func (d *mockDB) Begin() (*sql.Tx, error) {
	d.Called()
	return nil, nil
}

func (d *mockDB) QueryRow(query string, args ...interface{}) *sql.Row {
	d.Called(query, args)
	return nil
}

func Test_DeletePGRepo(t *testing.T) {
	repoName := "test"
	m := &mockDB{&mock.Mock{}}
	tables := []string{chartTable, repositoryTable, chartFilesTable}
	for _, table := range tables {
		q := fmt.Sprintf("DELETE FROM %s WHERE info -> 'repo' ->> 'name' = '%s'", table, repoName)
		// Since we are not specifying any argument, Query is called with []interface{}(nil)
		m.On("Query", q, []interface{}(nil))
	}

	pgManager := &postgresAssetManager{m}
	err := pgManager.Delete(repoName)
	if err != nil {
		t.Errorf("failed to delete chart repo test: %v", err)
	}
	m.AssertExpectations(t)
}

// TODO(andresmgot): Write tests for Sync method