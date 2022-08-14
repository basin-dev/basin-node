<script lang="ts">
    import Table from "@lib/Table.svelte";
    import {BasinSDK} from "@sdk/index";
    let resources: any[] = [];

    (async () => {
        let basin = new BasinSDK("http://localhost:8555");
        let url = "basin://tydunn.com.twitter.followers";
        resources = await basin.read(url);
        console.log('resources: ', resources.length);
    })()
</script>

<h2>Resources</h2>

<Table cols={["URL", "Health", "Revenue", "Subscribers"]} data={resources.map((resource) => {
    return [resource.follower.userLink, resource.follower.accountId, "$100", 12];
})}></Table>