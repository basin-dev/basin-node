<script lang="ts">
async function registerResource() {
    
}
</script>

<form on:submit|preventDefault={registerResource}>
    <span>New Resource > 
        <input type="text" placeholder="Resource Title" name="title">
    </span>

    <label for="schema-file">Schema File</label>
    <input type="file" name="schema-file" id="schema-file">

    <h3>Permissions</h3>
    No permissions for now, UCANs first

    <h3>Adapter</h3>
    <label for="adapter-http-url"></label>
    <input type="text" name="adapter-http-url" id="adapter-http-url">

    <input type="submit" value="Register">
</form>