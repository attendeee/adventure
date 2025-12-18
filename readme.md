# Adventure app

**Compile**

```bash
    go build -o ./adventure -ldflags="-w -s"

```

**Usage**

Any story must have an **"introduction"** arc. This arc is starting point.

```bash
    ./adventure -path ./assets/gopher.json
```

**Template**

```json
{
    // Each story arc will have a unique key that represents
    // the name of that particular arc.
    "story-arc": {
        "title": "A title for that story arc. Think of it like a chapter title.",
        "story": [
            "A series of paragraphs, each represented as a string in a slice.",
            "This is a new paragraph in this particular story arc."

        ],
        // Options will be empty if it is the end of that
        // particular story arc. Otherwise it will have one or
        // more JSON objects that represent an "option" that the
        // reader has at the end of a story arc.
        "options": [
            {
                "text": "the text to render for this option. eg 'venture down the dark passage'",
                "arc": "the name of the story arc to navigate to. This will match the story-arc key at the very root of the JSON document"

            }
        ]
    },
    ...
}
```
