package main

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
	"net"
	"strconv"
	"strings"
)

func handleRequest(conn net.Conn) {
	// Daten vom Client empfangen
	requestLines := make([]string, 0) //liste von strings mit 0 installisiert
	for {
		line, err := readLine(conn) //readline ist neue methode
		if err != nil {
			//fmt.Println("Error reading request:", err)
			return
		}
		if line == "" {
			break
		}
		requestLines = append(requestLines, line)
	}
	//for i, line := range requestLines {
	//	println(i, line)
	//} zum testen

	headers = make(map[string]string)
	lines := strings.Split(requestLines[0], "\n")
	firstLineParts := strings.Split(lines[0], " ") //erste zeile so teilen GET  /  HTTP1.1
	method = firstLineParts[0]                     //lese das erste Part
	path = firstLineParts[1]
	_ = firstLineParts[2] //httpVersion protokoll
	for _, line := range requestLines {
		parts := strings.SplitN(line, ": ", 2)
		if len(parts) == 2 {
			key := parts[0]
			value := parts[1]
			headers[key] = value
		}
	}
	//for key, value := range headers {
	//	fmt.Printf("Key: %s: value: %s\n", key, value)
	//} map ausgabe zum testen
	if method == "POST" {
		//println("POST")
		contentLength := -1
		lengthStr := strings.TrimSpace(headers["Content-Length"]) //Content-Length: 123
		length, err := strconv.Atoi(lengthStr)                    // convert string to int
		if err != nil {
			fmt.Println("Invalid Content-Length:", lengthStr)
			return
		}
		contentLength = length
		if contentLength == -1 {
			fmt.Println("Content-Length not found in request")
			return
		}
		// Lese den POSTBody und zählen Sie die gelesenen Bytes
		buffer := make([]byte, contentLength)
		bytesRead, err := io.ReadFull(conn, buffer) //Post HTML /main /r/r/n

		if err != nil {
			fmt.Println("Error reading request body:", err)
			return
		}

		if bytesRead != contentLength {
			fmt.Println("Incomplete request body")
			return
		}
		//println("BOdyBuffer: ", string(buffer), " Bufferlegth: ", len(buffer), " bytesRead: ", bytesRead) //zum testen
		handleAnfrage(string(buffer), conn)
	} else if method == "GET" {
		handleAnfrage("", conn) // "" Es Gibt kein POST Data
	} else {
		for i, line := range requestLines {
			println("receive: ", i, line)
		}
		badRequest(conn, "Bad Request")
	}
	defer closeConnection(conn)
}
func closeConnection(conn net.Conn) {
	err := conn.Close()
	if err != nil {
		return
	}
}

func handleKunde(conn net.Conn) {
	t := template.Must(template.ParseFiles("./html/kunde.html"))
	data := Kunde{
		Kontostand: KontoStand, //fmt.Sprintf("%.2f", )
	}
	buf := new(bytes.Buffer)
	err := t.Execute(buf, data)
	//err = t.Execute(os.Stdout, data) //for test stdout print on html page on Screen
	htmlStr := buf.String()
	response := "HTTP/1.1 200 OK\r\n" +
		"Content-Type: text/html\r\n" +
		"\r\n" + htmlStr
	// Antwort senden
	_, err = conn.Write([]byte(response))
	if err != nil {
		panic(err)
	}
}

func handleInventory(conn net.Conn) {

	t := template.Must(template.ParseFiles("./html/inventory.html"))
	data := struct {
		Inventar string
	}{Inventar: fmt.Sprintf("%.2f", Inventor)}
	buf := new(bytes.Buffer)
	err := t.Execute(buf, data)
	//err = t.Execute(os.Stdout, data) //for test stdout print on html page on Screen
	htmlStr := buf.String()
	response := "HTTP/1.1 200 OK\r\n" +
		"Content-Type: text/html\r\n" +
		"\r\n" + htmlStr
	// Antwort senden
	_, err = conn.Write([]byte(response))
	if err != nil {
		panic(err)
	}
}

