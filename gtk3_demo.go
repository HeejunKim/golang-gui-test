package main

import (
	"fmt"
	"log"

	"github.com/gotk3/gotk3/gtk"
)

func setup_window(title string) *gtk.Window {
	win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	if err != nil {
		log.Fatal("Unable to create window:", err)
	}
	win.SetTitle(title)
	win.Connect("destroy", func() {
		gtk.MainQuit()
	})
	win.SetDefaultSize(400, 100)
	win.SetPosition(gtk.WIN_POS_CENTER)
	return win
}

func setup_box(orient gtk.Orientation) *gtk.Box {
	box, err := gtk.BoxNew(orient, 0)
	if err != nil {
		log.Fatal("Unable to create box:", err)
	}
	return box
}

func setup_tview() *gtk.TextView {
	tv, err := gtk.TextViewNew()
	if err != nil {
		log.Fatal("Unable to create TextView:", err)
	}
	return tv
}

func setup_label(label string) *gtk.Label {
	l, err := gtk.LabelNew(label)
	if err != nil {
		log.Fatal("Unable to create label:", err)
	}
	return l
}

func setup_btn(label string, onClick func()) *gtk.Button {
	btn, err := gtk.ButtonNewWithLabel(label)
	if err != nil {
		log.Fatal("Unable to create button:", err)
	}
	btn.Connect("clicked", onClick)
	return btn
}

func get_buffer_from_tview(tv *gtk.TextView) *gtk.TextBuffer {
	buffer, err := tv.GetBuffer()
	if err != nil {
		log.Fatal("Unable to get buffer:", err)
	}
	return buffer
}

func get_text_from_tview(tv *gtk.TextView) string {
	buffer := get_buffer_from_tview(tv)
	start, end := buffer.GetBounds()

	text, err := buffer.GetText(start, end, true)
	if err != nil {
		log.Fatal("Unable to get text:", err)
	}
	return text
}

func set_text_in_tview(tv *gtk.TextView, text string) {
	buffer := get_buffer_from_tview(tv)
	buffer.SetText(text)
}

func main() {
	gtk.Init(nil)

	win := setup_window("Simple Example")
	box := setup_box(gtk.ORIENTATION_VERTICAL)
	win.Add(box)

	label1 := setup_label("label1")
	box.Add(label1)

	tv1 := setup_tview()
	//set_text_in_tview(tv1, "Hello there 1!")
	box.PackStart(tv1, true, true, 0)

	label2 := setup_label("label2")
	box.Add(label2)

	tv2 := setup_tview()
	//set_text_in_tview(tv2, "Hello there 2!")
	box.PackStart(tv2, true, true, 0)

	btn := setup_btn("Submit", func() {
		text1 := get_text_from_tview(tv1)
		text2 := get_text_from_tview(tv2)
		fmt.Println(text1)
		fmt.Println(text2)
	})
	box.Add(btn)

	win.ShowAll()
	gtk.Main()
}
