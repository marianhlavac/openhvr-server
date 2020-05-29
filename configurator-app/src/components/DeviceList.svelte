<script>
    import { createEventDispatcher } from 'svelte'
    import Panel from './Panel.svelte'
    import Button from './Button.svelte'
    import DeviceItem from './DeviceItem.svelte'
    import { fetchDeviceList } from '../api.js'

    const dispatch = createEventDispatcher()
    var handleNew = () => dispatch('newClicked', true)
    let devices = []

    $: (async () => {
        devices = await fetchDeviceList()
    })()
</script>

<style>
    h1 {
        font-size: 1.5em;
        font-weight: 300;
        margin: 0;
        padding: 0;
        margin-bottom: 1em;
    }

    ul.device-list {
        list-style: none;
        margin: 0;
        padding: 0;
    }

    li.device-item {
        margin: 0;
        padding: 0.5rem 0;
    }

    li.device-item:not(:last-child) {
        border-bottom: solid thin #ddd;
    }

    .new-device {
        position: absolute;
        top: 1.5rem;
        right: 1.5rem;
    }
</style>

<Panel>
    <h1>Devices</h1>
    <div class="new-device">
        <Button click={handleNew}>Add Device</Button>
    </div>
    <ul class="device-list">
        {#each devices as device}
            <DeviceItem {device} />
        {:else}
        Loading device list...
        {/each}
    </ul>
</Panel>

