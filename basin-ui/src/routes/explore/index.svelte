<script lang="ts">
    import Card from "@lib/Card.svelte";
    import Basin from "@util/basin";
    // Really want to be getting this list of resources from a curating transformer, but this is a demo
    let featuredResources = [
        "tydunn.com.twitter.followers",
        "tydunn.com.twitter.following"
    ];

    // Get info like icon, description, title, etc... from the sources.

</script>

<div class="grid grid-cols-12 gap-4">
    {#each featuredResources as resourceUrl}
        <Card className="col-span-2 w-full" title={resourceUrl}>
            
        </Card>
    {/each}
</div>

<!-- <h2 class="text-2xl">Popular Transformers</h2>

<h2 class="text-2xl">Popular Schemas</h2>

<h2 class="text-2xl">Popular Datasets</h2> -->

