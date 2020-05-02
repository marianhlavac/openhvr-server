<script>
    import { onMount } from 'svelte'
    import Panel from './Panel.svelte'
    import { fetchDevice, fetchDrivers, fetchTypes, updateDevice } from '../api.js'

    export let deviceId
    let device = {}
    let drivers = []
    let types = []

    $: (async() => {
        drivers = await fetchDrivers()
        types = await fetchTypes()
        if (deviceId != null) {
            device = await fetchDevice(deviceId)
        }
    })()

    function handleSubmit() {
        device.EffectType = parseInt(device.EffectType)
        updateDevice(deviceId, device)
        window.location.reload()
    }
</script>

<style>
    ul.device-list {
        list-style: none;
    }
</style>

<Panel>
    <h1>
    {#if deviceId}
    Editing {device.Name} (ID {deviceId})
    {:else}
    Creating new device
    {/if}
    </h1>
    <form on:submit|preventDefault={handleSubmit}>
        <label for="name">Device Name</label>
        <input id="name" bind:value={device.Name} />
        <label for="etype">Effect Type</label>
        <select id="etype" bind:value={device.EffectType}>
            {#each types as type}
            <option value="{type.Id}">Type {type.Name}</option>
            {/each}
        </select>
        <label for="type">Device Type</label>
        <select id="type" bind:value={device.Type}>
            {#each drivers as driver}
            <option value="{driver}">{driver}</option>
            {/each}
        </select>
        <label for="type">Connection URI</label>
        <input id="curi" type="text" bind:value={device.ConnectorUri} />
        <label for="type">Connection Param</label>
        <input id="curi" type="text" bind:value={device.ConnectorParam} />
        <label for="px">Position Vector</label>
        <input id="px" type="number" step="any" bind:value={device.LocationX} />
        <input id="py" type="number" step="any" bind:value={device.LocationY} />
        <input id="pz" type="number" step="any" bind:value={device.LocationZ} />
        <label for="px">Direction Vector</label>
        <input id="dx" type="number" step="any" bind:value={device.DirectionX} />
        <input id="dy" type="number" step="any" bind:value={device.DirectionY} />
        <input id="dz" type="number" step="any" bind:value={device.DirectionZ} />
        <label for="px">Direction Spread</label>
        <input id="px" type="number" step="any" bind:value={device.DirectionSpread} />
        &deg; (set to 0&deg; when non-directional)
        <br />
        <button type="submit">Save</button>
    </form>
</Panel>

