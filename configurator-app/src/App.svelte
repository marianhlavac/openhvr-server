<script>
	import Header from './components/Header.svelte'
	import DeviceList from './components/DeviceList.svelte'
	import SystemStatus from './components/SystemStatus.svelte'
	import NewDevice from './components/NewDevice.svelte'

	var creatingNewDevice = false

	function handleNewDeviceMode(ev) {
		creatingNewDevice = ev.detail
	}
</script>

<style>
	:global(body) {
		background: #f6f6f6;
	}

	main {
		max-width: 960px;
		margin: 4rem auto;
	}
</style>


<main>
	<Header />
	{#if creatingNewDevice}
	<NewDevice on:newClicked={handleNewDeviceMode} />
	{:else}
	<SystemStatus />
	<DeviceList on:newClicked={handleNewDeviceMode} />
	{/if}
</main>
