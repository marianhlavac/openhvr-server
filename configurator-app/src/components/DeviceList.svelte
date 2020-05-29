<script>
    import { createEventDispatcher } from 'svelte'
    import Panel from './Panel.svelte'
    import Button from './Button.svelte'
    import DeviceItem from './DeviceItem.svelte'
    import { fetchDeviceList } from '../api.js'

    const dispatch = createEventDispatcher()
    var handleNew = () => dispatch('newClicked', true)
    let devices = null

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

    .no-devices {
        text-align: center;
        margin: 6em 0 3em 0;
    }

    .no-devices p {
        color: #999;
        font-size: 1.7em;
        font-weight: 500;
    }
</style>

<Panel>
    <h1>Devices</h1>
    <div class="new-device">
        <Button click={handleNew}>Add Device</Button>
    </div>
    <ul class="device-list">
        {#if devices == null}
            Loading device list...
        {:else}
            {#each devices as device}
                <DeviceItem {device} />
            {:else}
                <div class="no-devices">
                    <p>No devices configured.</p>
                    <Button click={handleNew}>Add Device</Button>
                </div>
            {/each}
        {/if}
    </ul>
</Panel>