func handlePost(postBody string, path string) {
	//println(postBody)
	if path == "/" {
		m := handlePostBody(postBody) //map[string]string
		calculate(m["Aktie"], m["Price"], m["Anzahl"])
	} else if path == "/kunde" {
		m := handlePostBody(postBody) //map[string]string
		//fmt.Println(m["einzahlen"])
		f, _ := strconv.ParseFloat(m["einzahlen"], 8) //typ converting string to float
		//fmt.Println(f, reflect.TypeOf(f))
		KontoStand += f
		Inventor += f
	} else if path == "/inventory" {
		m := handlePostBody(postBody) //map[string]string
		//println("id: ", m["id"], "iban: ", m["iban"], " localhost", m["betrag"])
		if _, ok := m["id"]; ok { //ok means key id existiert => stonieren POST request
			RPCStonierung(m["bank"], m["id"])
		} else { //Überweisung POST request m["bank"]
			f, _ := strconv.ParseFloat(m["betrag"], 8)
			Inventor = Inventor - f
			KontoStand = KontoStand - f
			RPCuberweisung(m["betrag"], m["bank"])
		}
	}
}

// Die methode bearbeitet das PostBody und gibt map zurück
func handlePostBody(postBody string) map[string]string {
	// Füge Einträge zur Map basierend auf dem Eingabe-String hinzu //
	fields := strings.Split(postBody, "&")
	m := make(map[string]string)
	for _, f := range fields {
		kv := strings.Split(f, "=")
		m[kv[0]] = kv[1]
	}
	return m
}

// diese mehtode bearbeitet die main anfrage localhost:6543
func HandlGet(conn net.Conn) {
	t := template.Must(template.ParseFiles("./html/index.html"))
	var data []Data
	for k, v := range myAktienPrice {
		d := Data{Aktie: k, Price: v, LastUpdate: myAktienTime[k]}
		data = append(data, d)
	}
	buf := new(bytes.Buffer)
	err := t.Execute(buf, data)
	//err = t.Execute(os.Stdout, data) //for test stdout print on html page on Screen
	htmlStr := buf.String()
	response := "HTTP/1.1 200 OK\r\n" +
		"Content-Type: text/html\r\n" +
		"\r\n" + htmlStr
	// Antwort senden
	_, err = conn.Write([]byte(response))
	if err != nil {
		panic(err)
	}
}

// this Mehtod Decides if the request is GET or POST
// and saves the headers from the HTTP-request in a private map namens headers
// see info.go all Variable are there
func handleAnfrage(postData string, conn net.Conn) {

	//prafix := "\"Chromium\";v=\"112\", \"Google Chrome\";v=\"112\", \"Not:A-Brand\";v=\"99\""
	//prafix := "Google Chrome" && strings.Contains(headers["sec-ch-ua"], prafix)
	if method == "GET" {
		if path == "/" {
			HandlGet(conn)
		} else if path == "/kunde" { //
			handleKunde(conn)
		} else if path == "/inventory" {
			handleInventory(conn)
		} else {
			response := "400 Bad Request - Unsupported Request Type"
			badRequest(conn, response)
		}
	} else if method == "POST" {
		handlePost(postData, path)
		if path == "/" {
			HandlGet(conn)
		} else if path == "/kunde" {
			handleKunde(conn)
		} else if path == "/inventory" {
			handleInventory(conn)
		}
	} else {
		response := "Browser or Methode is not supported! just Google Chrome"
		badRequest(conn, response)
		println("Browser or Methode is not supported! just Google Chrome: ")
	}
	defer closeConnection(conn)

}

func badRequest(conn net.Conn, response string) {
	_, err := fmt.Fprintf(conn, "HTTP/1.1 400 Bad Request\r\n"+
		"Content-Type: text/plain\r\n"+
		"Content-Length: %d\r\n"+
		"Connection: close\r\n"+
		"\r\n"+
		"%s", len(response), response)
	if err != nil {
		return
	}
}

func readLine(conn net.Conn) (string, error) {
	line := ""
	buffer := make([]byte, 1)
	for {
		_, err := conn.Read(buffer)
		if err != nil {
			return "", err
		}
		if buffer[0] == '\r' {
			continue
		}
		if buffer[0] == '\n' {
			break
		}
		line += string(buffer[0])
	}
	return line, nil
}
