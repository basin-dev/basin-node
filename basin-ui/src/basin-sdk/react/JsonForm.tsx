import React from "react";
import {JsonForms} from "@jsonforms/react"
import {materialRenderers, materialCells} from "@jsonforms/material-renderers"

const schema = {
    "type": "array",
    "items": {
        "type": "object",
        "properties": {
            "follower": {
                "type": "object",
                "properties": {
                    "accountId": {
                        "type": "string"
                    },
                    "userLink": {
                        "type" : "string"
                    }
                },
                "required": ["accountId", "userLink"]
            }
        },
        "required": ["follower"]
    }
};

const Forms = () => {
  return (
    <div>
        <JsonForms
            schema={schema}
            // uischema={uischema}
            data={[]}
            renderers={materialRenderers}
            cells={materialCells}
            // onChange={({ errors, data }) => setData(data)}
        />
    </div>
  )
}

export default Forms
