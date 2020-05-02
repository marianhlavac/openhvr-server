<script>
    import { onMount, createEventDispatcher } from 'svelte'
    import Panel from './Panel.svelte'
    import Button from './Button.svelte'
    import { fetchDeviceList } from '../api.js'
    import { deleteDevice } from '../api.js'

    const dispatch = createEventDispatcher()
    let devices = []

    $: (async () => {
        devices = await fetchDeviceList()
    })()

    function handleEdit(id) {
		dispatch('deviceClicked', {id})
    }

    function handleDelete(id) {
        deleteDevice(id)
        window.location.reload()
    }
</script>

<style>
    h1 {
        font-size: 1.5em;
        font-weight: 300;
        text-transform: lowercase;
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
        <Button click={() => {handleEdit(null)}}>New Device</Button>
    </div>
    <ul class="device-list">
        {#each devices as device}
            <li class="device-item">
                <h2>{device.Name} (ID {device.Id})</h2>
                <dl>
                    <dt>Type</dt>
                    <dd>{device.Type}</dd>
                    <dt>URI</dt>
                    <dd>{device.ConnectorUri}</dd>
                    <dt>Pose</dt>
                    <dd>
                        [{device.LocationX}, {device.LocationY}, {device.LocationZ}],
                        [{device.DirectionX}, {device.DirectionY}, {device.DirectionZ}]
                    </dd>
                    <dt>Spread</dt>
                    <dd>
                        {device.DirectionSpread}&deg;
                    </dd>
                </dl>
                <button on:click={() => {handleEdit(device.Id)}}>Edit</button>
                <button on:click={() => {handleDelete(device.Id)}}>Delete</button>
            </li>
        {:else}
        Loading device list...
        {/each}
    </ul>
</Panel>

