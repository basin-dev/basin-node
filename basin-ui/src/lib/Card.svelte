<script lang="ts">
    export let title: string;
    export let className: string;
</script>

<div class={`p-4 bg-white w-fit rounded-md border-black border shadow-xl${' '}${className}`}>
    <h2>{title}</h2>
    <slot></slot>
</div>