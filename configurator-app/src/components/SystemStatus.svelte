<script>
    import Panel from './Panel.svelte'
    import { fetchStatus } from '../api.js'

    var status = 'loading'
    var panelVariant = 'loading'

    $: {
        fetchStatus().then(systemStatus => {
            var systemWorking = systemStatus.result == 'done'
            status = systemWorking ? 'responsive' : 'unreachable'
            panelVariant = systemWorking ? 'positive' : 'error'
        }).catch(err => {
            status = 'unreachable'
            panelVariant = 'error'
        })
    }
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

<Panel variant={panelVariant}>
    <div class="system-status">
        System is <strong>{status}</strong>.
    </div>
    <div class="system-props">
        localhost:47023
    </div>
</Panel>

