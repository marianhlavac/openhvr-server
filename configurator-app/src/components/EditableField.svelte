<script>
    export let editing = false
    export let value
    export let rdown = false

    let displayValue = rdown ? Math.round(value * 100) / 100 : value
</script>

{#if editing}
<slot />
{:else}
<span>{displayValue}</span>
{/if}
