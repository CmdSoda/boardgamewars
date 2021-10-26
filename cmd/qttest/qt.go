package main

import "os"
import "github.com/therecipe/qt/widgets"

func main() {
	//1. Eröffne eine neue Applikation (also eine Sammlung von Widgets)
	widgets.NewQApplication(len(os.Args), os.Args) // ... übergebe Kommandozeilenparameter (ansonsten "nil")

	//2. Erzeuge das Haupt-Fenster
	window := widgets.NewQMainWindow(nil, 0)
	window.SetWindowTitle("Titel des Hauptfensters")
	window.SetMinimumSize2(400, 100)

	//3. Erzeuge ein Layout ⇒ https://doc.qt.io/qt-5/qlayout.html
	layout := widgets.NewQHBoxLayout()

	//4. Erzeuge ein Widget und weise dem Widget ein Layout zu
	widget := widgets.NewQWidget(nil, 0)
	widget.SetLayout(layout)

	//5-1. Erzeuge eine Eingabezeile
	input := widgets.NewQLineEdit(nil)
	input.SetPlaceholderText("Schreib mal wieder!")

	//5-2. Erzeuge einen Knopf
	button := widgets.NewQPushButton2("Hier klicken", nil)
	button.ConnectClicked(func(checked bool) {
		//6. Weise dem Knopf ein Signal zu, das beim Drücken des Knopfes eine
		//   Funktion ausführen soll
		widgets.QMessageBox_Information(nil, "Titel des Infofensters", input.Text(), widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
	})

	//7.  Weise die Input-Widges dem Layout zu
	layout.AddWidget(input, 0, 0)  // ⇒ http://doc.qt.io/qt-5/qboxlayout.html#addWidget
	layout.AddWidget(button, 0, 0) // ⇒ https://doc.qt.io/qt-5/qboxlayout.html#addWidget

	//8. Dem Haupt-Fenster muß ein Haupt-Widget zugewiesen werden
	window.SetCentralWidget(widget)

	//9. Zeichne das Haupt-Fenster
	window.Show()

	//9. Damit die Applikation dynamisch auf Benutzereingaben/Events reagieren
	//   kann, muß sie in einer Endlosschleife laufen
	widgets.QApplication_Exec()
}
