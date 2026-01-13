<script lang="ts">
    export let title: string;
</script>

<div class="p-4 bg-white w-fit rounded-md border-black border shadow-xl">
    <h2>{title}</h2>
    <slot></slot>
</div>