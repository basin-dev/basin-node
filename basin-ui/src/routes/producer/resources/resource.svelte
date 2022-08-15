<script lang="ts">
    import Card from "@lib/Card.svelte";
    import Basin from "@util/basin";
    import {page} from "$app/stores";
import Loader from "@lib/Loader.svelte";
    let url = $page.url.searchParams.get("url") as string;

    let schema: Promise<any> = Basin.readMetadata(url, "schema");
    let permissions: Promise<any> = Basin.readMetadata(url, "permissions");
    let adapter: Promise<any> = Basin.readMetadata(url, "adapter");
    let data: Promise<any> = Basin.read(url);
</script>

<h1>Resource > Twitter Followers</h1>

<div class="flex gap-8 mt-4">
    <Card title="Health"></Card>
    <Card title="Royalties"></Card>
    <Card title="Subscribers"></Card>
</div>

{#await schema}
    <Loader></Loader>
{:then schema}
<h3 class="text-2xl">Schema</h3>
    <pre>{JSON.stringify(schema, null, 2)}</pre>
{/await}

{#await permissions}
    <Loader></Loader>
{:then permissions} 
<h3 class="text-2xl">Permissions</h3>
    <pre>{JSON.stringify(permissions, null, 2)}</pre>
{/await}

{#await adapter}
    <Loader></Loader>
{:then adapter}
<h3 class="text-2xl">Adapter</h3> 
    <pre>{JSON.stringify(adapter, null, 2)}</pre>
{/await}


{#await data}
<Loader></Loader>
{:then data} 
<h3 class="text-2xl">Data</h3>
<pre>{JSON.stringify(data, null, 2)}</pre>
{/await}