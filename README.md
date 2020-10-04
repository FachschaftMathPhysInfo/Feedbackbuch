# Feedbackbuch

In den letzten Jahren (vor 2020) gab es beim Mathematischen Vorkurs der Fachschaft MathPhysInfo immer ein großes grünes Papierbuch, in das die Teilnehmenden Kritik, Lob und Anerkennung an die Organisator:innen und Vortragenden loswerden konnten.
Da in der aktuellen Pandemiesituation weder diese Veranstaltungen in Präsenz stattfinden können noch es verantwortungsvoll wäre ein Buch durch die Reihen gehen zu lassen wird diese Anwendung den Ersatz für die bisher analoge Variante darstellen.

Anforderungen die Dabei in der Reihenfolge für die Anwendung wichtig sind:
1. **Anonymität der Teilnehmer:innen:**
   Um möglichst viel ehrliches Feedback zu erhalten wollen wir nicht, dass die Teilnehmer:innen sich vor ihren Komiliton:innen outen müssen.

2. **Darstellung von beliebigem Feedback:**
   Handschriftlich kann man in einem Buch vieles darstellen, was in Textform vielleicht nicht ganz so einfach ist. Da die Teilnehmerinnen aber potentiell auch Smileys malen oder (fachliches) Feedback zu den Mathematischen Vorträgen geben, ist es wichtig, dass das digitale Feedbackbuch auch das unterstützt.

3. **Interaktion zwischen den Teilnehmer:innen:**
   Viele Tools würden uns die bisherigen Features erlauben, würden aber noch eine zusätzliche Auswertung und aggregation der Ergebnisse erfordern. Wir lassen in diesem Fall die Teilnehmer:innen selbst die Gewichtung vornehmen, in dem es die Möglichkeit gibt andere Kommentare hochzuvoten. Da das Feedbackbuch grundsätzlich anonym sein soll, verlassen wir uns an dieser Stelle natürlich darauf, dass kein Teilnehmer sich die Anonymität zu nutze macht und seinem eigenen Kommentar 1000 upvotes gibt.


## Project Setup

### Project setup für Production (via Docker :whale::package:)
```
docker-compose up
```

Besuche `http://localhost:8080` und die Anwendung sollte für dich laufen :smile:

### Project setup for development (via :construction: `npm` :construction:)
```
npm install
```

#### Compiles and hot-reloads for development
```
npm run serve
```