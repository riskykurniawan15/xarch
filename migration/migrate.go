package migration

func ExecSchema() []string {
	// type file name in directory schema
	return []string{
		"20220522135600_users.sql",
		"20220523184621_notes.sql",
	}
}

func ExecSeeder() []string {
	// type file name in directory seeder
	return []string{
		"20220522143404_users_seed.sql",
	}
}
