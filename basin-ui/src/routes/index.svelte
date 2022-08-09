<script lang="ts">

    export let data: any;
</script>

<h1> Basin </h1>

<!-- <h3>Producer Sources</h3> -->
<h3>Ty's Followers</h3>
<p>{JSON.stringify(JSON.parse(data || ""), null, '\t')}</p>