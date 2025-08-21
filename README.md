# dredger-todos-async

_dredger-todos_ dient zum Demonstrieren und Testen von _dredger_. _dredger_ generiert aus OpenAPI- und AsyncAPI-Spezifikationen Go-Code, der ein Gerüst für einen Micro-Service darstellt. Entsprechend den Spezifikationen werden die definierten REST-Endpoints und publish-/subscriber-Funktionen für asynchrone Kommunikation automatisch generiert. Darüber hinaus werden auch auch Funktionalitäten generiert, die ein typischer Microservice benötigt:

- Konfiguration
- Commandline
- Sicherheitsfunktionen, wie ein Policy-Enforment-Point
- Logging
- Monitoring
- Tracing
- Nutzung von templ für HTML-Seiten und deren Lokalisierung
- Support für HTMX, SSE


Die Basis-Funktionalität ist es neue Aufgaben __Hinzufügen__ zu können, diese als erledigt (√) oder nicht erledigt ([]) zu markieren, sowie einzelne Aufgaben zu löschen (🗑️).

Man kann aber Aufgabenliste auch gemeinsam editieren, d.h. mit mehreren Instanzen des Micro-Services, bspw. auf verschiedenen HTTP-Ports, bedient auf zwei oder mehr Webbrowser(-Seiten). Dazu muss man für alle Instanzen das gleiche Label __Setzen__.

## Installation

Zuerst sollte man _dredgerTodos_ vom Quellcode installieren, wobei eine funktionierende _Go_-Umgebung und das Werkzeug _just_ vorausgesetzt werden:

```bash
just required
just install
```

Nun kann man mehrere Instanzen in mehreren Shells (Terminals) starten:

```bash
dredgerTodos -p 9990
dredgerTodos -p 9991
dredgerTodos -p 9992
```

In 3 verschiedenen Browser-Fenstern kann man sich die Aufgabenlisten anschauen:


```bash
http://localhost:9990
http://localhost:9991
http://localhost:9992
```

Als erstes setzt man in allen Fenstern das gleiche Label. Danach kann man in einem beliebigen Fenster die gemeinsame Aufgabenliste pflegen.

## Umsetzung

Initial wurden die OpenAPI- und die AsyncAPI-Spezifikationen erstellt und mit `dredger  generate OpenAPI.yaml AsyncAPI.yaml -o . -f -n dredger-todos` der initiale Code generiert. Schrittweise wurden REST-Handler, Async-Publisher, Async-Subscriber und weitere Anpassungen vorgenommen. Ein besonderer Fokus bilden die _usecases_, die die Business-Logik im Kern abbilden. Der Quellcode orientiert sich generell an der _Clean Architecture_ von _Uncle Bob_.

Bei den _usecases_ wird ersichtlich, dass _dredgerTodos_ eigentlich 2 Micro-Services in einem sind. Mit den REST-Services wird vor allem die Web-Oberfläche mittels HTMX realisiert. Mit SSE kann diese auch asynchron aktualisiert werden. Aktionen auf der Web-Oberfläche ändern den internen Zustand der Aufgabenliste NICHT. Stattdessen werden asynchron Messages an den Message Broker (NATS) _published_. Der andere Teil des Micro-Services hat die Messages _subscribed_ und erst beim Empfang einer Message wird der lokale Zustand der Aufgabenliste aktualisiert und ein _refresh_ der Aufgabenliste auf der Web-Seite mittels SSE signalisiert. Mittels des Session-Management können auch mehrere Nutzende (Nutzung verschiedener Browser) gleichzeitig dredgerTodos nutzen.
