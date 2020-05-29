<script>
    import { onMount } from 'svelte'
    import Button from './Button.svelte'
    import EditableField from './EditableField.svelte'
    import { fetchDevice, deleteDevice, fetchDrivers, fetchTypes, updateDevice, testDevice } from '../api.js'

    export let device
    let drivers = []
    let types = []
    let typesMap = {}

    function handleSubmit() {
        device.EffectType = parseInt(device.EffectType)
        updateDevice(device.Id, device)
        handleChangeEditMode(false)()
    }

    function handleDelete(id) {
        deleteDevice(id)
        window.location.reload()
    }

    var editing = false;

    var handleChangeEditMode = (to) => () => {
        editing = to;
    }
    var handleTestDevice = (id) => () => testDevice(id)

    $: (async() => {
        types = await fetchTypes()
        drivers = await fetchDrivers()
    })()

    $: typesMap = types.reduce((acc, i) => ({ ...acc, [i.Id]: i.Name }), {})
</script>

<style>
    li {
        margin: 0;
        padding: 1rem 0;
    }

    img {
        vertical-align: middle;
    }

    h2 {
        display: inline;
        margin-left: 0.6em;
        vertical-align: middle;
    }

    dl {
        display: flex;
        flex-wrap: wrap;
        justify-content: space-between;
    }

    .prop-col {
        display: flex;
        flex: 0 0 47%;
    }

    @media (max-width: 760px) {
        dl {
            display: block;
        }
    }

    dt {
        flex: 1 0;
        padding: 0.5em 0;
        font-weight: 600;
    }

    dd {
        padding: 0.5em 0;
    }

    .short-inputs input {
        width: 3em;
    }
</style>

<li class="device-item">
    <form on:submit|preventDefault={handleSubmit}>
        <img src="icons/device.png" width="24" alt="Device Icon" />
        {#if editing}
        <input id="name" bind:value={device.Name} />
        {:else}
        <h2>{device.Name} (ID {device.Id})</h2>
        {/if}
        <dl>
            <span class="prop-col">
                <dt><label for="etype">Effect Type</label></dt>
                <dd>
                    {#if editing}
                    <EditableField {editing} value={typesMap[device.EffectType]}>
                        <select id="etype" bind:value={device.EffectType}>
                        {#each types as type}
                        <option value="{type.Id}">Type {type.Name}</option>
                        {/each}
                        </select>
                    </EditableField>
                    {:else}
                    <span>{typesMap[device.EffectType]}</span>
                    {/if}
                </dd>
            </span>
            <span class="prop-col">
                <dt><label for="dtype">Device Type</label></dt>
                <dd>
                    <EditableField {editing} value={device.Type}>
                        <select id="dtype" bind:value={device.Type}>
                        {#each drivers as driver}
                        <option value="{driver}">{driver}</option>
                        {/each}
                        </select>
                    </EditableField>
                </dd>
            </span>
            <span class="prop-col">
                <dt><label for="curi">Connection URI</label></dt>
                <dd>
                    <EditableField {editing} value={device.ConnectorUri}>
                        <input id="curi" type="text" bind:value={device.ConnectorUri} />
                    </EditableField>
                </dd>
            </span>
            <span class="prop-col">
                <dt><label for="cparam">Connection Param</label></dt>
                <dd>
                    <EditableField {editing} value={device.ConnectorParam}>
                        <input id="cparam" type="text" bind:value={device.ConnectorParam} />
                    </EditableField>
                </dd>
            </span>
            <span class="prop-col">
                <dt><label for="px">Location Vector</label></dt>
                <dd class="short-inputs">
                    <EditableField {editing} value={device.LocationX} rdown>
                        <input id="px" type="number" step="any" bind:value={device.LocationX} />
                    </EditableField>
                    <EditableField {editing} value={device.LocationY} rdown>
                        <input id="py" type="number" step="any" bind:value={device.LocationY} />
                    </EditableField>
                    <EditableField {editing} value={device.LocationZ} rdown>
                        <input id="pz" type="number" step="any" bind:value={device.LocationZ} />
                    </EditableField>
                </dd>
            </span>
            <span class="prop-col">
                <dt><label for="dx">Direction Vector and Spread</label></dt>
                <dd class="short-inputs">
                    <EditableField {editing} value={device.DirectionX} rdown>
                        <input id="dx" type="number" step="any" bind:value={device.DirectionX} />
                    </EditableField>
                    <EditableField {editing} value={device.DirectionY} rdown>
                        <input id="dy" type="number" step="any" bind:value={device.DirectionY} />
                    </EditableField>
                    <EditableField {editing} value={device.DirectionZ} rdown>
                        <input id="dz" type="number" step="any" bind:value={device.DirectionZ} />
                    </EditableField>
                    <EditableField {editing} value={device.DirectionSpread}>
                        <input id="ds" type="number" step="any" bind:value={device.DirectionSpread} />
                    </EditableField>
                    &deg;
                </dd>
            </span>
        </dl>
        {#if editing}
        <Button type="submit">Save</Button>
        <Button click={() => {handleDelete(device.Id)}} appearance="light">Delete</Button>
        <Button click={handleChangeEditMode(false)} appearance="light">Cancel</Button>
        {:else}
        <Button click={handleTestDevice(device.Id)} appearance="light">Test</Button>
        <Button click={handleChangeEditMode(true)}>Edit</Button>
        {/if}
    </form>
</li>
