<script lang="ts">
    import Loader from "@lib/Loader.svelte";
    import Table from "@lib/Table.svelte";
    import Basin from "@util/basin";
    import {parseUrl} from "@sdk/helpers";
    
    let resources: Promise<string[]> = Basin.read("basin://did:key:z6MkoYyGsB9WLBmf12RrcBdai1UPcDcyvNWcMQdRpXzzfo4H.basin.producer.sources");
</script>

<h2>Resources</h2>

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