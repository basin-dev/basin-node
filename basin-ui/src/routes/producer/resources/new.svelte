<script lang="ts">
import Basin from "@util/basin";
let schemaFileInput: HTMLInputElement;

async function registerResource(event: any) {
    let formData = new FormData(event.target);

    if (!schemaFileInput.files || schemaFileInput.files.length < 1) {
        return;
    }

    let adapter = {
    "adapterName": "http",
        "config": {
            "read": {
                "body": "nothing",
                "method": "GET",
                "url": formData.get("adapter-http-url")
            },
            "write": {
                "body": "nothing",
                "method": "PUT",
                "url": formData.get("adapter-http-url")
            }
        }
    }

    let permissions: any[] = [];
    let file = await schemaFileInput.files.item(0)!.text();
    let schema = JSON.parse(file);

    let resp = await Basin.register(formData.get("url") as string, adapter, permissions, schema);
    console.log("Regsiter Response: ", resp);
}
</script>

<form on:submit|preventDefault={registerResource} class="child:mt-4">
    <span>New Resource > 
        <input type="text" placeholder="basin://..." name="url">
    </span>
    <br>

    <label for="schema-file">Schema File</label>
    <input type="file" name="schema-file" id="schema-file" bind:this={schemaFileInput}>

    <h3>Permissions</h3>
    No permissions for now, UCANs first

    <h3 class="text-2xl">Adapter</h3>
    <label for="adapter-http-url">Endpoint URL</label>
    <input type="text" name="adapter-http-url" id="adapter-http-url" placeholder="http://localhost:8000">
    <br>

    <input type="submit" class="cursor-pointer border border-black rounded-lg px-4 py-2" value="Register">
</form>