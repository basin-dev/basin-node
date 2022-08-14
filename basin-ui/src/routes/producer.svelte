<script lang="ts">
    export let resources: ProducerResource;
</script>

<h2>Resources</h2>

<table>
    <thead>
        <tr>
            <th>URL</th>
            <th>Health</th>
            <th>Revenue</th>
            <th>Subscribers</th>
        </tr>
    </thead>
    <tbody>
        {#each resources as resource}
            
        {/each}
    </tbody>
</table>