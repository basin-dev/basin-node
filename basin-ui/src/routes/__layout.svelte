<script lang="ts">
    import TopBar from "@lib/TopBar.svelte";
import "../app.css";
</script>

<div>
    <TopBar></TopBar>
    <div class="p-8 overflow-x-hidden"><slot></slot></div>
</div>