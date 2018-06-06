package main

import (
	"strings"
	"testing"
)

var data = initStorage()

func TestWhatShouldIdoSelect(t *testing.T) {
	test := whatShouldIDo("select 1", data)
	if test != 0 {
		t.Errorf("Select do not work")
	}
}

func TestWhatShouldIdoTestInsert(t *testing.T) {
	test := whatShouldIDo("insert bob", data)
	if test != 0 {
		t.Errorf("Insert do not work")
	}
}

func TestWhatShouldIdoShowAll(t *testing.T) {
	newStorage := initStorage()
	for range [4]int{1, 2, 3, 4} {
		newStorage.insertItem("insert Bob")
	}
	test := whatShouldIDo("select all", newStorage)
	if test != 0 {
		t.Error("Show all do not work")
	}
}

func TestSelect(t *testing.T) {
	newStorage := initStorage()
	newStorage.insertItem("insert bob")
	selectSmt := data.selectItem(0)
	if strings.Compare(selectSmt, "bob") != 0 {
		t.Error("Select do not work")
	}
}

func TestInsert(t *testing.T) {
	newStorage := initStorage()
	code := newStorage.insertItem("insert bob")
	if code != 0 {
		t.Error("Error with insert")
	}
}

func TestWhatShouldIdoSyntaxError(t *testing.T) {
	newStorage := initStorage()
	test := whatShouldIDo("bob", newStorage)
	if test != syntaxError {
		t.Error("Syntax parser do not work")
	}
}

func TestCreateTable(t *testing.T) {
	newStorage := initStorage()
	test := whatShouldIDo("create table name", newStorage)
	if test != 0 {
		t.Error("Cant create table")
	}

}
