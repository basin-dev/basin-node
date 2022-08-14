<script lang="ts">
    import Loader from "@lib/Loader.svelte";
    import Basin from "@util/basin";
    
    let data: Promise<any> = Basin.read("basin://tydunn.com.twitter.followers");
</script>

<h1> Basin </h1>

<h3>Ty's Followers</h3>
<h4>Basin resource: basin://tydunn.com.twitter.followers</h4>

{#await data then data}
    <pre id="json">{JSON.stringify(data, null, 2)}</pre>
{/await}