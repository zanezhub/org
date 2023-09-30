package main

func main() {
	var org organizer
	org.From = "D:\\Go\\org\\test\\from"
	org.To = "D:\\Go\\org\\test\\des"

	org.Regex = "2023(\\w\\w)"

	org.DirExists()
	org.GetEntries()
	org.ParseFiles()
	org.MakeDir()
	org.Move()
}
