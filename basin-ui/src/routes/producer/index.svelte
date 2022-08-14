<script lang="ts">
    import Loader from "@lib/Loader.svelte";
    import Table from "@lib/Table.svelte";
    import Basin from "@util/basin";
    
    let resources: Promise<any[]> = Basin.read("basin://tydunn.com.twitter.followers");
</script>

<h2>Resources</h2>

{#await resources}
    <Loader></Loader>
{:then resources} 
<Table cols={["URL", "Health", "Revenue", "Subscribers"]} data={resources.map((resource) => {
    return [resource.follower.userLink, resource.follower.accountId, "$100", 12];
})}></Table>
{/await}