package main

var database []Agenda

func getDatabase() []Agenda {
	if database == nil {
		database = make([]Agenda, 0)
	}
	return database
}

func ListAll() []Agenda {
	return getDatabase()
}

func Save(agenda Agenda) {
	updateData(append(getDatabase(), agenda))
}

func GetById(id int) Agenda {
	return getDatabase()[id]
}

func Delete(id int) bool {
	slice := getDatabase()
	if len(getDatabase()) > id {
		copy(slice[id:], slice[id+1:])
		slice = slice[:len(slice)-1]
		updateData(slice)
		return true
	} else {
		return false
	}
}

func updateData(newDatabase []Agenda) {
	database = newDatabase
}
