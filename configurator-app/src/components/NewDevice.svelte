<script>
    import { createEventDispatcher } from 'svelte'
    import Panel from './Panel.svelte'
    import Button from './Button.svelte'
    import EditableField from './EditableField.svelte'
    import { updateDevice, fetchDrivers, fetchTypes } from '../api.js'

    const dispatch = createEventDispatcher()
    let device = { 
        EffectType: 1,
        LocationX: 0, LocationY: 0, LocationZ: 0,
        DirectionX: 1, DirectionY: 0, DirectionZ: 0,
        DirectionSpread: 45,
    }
    let drivers = []
    let types = []

    $: (async() => {
        drivers = await fetchDrivers()
        types = await fetchTypes()
    })()

    function handleSubmit() {
        device.EffectType = parseInt(device.EffectType)
        updateDevice(null, device)
        dispatch('newClicked', false)
    }

    function handleCancel() {
        dispatch('newClicked', false)
    }

    const typesAssetsMap = {
        0: 'icons/type-generic.png',
        2: 'icons/type-wind.png',
        3: 'icons/type-heat.png'
    }

    function getTypeAssetIcon(id) {
        return id in typesAssetsMap ? typesAssetsMap[id] : typesAssetsMap[0]
    }
</script>

<style>
    h1 {
        font-size: 1.5em;
        font-weight: 300;
        margin: 0;
        padding: 0;
        margin-bottom: 1em;
    }

    dl {
        display: flex;
        flex-wrap: wrap;
        justify-content: space-between;
    }

    .prop-col {
        flex: 0 0 47%;
    }

    .prop-row {
        flex: 0 0 100%;
    }

    @media (max-width: 760px) {
        dl {
            display: block;
        }
    }

    dt {
        margin-top: 2em;
        padding: 0.25em 0;
        font-weight: 600;
    }

    dd {
        padding: 0.25em 0;
        margin: 0;
    }

    dd input,
    dd select {
        width: 100%;
    }

    .short-inputs input {
        width: 4.5em;
    }

    .type-choices {
        display: flex;
    }

    .type-choice input {
        visibility: hidden;
    }

    .type-choice label {
        display: block;
        width: 7rem;
        height: 10rem;
        margin: 0 1rem 1rem 0;
        border: thin solid #eee;
        border-radius: 1rem;
        padding: 1rem;
        text-align: center;
    }

    .type-choice label img {
        margin: 1.5em 0;
    }

    .type-choice input:checked + label {
        border: thin solid #333;
        background: #fcfcfc;
        font-weight: 600;
    }

    @media (max-width: 760px) {
        .type-choices {
            display: block;
        }

        .type-choice label {
            margin: 0;
            margin-bottom: 0.1em;
            width: auto;
            height: 2rem;
            line-height: 2rem;
            text-align: left;
        }

        .type-choice label img {
            float: left;
            margin: 0em;
            margin-right: 2rem;
            height: 2rem;
            width: 2rem;
        }
    }
</style>

<Panel>
    <h1>Add Device</h1>
    <form on:submit|preventDefault={handleSubmit}>
        <dl>
            <span class="prop-row">
                <dt><label for="name">Device Name</label></dt>
                <dd><input id="name" bind:value={device.Name} required /></dd>
            </span>
            <span class="prop-row">
                <dt><label for="etype">Effect Type</label></dt>
                <dd class="type-choices">
                    {#each types as type}
                    <div class="type-choice">
                        <input id={type.Id+'-type'} type="radio" bind:group={device.EffectType} value={type.Id}>
                        <label for={type.Id+'-type'}>
                            <img src={getTypeAssetIcon(type.Id)} alt="Type Icon" width="64" />
                            <div>Type {type.Name}</div>
                        </label>
                    </div>
                    {/each}
                </dd>
            </span>
            <span class="prop-col">
                <dt><label for="dtype">Device Type (Driver)</label></dt>
                <dd>
                    <select id="dtype" bind:value={device.Type} required>
                    {#each drivers as driver}
                    <option value="{driver}">{driver}</option>
                    {/each}
                    </select>
                </dd>
            </span>
            <span class="prop-col">
                <dt><label for="curi">Connection URI (IP Address)</label></dt>
                <dd>
                    <input id="curi" type="text" bind:value={device.ConnectorUri} required />
                </dd>
            </span>
            <span class="prop-col">
                <dt><label for="cparam">Connection Param</label></dt>
                <dd>
                    <input id="cparam" type="text" bind:value={device.ConnectorParam} />
                </dd>
            </span>
            <span class="prop-col">
                <dt><label for="px">Location Vector</label></dt>
                <dd class="short-inputs">
                    <input id="px" type="number" step="any" bind:value={device.LocationX} required />
                    <input id="py" type="number" step="any" bind:value={device.LocationY} required />
                    <input id="pz" type="number" step="any" bind:value={device.LocationZ} required />
                </dd>
            </span>
            <span class="prop-col">
                <dt><label for="dx">Direction Vector and Spread</label></dt>
                <dd class="short-inputs">
                    <input id="dx" type="number" step="any" bind:value={device.DirectionX} required />
                    <input id="dy" type="number" step="any" bind:value={device.DirectionY} required />
                    <input id="dz" type="number" step="any" bind:value={device.DirectionZ} required />
                    <input id="ds" type="number" step="any" bind:value={device.DirectionSpread} />
                    &deg;
                </dd>
            </span>
        </dl>
        <Button click={handleCancel} appearance="light">Cancel</Button>
        <Button type="submit">Save</Button>
    </form>
</Panel>
