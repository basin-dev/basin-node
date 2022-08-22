<script lang="ts">
    export let cols: string[];
    export let data: (any[] | null)[];
    export let onRowClick: (rowData: any) => void | undefined;
</script>

<table class="border border-black rounded-md border-collapse">
    <thead>
        <tr>
            {#each cols as col}
                <th class="py-2 px-4">{col}</th>
            {/each}
        </tr>
    </thead>
    <tbody>
        {#each data as tr}
            {#if tr}
                <tr class="border border-black hover:bg-blue-50 cursor-pointer" on:click={() => {
                    if (typeof onRowClick !== "undefined") onRowClick(tr);
                }}>
                    {#each tr as td}
                        <td class="py-2 px-4 text-center">{td}</td>                    
                    {/each}
                </tr>
            {/if}
        {/each}
    </tbody>
</table>