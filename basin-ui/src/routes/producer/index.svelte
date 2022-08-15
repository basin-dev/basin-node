<script lang="ts">
    import Loader from "@lib/Loader.svelte";
    import Table from "@lib/Table.svelte";
    import Basin from "@util/basin";
    import {parseUrl} from "@sdk/helpers";
    
    let resources: Promise<string[]> = Basin.read("basin://did:key:z6MkoYyGsB9WLBmf12RrcBdai1UPcDcyvNWcMQdRpXzzfo4H.basin.producer.sources");
</script>

<div class="flex">
    <h2 class="text-2xl">Resources</h2>
    <a class="ml-auto" href="/producer/resources/new"><button class="border border-black cursor-pointer rounded-full px-4 py-2 hover:bg-black hover:border-white hover:text-white">New Resource +</button></a>
</div>
<br>

{#await resources}
    <Loader></Loader>
{:then resources} 
<Table cols={["URL", "Health", "Revenue", "Subscribers"]} data={resources.map((resource) => {
    let parsed = parseUrl(resource);
    // Filtering out metadata resources
    if (parsed.domain.startsWith("meta.")) {
        return null;
    }
    return [resource, "?", "$100", 12];
})} onRowClick={(rowData) => {
    document.location = `/producer/resources/resource?url=${encodeURIComponent(rowData[0])}`;
}}></Table>
{/await}