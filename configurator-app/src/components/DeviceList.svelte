<script>
    import { onMount, createEventDispatcher } from 'svelte'
    import Panel from './Panel.svelte'
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
    ul.device-list {
        list-style: none;
    }
</style>

<Panel>
    <h1>Devices</h1>
    <button on:click={() => {handleEdit(null)}}>Create New Device</button>
    <ul class="device-list">
        {#each devices as device}
            <li>
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

