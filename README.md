# jsonschemate

Infer a jsonschema from a json file


## usage

```bash
go install github.com/JackKCWong/go-jsonschema/...@latest

echo '{"hello": "world"}' | jsonsche

# or just

jsonsche input.json
```

output:

```json
{
    "$schema": "https://json-schema.org/draft/2020-12/schema",
    "properties": {
        "hello": {
            "type": "string"
        }
    },
    "additionalProperties": false,
    "type": "object",
    "required": [
        "hello"
    ]
}
```


## how it works

It converts a json to Go struct and then generates a schema from the Go struct.

## caveats

* It marks all properties as required
* It doesn't do well with "null" values
