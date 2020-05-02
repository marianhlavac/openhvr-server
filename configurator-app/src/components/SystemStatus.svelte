<script>
    import Panel from './Panel.svelte'

    status = 'loading'
</script>

<style>
    .system-status {
        font-size: 1.5em;
        font-weight: 300;
    }
    .system-props {
        font-size: 0.9em;
        padding: 0.5em 0;
    }
</style>

<Panel {status}>
    <div class="system-status">
        system is <strong>{status}</strong>
    </div>
    <div class="system-props">
        localhost:47023
    </div>
</Panel>

